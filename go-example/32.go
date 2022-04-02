package main

import (
	"fmt"
	"time"
)

type Project struct{}

func (p *Project) deferError() {
	if err := recover(); err != nil {
		fmt.Println("recover: ", err)
	}
}

func (p *Project) exec(msgchan chan interface{}) {
	for msg := range msgchan {
		//op, ok := msg.(int)
		//fmt.Println(op, ok)
		op := msg.(int)
		fmt.Println(len(msgchan))
		fmt.Println("msg: ", op)
	}
}

func (p *Project) run(msgchan chan interface{}) {
	for {
		defer p.deferError()
		go p.exec(msgchan)
		time.Sleep(time.Second * 2)
	}
}

func (p *Project) Main() {
	a := make(chan interface{}, 10)
	go p.run(a)
	go func() {
		for {
			a <- 1
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 10000000)
}

func main() {
	p := new(Project)
	p.Main()
}
