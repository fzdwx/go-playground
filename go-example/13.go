package page

func main() {

	listen(myConsumer)

}

func listen(consumer Consumer2[int]) {
	consumer(123)
}

func myConsumer(val int) {
	println(val)
}

type Consumer2[T any] func(val T)
