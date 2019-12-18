package value

import "fmt"

//AuthInfomation 認証情報を格納する
type AuthInfomation struct {
	Type,
	Token string
}

//NewAuthInfomation 認証情報を保存する
func NewAuthInfomation(ty, token string) *AuthInfomation {
	return &AuthInfomation{ty, token}
}

//AuthorizationHeader 認証済みのヘッダーを生成
func (ai *AuthInfomation) AuthorizationHeader() (name, token string) {
	return "Authorization", fmt.Sprintf("%v %v", ai.Type, ai.Token)
}
