package entity

import (
	"encoding/json"
	"errors"
	"io"
	"net/url"
	"strings"

	"github.com/MatchlockIsDAST/httpmesconv/stringto"

	"github.com/MatchlockIsDAST/httpmesconv/tostring"

	"github.com/MatchlockIsDAST/MatchlockCore/src/domain/model/request/value"
)

/*
対応する Content Type

Formからの送信
- application/x-www-form-urlencoded
- multipart/form-data

JSON形式
- application/json

マルチパートは後日実装
*/

//Body HTTP requestのbodyを提供します
type Body struct {
	ContentType string
	Body        io.ReadCloser
	ParseValue  url.Values
	Param       *value.Param
}

//CreateBody 新規でBodyを提供する
func CreateBody(ctype string, body io.ReadCloser, pvalue url.Values) *Body {
	var (
		params *value.Param
		err    error
	)

	switch ctype {
	case "application/x-www-form-urlencoded":
	case "application/json":
		stringJSON := tostring.Body(body)
		params, err = purseJSON(stringJSON)
		if err != nil {
			panic(err)
		}
	}
	return &Body{
		ContentType: ctype,
		Body:        body,
		ParseValue:  pvalue,
		Param:       params,
	}
}

//Fetch Request内部のデータを*http.Requestに変換して取得します
func (b *Body) Fetch() io.ReadCloser {
	switch b.ContentType {
	case "application/x-www-form-urlencoded":
		return stringto.IoReadCloser(b.ParseValue.Encode())
	case "application/json":
		return stringto.IoReadCloser(b.Param.Fetch())
	}
	return nil
}

//ChangeParam Keyを元にBodyのvalueを変更します
func (b *Body) ChangeParam(key, value string) {
	switch b.ContentType {
	case "application/x-www-form-urlencoded":

	case "application/json":

	}
}

//JSON形式もしくはJSON内のパラメータの文字列をパースします
func purseJSON(jsondata string) (*value.Param, error) {
	var child = map[string]*value.Param{}
	//JSON内部の "name:val"をパースする
	if !isJSON(jsondata) {
		name, val, err := splitJSONColon(jsondata)
		if err != nil {
			return nil, err
		}
		return value.CreateParam(name, val), nil
	}
	//JSONをパースする
	purseParams := strings.Split(jsondata[1:len(jsondata)-1], ",")

	for _, param := range purseParams {
		pursedParam, err := purseJSON(param)
		if err != nil {
			return nil, err
		}
		child[pursedParam.Name] = pursedParam
	}

	params := value.CreateParam("object", "", child)
	return params, nil
}

//JSON内の "name:value" 形式をArray形式に変更します
func splitJSONColon(param string) (name, value string, err error) {
	err = errors.New("JSON形式ではないデータが引数に渡されています")
	purseArray := strings.Split(param, ":")
	if len(purseArray) != 2 {
		return "", "", err
	}
	name, value = purseArray[0], purseArray[1]
	return name, value, nil
}

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}
