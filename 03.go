package main

import "fmt"

//## 概述
//结构体是将零个或多个任意类型的变量，组合在一起的聚合数据类型，也可以看做是数据的集合。

//type Person struct {
//	Name string
//	Age int
//}
//func main() {
//	var p1 Person
//	p1.Name = "Tom"
//	p1.Age = 30
//	fmt.Println("p1 = ",p1)
//
//	var p2 = Person{Name:"starkwang",Age: 27}
//	fmt.Println("p2=",p2)
//
//	p3 := Person{Name:"starkwang",Age:27}
//	fmt.Println("p2 =",p3)
//
//	p4 := struct {
//		Name string
//		Age int
//	}{ Name: "匿名", Age:27}
//	fmt.Println("p4 = ", p4)
//}

type Result struct {
	Code int `json :"code"`
	Message string `json:"msg"`
}

//func main() {
//	var res Result
//	res.Code = 200
//	res.Message = "success"
//
//	jsons,errs := json.Marshal(res)
//
//	if errs != nil {
//		fmt.Printf("json marshal error:",errs)
//
//	}
//
//	fmt.Println("json data:", string(jsons))
//
//	//反序列化
//
//	var res2 Result
//	errs = json.Unmarshal(jsons, &res2)
//
//	if errs != nil {
//		fmt.Println("json unmarshall error :", errs)
//	}
//
//	fmt.Println("res2 :", res2)
//}

//func main() {
//	var res Result
//	res.Code    = 200
//	res.Message = "success"
//	toJson(&res)
//
//	setData(&res)
//	toJson(&res)
//}
//
//func setData (res *Result) {
//	res.Code    = 500
//	res.Message = "fail"
//}
//
//func toJson (res *Result) {
//	jsons, errs := json.Marshal(res)
//	if errs != nil {
//		fmt.Println("json marshal error:", errs)
//	}
//	fmt.Println("json data :", string(jsons))
//}
type Person struct {
	FirstName, LastName string
	//leaves     map[string]int
}

func main() {
	p := Person{
		FirstName: "Aaron",
		LastName:  "Chen",
	}

	a := Person{
		FirstName: "Aaron",
		LastName:  "Chen",
	}

	fmt.Println(p == a)  // true
}