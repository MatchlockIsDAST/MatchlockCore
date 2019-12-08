package value

import (
	"fmt"
	"log"
	"strings"

	"github.com/MatchlockIsDAST/MatchlockCore/src/domain/model/request/service"
)

//Param 値を管理するためのStruct
type Param struct {
	Name     string
	isChilde bool
	Param    string
	Childe   map[string]*Param
}

//CreateParam Paramを作成します。childeは全てマージしてChildeに格納します
func CreateParam(name, value string, childe ...map[string]*Param) (param *Param) {
	param = &Param{}
	//childeが存在した場合
	if len(childe) > 0 {
		param.isChilde = true
		paramChilde := childe[0]
		//1個より大きい場合
		if len(childe) != 1 {
			for _, c := range childe[1:] {
				paramChilde = merge(paramChilde, c)
			}
		}
		param.Childe = paramChilde
	}
	param.Name = name
	param.Param = value
	return param

}

//ChangeParam Paramの変更をする
//key	index=0から変更するParamの改装までArray stringで指定する
//val	string型で単体
func (p *Param) ChangeParam(key []string, val string) {
	if len(key) > 1 {
		p.Childe[key[0]].ChangeParam(key[1:], val)
		return
	}
	if p.Name == key[0] {
		p.Param = val
		return
	}
	if p.Childe[key[0]] != nil {
		p.Childe[key[0]].Param = val
		return
	}
	log.Println("値の変更に失敗しました")
	return
}

//Fetch Paramに設定されたJSONデータの取得をします
func (p *Param) Fetch() (stringParam string) {
	if p.isChilde {
		var rawParamArray = []string{}
		for _, c := range p.Childe {
			rawParamArray = append(rawParamArray, c.Fetch())
		}
		if p.Name != "" && p.Name != "object" {
			return putColonBetween(addQuarto(p.Name), changeToJSONFormat(strings.Join(rawParamArray, ",")))
		}
		return changeToJSONFormat(strings.Join(rawParamArray, ","))
	}
	return putColonBetween(addQuarto(p.Name), addQuarto(p.Param))
}

//Map Param のマージを行う
func merge(m1, m2 map[string]*Param) map[string]*Param {
	ans := map[string]*Param{}

	for k, v := range m1 {
		ans[k] = v
	}
	for k, v := range m2 {
		ans[k] = v
	}
	return (ans)
}

//コロンを挟むように nameとvalを配置する
func putColonBetween(name, val string) string {
	return fmt.Sprintf("%s:%s", name, val)
}

//JSON形式の "{}"を生成する
func changeToJSONFormat(param string) (enclosedjson string) {
	enclosedjson = fmt.Sprintf("{%s}", param)
	return enclosedjson
}

//ダブルクォートで囲い込む
func addQuarto(str string) string {
	return fmt.Sprintf("\"%s\"", str)
}

//JSON形式もしくはJSON内のパラメータの文字列をパースします
func PurseJSON(jsondata string) (*Param, error) {
	var child = map[string]*Param{}
	//JSON内部の "name:val"をパースする
	if !service.IsJSON(jsondata) {
		name, val, err := service.SplitJSONColon(jsondata)
		if err != nil {
			return nil, err
		}
		return CreateParam(name, val), nil
	}
	//JSONをパースする
	purseParams := strings.Split(jsondata[1:len(jsondata)-1], ",")

	for _, param := range purseParams {
		pursedParam, err := PurseJSON(param)
		if err != nil {
			return nil, err
		}
		child[pursedParam.Name] = pursedParam
	}

	params := CreateParam("object", "", child)
	return params, nil
}
