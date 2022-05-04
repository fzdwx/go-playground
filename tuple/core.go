package tuple

type Tuple interface {
	AsList() []any
}

type Tuple2[V1 any, V2 any] struct {
	V1 V1
	V2 V2
}

type Tuple3[V1 any, V2 any, V3 any] struct {
	V1 V1
	V2 V2
	V3 V3
}

type Tuple4[V1 any, V2 any, V3 any, V4 any] struct {
	V1 V1
	V2 V2
	V3 V3
	V4 V4
}

func Of2[V1 any, V2 any](v1 V1, v2 V2) Tuple2[V1, V2] {
	return Tuple2[V1, V2]{
		v1,
		v2,
	}
}

func Of3[V1 any, V2 any, V3 any](v1 V1, v2 V2, v3 V3) Tuple3[V1, V2, V3] {
	return Tuple3[V1, V2, V3]{
		v1,
		v2,
		v3,
	}
}

func Of4[V1 any, V2 any, V3 any, V4 any](v1 V1, v2 V2, v3 V3, v4 V4) Tuple4[V1, V2, V3, V4] {
	return Tuple4[V1, V2, V3, V4]{
		v1,
		v2,
		v3,
		v4,
	}
}

//func (t Tuple2) AsList() []any {
//	return []any{t.V1, t.V2}
//}

//func (t Tuple3) AsList() []any {
//	return []any{t.V1, t.V2, t.V3}
//}

//func (t Tuple4) AsList() []any {
//	return []any{t.V1, t.V2, t.V3, t.V4}
//}
