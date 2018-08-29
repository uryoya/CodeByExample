// goto文のテスト
package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("== main関数冒頭 ==")

	foo()

	fmt.Println("main()からreturnError()を呼び出し")
	if err := returnError(); err != nil {
		goto CATCH
	}

	fmt.Println("ここは実行されない@main")

CATCH:
	fmt.Println("main()でgoto CATCH:")

	fmt.Println("== main関数末尾 ==")
}

func returnError() error {
	return errors.New("Something exception!")
}

func foo() {
	fmt.Println("  == foo関数冒頭 ==")
	{
		fmt.Println("  foo()からreturnError()を呼び出し")
		if err := returnError(); err != nil {
			goto CATCH
		}
	}

	fmt.Println("  ここは実行されない@main")

CATCH:
	fmt.Println("  foo()でgoto CATCH:")

	fmt.Println("  == foo関数末尾 ==")
}
