// guessNumber project main.go
package main

import (
	"fmt"
	"sort"
)

func main() {
	//这是一段来自sort.Search方法源码注释的代码，有修改，原来有点bug
	fmt.Println("Pick a number from 0 to 100.")
	fmt.Printf("Your number is %d\n",
		sort.Search(100, func(i int) bool {
			fmt.Printf("Is your number <= %d? ", i)
			var s string
			fmt.Scanf("%s\n", &s)
			return s != "" && s[0] == 'y'
		}))
}
