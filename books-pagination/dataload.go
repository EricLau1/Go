package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
	"books/api/models"
)

func Random(max float64) float32 {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	x := r.Float64() * max
	var y float64 = 100
	return float32(math.Round(x * y) / y)
}

func main() {
	for i := 0; i < 100; i++ {
		models.NewBook(models.Book{Title:"The Wonderful History", Rating: Random(5.0)})
	}
	fmt.Println("dataload success!")
}