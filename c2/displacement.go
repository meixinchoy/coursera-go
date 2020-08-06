package main

import (
	"fmt"
	"math"
)
	
func GenDisplaceFn(a,v_0,s_0 float64) func(float64)float64{
	return func(t float64) float64{
		return 0.5*a*(math.Pow(t,2)) + v_0*t + s_0
	}
}

func main(){
	var a, v_0, s_0, t float64
	fmt.Print("Enter acceleration: ")
	fmt.Scan(&a)
	fmt.Print("Enter initial velocity: ")
	fmt.Scan(&v_0)
	fmt.Print("Enter initial displacement: ")
	fmt.Scan(&s_0)

	fn := GenDisplaceFn(a, v_0, s_0)

	for {
		fmt.Print("Enter time (999 to stop program): ")
		fmt.Scan(&t)
		if t==999{
			break;
		}
		fmt.Println(fn(t))		
	}
}