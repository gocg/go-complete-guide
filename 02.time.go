package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//var x := 200

func main() {
	now := time.Now()
	fmt.Printf("current  time:%v\n", now)

	year := now.Year() // 年
	month := now.Month() //约
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n",year, month,day,hour,minute , second)
	//tickDemo()
	fmt.Println(arraySum([3]int{1,2,3}))
	//var a = [3]int{3,4,9}
	var numArray = [...]int{1, 2}
	a := [...]int{1: 1, 3: 5}
	fmt.Println("xxx",numArray)
	fmt.Println("xxxx",a)


	var city = [...]string{"北京", "上海", "深圳"}
	// 方法1：for循环遍历
	for i := 0; i < len(city); i++ {
		fmt.Println(city[i])
	}

	// 方法2：for range遍历
	for index, value := range city {
		fmt.Println(index, value)
	}

	mutilCity := [3][2]string{
		{"北京","上海"},
		{"广州","深圳"},
		{"成都","重庆"},
	}

	fmt.Println("str",mutilCity)

	//scoreMap := make(map[string]int, 8)
	//scoreMap["张三"] = 90
	//scoreMap["小明"] = 100
	//fmt.Println(scoreMap)
	//fmt.Println(scoreMap["小明"])
	//fmt.Printf("type of a:%T\n", scoreMap)

	//scoreMap := make(map[string]int,8)
	//scoreMap["starkwang"] = 90
	//scoreMap["shudong"] = 100
	//fmt.Println(scoreMap)
	//fmt.Println(scoreMap["shudong"])
	//fmt.Printf("type of a:%T\n", scoreMap)
	//
	//v,ok := scoreMap["shudong"]
	//
	//if ok {
	//	fmt.Println(v)
	//} else {
	//	fmt.Println("查无此人")
	//}
	//
	//rand.Seed(time.Now().UnixNano())

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprint("stu%02d",i)

		value := rand.Intn(100)
		scoreMap[key] = value
	}


	var keys = make([]string,0,200)
	for key := range scoreMap {
		keys = append(keys,key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}



}



func arraySum(x [3]int) int{
	sum := 0
	for _, v := range x{
		sum = sum + v
	}
	return sum
}

func tickDemo() {
	ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Println(i)//每秒都会执行的任务
	}
}


//输出结果：current  time:2021-04-01 12:29:40.093922 +0800 CST m=+0.000076300
//2021-04-01 12:29:40

