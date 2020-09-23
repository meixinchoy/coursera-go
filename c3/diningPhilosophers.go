package main

import (
	"fmt"
	"strconv"
	"sync"
)

//ChopS struct
type ChopS struct{ sync.Mutex }

//Philo struct
type Philo struct {
	num             int
	leftCS, rightCS *ChopS
}

var wg sync.WaitGroup
var on sync.Once
var fullcounter int = 0
var philos []*Philo = make([]*Philo, 5)
var csticks []*ChopS = make([]*ChopS, 5)

func (p *Philo) eat() {
	p.leftCS.Lock()
	p.rightCS.Lock()
	fmt.Println("starting to eat <" + strconv.Itoa(p.num) + ">")
	p.leftCS.Unlock()
	p.rightCS.Unlock()
	fmt.Println("finishing eating <" + strconv.Itoa(p.num) + ">")
	wg.Done()
}

func setup() {
	wg.Add(10)
	for i := 0; i < 5; i++ {
		csticks[i] = new(ChopS)
		wg.Done()
	}
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i + 1, csticks[i], csticks[(i+1)%5]}
		wg.Done()
	}
	wg.Wait()
	return
}

func main() {
	on.Do(setup)
	for i := 0; i < 3; i++ {
		wg.Add(5)
		for i := 0; i < 5; i++ {
			go philos[i].eat() //philosophers eating
		}
		wg.Wait()
	}
}

/*peer review ans by: pujita

package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct {
	sync.Mutex
}
type Philos struct {
	num, count      int
	leftcs, rightcs *ChopS
}

func (p Philos) eat(c chan *Philos, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		c <- &p
		if p.count < 3 {
			p.leftcs.Lock()
			p.rightcs.Lock()

			fmt.Println("starting to eat ", p.num)
			p.count = p.count + 1
			fmt.Println("finishing eating", p.num)
			p.rightcs.Unlock()
			p.leftcs.Unlock()
			wg.Done()
		}

	}
}

func host(c chan *Philos, wg *sync.WaitGroup) {
	for {
		if len(c) == 2 {
			<-c
			<-c
			time.Sleep(20 * time.Millisecond)
		}
	}
}

func main() {
	var i int
	var wg sync.WaitGroup
	c := make(chan *Philos, 2)

	wg.Add(15)

	ChopSticks := make([]*ChopS, 5)
	for i = 0; i < 5; i++ {
		ChopSticks[i] = new(ChopS)
	}

	Philosophers := make([]*Philos, 5)
	for i = 0; i < 5; i++ {
		Philosophers[i] = &Philos{i + 1, 0, ChopSticks[i], ChopSticks[(i+1)%5]}
	}

	go host(c, &wg)
	for i = 0; i < 5; i++ {
		go Philosophers[i].eat(c, &wg)
	}
	wg.Wait()
}

*/
