package main

import "fmt"

func getUniqueNumbers(input string) []int {
    // inisialisasi map untuk menyimpan frekuensi setiap angka
    freq := make(map[rune]int)

    // menghitung frekuensi setiap angka dalam input
    for _, char := range input {
        freq[char]++
    }

    // membuat array untuk menyimpan angka-angka yang hanya muncul sekali
    var uniqueNumbers []int

    // mengecek setiap angka dalam input, dan jika frekuensinya sama dengan 1,
    // maka menambahkan angka tersebut ke dalam array uniqueNumbers
    for _, char := range input {
        if freq[char] == 1 {
            uniqueNumbers = append(uniqueNumbers, int(char-'0'))
        }
    }

    return uniqueNumbers
}

func main() {
    input1 := "76523752"
    fmt.Printf("Input: %s\nOutput: %v\n", input1, getUniqueNumbers(input1))

    input2 := "1122"
    fmt.Printf("Input: %s\nOutput: %v\n", input2, getUniqueNumbers(input2))
}

