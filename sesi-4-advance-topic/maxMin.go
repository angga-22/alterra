package main

import "fmt"

func findMinMax(nums *[6]int) (int, int) {
    min := nums[0]
    max := nums[0]
    for _, num := range nums {
        if num < min {
            min = num
        }
        if num > max {
            max = num
        }
    }
    return min, max
}

func main() {
    var nums [6]int
    for i := 0; i < 6; i++ {
        fmt.Printf("Masukkan angka ke-%v: ", i+1)
        fmt.Scan(&nums[i])
    }

    min, max := findMinMax(&nums)

    fmt.Printf("Nilai terkecil: %v\n", min)
    fmt.Printf("Nilai terbesar: %v\n", max)
}

