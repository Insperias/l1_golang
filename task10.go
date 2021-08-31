package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{
		-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5,
	}

	groups := make(map[float64][]float64) // Карта с ключами float64 и значениями []float64

	for _, t := range temperatures {
		g := math.Trunc(t/10) * 10 // Округляет температуры вниз -20, -30 и так далее
		groups[g] = append(groups[g], t)
	}

	for g, temperatures := range groups {
		fmt.Printf("%v: %v\n", g, temperatures)
	}
}
