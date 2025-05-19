// Упражнение 1.1. Измените программу echo так, чтобы она выводила также os.Args[0], имя выполняемой команды
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Name:", os.Args[0])
	fmt.Println("Args:", strings.Join(os.Args[1:], " "))
}
