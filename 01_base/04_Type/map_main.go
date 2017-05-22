/**
 *  datetime: 2017-05-22
 *  Go语言-map字典
 *  Go中字典类型是一个关联数组结构。关联数组是用来保存键值对的无序集合。
 *  字典类型属于引用类型， 其长度不是固定的
 *  字典类型的零值是nil, 零值的长度是0
 *  可以使用len获取字典长度
 *  字典类型是非线程安全，在多个线程或者协程存储时，要加锁
 */
package main

import (
	"fmt"
)

func main() {

	//初始化字典类型 map[keyType]ValueType{}
	var m1 = map[string]bool{"isGood": true, "isActive": true, "less10": false}
	fmt.Println(m1)
	//获取字典长度
	fmt.Println("len(m1) = ", len(m1)) //len(m1) =  3
	//取值
	fmt.Println("m1[\"isGood\"] = ", m1["isGood"]) //m1["isGood"] =  true
	//更改值
	m1["isGood"] = false
	fmt.Println("m1[\"isGood\"] = ", m1["isGood"]) //m1["isGood"] =  false
	//增加值
	m1["isStart"] = true
	fmt.Println("m1=", m1) //map[isGood:false isActive:true less10:false isStart:true]
	//删除
	delete(m1, "isStart")
	fmt.Println("m1=", m1) //m1= map[isGood:false isActive:true less10:false]

	//判断key是否存在
	isGood, ok := m1["isGood"]
	if ok {
		fmt.Println("isGood exists")
	} else {
		fmt.Println("isGood not exists")
	}
	_ = isGood

	//遍历map

	// return
	// isGood = false
	// isActive = true
	// less10 = false

	for key, value := range m1 {
		fmt.Printf("%v = %v\n", key, value)
	}

	//引用类型  - 一个改变，另一个跟着改变
	m2 := m1
	m2["isGood2"] = true
	fmt.Println("m1=", m1) //m1= map[isGood:false isActive:true less10:false isGood2:true]
	fmt.Println("m2=", m2) //m2= map[isGood:false isActive:true less10:false isGood2:true]

	// 通过make 创建map  - make用于内建类型（map、slice 和channel）的内存分配
	var m3 = make(map[string]string)
	m3["username"] = "jsser"
	m3["domain"] = "jsser.com"
	fmt.Println(m3)

}
