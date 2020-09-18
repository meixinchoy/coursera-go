package main

import (
	"fmt"
	"time"
)

func first() {
	for i := 1; i < 9; i++ {
		fmt.Println("first", i)
		time.Sleep(time.Millisecond)
	}
}

func second() {
	for i := 1; i < 9; i++ {
		fmt.Println("second", i)
		time.Sleep(time.Millisecond) //program sleeps for 80miliseconds
	}
}

func main() {
	go first()
	go second()
	time.Sleep(time.Millisecond * 50) //program sleeps for 100miliseconds
	fmt.Println("program ended")
}

/*Race condition is a behavior which occurs where the output depends on non-deterministic ordering. In other words, it is dependent on the sequence of the events. Race condition occur due to the change of sequence when communicating between tasks*/

/*
 code from peer evaluation (by Amsavarthan Lv):

package main

import (
  "fmt"
)

func add_one(pt *int) {
  (*pt)++
  fmt.Println(*pt)
}

func sub_one(pt *int) {
  (*pt)--
  fmt.Println(*pt)
}
func main(){
  i := 0
  go add_one(&i)
  go sub_one(&i)

  i++
  fmt.Println()

}






Thierry Chantier's code

package main

import (
	"fmt"
	"time"
)

var counter int

func raceThing(who string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s %d\n", who, counter)
		counter++
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go raceThing("One")
	go raceThing("Two")
	// add some sleep on the main thread to allow goroutines to be executed
	time.Sleep(time.Second * 10)
}
*/
