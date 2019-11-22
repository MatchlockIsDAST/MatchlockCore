package request

import (
	"net/http"

	"github.com/MatchlockIsDAST/MatchlockCore/src/domain/model/request/entity"
)

//Request Core内部で利用するHTTP Requestの定義です
type Request struct {
	*http.Request
	URL    *entity.URL
	Header *entity.Header
	Body   *entity.Body
}

//New *model.Requestを返します
func New(req *http.Request) (request *Request) {
	request.Request = req
	request.URL = entity.CreateURL(req.URL)
	request.Header = entity.CreateHeader(req.Header)
	e := req.ParseForm()
	if e != nil {
		panic(e)
	}
	request.Body = entity.CreateBody(req.Header.Get("Content-Type"), req.Body, req.Form)
	return request
}
