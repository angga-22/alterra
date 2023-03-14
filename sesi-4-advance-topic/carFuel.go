package main

import "fmt"

type Car struct {
    Type   string
    FuelIn float64
}

func (car Car) EstimateDistance() float64 {
    // Menghitung perkiraan jarak yang bisa ditempuh berdasarkan bahan bakar yang sedang terisi
    return car.FuelIn * 1.5
}

func main() {
    car := Car{Type: "sedan", FuelIn: 20.0}
    distance := car.EstimateDistance()
    fmt.Printf("Perkiraan jarak yang bisa ditempuh oleh mobil %s dengan bahan bakar %v L: %v km", car.Type, car.FuelIn, distance)
}

