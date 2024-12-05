// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var a float64
// 	var b float64
// 	var s float64
// 	fmt.Println("Реши квадратное уровнение")
// 	fmt.Println("Введи a:")
// 	fmt.Scan(&a)

// 	fmt.Println("Введи b:")
// 	fmt.Scan(&b)

// 	fmt.Println("Введи c:")
// 	fmt.Scan(&c)
// 	D := (b * b) - 4*(a*c)
// 	if D > 0 {
// 		var x1 float64
// 		var x2 float64
// 		x1 = (-b + math.Sqrt(D)) / (2 * a)
// 		x2 = (-b - math.Sqrt(D)) / (2 * a)
// 	}
// }

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var lenght uint8
	charset := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println("Какую длину пароля вы хотите?")
	fmt.Scan(&lenght)

	useNumbers := getUserInput("Добавить ли цифры в пароль?")
	useUppercase := getUserInput("Добавить ли большие буквы?")
	useSpecials := getUserInput("Добавить ли символы?")

	if useNumbers {
		charset = charset + "0123456789"
	}

	if useUppercase {
		charset = charset + "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	if useSpecials {
		charset = charset + "!@#$%^&*()-_=+[]{}|;:,.<>?/"
	}

	rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел
	password := make([]byte, lenght)
	for i := uint8(0); i < lenght; i++ {
		randomIndex := rand.Intn(len(charset))
		password[i] = charset[randomIndex]
	}
	fmt.Println(string(password))
	fmt.Scan(&charset)
}

func getUserInput(text string) bool {
	var input string
	var result bool
	fmt.Println(text)
	for {
		fmt.Scan(&input)
		if input != "Y" && input != "y" && input != "n" && input != "N" {
			fmt.Println("Ошибка, попробуй снова")
			continue
		}
		break
	}
	result = input == "Y" || input == "y"
	return result

}
