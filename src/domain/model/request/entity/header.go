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

//Fetch Request内部のデータを*http.Requestに変換して取得します
func (h *Header) Fetch() http.Header {
	return h.Header
}
