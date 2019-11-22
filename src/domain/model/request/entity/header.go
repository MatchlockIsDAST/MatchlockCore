package entity

import "net/http"

//Header http.Headerのwrap
type Header struct {
	http.Header
}

//CreateHeader 新規でHeaderを返します
func CreateHeader(header http.Header) *Header {
	return &Header{header}
}
