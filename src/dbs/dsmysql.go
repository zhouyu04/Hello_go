package dbs

import (
	"database/sql"
	"entity"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zhuxiujia/GoMybatis"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

const (
	DB_Driver2 = "root:zhouyu920414@(127.0.0.1:3306)/app_store?charset=utf8&parseTime=True&loc=Local"
)

type UserInfo struct {
	username string
	pwd      string
	nickname string
}

var Dbconn *sql.DB

var Engine = GoMybatis.GoMybatisEngine{}.New()

var regStruct map[string]interface{}

func initModel() {
	regStruct = make(map[string]interface{})
	regStruct["TerminalsMapper"] = new(entity.TerminalsMapper)
	regStruct["userMapper"] = new(entity.UserMapper)
}

func InitDB() {

	fmt.Println("初始化模板........")
	initModel()

	fmt.Println("初始化数据库连接..................")
	var engine = GoMybatis.GoMybatisEngine{}.New()
	//mysql链接格式为         用户名:密码@(数据库链接地址:端口)/数据库名称   例如root:123456@(***.mysql.rds.aliyuncs.com:3306)/test
	_, err := engine.Open("mysql", DB_Driver2)
	if err != nil {
		panic(err.Error())
	}

	//读取mapper xml文件
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println("项目路径:", dir)
	mapperPath := dir + "/src/mapper"
	fmt.Println("mapper文件夹路径:", mapperPath)
	//获取文件夹下所有mapper文件
	infos, _ := ioutil.ReadDir(mapperPath)
	for _, file := range infos {
		name := file.Name()
		fmt.Println("文件名:", name)
		if strings.Contains(name, "Mapper.xml") {
			bytes, _ := ioutil.ReadFile(mapperPath + "/" + name)
			//var terminalsMapper TerminalsMapper
			mapper := strings.Replace(name, ".xml", "", -1)
			//t := reflect.ValueOf(regStruct[mapper]).Type()
			t := regStruct[mapper]
			of := reflect.TypeOf(t)
			//设置对应的mapper xml文件
			fmt.Println("of:", of)
			engine.WriteMapperPtr(t, bytes)
		}

	}
}
