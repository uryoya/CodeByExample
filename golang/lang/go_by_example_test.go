/**
 * Go By Example
 * https://gobyexample.com/
 */
package lang

import (
	"fmt"
	"math"
	"reflect"
	"testing"
	"time"
	"uryoya/code_by_example/test"
)

// https://gobyexample.com/values
func TestValues(t *testing.T) {
	test.Expect(t, "go"+"lang" == "golang") // string
	test.Expect(t, 1+1 == 2)                // int
	test.Expect(t, 9.0/3.0 == 3.0)          // float
	test.Expect(t, true && false == false)  // bool
}

// https://gobyexample.com/variables
func TestVariables(t *testing.T) {
	// 初期化された値によって型を推論する
	var a = "initial"
	test.Expect(t, reflect.TypeOf(a).Name() == "string")
	test.Expect(t, a == "initial")

	// 複数の変数を定義できる
	var b, c int = 1, 2
	test.Expect(t, b == 1)
	test.Expect(t, c == 2)

	// := は変数宣言と初期化の省略形
	d := "apple"
	test.Expect(t, reflect.TypeOf(d).Name() == "string")
	test.Expect(t, d == "apple")

	// 初期化なしで宣言された変数は zero-valued で初期化される
	var e int
	test.Expect(t, e == 0)
	var f string
	test.Expect(t, f == "")
	var g bool
	test.Expect(t, g == false)
	var h float32
	test.Expect(t, h == 0.0)
}

// https://gobyexample.com/constants
func TestConst(t *testing.T) {
	const n = 500000000
	test.Expect(t, n == 500000000)

	// 定数式は任意の精度で計算を実行する
	const d = 3e20 / n
	test.Expect(t, d == 6e+11)
	// 数値定数は明示的な変換などで指定されるまで型を持たない
	test.Expect(t, reflect.TypeOf(d).Name() != "int64")
	test.Expect(t, reflect.TypeOf(int64(d)).Name() == "int64")

	// 変数への割り当てや関数呼び出しなど、数値を必要にする文脈によって数値に型が与えられる
	// math.Sin は float64 を使っている
	sin_n := math.Sin(n)
	test.Expect(t, reflect.TypeOf(sin_n).Name() == "float64")
	test.Expect(t, sin_n == -0.28470407323754404)
}

// https://gobyexample.com/for
func TestFor(t *testing.T) {
	// C言語のfor風
	actual := 0
	for i := 0; i < 10; i++ {
		actual += 2
	}
	test.Expect(t, actual == 20)

	// C言語のwhile風
	ii := 0
	actual = 0
	for ii < 10 {
		ii++
		actual += 3
	}
	test.Expect(t, actual == 30)

	// array, slice, string, map, channelの読み込み時は `range` を使う
	lastKey := 0
	lastValue := '0'
	for key, value := range "abcdefg" {
		lastKey = key
		lastValue = value
	}
	test.Expect(t, lastKey == 6)
	test.Expect(t, lastValue == 'g')
}

// https://gobyexample.com/if-else
func TestIfElse(t *testing.T) {
	// 基本
	actual := ""
	if 7%2 == 0 {
		actual = "even"
	} else {
		actual = "odd"
	}
	test.Expect(t, actual == "odd")

	// if文の中で変数宣言をしてスコープ内で使うことができる
	if num := 9; num < 0 {
		actual = "is negative"
	} else if num < 10 {
		actual = "has 1 digit"
	} else {
		actual = "has multiple digits"
	}
	test.Expect(t, actual == "has 1 digit")
}

// https://gobyexample.com/switch
func TestSwitch(t *testing.T) {
	// 基本
	i := 2
	actual := ""
	switch i {
	case 1:
		actual = "1"
	case 2:
		actual = "2"
	case 3:
		actual = "3"
	}
	test.Expect(t, actual == "2")

	// カンマ区切りで複数の条件を指定することができる
	// defaultも使える
	now := time.Monday
	actual = ""
	switch now {
	case time.Saturday, time.Sunday:
		actual = "It's the weekend"
	default:
		actual = "It's a weekday"
	}
	test.Expect(t, actual == "It's a weekday")

	// 式のないswitchでif/elseを表現することもできる
	actual = ""
	switch {
	case 7%2 == 0:
		actual = "even"
	default:
		actual = "odd"
	}
	test.Expect(t, actual == "odd")

	// 任意の値の型によってswitchすることもできる
	whatAmI := func(i interface{}) string {
		switch t := i.(type) {
		case bool:
			return "bool"
		case int:
			return "int"
		default:
			return fmt.Sprintf("unknown(%T)", t)
		}
	}
	test.Expect(t, whatAmI(true) == "bool")
	test.Expect(t, whatAmI(1) == "int")
	test.Expect(t, whatAmI("hey") == "unknown(string)")
}

// https://gobyexample.com/arrays
func TestArray(t *testing.T) {
	// 配列を宣言すると zero-valued で埋めた任意長の配列が作られる
	var a [5]int
	test.Expect(t, fmt.Sprint(a) == "[0 0 0 0 0]")

	a[4] = 100
	test.Expect(t, fmt.Sprint(a) == "[0 0 0 0 100]")

	// 配列長の取得は len(array)
	test.Expect(t, len(a) == 5)

	// 任意の値で初期化
	b := [5]int{1, 2, 3, 4, 5}
	test.Expect(t, fmt.Sprint(b) == "[1 2 3 4 5]")

	// array型は1次元だけど型を組み合わせて多次元配列を作れる
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	test.Expect(t, fmt.Sprint(twoD) == "[[0 1 2] [1 2 3]]")
}

// https://gobyexample.com/slices
// Goの開発チームによるスライスの解説Blog: https://go.dev/blog/slices-intro
func TestSlice(t *testing.T) {
	// 配列とは違い、スライスは（要素の数ではなく）含まれる要素によってのみ型付けされる (?)
	// 長さがゼロではない空のスライスを作るには組み込みのmakeを使う
	s := make([]string, 3)
	test.Expect(t, fmt.Sprint(s) == "[  ]") // zero-valued で初期化される

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	test.Expect(t, fmt.Sprint(s) == "[a b c]")
	test.Expect(t, len(s) == 3)

	// 配列とは違い、新しい要素を追加することもできる
	s = append(s, "d")
	s = append(s, "e", "f")
	test.Expect(t, fmt.Sprint(s) == "[a b c d e f]")
	test.Expect(t, len(s) == 6)

	// 空のスライスを作ってからcopyを使ってスライスをコピーすることもできる
	c := make([]string, len(s))
	copy(c, s)
	test.Expect(t, fmt.Sprint(c) == "[a b c d e f]")

	// 補足：コピー元より短いスライスにコピーすると余剰分は無くなるっぽい
	cc := make([]string, 3)
	copy(cc, s)
	test.Expect(t, fmt.Sprint(cc) == "[a b c]")
	test.Expect(t, len(cc) == 3)

	// スライスをスライスする構文 slice[low:high] も使える
	l := s[2:5] // s[2] ~ s[4]
	test.Expect(t, fmt.Sprint(l) == "[c d e]")
	l = s[:5] // s[0] ~ s[4]
	test.Expect(t, fmt.Sprint(l) == "[a b c d e]")
	l = s[2:] // s[2] ~ s[5]
	test.Expect(t, fmt.Sprint(l) == "[c d e f]")

	// 1行でスライスの宣言と初期化をすることができる
	foo := []string{"g", "h", "i"}
	test.Expect(t, fmt.Sprint(foo) == "[g h i]")

	// 多次元スライスも作ることができる
	// 配列とは違い、要素の配列の長さは固定されない
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	test.Expect(t, fmt.Sprint(twoD) == "[[0] [1 2] [2 3 4]]")
}

// https://gobyexample.com/maps
func TestMap(t *testing.T) {
	// 組み込み関数の make で初期化する
	m := make(map[string]int)

	// name[key] = val 構文を使ってkey/valueをセットする
	m["k1"] = 7
	m["k2"] = 13
	test.Expect(t, fmt.Sprint(m) == "map[k1:7 k2:13]")

	// name[key] で値を取れる
	test.Expect(t, m["k1"] == 7)

	// len(map) で長さも取れる
	test.Expect(t, len(m) == 2)

	// 組み込み関数の delete で key/value ペアをmapから削除できる
	delete(m, "k2")
	test.Expect(t, fmt.Sprint(m) == "map[k1:7]")

	// mapから値を取得するとき2番めの戻り値を利用することもできる
	// 2番めの戻り値は値の存在の有無を表している
	val, prs := m["k2"]
	test.Expect(t, prs == false)
	test.Expect(t, val != 13)

	// 宣言と初期化をまとめて書くこともできる
	n := map[string]int{"foo": 1, "bar": 2}
	test.Expect(t, fmt.Sprint(n) == "map[bar:2 foo:1]")
}

// https://gobyexample.com/range
func TestRange(t *testing.T) {
	// pass
}

// https://gobyexample.com/functions
func TestFunction(t *testing.T) {
	// pass
}

// https://gobyexample.com/multiple-return-values
func TestMultipleReturnValues(t *testing.T) {
	// Goは組み込みで多値返却ができる
	// 主に値とエラー値を戻り値として表現するときに使っている
	vals := func() (int, int) {
		return 3, 7
	}

	a, b := vals()
	test.Expect(t, a == 3)
	test.Expect(t, b == 7)
	_, c := vals() // 省略記法
	test.Expect(t, c == 7)
}

// https://gobyexample.com/variadic-functions
func TestVariadicFunctions(t *testing.T) {
	// 可変長引数関数
	sum := func(nums ...int) int {
		total := 0
		for _, num := range nums {
			total += num
		}
		return total
	}

	actual := sum(1, 2, 3)
	test.Expect(t, actual == 6)

	// 配列を可変長引数に展開することもできる
	nums := []int{1, 2, 3, 4}
	actual = sum(nums...)
	test.Expect(t, actual == 10)
}

// https://gobyexample.com/closures
// クロージャー
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func TestClosures(t *testing.T) {
	// intSeq はクロージャーを返し、内部変数の i はクロージャーの中で独立している
	nextInt := intSeq()
	test.Expect(t, nextInt() == 1)
	test.Expect(t, nextInt() == 2)
	test.Expect(t, nextInt() == 3)

	// i は独立しているので、新しいクロージャーを生成するとリセットされる
	nextInt2 := intSeq()
	test.Expect(t, nextInt2() == 1)
	test.Expect(t, nextInt() == 4)
}

// https://gobyexample.com/recursion
func TestRecursion(t *testing.T) {
	// pass
}

// https://gobyexample.com/pointers
func TestPointers(t *testing.T) {
	zeroval := func(ival int) {
		ival = 0
	}
	zeroptr := func(iptr *int) {
		*iptr = 0
	}

	i := 1

	zeroval(i)
	test.Expect(t, i == 1)

	zeroptr(&i)
	test.Expect(t, i == 0)
}

// https://gobyexample.com/strings-and-runes
func TestStringsAndRunes(t *testing.T) {
	// pass
}

// https://gobyexample.com/structs
func TestStructs(t *testing.T) {
	type person struct {
		name string
		age  int
	}
	// クラスはないので関数でコンストラクタを作る
	newPerson := func(name string) *person {
		p := person{name: name}
		p.age = 42
		// ローカル変数は関数のスコープを超えて存続するので、ローカル変数へのポインターを安全に帰すことができる（原文ママ）
		return &p
	}

	p := person{"Bob", 20} // 新しい構造体の作成
	test.Expect(t, fmt.Sprint(p) == "{Bob 20}")
	p = person{name: "Alice", age: 30} // フィールドを明示することもできる
	test.Expect(t, fmt.Sprint(p) == "{Alice 30}")
	p = person{name: "Fred"} // フィールドを省略すると zero-valued になる
	test.Expect(t, fmt.Sprint(p) == "{Fred 0}")
	pp := &person{name: "Ann", age: 40} // 構造体へのポインタを得る
	test.Expect(t, fmt.Sprint(pp) == "&{Ann 40}")
	pp = newPerson("Jon") // 慣用的なコンストラクタ関数で構造体の初期化をカプセル化する例
	test.Expect(t, fmt.Sprint(pp) == "&{Jon 42}")

	s := person{name: "Sean", age: 50}
	test.Expect(t, s.name == "Sean") // ドットでフィールドへアクセス
	sp := &s
	test.Expect(t, sp.age == 50) // 構造体ポインタでも同じ
	sp.age = 51
	test.Expect(t, sp.age == 51) // 構造体のフィールドもmutable ¯\_(ツ)_/¯

	// ポインタを使わないと構造体も値渡しになるっぽい
	ss := s
	ss.name = "Taro"
	test.Expect(t, fmt.Sprint(s) == "{Sean 51}")
	test.Expect(t, fmt.Sprint(ss) == "{Taro 51}")
}

// https://gobyexample.com/methods
// メソッド
type rect struct {
	width, height int
}

// レシーバーを *rect で定義
func (r *rect) area() int {
	return r.width * r.height
}

// メソッドはポインター型または値レシーバー型のいずれかに対して定義できる
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}
func TestMethods(t *testing.T) {
	r := rect{width: 10, height: 5}

	test.Expect(t, r.area() == 50)
	test.Expect(t, r.perim() == 30)

	// ポインタでもいける
	rp := &r
	test.Expect(t, rp.area() == 50)
	test.Expect(t, rp.perim() == 30)
}

// https://gobyexample.com/interfaces
// インターフェース
type geometry interface {
	area() float64
	perim() float64
}
type rect2 struct {
	width, height float64
}

func (r rect2) area() float64 {
	return r.width * r.height
}
func (r rect2) perim() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func TestInterfaces(t *testing.T) {
	tests := []struct {
		g           geometry
		area, perim float64
	}{
		{g: rect2{width: 3, height: 4}, area: 12, perim: 14},
		{g: circle{radius: 5}, area: 78.53981633974483, perim: 31.41592653589793},
	}
	for _, testcase := range tests {
		test.Expect(t, testcase.g.area() == testcase.area)
		test.Expect(t, testcase.g.perim() == testcase.perim)
	}
}
