package auditinfomation

import (
	"errors"

	"github.com/MatchlockIsDAST/MatchlockCore/src/domain/model/auditinfo"
	"github.com/getkin/kin-openapi/openapi3"
)

//AuditInfomationCreater AuditInfomationの内部で組み立てるためのmethodを提供します
type AuditInfomationCreater interface {
}

type auditinfomationcretater struct {
	*auditinfo.AuditInfo
}

//NewAuditInfomationCreater AuditInfomationを生成するためのインターフェースを生成します
func NewAuditInfomationCreater(servers []string, replacestring string) (auditinfocreater AuditInfomationCreater) {
	return auditinfocreater
}

func (aic *auditinfomationcretater) AddPoints(apc AuditPointsCreater) {

}

func (aic *auditinfomationcretater) Authorization(authtype, token string) {

}

func (aic *auditinfomationcretater) Fetch() (auditinfomation *auditinfo.AuditInfo) {
	return auditinfomation
}

//AuditPointsCreater 検査を行うエンドポイントの情報を定義情報へ変換する
//現状は最も求めている形式に近しいOpenAPI3のPathsを用いているが今後変更をする
type AuditPointsCreater interface {
}

type auditpointscreater struct {
	openapi3.Paths
}

//NewAuditPointCreater 検査を行うエンドポイントを整形するためのインターフェースを生成します
func NewAuditPointCreater() (auditpointcreater AuditPointsCreater) {
	return
}

//AddOpenAPIPaths openapiのpathsを入れることで動作を短縮できるmethod
func (apc *auditpointscreater) AddOpenAPIPaths(paths openapi3.Paths) {

}

func (apc *auditpointscreater) AddAuditPoint(path string, apic AuditPointItemCreater) {

}

//AuditPointItemCreater 内部のパラメーターやBodyを整形する
type AuditPointItemCreater interface {
}
type auditpointitemcreater struct {
	*openapi3.PathItem
}

//NewAuditPointItemCreater Itemを生成するためのインターフェースを生成します
func NewAuditPointItemCreater() (apic AuditPointItemCreater) {
	return
}

type AuditOperationCreater interface {
}

type auditoperationcreater struct {
	*openapi3.Operation
}

func NewAuditOperationCreater() (aoc AuditOperationCreater) {
	return
}

/*
Parameter関連の処理
*/

//AuditParametersCreater パラメーター(QueryやPath,Header)などの処理を定義情報を問してまとめる
type AuditParametersCreater interface {
	AddParameter(param AuditParameterCreater)
}

type auditparameterscreater struct {
	openapi3.Parameters
}

func NewAuditParametersCreater(params ...AuditParameterCreater) AuditOperationCreater {
	paras := &auditparameterscreater{}
	for i := 0; i < len(params); i++ {
		paras.AddParameter(params[i])
	}
	return paras
}
func (apc *auditparameterscreater) AddParameter(param AuditParameterCreater) {
	apc.Parameters = append(apc.Parameters, param.Fetch())
}

type AuditParameterCreater interface {
	Fetch() *openapi3.ParameterRef
}

type auditparametercreater struct {
	*openapi3.Parameter
}

func NewAuditParameterCreater(name, in, style string, required bool, asc AuditSchmaCreater) (aoc AuditOperationCreater, err error) {
	var sch *openapi3.SchemaRef
	if name == "" || in == "" {
		return nil, errors.New("nameもしくはinが空です")
	}
	if asc != nil {
		sch = asc.Fetch()
	}
	aoc = &openapi3.Parameter{
		Name:     name,
		In:       in,
		Style:    style,
		Schema:   sch,
		Required: required,
	}
	return aoc, nil
}
func (apc *auditparametercreater) Fetch() *openapi3.ParameterRef {
	return &openapi3.ParameterRef{Value: apc.Parameter}
}

//AuditSchmaCreater 型と型の形式の定義を行う
// ty  : stringやintegerなどの大まかな方定義
// format : int64などの細かな情報
type AuditSchmaCreater interface {
	Fetch() *openapi3.SchemaRef
}

type auditschmacreater struct {
	*openapi3.Schema
}

//NewAuditSchmaCreater 型の定義情報を取得する

func NewAuditSchmaCreater(ty string, format string, items AuditSchmaCreater) AuditSchmaCreater {
	if ty == "" {
		ty, format = "string", "string"
	}
	as := &auditschmacreater{&openapi3.Schema{Type: ty, Format: format}}
	if items != nil {
		as.Schema.Items = items.Fetch()
		return as
	}
	return as
}

func (asc *auditschmacreater) Fetch() *openapi3.SchemaRef {
	return &openapi3.SchemaRef{Value: asc.Schema}
}
