package chat

import "reflect"

func PrepareJSON(obj interface{}) map[string]interface{} {
	jsonMap := make(map[string]interface{})
	jsonMap["Type"] = reflect.TypeOf(obj).Name()
	jsonMap["Object"] = obj
	return jsonMap
}
