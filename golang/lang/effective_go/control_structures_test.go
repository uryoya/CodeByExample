/**
 * Effective Go
 * https://go.dev/doc/effective_go#control-structures
 */
package effective_go

import (
	"testing"
	"uryoya/code_by_example/test"
)

func maybeError() error {
	return nil
}

func maybeErrorOrVal() (string, error) {
	return "ok", nil
}

func TestIf(t *testing.T) {
	// 基本的な構文
	var actual string
	x := 3
	if x%2 == 0 {
		actual = "even"
	} else {
		actual = "odd"
	}
	test.Expect(t, actual == "odd")

	// 初期化をif, switch内で宣言することができる
	actual = "ok"
	if err := maybeError(); err != nil {
		actual = "error"
	}
	test.Expect(t, actual == "ok")
	if _, err := maybeErrorOrVal(); err != nil {
		actual = "error"
	}
	test.Expect(t, actual == "ok")
}

func codeUsing(a, b string) string {
	return a + b
}

func TestRedeclarationAndReassignment(t *testing.T) {
	// 再宣言と再代入
	// := を使って変数宣言をする時、新しい変数を伴って宣言する場合はスコープ内で変数名が重複していても再宣言・再代入することができる
	//　長いif-elseチェーンでerr値を簡単に利用したい時に使われる
	actual := "ok"
	s1, err := maybeErrorOrVal()
	if err != nil {
		actual = "error"
	}
	s2, err := maybeErrorOrVal()
	if err != nil {
		actual = "error"
	}
	codeUsing(s1, s2)
	test.Expect(t, actual == "ok")
}

func TestSwitch(t *testing.T) {
	// GoのswitchはCのように定数や整数を与える必要はなく、ケースが一致するまで上から下に評価される。
	actual := ""
	x := 2
	switch {
	case x < 1:
		actual = "less than 1"
	case 1 <= x && x <= 10:
		actual = "1 to 10"
	case 11 <= x && x <= 20:
		actual = "11 to 20"
	}
	test.Expect(t, actual == "1 to 10")

	// カンマ区切りで羅列することも可能
	actual = ""
	c := '?'
	switch c {
	case '!', '?', '&':
		actual = "symbol"
	default:
		actual = "character"
	}
	test.Expect(t, actual == "symbol")
}

func TestTypeSwitch(t *testing.T) {
	// インターフェースの型を動的にswitchすることもできる
	var tt interface{}
	tt = "Hello, World!"
	actual := ""
	switch tt.(type) {
	default:
		actual = "unknown type"
	case bool:
		actual = "bool"
	case string:
		actual = "string"
	case int:
		actual = "int"
	case *bool:
		actual = "pointer to bool"
	case *string:
		actual = "pointer to string"
	case *int:
		actual = "pointer to int"
	}
	test.Expect(t, actual == "string")
}
