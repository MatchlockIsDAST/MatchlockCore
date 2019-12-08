package service

import (
	"encoding/json"
	"errors"
	"strings"
)

func IsJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

//JSON内の "name:value" 形式をArray形式に変更します
func SplitJSONColon(param string) (name, value string, err error) {
	err = errors.New("JSON形式ではないデータが引数に渡されています")
	purseArray := strings.Split(param, ":")
	if len(purseArray) != 2 {
		return "", "", err
	}
	name, value = purseArray[0], purseArray[1]
	return name, value, nil
}
