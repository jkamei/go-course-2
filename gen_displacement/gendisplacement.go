package main

import ("fmt"
"math"
)

func main() {
	var acc, initVel, initDist, time float64
	fmt.Printf("accleration: ")
	fmt.Scanln(&acc)

	fmt.Printf("initial velocity: ")
	fmt.Scanln(&initVel)

	fmt.Printf("initial displacement: ")
	fmt.Scanln(&initDist)

	DisplaceFunc := GenDisplaceFn(acc, initVel, initDist)
	for {
		fmt.Scanln(&time)

		fmt.Printf("final position: %f\n\n", DisplaceFunc(time))
	}
}

func GenDisplaceFn(acc, initVel, initDist float64) func (float64) float64 {
	fmt.Println(acc, initVel, initDist)
	fn := func (time float64) float64 {
		return 0.5 * acc * math.Pow(time, 2) + initVel * time + initDist
	}
	return fn
}