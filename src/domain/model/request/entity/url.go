package entity

import (
	"net/url"
)

//URL Core内部で利用するURLを定義している
type URL struct {
	*url.URL
	Path     string
	Query    url.Values
	Fragment string
}

//CreateURL URLをwrapします
func CreateURL(u *url.URL) *URL {
	return &URL{
		URL:      u,
		Path:     u.Path,
		Query:    u.Query(),
		Fragment: u.Fragment,
	}
}

//SetQuery urlのquery stringを追加する
func (u *URL) SetQuery(key, val string) {
	u.Query.Set(key, val)
}

//AddQuery urlのQuery を新規追加します
func (u *URL) AddQuery(key, val string) {
	u.Query.Add(key, val)
}

//Fetch Request内部のデータを*http.Requestに変換して取得します
func (u *URL) Fetch() *url.URL {
	u.URL.Path = u.Path
	u.URL.RawQuery = u.Query.Encode()
	u.URL.Fragment = u.Fragment
	return u.URL
}
