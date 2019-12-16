package judgresult

import (
	"github.com/MatchlockIsDAST/MatchlockCore/src/domain/model/request"
	"github.com/MatchlockIsDAST/MatchlockCore/src/domain/model/response"
	"github.com/MatchlockIsDAST/MatchlockCore/src/domain/service"
)

//JudgResult service.Judgmentの結果
type JudgResult struct {
	ID           string //検査結果のID
	SignatureID  string
	Payloads     []string //検査に利用されたPayload
	IsDifference bool     //検査の結果 差があったか
	ShouldBe     bool     //本来のbool値
	IsNormal     bool     //IsNormal 正常な結果であるか？
	Requests     map[string]*request.Request
	Response     map[string]*response.Response
}

//New Judgの結果を生成して返す
func New(signatureID string, payloads []string, isDifference, shouldbe bool) (judgresult *JudgResult) {
	judgresult = &JudgResult{
		ID:           service.UUID(),
		Payloads:     payloads,
		IsDifference: isDifference,
		ShouldBe:     shouldbe,
		IsNormal:     !(isDifference == shouldbe),
		Requests:     map[string]*request.Request{},
		Response:     map[string]*response.Response{},
	}
	return judgresult
}

//AddRequests 検査のRequestsを追加する
func (judg *JudgResult) AddRequests(requests ...*request.Request) {
	for i := 0; i < len(requests); i++ {
		judg.Requests[requests[i].ID] = requests[i]
	}
}

//AddResponse 検査の結果を格納する
func (judg *JudgResult) AddResponse(responses ...*response.Response) {
	for i := 0; i < len(responses); i++ {
		judg.Response[responses[i].ID] = responses[i]
	}
}
