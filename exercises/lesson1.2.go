// Упражнение 1.2. Измените программу echo так, чтобы она выводила индекс и
// значение каждого аргумента по одному аргументу в строке.
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[1:] {
		fmt.Println("Index argument", index)
		fmt.Println("Name argument:", arg)
		fmt.Println("")
	}
}
