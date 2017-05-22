/**
 *  datetime: 2017-05-21
 *  Go语言-切片
 *  切片是针对底层数组中某个连续片段的描述 ^-^
 *  0. 切片属于引用类型, 数组属于值类型
 *  1. 切片的长度是可变的
 *  2. 只要切片元素类型相同，两个切片的类型就相同.
 *  3. 切片的零值是nil, 且长度(len)和容量(cap) 都是0
 *  4. 切片的容量和长度理解
 *       * 当不变更底层数组的情况下，len和cap值是相当的
 *       * 如果通过append函数，增加元素个数，超过了设定的长度，那么生成新的切片的容量就会增加。
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
	var arr1 = [10]int{}
	slice1 := arr1[:] // [startIndex:endIndex], 如果缺省startIndex默认从第一个开始， 缺省endIndex默认到最后一个
	slice2 := arr1[:]
	_, _, _ = arr1, slice1, slice2

	fmt.Println("arr1 = ", arr1)
	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)

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
	fmt.Println("len(slice6) = ", len(slice5)) //len(slice6) =  26
	fmt.Println("cap(slice6) = ", cap(slice5)) //cap(slice6) =  40
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
	fmt.Println("for  slice6", slice6) //for range slice6 [11 21 31 41 51]

}

func test(t []int) []int {
	if len(t) > 1 {
		t[0] = 100
	}
	return t
}
