package function

type Consumer[V1 any] func(v1 V1)

type BiConsumer[V1 any, V2 any] func(v1 V1, v2 V2)
