package main

import "fmt"

func countString(slice []string, target string) int {
    count := 0
    for _, s := range slice {
        if s == target {
            count++
        }
    }
    return count
}

func main() {
    slice := []string{"foo", "bar", "baz", "foo", "qux", "foo", "baz"}
    target := "foo"
    count := countString(slice, target)
    fmt.Printf("The string \"%s\" appears %d times in the slice\n", target, count)
}

