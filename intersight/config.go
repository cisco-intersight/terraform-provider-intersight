package intersight

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	u "net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	ApiKey        string
	SecretKeyFile string
	Endpoint      string
}

// getComputedDigest accepts json in bytes format
// Returns an encoded form of the json using sha256 hash algorithm.
func getComputedDigest(jsonPayload []byte) string {
	digest := sha256.New()
	digest.Write(jsonPayload)
	finalDigest := "SHA-256=" + base64.StdEncoding.EncodeToString(digest.Sum(nil))
	log.Println("Payload digest", finalDigest)
	return finalDigest
}

// readPrivateKey reads RSA PRIVATE FROM a file, parses and returns an object of rsa PrivateKey
func readPrivateKey(filePath string) (*rsa.PrivateKey, error) {
	privateKeyFile, err := os.Open(filePath)
	if err != nil {
		log.Printf("failed to open secret key file: %s\n", err)
		return nil, err
	}
	privateKey, err := ioutil.ReadAll(privateKeyFile)
	if err != nil {
		log.Printf("failed to read secret key: %s\n", err)
		return nil, err
	}
	err = privateKeyFile.Close()
	if err != nil {
		log.Println("failed while closing file: ", err)
		return nil, err
	}
	block, _ := pem.Decode(privateKey)
	if block == nil {
		log.Println("failed to parse PEM block containing the public key")
		return nil, err
	}
	privateKeyParsed, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic("failed to parse DER encoded public key: " + err.Error())
	}
	return privateKeyParsed, nil
}

// method will be POST, GET, PATCH or DELETE
// target url is the request endpoint e.g. /api/v1/sol/Policies
// return authorization header content and current date if no error occurred, else empty string
func (s *Config) getHTTPSign(method string, targetPath string, digest string) (string, string, error) {
	privateKey, err := readPrivateKey(s.SecretKeyFile)
	if err != nil {
		return "", "", err
	}
	currentDate := strings.Replace(time.Now().UTC().Format(time.RFC1123), "UTC", "GMT", -1)
	host := strings.TrimPrefix(strings.TrimPrefix(s.Endpoint, "https://"), "http://")
	lStringToSign := "(request-target): " + strings.ToLower(method) + " " + strings.ToLower(targetPath) + "\ncontent-type: application/json\ndate: " + currentDate + "\ndigest: " + digest + "\nhost: " + host
	h := sha256.New()
	_, shaerr := h.Write([]byte(lStringToSign))
	if shaerr != nil {
		return "", "", shaerr
	}
	hashed := h.Sum(nil)

	rng := rand.Reader
	signature, signerr := rsa.SignPKCS1v15(rng, privateKey, crypto.SHA256, hashed[:])
	if signerr != nil {
		return "", "", signerr
	}
	encSignature := base64.StdEncoding.EncodeToString(signature)
	authorisationHeader := fmt.Sprintf(`Signature keyId="%s",algorithm="rsa-sha256",headers="(request-target) content-type date digest host",signature="%s"`, s.ApiKey, encSignature)
	return authorisationHeader, currentDate, nil
}

//SendRequest accepts url and payload json. Sends a POST request.
func (s *Config) SendRequest(url string, data []byte) ([]byte, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("URL:>", s.Endpoint+"/api/v1/"+url)

	payloadBytes, err := sanitizeJson(data)
	if err != nil {
		log.Printf("error in sanitizing data. error: %s", err.Error())
		return []byte(""), err
	}
	digest := getComputedDigest(payloadBytes)
	authorizationHeader, currentDate, authErr := s.getHTTPSign(http.MethodPost, "/api/v1/"+url, digest)
	if authErr != nil {
		panic(authErr)
	}
	req, err := http.NewRequest(http.MethodPost, s.Endpoint+"/api/v1/"+url, bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Digest", digest)
	req.Header.Set("Origin", s.Endpoint)
	req.Header.Set("Authorization", authorizationHeader)
	req.Header.Set("Date", currentDate)

	log.Println("===========>", string(payloadBytes))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func() { _ = resp.Body.Close() }()

	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}
	log.Println("response Body:", string(body))
	if resp.StatusCode != http.StatusOK {
		log.Println("response Status:", resp.Status)
		log.Println("response Headers:", resp.Header)
		return body, fmt.Errorf("SendRequest failed. Url %s Status Code: %d Message: %v", url, resp.StatusCode, string(body))
	}
	return body, nil
}

func (s *Config) SendDeleteRequest(url string) ([]byte, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("URL:>", s.Endpoint+"/api/v1/"+url)

	authorizationHeader, currentDate, authErr := s.getHTTPSign(http.MethodDelete, "/api/v1/"+url, getComputedDigest([]byte("")))
	if authErr != nil {
		panic(authErr)
	}
	req, err := http.NewRequest(http.MethodDelete, s.Endpoint+"/api/v1/"+url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Digest", getComputedDigest([]byte("")))
	req.Header.Set("Origin", s.Endpoint)
	req.Header.Set("Authorization", authorizationHeader)
	req.Header.Set("Date", currentDate)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func() { _ = resp.Body.Close() }()

	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}

	log.Println("response Body:", string(body))
	if resp.StatusCode != http.StatusOK {
		log.Println("response Status:", resp.Status)
		log.Println("response Headers:", resp.Header)
		return body, fmt.Errorf("SendDeleteRequest failed. Url %s Status Code: %d Message: %v", url, resp.StatusCode, string(body))
	}

	return body, nil
}

func (s *Config) SendUpdateRequest(url string, data []byte) ([]byte, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("URL:>", s.Endpoint+"/api/v1/"+url)

	payloadBytes, err := sanitizeJson(data)
	if err != nil {
		log.Printf("error in sanitizing data. error: %s", err.Error())
		return []byte(""), err
	}
	digest := getComputedDigest(payloadBytes)
	authorizationHeader, currentDate, authErr := s.getHTTPSign(http.MethodPatch, "/api/v1/"+url, digest)
	if authErr != nil {
		panic(authErr)
	}
	req, err := http.NewRequest(http.MethodPatch, s.Endpoint+"/api/v1/"+url, bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Digest", digest)
	req.Header.Set("Origin", s.Endpoint)
	req.Header.Set("Authorization", authorizationHeader)
	req.Header.Set("Date", currentDate)

	log.Println("===========>", string(payloadBytes))
	log.Printf("Request %+v\n", req)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}

	log.Println("response Body:", string(body))
	if resp.StatusCode != http.StatusOK {
		log.Println("response Status:", resp.Status)
		log.Println("response Headers:", resp.Header)
		return body, fmt.Errorf("SendUpdateRequest failed. Url %s Status Code: %d Message: %v", url, resp.StatusCode, string(body))
	}

	return body, nil
}

// SendGetRequest sends Get request to appliance with or without payload
func (s *Config) SendGetRequest(url string, data []byte) ([]byte, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("SendGetRequest URL:>", s.Endpoint+"/api/v1/"+url)
	var err error
	req, err := http.NewRequest(http.MethodGet, s.Endpoint+"/api/v1/"+url, nil)
	if err != nil {
		return []byte(""), err
	}
	// add GET request query, if exists
	payloadBytes := []byte("")
	if string(data) != "" {
		data, err = sanitizeJson(data)
		if err != nil {
			log.Printf("error in sanitizing data. error: %s", err.Error())
			return []byte(""), err
		}
		req.URL.RawQuery = "$filter=" + (&u.URL{Path: getRequestParams(data)}).String()
	}
	targetURL := strings.TrimPrefix(req.URL.String(), s.Endpoint)
	log.Println("get data source URL", targetURL)
	digest := getComputedDigest(payloadBytes)
	authorizationHeader, currentDate, authErr := s.getHTTPSign(http.MethodGet, targetURL, digest)
	if authErr != nil {
		panic(authErr)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Digest", digest)
	req.Header.Set("Origin", s.Endpoint)
	req.Header.Set("Authorization", authorizationHeader)
	req.Header.Set("Date", currentDate)

	log.Printf("Request: %+v", req)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func() { _ = resp.Body.Close() }()

	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}

	log.Println("response Body:", string(body))
	if resp.StatusCode != http.StatusOK {
		log.Println("response Status:", resp.Status)
		log.Println("response Headers:", resp.Header)
		return body, fmt.Errorf("SendGetRequest failed. Url %s Status Code: %d Message: %v", url, resp.StatusCode, string(body))
	}

	return body, nil
}

func sanitizeJson(in []byte) ([]byte, error) {
	log.Println("Data to be sanitized", string(in))
	var s map[string]interface{}
	err := json.Unmarshal(in, &s)
	if err != nil {
		return []byte(""), err
	}

	readOnlyProps := []string{
		"Ancestors",
		"ClaimedTime",
		"ConnectionStatusLastChangeTime",
		"CreateTime",
		"ModTime",
		"ReleaseTime",
		"ReleaseDate",
		"ImportedTime",
		"LastAccessTime",
		"Owners",
		"ConfigChangeDetails",
		"RunningWorkflows",
		"TimeZone",
		"CleanupTime",
		"EndTime",
		"StartTime",
	}
	for k, v := range s {
		if v == nil || v == "" {
			delete(s, k)
			continue
		}
		for _, p := range readOnlyProps {
			if k == p {
				delete(s, k)
			}
		}
	}
	b, err := json.Marshal(s)
	if err == nil {
		b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
		b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
		b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	}
	log.Println("Sanitized data", string(b))
	return b, err
}

func getRequestParams(in []byte) string {
	var o string
	var s map[string]interface{}
	err := json.Unmarshal(in, &s)
	if err != nil {
		return ""
	}
	for k, v := range s {
		log.Printf("Type: %+v", reflect.TypeOf(v).Kind())
		switch reflect.TypeOf(v).Kind() {
		case reflect.String:
			o += k + " eq '" + v.(string) + "'"
		case reflect.Bool:
			o += k + " eq " + strconv.FormatBool(v.(bool))
		case reflect.Int:
			o += k + " eq " + strconv.FormatInt(v.(int64), 10)
		case reflect.Float64:
			o += k + " eq " + fmt.Sprintf("%f", v.(float64))
		}
		o += " and "
	}
	o = strings.TrimSuffix(o, " and ")
	return o
}
