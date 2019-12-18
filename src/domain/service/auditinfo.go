package service

import (
	"net/http"

	"github.com/MatchlockIsDAST/MatchlockCore/src/domain/model/auditinfo"
	"github.com/MatchlockIsDAST/MatchlockCore/src/domain/model/auditinfo/value"
)

/*
定義情報をAuditInfomation構造体にして返す
*/

//AuditInfomationConverter 定義情報の整形用
type AuditInfomationConverter interface {
	SetAuth(ty, token string)
	AddPoints(points ...*auditinfo.AuditPoint)
	Fetch() *auditinfo.AuditInfo
}

type auditinfomationconverter struct {
	*auditinfo.AuditInfo
}

//NewAuditInfomationConverter 検査用の情報を整形するサービスを生成します
func NewAuditInfomationConverter(server []string, replacestring string, auditpoints ...*auditinfo.AuditPoint) AuditInfomationConverter {
	return &auditinfomationconverter{auditinfo.NewAuditInfo(server, replacestring, auditpoints...)}
}

func (aic *auditinfomationconverter) Fetch() *auditinfo.AuditInfo { return aic.AuditInfo }

//AuditPointConverter 定義情報へのコンバーターを提供する
type AuditPointConverter interface {
	AddParams(params ...*value.Param)
	AddParam(name, ty string, required bool, pi ...*value.Param)
	Fetch() *auditinfo.AuditPoint
}

type auditpointconverter struct {
	*auditinfo.AuditPoint
}

//NewAuditPointConverter 検査用のエンドポイントを生成します
func NewAuditPointConverter(method, path, contenttype string, header http.Header) AuditPointConverter {
	return &auditpointconverter{auditinfo.NewAuditPoint(method, path, contenttype, header)}
}

func (apc *auditpointconverter) Fetch() *auditinfo.AuditPoint { return apc.AuditPoint }
