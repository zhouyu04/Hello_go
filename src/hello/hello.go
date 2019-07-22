package main

import (
	"customerMap"
	"fmt"
	"unsafe"
)

var (
	name = "zzyy"
	age  = 18
)

const (
	aaa = "abc"
	bbb = len(aaa)
	ccc = unsafe.Sizeof(aaa)
)

func main() {

	var key = "b"
	var hashcode = 0

	for i := range key {

	}

}

//hello world
//func main()  {
//	fmt.Println("hello world!")
//	fmt.Println("zyy：runoob.com")
//}

//变量声明
//func main()  {
//
//	var a = "Runoob"
//	fmt.Println(a)
//
//	var b, c int = 1, 2
//	fmt.Println(b, c)
//}

//变量声明三种方式
//func main()  {
//
//	var a = "Runoob"
//	fmt.Println(a)
//
//	var b string
//	fmt.Println(b)
//
//	c := "zzyy"
//	fmt.Println(c)
//}

//多变量申明
//func main(){
//
//	var a,b,c int
//	fmt.Println(a,b,c)
//
//	a,b,c = 1,2,3
//	fmt.Println(a,b,c)
//
//	d,e,f := 4,5,6
//	fmt.Println(d,e,f)
//
//	//打印的是最上面申明的全局变量
//	fmt.Println(name,age)
//}

//常量声明
//func main() {
//
//	const LENGTH = 3
//	const WIDTH = 5
//	var area int
//	const a, b, c = 1, false, "hehe"
//	const name = "one"
//	area = LENGTH * WIDTH
//
//	fmt.Printf("面积为：%d", area)
//	fmt.Println()
//	fmt.Println(name)
//	fmt.Println(a, b, c)
//
//	fmt.Println(aaa,bbb,ccc)
//}

//其他运算符
//func main(){
//
//	var a int = 4
//	var b int32
//	var c float32
//	var ptr *int
//
//	/* 运算符实例 */
//	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a )
//	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b )
//	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c )
//
//	/*  & 和 * 运算符实例 */
//	ptr = &a     /* 'ptr' 包含了 'a' 变量的地址 */
//	fmt.Printf("a 的值为  %d\n", a)
//	fmt.Printf("ptr 为 %d\n", ptr)
//	fmt.Printf("*ptr 为 %d\n", *ptr)
//
//}

//func main() {
//	var C, c int //声明变量
//	C = 1        /*这里不写入FOR循环是因为For语句执行之初会将C的值变为1，当我们goto A时for语句会重新执行（不是重新一轮循环）*/
////A:
//	for C < 100 {
//		C++ //C=1不能写入for这里就不能写入
//		for c = 2; c < C; c++ {
//			if C%c == 0 {
//				//goto A //若发现因子则不是素数
//				continue
//			}
//		}
//		fmt.Println(C, "是素数")
//	}
//}

//func main() {
//	for i := 1; i <= 9; i++ { // i 控制行，以及计算的最大值
//		for j := 1; j <= i; j++ { // j 控制每行的计算个数
//			fmt.Printf("%d*%d=%d ", j, i, j*i)
//		}
//		fmt.Println("")
//	}
//}

//func main()  {
//
//	number := maxNumber(3, 2)
//
//	fmt.Println("函数返回值：",number)
//
//	s1, s2 := changeStr("zzyy", "hello")
//	fmt.Println("交换字符结果：",s1,s2)
//}

//交换字符
func changeStr(str1, str2 string) (string, string) {

	return str2, str1
}

//返回两个函数的最大值
func maxNumber(num1, num2 int) int {

	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

//func main(){
//	var a int
//	fmt.Println("for start")
//	for a=0; a < 10; a++ {
//		fmt.Println(a)
//	}
//	fmt.Println("for end")
//
//	fmt.Println(a)
//}

//func main(){
//
//	var _a *int
//	a := 3
//
//	_a = &a
//
//	fmt.Println("a内存地址：",&a)
//
//	fmt.Println("_a变量存储指针地址:",_a)
//
//	//使用*xx用来访问xx指针变量存储的值！！！
//	fmt.Println("_a变量的值：",*_a)
//
//	fmt.Println("指针变量内存地址：",&_a)
//
//	fmt.Println("指针变量内存地址指针：",*&_a)
//
//	var ptr *int
//
//	fmt.Println("未赋值指针:",ptr)
//
//}

type Books struct {
	title   string
	author  string
	subject string
	id      int
}

type Rect struct {
	//定义矩形类
	x, y          float64 //类型只包含属性，并没有方法
	width, height float64
}

func (r *Rect) Area() float64 { //为Rect类型绑定Area的方法，*Rect为指针引用可以修改传入参数的值
	return r.width * r.height //方法归属于类型，不归属于具体的对象，声明该类型的对象即可调用该类型的方法
}

//func main(){
//
//	books := Books{title: "go语言从入门到放弃", author: "zzyy", subject: "go", id: age}
//
//	fmt.Println("书结构体：",books)
//	bookAuthor := books.author
//	fmt.Println(bookAuthor)
//
//	fmt.Println("------------------------------------------------------")
//
//	rect := Rect{x: 1, y: 2, width: 3, height: 4}
//
//	area := rect.Area()
//
//	fmt.Println("面积：",area)
//}

//func main(){
//
//	ints := make([]int, 5, 10)
//
//	i := len(ints)
//
//	caps := cap(ints)
//
//	fmt.Println(ints,"-",i,"-",caps)
//
//}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

//func main()  {
//
//	var numbers []int
//	printSlice(numbers)
//
//	/* 允许追加空切片 */
//	numbers = append(numbers, 0)
//	printSlice(numbers)
//
//	/* 向切片添加一个元素 */
//	numbers = append(numbers, 1)
//	printSlice(numbers)
//
//	/* 同时添加多个元素 */
//	numbers = append(numbers, 2,3,4)
//	printSlice(numbers)
//
//	/* 创建切片 numbers1 是之前切片的两倍容量*/
//	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
//	printSlice(numbers1)
//
//	/* 拷贝 numbers 的内容到 numbers1 */
//	copy(numbers1,numbers)
//	printSlice(numbers1)
//
//
//}

func sliceTest() {
	arr := []int{1, 2, 3, 4, 5}
	s := arr[2:4]
	printSlice(arr)
	printSlice(s)
	for e := range s {
		fmt.Println(s[e])
	}
	s1 := make([]int, 3)
	for e := range s1 {
		fmt.Println(s1[e])
	}
}

//func main(){
//	sliceTest()
//}

//func main() {
//
//	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
//	nums := []int{2, 3, 4,5,6,7,8,9}
//	sum := 0
//	for _, num := range nums {
//		sum += num
//	}
//	fmt.Println("sum:", sum)
//	//在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
//	for i, num := range nums {
//		if num == 3 {
//			fmt.Println("index:", i)
//		}
//	}
//	//range也可以用在map的键值对上。
//	kvs := map[string]string{"a": "apple", "b": "banana"}
//	for k, v := range kvs {
//		fmt.Printf("%s -> %s\n", k, v)
//	}
//	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
//	for i, c := range "go" {
//		fmt.Println(i, c)
//	}
//
//}

//func main() {
//	fmt.Println(len(os.Args))
//	for _, arg := range os.Args {
//		fmt.Println(arg)
//	}
//}

//func main(){
//	var countryCapitalMap map[string]string /*创建集合 */
//	countryCapitalMap = make(map[string]string)
//
//	countryCapitalMap ["China"] = "中国"
//	fmt.Println(countryCapitalMap)
//
//	countryCapitalMap ["China"] = "中华人民共和国"
//	fmt.Println(countryCapitalMap)
//
//	value,ok := countryCapitalMap [ "China" ]
//
//	fmt.Println("ok:",ok)
//	fmt.Println("value:",value)
//}
