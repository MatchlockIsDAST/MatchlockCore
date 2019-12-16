package repository

import "github.com/MatchlockIsDAST/MatchlockCore/src/domain/model/judgresult"

//JudgmentResult 結果を取得する際に利用される
type JudgmentResult interface {
	Save(judgmentresult *judgresult.JudgResult) (savesuccess bool, err error)
	Find(id string) *judgresult.JudgResult
}
