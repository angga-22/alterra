package main

import (
    "fmt"
    "sort"
)

type Student struct {
    Name  string
    Score []float64
}

func (s *Student) AddScore(score float64) {
    s.Score = append(s.Score, score)
}

func (s Student) AverageScore() float64 {
    var total float64
    for _, score := range s.Score {
        total += score
    }
    return total / float64(len(s.Score))
}

func (s Student) MinScore() float64 {
    minScore := s.Score[0]
    for _, score := range s.Score {
        if score < minScore {
            minScore = score
        }
    }
    return minScore
}

func (s Student) MaxScore() float64 {
    maxScore := s.Score[0]
    for _, score := range s.Score {
        if score > maxScore {
            maxScore = score
        }
    }
    return maxScore
}

func main() {
    var students []Student
    for i := 0; i < 5; i++ {
        var name string
        var scores []float64
        fmt.Printf("Masukkan nama siswa ke-%v: ", i+1)
        fmt.Scan(&name)
        for j := 0; j < 3; j++ {
            var score float64
            fmt.Printf("Masukkan skor ke-%v: ", j+1)
            fmt.Scan(&score)
            scores = append(scores, score)
        }
        student := Student{Name: name, Score: scores}
        students = append(students, student)
    }

    var totalScore float64
    var allScores []float64
    for _, student := range students {
        totalScore += student.AverageScore()
        allScores = append(allScores, student.Score...)
    }
    averageScore := totalScore / float64(len(students))
    sort.Float64s(allScores)

    fmt.Println("Rata-rata skor: ", averageScore)
    fmt.Println("Siswa dengan skor terendah:")
    for _, student := range students {
        if student.MinScore() == allScores[0] {
            fmt.Printf("- %s dengan skor %v\n", student.Name, student.MinScore())
        }
    }
    fmt.Println("Siswa dengan skor tertinggi:")
    for _, student := range students {
        if student.MaxScore() == allScores[len(allScores)-1] {
            fmt.Printf("- %s dengan skor %v\n", student.Name, student.MaxScore())
        }
    }
}

