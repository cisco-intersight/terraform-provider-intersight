package intersight

import (
	"encoding/json"
	"log"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// SuppressDiffAdditionProps Suppress Difference functions for additional properties
func SuppressDiffAdditionProps(k, old, new string, d *schema.ResourceData) bool {
	if old == "null" && new == "" {
		return true
	}
	var oldJson = make(map[string]interface{})
	var newJson = make(map[string]interface{})
	err := json.Unmarshal([]byte(old), &oldJson)
	err1 := json.Unmarshal([]byte(new), &newJson)
	if err != nil || err1 != nil {
		log.Printf("Error while json parsing ERR: %+v ERR1: %+v", err, err1)
		return false
	}
	different := true
	for k, v := range newJson {
		different = different && recursiveValueCheck(oldJson, k, v)
	}
	return different
}

func recursiveValueCheck(oldM map[string]interface{}, k string, v interface{}) bool {
	x := reflect.TypeOf(v).String()
	b := true
	simpleDatatypes := "string int int32 int64 float bool float64"
	complexDatatypes := "map[string]interface {}"
	if strings.Contains(simpleDatatypes, x) {
		if oldM[k] != v {
			b = false
		}
	} else if strings.Contains(complexDatatypes, x) {
		if oldM[k] == nil || v == nil {
			b = false
		} else {
			oldMap := oldM[k].(map[string]interface{})
			newMap := v.(map[string]interface{})
			for k1, v1 := range newMap {
				b = b && recursiveValueCheck(oldMap, k1, v1)
			}
		}
	} else {
		log.Printf("Type did not match: %s", x)
	}
	return b
}
