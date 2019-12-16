package signature

import (
	"log"

	"github.com/MatchlockIsDAST/MatchlockCore/src/domain/model/judgresult"
	"github.com/ghodss/yaml"
)

//Signature シグネチャの定義
type Signature struct {
	Level      int      `label:"検査レベル" yaml:"level" json:"level" validate:"unrequired"`
	Type       string   `label:"検査タイプ" yaml:"type" json:"type" validate:"required"`
	Signatures []string `label:"シグネチャー" yaml:"signatures" json:"signatures" validate:"required"`
	Time       TimeDiff `label:"許容時間" yaml:"time" json:"time" validate:"unrequired,time"`
	ShouldBe   bool     `label:"推定される結果" yaml:"shouldbe"  json:"shouldbe" validate:"required"`
	Judgment   *judgresult.JudgResult
}

//TimeDiff 時間差を定義
type TimeDiff struct {
	Max int `label:"最大許容時間" yaml:"max" json:"max" validate:"unrequired,second"`
	Min int `label:"最小許容時間" yaml:"min" json:"min" validate:"unrequired,second"`
}

//New シグネチャ間利用のmodelを生成します
func New(signatureByte []byte) (signature []*Signature, err error) {
	signature = []*Signature{}
	err = yaml.Unmarshal(signatureByte, signature)
	if err != nil {
		log.Println("Signetureのunmarchalに失敗しました")
		return nil, err
	}
	return signature, err
}

//SetJudgment 結果を設定するための関数です
func (sig *Signature) SetJudgment(judg *judgresult.JudgResult) {
	sig.Judgment = judg
}
