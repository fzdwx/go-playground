package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	js, _ := json.Marshal(vc)

	fmt.Printf("%s", js)

	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
	var f interface{}
	fmt.Println("\n===============")
	_ = json.Unmarshal(b, &f)

	fmt.Println(f)

	hash := md5.New()

	s, _ := io.WriteString(hash, "Hello world")

	fmt.Println(s)

	sum := hash.Sum(nil)

	fmt.Println(sum)

	// 设置路由，如果访问/，则调用index方法
	http.HandleFunc("/", index)

	// 启动web服务，监听9090端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	user := User2{
		name: "fzdwx",
		age:  18,
	}
	b, _ := json.Marshal(user)
	fmt.Fprint(writer, b)
}

// User is custom struct
type User2 struct {
	name string "this is user name"
	age  int    "this is user age"
}

type Point struct {
	X int
	Y int
	Z int
}

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}
