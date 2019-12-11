package judgresult

import (
	"fmt"

	"github.com/google/uuid"
)

//JudgResult service.Judgmentの結果
type JudgResult struct {
	ID           string   //検査結果のID
	Payloads     []string //検査に利用されたPayload
	IsDifference bool     //検査の結果 差があったか
	ShouldBe     bool     //本来のbool値
	IsNormal     bool     //IsNormal 正常な結果であるか？
}

//New Judgの結果を生成して返す
func New(payloads []string, isDifference, shouldbe bool) (judgresult *JudgResult) {
	judgresult = &JudgResult{
		ID:           newUUID(),
		Payloads:     payloads,
		IsDifference: isDifference,
		ShouldBe:     shouldbe,
		IsNormal:     !(isDifference == shouldbe),
	}
	return judgresult
}

func newUUID() (uu string) {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return
	}
	uu = u.String()
	return uu
}
