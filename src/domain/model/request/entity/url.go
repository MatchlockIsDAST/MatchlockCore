package entity

import (
	"net/url"

	"github.com/MatchlockIsDAST/MatchlockCore/src/domain/model/request/value"
)

//URL Core内部で利用するURLを定義している
type URL struct {
	*url.URL
	Path     string
	Query    map[string]value.Param
	Fragment string
}

//CreateURL URLをwrapします
func CreateURL(u *url.URL) *URL {
	return &URL{}
}
