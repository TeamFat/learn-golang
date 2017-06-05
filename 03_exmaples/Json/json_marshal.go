/**
 * datetime: 2017-06-05
 * Go Json操作 - 编码
 * 使用encoding/json 中Marshal进行JSON编码
 * Go语言结构体的成员名字作为JSON对象(通过反射), 只有导出的结构体成员才会被编码(大写字母开头的成员)
 * 一个构体成员Tag是和在编译阶段关联到该成员的元信息字符串
 * 结构体的成员Tag可以是任意的字符串面值，但是通常是一系列用空格分隔的key:"value"键值对序列
 * json开头键名对应的值用于控制encoding/json包的编码和解码的行为，并且encoding/...下面其它的包也遵循这个约定
 * tag成员中可以有omitempty选项，表示当该结构体成员为零值时，JSON对象不包含该成员。
 *
 */
package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

//Response json字段别名
type Response struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
	Age    int      `json:"age,omitempty"`
	name   string
}

func main() {

	//布尔值
	Jbool, _ := json.Marshal(true)
	fmt.Println(string(Jbool))
	//int
	Jint, _ := json.Marshal(1)
	fmt.Println(string(Jint))
	//float
	Jfloat, _ := json.Marshal(3.14)
	fmt.Println(string(Jfloat))
	//string
	Jstring, _ := json.Marshal("hello 中国")
	fmt.Println(string(Jstring))
	//切片
	s1 := []string{"hello", "world"}
	Js1, _ := json.Marshal(s1)
	fmt.Println(string(Js1))
	//hash map
	m1 := map[string]string{"username": "jsser", "from": "china"}
	Jm1, _ := json.Marshal(m1)
	fmt.Println(string(Jm1))
	//struct
	st1 := Student{Name: "jsser", Age: 20}
	Jst1, _ := json.Marshal(st1)
	fmt.Println(string(Jst1))

	res2D := Response{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
		Age:    100,
		name:   "jsser"} //小写字母name成员，不会被编码
	res2B, _ := json.Marshal(res2D)
	fmt.Printf("%s\n", res2B)
	// fmt.Println(string(res2B))

}

//$ go run json_marshal.go

// true
// 1
// 3.14
// "hello 中国"
// ["hello","world"]
// {"from":"china","username":"jsser"}
// {"Name":"jsser","Age":20}
// {"page":1,"fruits":["apple","peach","pear"],"age":100}
