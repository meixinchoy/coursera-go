package main

import (
	"fmt"
	"sync"
)

func sort(arr []int, c1 chan int, wg *sync.WaitGroup) {
	//newSli := arr[0:len(arr)]
	for {
		isDirty := false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				temp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
				isDirty = true
			}
		}
		if !isDirty {
			wg.Done()
			break
		}
	}
}

func main() {
	const arrlength = 12 //change the length of array here (must be multiples of four)
	var wg sync.WaitGroup
	wg.Add(1)
	ori := [arrlength]int{}
	sorted := [arrlength]int{}
	for i := 0; i < arrlength; i++ {
		fmt.Print("Enter integer:")
		fmt.Scanln(&ori[i])
		sorted[i] = ori[i]
		if i == arrlength-1 {
			wg.Done()
		}
	}
	wg.Wait()
	wg.Add(4)
	c1 := make(chan int, 2)
	go sort(sorted[0:arrlength/4], c1, &wg)
	go sort(sorted[arrlength/4:arrlength/2], c1, &wg)
	go sort(sorted[arrlength/2:arrlength*3/4], c1, &wg)
	go sort(sorted[arrlength*3/4:arrlength], c1, &wg)

	wg.Wait()
	for j := 0; j < 4; j++ {
		fmt.Print("\nbefore: ")
		for i := 0; i < arrlength/4; i++ {
			fmt.Print(ori[i+j*arrlength/4])
			fmt.Print(" ")
		}
		fmt.Print("\nafter: ")
		for i := 0; i < arrlength/4; i++ {
			fmt.Print(sorted[i+j*arrlength/4])
			fmt.Print(" ")
		}
	}
	fmt.Print("\n\nbefore: ")
	for i := 0; i < arrlength; i++ {
		fmt.Print(sorted[i])
		fmt.Print(" ")
	}
	wg.Add(1)
	go sort(sorted[0:arrlength], c1, &wg)
	wg.Wait()
	fmt.Print("\nafter: ")
	for i := 0; i < arrlength; i++ {
		fmt.Print(sorted[i])
		fmt.Print(" ")
	}

}
