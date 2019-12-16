package response

import (
	"net/http"

	"github.com/MatchlockIsDAST/httpmesconv/stringto"

	"github.com/MatchlockIsDAST/httpmesconv/tostring"

	"github.com/MatchlockIsDAST/MatchlockCore/src/domain/model/request"
)

//Response Core内部で利用するHTTP Response の定義です
type Response struct {
	*http.Response
	Header http.Header
	Body   string
	*request.Request
	ID string
}

//New *model.Response を返します
func New(resp *http.Response, id string) (response *Response) {
	response = &Response{
		Response: resp,
		Header:   resp.Header,
		Body:     tostring.Body(resp.Body),
		Request:  request.New(resp.Request, id),
		ID:       id,
	}
	return response
}

//Fetch Response 内部のデータを*http.Requestに変換して取得します
func (r *Response) Fetch() *http.Response {
	response := r.Response
	response.Header = r.Header
	response.Body = stringto.IoReadCloser(r.Body)
	response.Request = r.Request.Fetch()
	return response
}
