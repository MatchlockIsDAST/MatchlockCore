package main

/*
import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func main() {
	testJson := "{\"test1\": \"testvalue1\", \"test2\": {\"test2c\": \"test2cValue\"}, \"test3\": [\"test3cArray1\",\"test3cArray2\",1]}"
	var jsonBody map[string]interface{}
	fmt.Println(testJson)
	json.Unmarshal([]byte(testJson), &jsonBody)
	fmt.Println(jsonBody)
	body := map[string]value{}
	for name, v := range jsonBody {
		body[name] = NewJson(name, v)
	}
	fmt.Println(body)
}

type value struct {
	Name  string
	Value interface{}
	Type  string
}

func NewJson(name string, v interface{}) value {
	var retv value
	retv.Name = name
	if strings.HasPrefix(reflect.TypeOf(v).String(), "map") {
		array := []value{}
		for n, vchild := range v.(map[string]interface{}) {
			array = append(array, NewJson(n, vchild))
		}
		retv.Value = array
		retv.Type = "Object"
	} else if strings.HasPrefix(reflect.TypeOf(v).String(), "[]") {
		array := []value{}
		for _, vchild := range v.([]interface{}) {
			array = append(array, NewJson(name, vchild))
		}
		retv.Value = array
		retv.Type = "Array"
	} else {
		retv.Value = v
		retv.Type = reflect.TypeOf(v).String()
	}
	return retv
}
*/
