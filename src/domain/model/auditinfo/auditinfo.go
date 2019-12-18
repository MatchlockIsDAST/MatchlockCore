package auditinfo

import (
	"net/http"

	"github.com/MatchlockIsDAST/MatchlockCore/src/domain/model/auditinfo/value"
)

//AuditInfo 検査を行うための情報をまとめた構造体
type AuditInfo struct {
	Server        []string
	ReplaceString string
	AuthInfo      *value.AuthInfomation
	Points        []*AuditPoint
}

//NewAuditInfo PointとサーバーをまとめたAuditInfoを返す
func NewAuditInfo(server []string, replace string, points ...*AuditPoint) (info *AuditInfo) {
	info.Server = server
	info.ReplaceString = replace
	if len(points) != 0 {
		info.Points = points
	}
	return info
}

//SetAuth 認証情報を設定します
func (audit *AuditInfo) SetAuth(ty, token string) {
	audit.AuthInfo = value.NewAuthInfomation(ty, token)
}

//AddPoints EndPointを追加することができます
func (audit *AuditInfo) AddPoints(points ...*AuditPoint) {
	audit.Points = append(audit.Points, points...)
}

//AuditPoint Path(endpoint)を定義した構造体
type AuditPoint struct {
	Method,
	Path,
	ContentType string
	Header http.Header
	value.ParamInfomation
}

//NewAuditPoint 検査を行うエンドポイントを入力します
func NewAuditPoint(method, path, contenttype string, header http.Header) *AuditPoint {
	return &AuditPoint{
		Method:          method,
		Path:            path,
		ContentType:     contenttype,
		Header:          header,
		ParamInfomation: value.NewParamInfomation(),
	}
}
