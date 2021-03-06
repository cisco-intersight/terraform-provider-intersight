/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-06-17T02:04:50-07:00.
 *
 * API version: 1.0.9-1867
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// ConnectorHttpRequest A HTTP request sent by a cloud service to be proxied through a device connector.
type ConnectorHttpRequest struct {
	ConnectorBaseMessage
	// Contents of the request body to send for PUT/PATCH/POST requests.
	Body *string `json:"Body,omitempty"`
	// The timeout for establishing the TCP connection to the target host. If not set the request timeout value is used.
	DialTimeout *int64 `json:"DialTimeout,omitempty"`
	// Collection of key value pairs to set in the request header.
	Header interface{} `json:"Header,omitempty"`
	// The request is for an internal platform API that requires authentication to be inserted by the platform implementation.
	Internal *bool `json:"Internal,omitempty"`
	// Method specifies the HTTP method (GET, POST, PUT, etc.). For client requests an empty string means GET.
	Method *string `json:"Method,omitempty"`
	// The timeout for the HTTP request to complete, from connection establishment to response body read complete. If not set a default timeout of five minutes is used.
	Timeout              *int64        `json:"Timeout,omitempty"`
	Url                  *ConnectorUrl `json:"Url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorHttpRequest ConnectorHttpRequest

// NewConnectorHttpRequest instantiates a new ConnectorHttpRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorHttpRequest() *ConnectorHttpRequest {
	this := ConnectorHttpRequest{}
	return &this
}

// NewConnectorHttpRequestWithDefaults instantiates a new ConnectorHttpRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorHttpRequestWithDefaults() *ConnectorHttpRequest {
	this := ConnectorHttpRequest{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *ConnectorHttpRequest) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorHttpRequest) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *ConnectorHttpRequest) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *ConnectorHttpRequest) SetBody(v string) {
	o.Body = &v
}

// GetDialTimeout returns the DialTimeout field value if set, zero value otherwise.
func (o *ConnectorHttpRequest) GetDialTimeout() int64 {
	if o == nil || o.DialTimeout == nil {
		var ret int64
		return ret
	}
	return *o.DialTimeout
}

// GetDialTimeoutOk returns a tuple with the DialTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorHttpRequest) GetDialTimeoutOk() (*int64, bool) {
	if o == nil || o.DialTimeout == nil {
		return nil, false
	}
	return o.DialTimeout, true
}

// HasDialTimeout returns a boolean if a field has been set.
func (o *ConnectorHttpRequest) HasDialTimeout() bool {
	if o != nil && o.DialTimeout != nil {
		return true
	}

	return false
}

// SetDialTimeout gets a reference to the given int64 and assigns it to the DialTimeout field.
func (o *ConnectorHttpRequest) SetDialTimeout(v int64) {
	o.DialTimeout = &v
}

// GetHeader returns the Header field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectorHttpRequest) GetHeader() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectorHttpRequest) GetHeaderOk() (*interface{}, bool) {
	if o == nil || o.Header == nil {
		return nil, false
	}
	return &o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *ConnectorHttpRequest) HasHeader() bool {
	if o != nil && o.Header != nil {
		return true
	}

	return false
}

// SetHeader gets a reference to the given interface{} and assigns it to the Header field.
func (o *ConnectorHttpRequest) SetHeader(v interface{}) {
	o.Header = v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *ConnectorHttpRequest) GetInternal() bool {
	if o == nil || o.Internal == nil {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorHttpRequest) GetInternalOk() (*bool, bool) {
	if o == nil || o.Internal == nil {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *ConnectorHttpRequest) HasInternal() bool {
	if o != nil && o.Internal != nil {
		return true
	}

	return false
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *ConnectorHttpRequest) SetInternal(v bool) {
	o.Internal = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *ConnectorHttpRequest) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorHttpRequest) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *ConnectorHttpRequest) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *ConnectorHttpRequest) SetMethod(v string) {
	o.Method = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *ConnectorHttpRequest) GetTimeout() int64 {
	if o == nil || o.Timeout == nil {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorHttpRequest) GetTimeoutOk() (*int64, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *ConnectorHttpRequest) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *ConnectorHttpRequest) SetTimeout(v int64) {
	o.Timeout = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ConnectorHttpRequest) GetUrl() ConnectorUrl {
	if o == nil || o.Url == nil {
		var ret ConnectorUrl
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorHttpRequest) GetUrlOk() (*ConnectorUrl, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ConnectorHttpRequest) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given ConnectorUrl and assigns it to the Url field.
func (o *ConnectorHttpRequest) SetUrl(v ConnectorUrl) {
	o.Url = &v
}

func (o ConnectorHttpRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedConnectorBaseMessage, errConnectorBaseMessage := json.Marshal(o.ConnectorBaseMessage)
	if errConnectorBaseMessage != nil {
		return []byte{}, errConnectorBaseMessage
	}
	errConnectorBaseMessage = json.Unmarshal([]byte(serializedConnectorBaseMessage), &toSerialize)
	if errConnectorBaseMessage != nil {
		return []byte{}, errConnectorBaseMessage
	}
	if o.Body != nil {
		toSerialize["Body"] = o.Body
	}
	if o.DialTimeout != nil {
		toSerialize["DialTimeout"] = o.DialTimeout
	}
	if o.Header != nil {
		toSerialize["Header"] = o.Header
	}
	if o.Internal != nil {
		toSerialize["Internal"] = o.Internal
	}
	if o.Method != nil {
		toSerialize["Method"] = o.Method
	}
	if o.Timeout != nil {
		toSerialize["Timeout"] = o.Timeout
	}
	if o.Url != nil {
		toSerialize["Url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorHttpRequest) UnmarshalJSON(bytes []byte) (err error) {
	type ConnectorHttpRequestWithoutEmbeddedStruct struct {
		// Contents of the request body to send for PUT/PATCH/POST requests.
		Body *string `json:"Body,omitempty"`
		// The timeout for establishing the TCP connection to the target host. If not set the request timeout value is used.
		DialTimeout *int64 `json:"DialTimeout,omitempty"`
		// Collection of key value pairs to set in the request header.
		Header interface{} `json:"Header,omitempty"`
		// The request is for an internal platform API that requires authentication to be inserted by the platform implementation.
		Internal *bool `json:"Internal,omitempty"`
		// Method specifies the HTTP method (GET, POST, PUT, etc.). For client requests an empty string means GET.
		Method *string `json:"Method,omitempty"`
		// The timeout for the HTTP request to complete, from connection establishment to response body read complete. If not set a default timeout of five minutes is used.
		Timeout *int64        `json:"Timeout,omitempty"`
		Url     *ConnectorUrl `json:"Url,omitempty"`
	}

	varConnectorHttpRequestWithoutEmbeddedStruct := ConnectorHttpRequestWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConnectorHttpRequestWithoutEmbeddedStruct)
	if err == nil {
		varConnectorHttpRequest := _ConnectorHttpRequest{}
		varConnectorHttpRequest.Body = varConnectorHttpRequestWithoutEmbeddedStruct.Body
		varConnectorHttpRequest.DialTimeout = varConnectorHttpRequestWithoutEmbeddedStruct.DialTimeout
		varConnectorHttpRequest.Header = varConnectorHttpRequestWithoutEmbeddedStruct.Header
		varConnectorHttpRequest.Internal = varConnectorHttpRequestWithoutEmbeddedStruct.Internal
		varConnectorHttpRequest.Method = varConnectorHttpRequestWithoutEmbeddedStruct.Method
		varConnectorHttpRequest.Timeout = varConnectorHttpRequestWithoutEmbeddedStruct.Timeout
		varConnectorHttpRequest.Url = varConnectorHttpRequestWithoutEmbeddedStruct.Url
		*o = ConnectorHttpRequest(varConnectorHttpRequest)
	} else {
		return err
	}

	varConnectorHttpRequest := _ConnectorHttpRequest{}

	err = json.Unmarshal(bytes, &varConnectorHttpRequest)
	if err == nil {
		o.ConnectorBaseMessage = varConnectorHttpRequest.ConnectorBaseMessage
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Body")
		delete(additionalProperties, "DialTimeout")
		delete(additionalProperties, "Header")
		delete(additionalProperties, "Internal")
		delete(additionalProperties, "Method")
		delete(additionalProperties, "Timeout")
		delete(additionalProperties, "Url")

		// remove fields from embedded structs
		reflectConnectorBaseMessage := reflect.ValueOf(o.ConnectorBaseMessage)
		for i := 0; i < reflectConnectorBaseMessage.Type().NumField(); i++ {
			t := reflectConnectorBaseMessage.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectorHttpRequest struct {
	value *ConnectorHttpRequest
	isSet bool
}

func (v NullableConnectorHttpRequest) Get() *ConnectorHttpRequest {
	return v.value
}

func (v *NullableConnectorHttpRequest) Set(val *ConnectorHttpRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorHttpRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorHttpRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorHttpRequest(val *ConnectorHttpRequest) *NullableConnectorHttpRequest {
	return &NullableConnectorHttpRequest{value: val, isSet: true}
}

func (v NullableConnectorHttpRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorHttpRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
