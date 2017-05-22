/**
 *  datetime: 2017-05-21
 *  Go语言-数组
 *  1. 数组是一组相同类型元素组成的序列。
 *  2. 数组一旦声明，数组类型和长度都不允许改变
 *  3. 通过数组下标可以访问数组的元素，数组下标从0开始。
 *  4. 切勿越界赋值或者访问数组元素
 *  5. 使用len获取数组长度
 *  6. 数组赋值操作相当于拷贝，一个数组的改变，不会影响到另一个数组
 *  7. 数组长度使用 ... 特殊符号，go编译器可以自动获取数组长度
 */
package main

import "fmt"

func main() {

	// 初始化数组
	var arr1 [4]int = [4]int{1, 2, 3, 4} // 依据值类型推倒	var arr1 = [4]int{1,2,3,4}, 短声明 arr1 := [4]int{1,2,3,4}

	// 初始化一个字符串数组，数组长度为2
	var arr2 = [2]string{"hello", "world"}

	// 数组赋值操作相当于值拷贝, 一个数组值的改变，不会影响到另一个数组的值
	arr3 := arr2
	//修改第二个元素的值
	arr2[1] = "hhhh"

	fmt.Println(arr2, arr3) // [hello hhhh] [hello world]

	//访问数组第二个元素
	fmt.Println(arr1[1])

	//查看数据元素个数- len, 当然使用cap 也可以获取数组长度，但是不推荐。cap 一般使用在获取slice容量上。
	fmt.Println(len(arr1))

	//遍历数组元素
	fmt.Println("开始遍历数组：arr1")
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}
	fmt.Println("-----")
	for index, value := range arr1 {
		fmt.Printf("index=%d,value=%d\n", index, value)
	}

	fmt.Println("arr1 遍历结束")

	// 数组长度使用 ... 标记， 标识该数组长度由编译器计算获得
	var arr4 = [...]int{3, 3, 3, 3, 3}
	var arr5 = [...]int{4, 4, 4, 4, 4}
	_ = arr5
	fmt.Println(arr4) // [3, 3, 3, 3, 3]
}

//$ go run array_main.go

/*
hello hhhh] [hello world]
2
4
开始遍历数组：arr1
1
2
3
4
-----
index=0,value=1
index=1,value=2
index=2,value=3
index=3,value=4
arr1 遍历结束
[3 3 3 3 3]

*/
