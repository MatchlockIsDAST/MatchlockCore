package entity

import (
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
		params, err = value.PurseJSON(stringJSON)
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
//JSON形式の値を変更する場合は key1/key2/key3 のような形でスラッシュでkeyをわたしてください
func (b *Body) ChangeParam(key, value string) {
	switch b.ContentType {
	case "application/x-www-form-urlencoded":
		b.ParseValue.Set(key, value)
	case "application/json":
		b.Param.ChangeParam(strings.Split(key, "/"), value)
	}
}
