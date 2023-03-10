package main

import "fmt"

func multipliedSeven(n){
	result := n % 7 === 0
	return result
}

func main(){
	num := 10
	if multipliedSeven(num) fmt.Println("%d is zero multiplied by 7", num)
	else fmt.Println("%d is not zero when multiplied by 7", num)
}
