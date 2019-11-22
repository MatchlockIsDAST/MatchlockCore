package entity

import (
	"io"
	"net/url"
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
}

//CreateBody 新規でBodyを提供する
func CreateBody(ctype string, body io.ReadCloser, pvalue url.Values) *Body {
	return &Body{
		ContentType: ctype,
		Body:        body,
		ParseValue:  pvalue,
	}
}
