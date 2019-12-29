package func5_test

import (
	"testing"
	"fmt"
)

//声明一个fire函数
func fire()  {
	fmt.Print("--->this is fire func")
}

func study(str string)  {
	fmt.Printf("--->I'm studying %s", str)
}

// 测试，函数是否赋值给变量
func TestFuncVariable(t *testing.T)  {
	//声明变量f，类型是func类型，函数类型
	//将变量 f 声明为 func() 类型，
	// 此时 f 就被俗称为“回调函数”，此时 f 的值为 nil
	var f func()

	f = fire

	f()
}

// 测试，函数是否赋值给变量
func TestStudy(t *testing.T)  {

	st := study

	st("English")
}

// 测试，函数类型是否可以作为参数进行传递
func learning(f func())  {
	fmt.Printf("--->this is running func")
	f()
}

func TestFuncParams(t *testing.T)  {
	f := fire
	// 将函数作为参数进行传递
	learning(f)
}
