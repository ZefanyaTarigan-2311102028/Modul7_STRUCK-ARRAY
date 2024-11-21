package main

import (
	"fmt"
	"math"
)

type Titik struct {
	X int
	Y int
}

type Lingkaran struct {
	Pusat  Titik
	Radius int
}

func jarak(p, q Titik) float64 {
	return math.Sqrt(float64((p.X-q.X)*(p.X-q.X) + (p.Y-q.Y)*(p.Y-q.Y)))
}

func didalam(c Lingkaran, p Titik) bool {
	return jarak(c.Pusat, p) <= float64(c.Radius)
}

func main() {
	var c1, c2 Lingkaran
	var p Titik

	fmt.Println("Masukkan titik pusat dan radius lingkaran 1 (cx cy r):")
	fmt.Scanln(&c1.Pusat.X, &c1.Pusat.Y, &c1.Radius)

	fmt.Println("Masukkan titik pusat dan radius lingkaran 2 (cx cy r):")
	fmt.Scanln(&c2.Pusat.X, &c2.Pusat.Y, &c2.Radius)

	fmt.Println("Masukkan titik sembarang (x y):")
	fmt.Scanln(&p.X, &p.Y)

	if didalam(c1, p) && didalam(c2, p) {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if didalam(c1, p) {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if didalam(c2, p) {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
