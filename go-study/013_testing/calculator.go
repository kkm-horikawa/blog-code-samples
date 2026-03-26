package main

import (
	"errors"
	"fmt"
	"math"
)

// Add は2つの整数を足す
func Add(a, b int) int {
	return a + b
}

// Subtract は2つの整数を引く
func Subtract(a, b int) int {
	return a - b
}

// Divide は割り算を行う。0除算はエラーを返す
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("0で割ることはできません")
	}
	return a / b, nil
}

// IsPrime は素数判定を行う
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Add(2, 3) =", Add(2, 3))
	fmt.Println("Subtract(10, 4) =", Subtract(10, 4))

	result, err := Divide(10, 3)
	if err != nil {
		fmt.Println("エラー:", err)
	} else {
		fmt.Printf("Divide(10, 3) = %.2f\n", result)
	}

	fmt.Println("IsPrime(7) =", IsPrime(7))
	fmt.Println("IsPrime(10) =", IsPrime(10))
}
