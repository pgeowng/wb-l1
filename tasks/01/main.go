package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Human struct {
	Name         string
	IQ           int
	Age          int
	Female       bool
	FitnessIndex float64
}

func (h *Human) Confidence() float64 {
	return rand.Float64()
}

type Action struct {
	Human
}

func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

func (a *Action) SuccessRate() float64 {
	age := sigmoid(float64(a.Age - 20))       // 14..20..26..
	iq := sigmoid(float64(a.IQ-100) / 10 * 3) // 100..110.120..

	fi := 0.0
	if a.Female {
		fi = sigmoid(float64(a.FitnessIndex - 56)) // 50..
	} else {
		fi = sigmoid(float64(a.FitnessIndex - 61)) // 55..
	}

	return age * iq * fi
}

func main() {
	rand.Seed(time.Now().UnixNano())

	actions := []Action{
		{
			Human{
				Name:         "Jane",
				IQ:           120,
				Age:          25,
				Female:       true,
				FitnessIndex: 76,
			},
		}, {
			Human: Human{
				Name:         "John",
				IQ:           110,
				Age:          24,
				Female:       false,
				FitnessIndex: 65,
			},
		}, {
			Human{
				Name:         "Jake",
				IQ:           130,
				Age:          19,
				Female:       false,
				FitnessIndex: 60,
			},
		},
	}

	fmt.Printf("Performer | Confidence | Success Rate\n")
	for _, action := range actions {
		fmt.Printf("%9s | %10.2f | %10.2f\n", action.Name, action.Confidence(), action.SuccessRate())
	}

}
