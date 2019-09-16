package test

import (
	"dbs"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //选择需要的数据库驱动导入
	"github.com/zhuxiujia/GoMybatis"
	"github.com/zhuxiujia/GoMybatis/example"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"reflect"
	"testing"
)

func TestHello(t *testing.T) {
	fmt.Println("hello test")
}

var xmlBytes = []byte(`
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
"https://raw.githubusercontent.com/zhuxiujia/GoMybatis/master/mybatis-3-mapper.dtd">
<mapper>
    <select id="SelectAll">
        select * from biz_activity where delete_flag=1 order by create_time desc
    </select>
</mapper>
`)

type ExampleActivityMapperImpl struct {
	SelectAll func() ([]example.Activity, error)
}

func TestGoMybaatis(t *testing.T) {

	var engine = GoMybatis.GoMybatisEngine{}.New()
	//Mysql链接格式 用户名:密码@(数据库链接地址:端口)/数据库名称,如root:123456@(***.com:3306)/test
	_, err := engine.Open("mysql", "root:zhouyu920414@(127.0.0.1:3306)/app_store?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	//var exampleActivityMapperImpl ExampleActivityMapperImpl
	exampleActivityMapperImpl := new(ExampleActivityMapperImpl)
	//加载xml实现逻辑到ExampleActivityMapperImpl
	engine.WriteMapperPtr(exampleActivityMapperImpl, xmlBytes)

	//使用mapper
	result, _ := exampleActivityMapperImpl.SelectAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

}

func TestCustomerMybatis(t *testing.T) {

	dbs.InitDB()

	userMapper := dbs.Engine.GetObj("userMapper")
	fmt.Println(userMapper)
}

/**
 * 功能描述:打印函数
 *
 * @auther: zhouyu
 * @date: 2019/8/21 16:54
 */
func TestPic(t *testing.T) {
	// 图片大小
	const size = 3000
	// 根据给定大小创建灰度图
	pic := image.NewGray(image.Rect(0, 0, size, size))
	// 遍历每个像素
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			// 填充为白色
			pic.SetGray(x, y, color.Gray{255})
		}
	}

	// 从0到最大像素生成x坐标
	for x := 0; x < size; x++ {
		// 让sin的值的范围在0~2Pi之间
		s := float64(x) * 2 * math.Pi / size
		// sin的幅度为一半的像素。向下偏移一半像素并翻转
		y := size/2 - math.Sin(s)*size/2
		// 用黑色绘制sin轨迹
		pic.SetGray(x, int(y), color.Gray{0})
	}

	// 创建文件
	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err)
	}
	// 使用png格式将数据写入文件
	png.Encode(file, pic) //将image信息写入文件中
	// 关闭文件
	file.Close()
}

// 定义商标结构
type Brand struct {
}

// 为商标结构添加Show()方法
func (t Brand) Show() {
}

// 为Brand定义一个别名FakeBrand
type FakeBrand = Brand

// 定义车辆结构
type Vehicle struct {
	// 嵌入两个结构
	FakeBrand
	Brand
}

func TestAlias(t *testing.T) {

	// 声明变量a为车辆类型
	var a Vehicle

	// 指定调用FakeBrand的Show
	a.FakeBrand.Show()
	// 取a的类型反射对象
	ta := reflect.TypeOf(a)
	// 遍历a的所有成员
	for i := 0; i < ta.NumField(); i++ {
		// a的成员信息
		f := ta.Field(i)
		// 打印成员的字段名和类型
		fmt.Printf("FieldName: %v, FieldType: %v\n", f.Name, f.Type.Name())
	}

}

func TestSlice(t *testing.T) {

	var a = []int{1, 2, 3}

	a = append(a, 4)
	fmt.Println(a)
	a = append([]int{0, -1}, a...) // 在开头添加1个元素
	fmt.Println(a)
	a = append([]int{-3, -2, -1}, a...) // 在开头添加1个切片
	fmt.Println(a)

	//ints := a[:2]
	//fmt.Println(ints)
	//
	//i := a[2:]
	//fmt.Println(i)

	i2 := append([]int{9}, a[2:]...)
	fmt.Println(i2)
}

func TestSliceCopy(t *testing.T) {
	// 设置元素数量为1000
	const elementCount = 1000
	// 预分配足够多的元素切片
	srcData := make([]int, elementCount)
	// 将切片赋值
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	// 引用切片数据
	refData := srcData
	// 预分配足够多的元素切片
	copyData := make([]int, elementCount)
	// 将数据复制到新的切片空间中
	copy(copyData, srcData)
	// 修改原始数据的第一个元素
	srcData[0] = 999
	// 打印引用切片的第一个元素
	fmt.Println(refData[0])
	// 打印复制切片的第一个和最后一个元素
	fmt.Println(copyData[0], copyData[elementCount-1])
	// 复制原始数据从4到6(不包含)
	copy(copyData, srcData[4:6])
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}
}
