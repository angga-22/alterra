package main

import "fmt"

func main(){
	var top, bottom, height float64
	// Read input values from the user
    fmt.Println("Enter the top length of the trapezium:")
    fmt.Scanln(&top)

    fmt.Println("Enter the bottom length of the trapezium:")
    fmt.Scanln(&bottom)

    fmt.Println("Enter the height of the trapezium:")
    fmt.Scanln(&height)

    // Calculate the area of the trapezium
    area := 0.5 * (top + bottom) * height

    // Print the result
    fmt.Printf("The area of the trapezium with top length %.2f, bottom length %.2f, and height %.2f is %.2f\n", top, bottom, height, area)
}
