package main

// 匿名函数

var Add = func(a, b int) int {
	return a + b
}

func sum(c int, f func(int, int) int) int {
	d := f(c, 3)
	return d
}

/*
func main() {
	re := sum(99, Add)
	fmt.Println(re)
	var sli []int = []int{3, 1, 2}
	var f = func(a, b int) bool {
		return a > b
	}
	sort.Slice(sli, f)
	fmt.Println(sli)
	sort.Slice(sli, func(i, j int) bool {
		return sli[i] < sli[j]
	})
	fmt.Println(sli)
}
*/
// 多个参数和多个返回值
func Swap(a, b int) (int, int) {
	return b, a
}

// 可变数量的参数
// more 对应 []int 切片类型
func Sum(a int, more ...int) int {
	for _, v := range more {
		a += v
	}
	return a
}

/*
Go 语言中的函数可以有多个参数和多个返回值，参数和返回值都是以传值的方式和被调用者交换数据。
在语法上，函数还支持可变数量的参数，可变数量的参数必须是最后出现的参数，可变数量的参数其实是一个切片类型的参数。
*/
//闭包
// 不仅函数的参数可以有名字，也可以给函数的返回值命名：
func Find(m map[int]int, key int) (value int, ok bool) {
	value, ok = m[key]
	return
}

// 如果返回值命名了，可以通过名字来修改返回值，也可以通过 defer 语句在 return 语句之后修改返回值：
func Inc() (v int) {
	defer func() { v++ }()
	return 42
}

/*
其中 defer 语句延迟执行了一个匿名函数，因为这个匿名函数捕获了外部函数的局部变量 v，这种函数我们一般叫闭包。i

	闭包对捕获的外部变量并不是传值方式访问，而是以引用的方式访问。

闭包的这种引用方式访问外部变量的行为可能会导致一些隐含的问题：
*/
func main0() {
	for i := 0; i < 3; i++ {
		defer func() { println(i) }()
	}
}

/*
因为是闭包，在 for 迭代语句中，每个 defer 语句延迟执行的函数引用的都是同一个 i 迭代变量，在循环结束后这个变量的值为 3，因此最终输出的都是3。

修复的思路是在每轮迭代中为每个 defer 函数生成独有的变量。可以用下面两种方式：
*/

func main2() {
	for i := 0; i < 3; i++ {
		i := i // 定义一个循环体内局部变量 i
		defer func() { println(i) }()
	}
}

func main4() {
	for i := 0; i < 3; i++ {
		// 通过函数传入 i
		// defer 语句会马上对调用参数求值
		defer func(i int) { println(i) }(i)
	}
}

/*
第一种方法是在循环体内部再定义一个局部变量，这样每次迭代 defer 语句的闭包函数捕获的都是不同的变量，这些变量的值对应迭代时的值。
第二种方式是将迭代变量通过闭包函数的参数传入，defer 语句会马上对调用参数求值。两种方式都是可以工作的。
不过一般来说,在 for 循环内部执行 defer 语句并不是一个好的习惯，此处仅为示例，不建议使用。
/*


/*
o 语言中，如果以切片为参数调用函数时，有时候会给人一种参数采用了传引用的方式的假象：因为在被调用函数内部可以修改传入的切片的元素。
其实，任何可以通过函数参数修改调用参数的情形，都是因为函数参数中显式或隐式传入了指针参数。
函数参数传值的规范更准确说是只针对数据结构中固定的部分传值，例如字符串或切片对应结构体中的指针和字符串长度结构体传值，
但是并不包含指针间接指向的内容。将切片类型的参数替换为类似 reflect.SliceHeader 结构体就很好理解切片传值的含义了：
*/

func twice1(x []int) {
	for i := range x {
		x[i] *= 2
	}
}

type IntSliceHeader struct {
	Data []int
	Len  int
	Cap  int
}

func twice2(x IntSliceHeader) {
	for i := 0; i < x.Len; i++ {
		x.Data[i] *= 2
	}
}
