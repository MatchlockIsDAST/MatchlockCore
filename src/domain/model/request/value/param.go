package value

//Param 値を管理するためのStruct
type Param struct {
	Name     string
	isChilde bool
	Param    string
	Childe   map[string]Param
}

func CreateParam() {

}
