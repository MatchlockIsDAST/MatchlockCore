package value

import "testing"

func testParam１() (param *Param) {
	mapChild1 := map[string]*Param{
		"mapChilde1": CreateParam("mapChilde1inchild", "fuga2"),
		"mapChilde2": CreateParam("mapC2", "", map[string]*Param{
			"mapC2inC": CreateParam("mapC2inC", "fuga3"),
		}),
	}
	param = CreateParam("object", "", mapChild1)
	return param
}
func testParam2() (param *Param) {
	mapChild1 := map[string]*Param{
		"mapChilde1": CreateParam("mapChilde1inchild", "fuga2"),
		"mapChilde2": CreateParam("mapC2", "", map[string]*Param{
			"mapC2inC": CreateParam("mapC2inC", "fuga3"),
		}),
	}
	mapChild2 := map[string]*Param{
		"mapChilde3": CreateParam("mapChilde1inchild", "fuga3"),
		"mapChilde4": CreateParam("mapC3", "", map[string]*Param{
			"mapC3inC": CreateParam("mapC3inC", "fuga3"),
		}),
	}
	param = CreateParam("object", "", mapChild1, mapChild2)
	return
}
func TestFetch(t *testing.T) {
	param := testParam１()
	if param.Fetch() != "{\"mapChilde1inchild\":\"fuga2\",\"mapC2\":{\"mapC2inC\":\"fuga3\"}}" {
		t.Error("正規のJSONが帰ってきていません")
	}
	param = testParam2()
	if param.Fetch() == "{\"mapC2\":{\"mapC2inC\":\"fuga3\"},\"mapChilde1inchild\":\"fuga3\",\"mapC3\":{\"mapC3inC\":\"fuga3\"},\"mapChilde1inchild\":\"fuga2\"}" {
		t.Error("複数childeのCreateParamに失敗しました")
	}
}
