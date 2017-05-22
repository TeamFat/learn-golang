/**
 *  datetime: 2017-05-22
 *  Go语言-切片
 *  切片是针对底层数组中某个连续片段的描述 ^-^
 *  0. 切片属于引用类型, 数组属于值类型
 *  1. 切片的长度是可变的
 *  2. 只要切片元素类型相同，两个切片的类型就相同.
 *  3. 切片的零值是[], 且长度(len)和容量(cap) 都是0
 *  4. 切片的容量和长度理解
 *  5. slice几个有用的内置函数：len, cap, append, cpoy
 */
package main

import (
	"fmt"
)

func main() {

	//声明一个切片类型的变量 - 长度为0, 容量为0
	var s1 = []int{} //可以声明直接赋值 var s1 = []int{1,2,3,4,5}
	//向切片s1追加元素
	s1 = append(s1, 10)
	s1 = append(s1, 20)
	s1 = append(s1, 30)

	_ = s1

	//通过数组来初始化切片 - slice1, slice2

	//slice有一些简便的操作
	//slice的默认开始位置是0，ar[:n]等价于ar[0:n]
	//slice的第二个序列默认是数组的长度，ar[n:]等价于ar[n:len(ar)]
	//如果从一个数组里面直接获取slice，可以这样ar[:]，因为默认第一个序列是0，第二个是数组的长度，即等价于ar[0:len(ar)]

	var arr1 = [10]int{}
	slice1 := arr1[:]
	slice2 := arr1[:]
	_, _, _ = arr1, slice1, slice2

	fmt.Println("arr1 = ", arr1)
	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)

	//关于切片容量和切片长度
	var arr01 = [10]int{0, 1, 2, 3, 4, 5} //因为数组的是类型和长度是不可变的，所以它的容量和长度都是相等的。len = cap
	//通过数组初始化切片
	//https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.2.md
	slice01 := arr01[2:5]                    //从arr01索引第二个元素开始，到第5个索引元素结束(不包含第5个)。
	fmt.Println("slice01 cap", cap(slice01)) //slice01的容量是10 - 2 = 8个 (slice开始位置到数组的最后位置的长度)
	fmt.Println("slice01 len", len(slice01)) //slice01的长度为切片的元素个数

	// 修改数组arr1的第一个元素的值
	arr1[0] = 100

	fmt.Println("arr1 = ", arr1)     //arr1 =  [100 0 0 0 0 0 0 0 0 0]
	fmt.Println("slice1 = ", slice1) //slice1 =  [100 0 0 0 0 0 0 0 0 0]
	fmt.Println("slice2 = ", slice2) //slice2 =  [100 0 0 0 0 0 0 0 0 0]

	//修改slice值
	slice1[1] = 200

	fmt.Println("arr1 = ", arr1)     //arr1 =  [100 200 0 0 0 0 0 0 0 0]
	fmt.Println("slice1 = ", slice1) //slice1 =  [100 200 0 0 0 0 0 0 0 0]
	fmt.Println("slice2 = ", slice2) //slice2 =  [100 200 0 0 0 0 0 0 0 0]

	//声明一个切片
	slice3 := []int{10, 20, 30, 40}
	//赋值
	slice4 := slice3

	//
	//append函数会改变slice所引用的数组的内容，从而影响到引用同一数组的其它slice。
	//但当slice中没有剩余空间（即(cap-len) == 0）时，此时将动态分配新的数组空间。
	//返回的slice数组指针将指向这个空间，而原数组的内容将保持不变；其它引用此数组的slice则不受影响
	//

	//重点：依据slice3的值，生成新的切片值，并把50,60追加到该值的背后,然后赋值操作会把这个新切片值在赋给slice3。所以新的slice3和旧的slice3的指向不同的底层数组。

	slice3 = append(slice3, 50, 60)
	fmt.Println("slice3=", slice3) //sice3= [10 20 30 40 50 60]
	fmt.Println("slice4=", slice4) //sice4= [10 20 30 40]

	/**
	 *  使用make初始化切片
	 *  make([]type,len,cap)   len <= cap
	 */
	var slice5 = make([]int, 20, 20)
	slice5 = append(slice5, slice3...)         //当append的值是切片的时，请带上...符号
	fmt.Println("len(slice5) = ", len(slice5)) //len(slice5) =  26
	fmt.Println("cap(slice5) = ", cap(slice5)) //cap(slice5) =  40 //容量cap不足时将按总需要长度的2倍扩容
	fmt.Println("slice5 ", slice5)             //slice5  [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 10 20 30 40 50 60]

	/**
	 * 遍历切片
	 */

	var slice6 = []int{10, 20, 30, 40, 50}
	for _, item := range slice6 {
		item++
	}
	fmt.Println("for range slice6", slice6) //for range slice6 [10 20 30 40 50]

	for index := 0; index < len(slice6); index++ {
		slice6[index]++
	}
	fmt.Println("for  slice6", slice6) //for slice6 [11 21 31 41 51]

	//slice 支持第三个参数 [startIndex:endIndex:cap] 并且 endIndex > cap
	var arr2 [10]int
	slice7 := arr2[2:5]
	_ = slice7
	fmt.Println(cap(slice7))
	fmt.Println(len(slice7))
	fmt.Println(slice7)

	//copy
	var tSlice = []int{1, 2, 3, 4, 5}
	var tmp = make([]int, len(tSlice))

	copyCount := copy(tmp, tSlice)
	tmp[0] = 10
	fmt.Println("tmp：", tmp)             //tmp： [10 2 3 4 5]
	fmt.Println("tSlice：", tSlice)       //tSlice： [1 2 3 4 5]  //不会随着改变
	fmt.Println("copyCount：", copyCount) //copyCount： 5

	//func
	_ = test(tSlice)
	// 引用传递
	fmt.Println("test tSlice", tSlice) //test tSlice [100 2 3 4 5]
}

func test(t []int) []int {
	if len(t) > 1 {
		t[0] = 100
	}
	return t
}
