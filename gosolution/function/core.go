package function

func Foreach[E any](s []E, consumer BiConsumer[int, E]) {
	for i, v := range s {
		consumer(i, v)
	}
}
