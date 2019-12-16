package repository

import "github.com/MatchlockIsDAST/MatchlockCore/src/domain/model/signature"

//Signature 検査用のシグネチャーを格納します
type Signature interface {
	Save(sig *signature.Signature) (savesuccess bool, err error)
	List() (sigs []*signature.Signature)
}
