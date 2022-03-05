package main

func main() {

	listen(myConsumer)

}

func listen(consumer Consumer[int]) {
	consumer(123)
}

func myConsumer(val int) {
	println(val)
}

type Consumer[T any] func(val T)
