package main

import "fmt"

func ArrayMerge(a, b []string) []string {
	for _, name := range b {
		found := false
		for _, nameExist := range a {
			if name == nameExist {
				found = true
				break
			}	
		}
		if !found {
			a = append(a, name)
		}
	}

	return a

}

func main(){
	fmt.Println("tes") 
}
