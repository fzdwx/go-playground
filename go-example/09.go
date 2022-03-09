package page

import "fmt"

//
// for condition {
//   // loop body
// }
//

//
// for pos,char := range str {
//		// loop body
// }
//
// for ix,val := range coll{
//  	// loop body
// }
//
// val 始终为集合中对应索引的值拷贝，因此它一般只具有只读性质，对它所做的任何修改都不会影响到集合中原有的值
// 如果 val 为指针，则会产生指针的拷贝，依旧可以修改集合中的原值
func main() {
	s := ""
	for s != "aaaaa" {
		fmt.Println("Value of s:", s)
		s = s + "a"
	}
}

func ex3() {
	for i := 0; i < 5; i++ {
		var v int
		fmt.Println(v)
		v = 5
	}
}

func f2() {
	str := "这是中国话"
	for i, r := range str {
		println(i, r)
		println(i, string(r))
	}
}

func f1() {
	i := 5

	for i > 0 {
		println(i)
		i--
	}
}
