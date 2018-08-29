package main

import (
	"fmt"
	"sort"
)

type IntArray []int

func (data IntArray) Len() int           { return len(data) }
func (data IntArray) Less(i, j int) bool { return data[i] < data[j] }
func (data IntArray) Swap(i, j int)      { data[i], data[j] = data[j], data[i] }

func main() {
	// 自分で定義したデータ型
	data := &IntArray{4, 2, 9, 1}
	fmt.Println(data)
	sort.Sort(data)
	fmt.Println(data)

	fmt.Println()

	// sortパッケージで定義されている標準データ型
	data2 := &sort.IntSlice{9, 3, 7, 1, 5}
	fmt.Println(data2)
	sort.Sort(data2)
	fmt.Println(data2)

	fmt.Println()

	// プリミティプ型
	data3 := []int{3, 6, 2, 10}
	fmt.Println(data3)
	sort.Ints(data3)
	fmt.Println(data3)
}
