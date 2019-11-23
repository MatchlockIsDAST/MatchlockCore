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

//Fetch Request内部のデータを*http.Requestに変換して取得します
func (r *Request) Fetch() *http.Request {
	r.Request.URL = r.URL.Fetch()
	r.Request.Header = r.Header.Fetch()
	r.Request.Body = r.Body.Fetch()
	return r.Request
}
