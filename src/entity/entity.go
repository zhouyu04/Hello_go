package entity

import (
	"dbs"
	"fmt"
	"github.com/zhuxiujia/GoMybatis"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Terminals struct {
	Id               int64  `json:"id" gm:"id"`
	Dbid             int64  `json:"dbid"`
	Pos              int64  `json:"pos"`
	Name             string `json:"name"`
	Shop_name        string `json:"shop_name"`
	Ip_addr          string `json:"ip_addr"`
	Run_env          string `json:"run_env"`
	Core_ver         string `json:"core_ver"`
	Db_ver           string `json:"db_ver"`
	Vertion          string `json:"vertion"`
	Release_id       int64  `json:"release_id"`
	Await_release_id int64  `json:"await_release_id"`
	Status           int    `json:"status"`
}

type TerminalsMapper struct {
	SelectAll  func() ([]Terminals, error)
	SelectById func(id string) (Terminals, error) `mapperParams:"id"`
	//SelectTemplete      func(Id string) ([]Terminals, error) `mapperParams:"Id"`
	//InsertTemplete      func(arg Terminals) (int64, error)
	//InsertTempleteBatch func(args []Terminals) (int64, error) `mapperParams:"args"`
	//UpdateTemplete      func(Id Terminals) (int64, error)     `mapperParams:"Id"`
	//DeleteTemplete      func(Id string) (int64, error)        `mapperParams:"Id"`
}

//初始化mapper文件和结构体
func InitMapperByLocalSession() TerminalsMapper {

	fmt.Println("初始化数据库连接..................")
	var engine = GoMybatis.GoMybatisEngine{}.New()
	//mysql链接格式为         用户名:密码@(数据库链接地址:端口)/数据库名称   例如root:123456@(***.mysql.rds.aliyuncs.com:3306)/test
	_, err := engine.Open("mysql", dbs.DB_Driver2) //此处请按格式填写你的mysql链接，这里用*号代替
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("初始化mapper................")
	//读取mapper xml文件
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	fmt.Println("路径:", dir)

	bytes, _ := ioutil.ReadFile(dir + "/src/mapper/TerminalsMapper.xml")
	var terminalsMapper TerminalsMapper
	//设置对应的mapper xml文件
	engine.WriteMapperPtr(&terminalsMapper, bytes)
	fmt.Println("mybatis初始化完成..................")
	return terminalsMapper
}
