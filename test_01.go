package main

import (
	"fmt"
	. "github.com/logrusorgru/aurora"
	"math/rand"
	"time"
)

func producer(ch chan <-int,name string){
   for{
   	 sleep()
   	 n:=rand.Intn(100)
   	 fmt.Printf("Channel %s -> %d  ",Yellow(name),Red(n))
   	 ch<-n
   }
}

func sleep() {
   time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
}

func consumer(ch <-chan int){
	for n := range ch{
		fmt.Printf("Channel Value is <- %d\n",Blue(n))
	}
}

func fanIb(chA,chB <- chan int,chC chan <-  int){
	var n int
	for{
		select {
		case n=<-chA:
			chC<-n
		case n=<-chB:
            chC <- n
		}
	}
}
func main() {

	chA:=make(chan int)
	chB:=make(chan int)
	chC:=make(chan int)
	go producer(chA,"A")
	go producer(chB,"B")
	go consumer(chC)

	fanIb(chA,chB,chC)
}
