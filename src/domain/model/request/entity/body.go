package entity

import (
	"io"
	"net/url"

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
	Param       map[string]value.Param
}

//CreateBody 新規でBodyを提供する
func CreateBody(ctype string, body io.ReadCloser, pvalue url.Values) *Body {
	switch ctype {
	case "application/x-www-form-urlencoded":
	case "multipart/form-data":
	case "application/json":
	}
	return &Body{
		ContentType: ctype,
		Body:        body,
		ParseValue:  pvalue,
	}
}

//Fetch Request内部のデータを*http.Requestに変換して取得します
func (b *Body) Fetch() io.ReadCloser {
	return nil
}
