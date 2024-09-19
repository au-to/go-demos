package main // 包声明,每个Go应用程序都包含一个名为main的包

import "fmt" // 导入fmt包，fmt包实现了格式化I/O函数

// 程序入口
func main() {
	// 打印，Println会自动换行，Print不会
	fmt.Println("hello world")
	fmt.Print("hello world\n")

	// Println和Print支持使用变量
	var a int = 10
	fmt.Println(a)
	fmt.Print(a)

	// 变量声明
	var str string = "hello world"

	// 变量声明简写，简写只能用于局部变量（函数内部）
	str1 := "hello world"

	// 声明多个变量
	var num1, num2, num3 int = 1, 2, 3

	// 声明常量
	const VALUE int = 10

	// 声明多个常量
	const (
		VALUE1 int = 10
		VALUE2 int = 20
		VALUE3 int = 30
	)

	/* 基本数据类型 */

	// 整数类型 int8 int16 int32 int64
	var intNum int = 10

	// 浮点数类型
	var floatNum float32 = 10.0
	var floatNum2 float64 = 20.0

	// 布尔类型
	val bol bool = true

	// 字符串类型
	var strType string = "hello world"

	// 字符类型
	var charType rune = 'a'

	// 复数类型
	var complexType complex64 = 10 + 10i

	// 指针类型
	var pointerType *int = &intNum

}