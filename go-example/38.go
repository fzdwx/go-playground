package main

func main() {
	var N = 10

	pending, done := make(chan *task), make(chan *task)
	go sendWork(pending)     // put tasks with work on the channel
	for i := 0; i < N; i++ { // start N goroutines to do work
		go Worker(pending, done)
	}
	consumeWork(done) // continue with the processed tasks
}

func consumeWork(done chan *task) {

}

func sendWork(pending chan *task) {
	pending <- &task{1, "one"}
}

type task struct {
	id   int
	name string
}

func Worker(in, out chan *task) {
	for {
		t := <-in
		process(t)
		out <- t
	}
}

func process(t *task) {

}
