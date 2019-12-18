package value

//ParamInfomation Paramの肩と名前,利用に関する構造体
type ParamInfomation struct {
	Params map[string]*Param
}

//NewParamInfomation Paramを突っ込んでよしなにする
func NewParamInfomation(params ...*Param) ParamInfomation {
	var paramsMap = map[string]*Param{}
	for i := 0; i < len(params); i++ {
		paramsMap[params[i].Name] = params[i]
	}
	return ParamInfomation{Params: paramsMap}
}

//AddParam NewParamを内部で行う
func (pi *ParamInfomation) AddParam(name, ty string, required bool, ps ...*Param) {
	param := NewParam(name, ty, required, ps...)
	pi.Params[name] = param
}

//AddParams paramを生成して
func (pi *ParamInfomation) AddParams(params ...*Param) {
	for i := 0; i < len(params); i++ {
		pi.Params[params[i].Name] = params[i]
	}
}

//Param パラメーターの情報をまとめる
type Param struct {
	Name     string
	Type     string
	Required bool
	ParamInfomation
}

//NewParam Paramの情報を利用しやすくする
func NewParam(name, ty string, required bool, pi ...*Param) *Param {
	return &Param{
		Name:            name,
		Type:            ty,
		Required:        required,
		ParamInfomation: NewParamInfomation(pi...),
	}
}
