package value

import (
	"fmt"
	"strings"
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

func (p *Param) Fetch() (stringParam string) {
	if p.isChilde {
		var rawParamArray = []string{}
		for _, c := range p.Childe {
			rawParamArray = append(rawParamArray, c.Fetch())
		}
		if p.Name != "" && p.Name != "object" {
			return addQuarto(p.Name) + ":{" + strings.Join(rawParamArray, ",") + "}"
		}
		return "{" + strings.Join(rawParamArray, ",") + "}"
	}
	return addQuarto(p.Name) + ":" + addQuarto(p.Param)
}

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

func addQuarto(str string) string {
	return fmt.Sprintf("\"%s\"", str)
}
