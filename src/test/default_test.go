package test

import (
	"dbs"
	"entity"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //选择需要的数据库驱动导入
	"github.com/zhuxiujia/GoMybatis"
	"github.com/zhuxiujia/GoMybatis/example"
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
	var exampleActivityMapperImpl ExampleActivityMapperImpl

	//加载xml实现逻辑到ExampleActivityMapperImpl
	engine.WriteMapperPtr(&exampleActivityMapperImpl, xmlBytes)

	//使用mapper
	result, _ := exampleActivityMapperImpl.SelectAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

}

func Test_select(t *testing.T) {

	//entity.InitMapperByLocalSession()
	dbs.InitDB()

	terminalsMapper := entity.Engine.GetObj("TerminalsMapper").(*entity.TerminalsMapper)
	//terminals, _ := terminalMapper.SelectAll()

	//var id = 1
	terminals, _ := terminalsMapper.SelectById("1")

	//i := terminals[0]

	fmt.Println(terminals)
	//fmt.Println(i)

	//s := i.Name
	//fmt.Println(s)
}
