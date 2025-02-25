package main

import "fmt"

type element interface {
	Density() float64
}

// IsDenser - compare density of two objects
func IsDenser(a, b element) bool {
	return a.Density() > b.Density()
}

// Metal - mass and volume information
type Metal struct {
	mass   float64
	volume float64
}

// Density - return density of metal
func (m *Metal) Density() float64 {
	// density = mass/volume
	return m.mass / m.volume
}

type Gas struct {
	pressure      float64
	temperature   float64
	molecularMass float64
}

func (g *Gas) Density() float64 {
	var density float64
	density = (g.molecularMass * g.pressure) / (0.0821 * (g.temperature + 273))
	return density
}

func main() {
	gold := Metal{478, 24}
	silver := Metal{100, 10}

	result := IsDenser(&gold, &silver)
	if result {
		fmt.Println("gold has higher density than silver")
	} else {
		fmt.Println("silver has higher density than gold")
	}

	oxygen := Gas{
		pressure:      5,
		temperature:   27,
		molecularMass: 32,
	}

	hydrogen := Gas{
		pressure:      1,
		temperature:   0,
		molecularMass: 2,
	}

	result = IsDenser(&oxygen, &hydrogen)
	if result {
		fmt.Println("oxygen has higher density than silver")
	} else {
		fmt.Println("hydrogen has higher density than gold")
	}
}
