package main

import (
	"errors"
	"fmt"
	"time"
)

var load = 100000000
var chBuf = 10000

// CSP Communicating Sequential Processes
func consumer(data chan int, done chan bool) {
	for i := range data {
		//fmt.Println("Consuming:", i)

		if i < 0 {
		}
	}

	done <- true
}

func producer(data chan int) {
	for i := 0; i < load; i++ {
		//fmt.Println("Producing:", i)
		data <- i
	}
	close(data)
}

func demoChannelRoutine() {
	start := time.Now()
	data := make(chan int, chBuf)
	done := make(chan bool, 1)

	go consumer(data, done)
	go producer(data)

	<-done

	elapsed := time.Since(start).Seconds()
	fmt.Printf("Elapsed: %.3fs\n", elapsed)
	tps := float64(load) / elapsed
	fmt.Printf("Transaction Per Second: %.0f\n", tps)
	// TPS: 37287; Load 1m with print in both consumer and producer
	// TPS: 35362; Load 100k with print in both consumer and producer
	// TPS: 64331; Load 100k with print in consumer only
	// TPS: 4532127; Load 10m without print

}

func task(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Runing %d -> %d\n", id, i)
		time.Sleep(time.Second)
	}
}

func demoGoroutine() {
	go task(1)
	go task(2)
	go task(3)
	go task(4)
}

type user struct {
	name string
	age  int
}

type manager struct {
	user
	dep string
}

// func (u *user) attachMethod() { // *user is also ok, diff?
func (u user) AttachMethod() {
	fmt.Println("Hello attached method", u.age)
}

func demoMethod() {
	var m manager
	m.AttachMethod()
}

// MyUser represents user type who implements its method
type MyUser interface {
	AttachMethod()
}

func demoInterface() {
	var u user
	var myUser MyUser = u
	myUser.AttachMethod()
}

func demoStruct() {
	var m manager
	m.name = "Mgr"
	m.age = 39
	m.dep = "Org"
	fmt.Println(m)
}

func demoMap() {
	m := make(map[string]int)
	m["k1"] = 100
	fmt.Println(m)
	v1, ok := m["k1"]

	if ok {
		fmt.Println(v1, ok)
	}

	v2, ok := m["k2"]
	if !ok {
		fmt.Println("no k2", v2, ok)
	}
	delete(m, "k1")
	fmt.Println(m)
}

func demoSlice() {
	slc := make([]int, 2, 5)
	fmt.Println(slc)
	for i := 1; i < 8; i++ {
		slc = append(slc, i)
	}
	fmt.Println(slc)

}

func demoReturnFunc(x int) func() {
	//closure
	f := func() { fmt.Println(x) }
	return f
}

func demoMultiReturn() (int, error) {
	a, b := 2, 0
	if b == 0 {
		return a, errors.New("divided by 0")
	}

	return a / b, nil
}

func main() {
	defer fmt.Println("Cleaning up..")
	//x := 5
	//fmt.Println(demoMultiReturn())
	//demoSlice()
	//demoReturnFunc(x)()
	//demoMap()
	//demoStruct()
	//demoMethod()
	//demoInterface()

	/*
		demoGoroutine()
		time.Sleep(time.Second * 6)
	*/

	demoChannelRoutine()
}
