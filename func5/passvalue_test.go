package func5_test

import (
	"testing"
	"fmt"
)

type Data struct {
	complax  []int      // 测试切片在参数传递中的效果
	instance InnerData  // 实例分配的innerData
	// 注意一下，下面的ptr的值，获取的是地址，具体地址里的值是什么，并不关心
	ptr      *InnerData // 将ptr声明为InnerData的指针类型
}

// 代表各种结构体字段
type InnerData struct {
	A int
}

// 声明一个全局变量
var in = Data{
	// 测试切片在参数传递中的效果
	complax: []int{
		2, 3, 4,
	},
	// 实例分配的innerData
	instance: InnerData{
		A: 4,
	},
	// 将ptr声明为InnerData的指针类型
	ptr: &InnerData{
		A: 6,
	},
}

//测试，结构体作为函数参数传入
func TestPassByValue(t *testing.T) {
	fmt.Printf("传入函数前, 打印结构体Data实例的值:\t%+v\n", in)
	fmt.Printf("传入函数前, 打印结构体Data实例中instance的地址:\t%p\n", &in.instance)
	fmt.Printf("传入函数前, 打印结构体Data实例的地址:\t%p\n", &in)
	fmt.Println()
	out := passByValue(in)
	fmt.Println()
	fmt.Printf("经过函数处理后, 打印结构体Data实例的值:\t%+v\n", out)
	fmt.Printf("经过函数处理后, 打印结构体Data实例中instance的地址:\t%p\n", &out.instance)
	fmt.Printf("经过函数处理后, 打印结构体Data实例的地址:\t%p\n", &out)

}

//函数内部的变量，是局部变量
func passByValue(inFunc Data) Data {

	fmt.Printf("在函数内部, 打印结构体Data实例的值:\t%+v\n", inFunc)
	fmt.Printf("在函数内部, 打印结构体Data实例中instance的地址:\t%p\n", &inFunc.instance)
	fmt.Printf("在函数内部, 打印结构体Data实例的地址:\t%p\n", &inFunc)
	return inFunc
}
