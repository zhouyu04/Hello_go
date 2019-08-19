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
	DB_Driver = "root:zhouyu920414@/jdy_publish?charset=utf8"

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
	regStruct["TerminalsMapper"] = entity.TerminalsMapper{}
}

func InitDB() {

	fmt.Println("初始化模板........")
	initModel()

	fmt.Println("初始化数据库连接..................")
	var engine = GoMybatis.GoMybatisEngine{}.New()
	//mysql链接格式为         用户名:密码@(数据库链接地址:端口)/数据库名称   例如root:123456@(***.mysql.rds.aliyuncs.com:3306)/test
	_, err := engine.Open("mysql", DB_Driver2) //此处请按格式填写你的mysql链接，这里用*号代替
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
			t := reflect.ValueOf(regStruct[mapper]).Type()
			//设置对应的mapper xml文件
			engine.WriteMapperPtr(&t, bytes)
		}

	}
}

func OpenDB() (success bool, db *sql.DB) {
	var isOPen = false
	db, err := sql.Open("mysql", DB_Driver)
	if err != nil {
		isOPen = false
		fmt.Println("faile...")
	} else {
		isOPen = true
		fmt.Println("success...")
	}
	checkErr(err)
	db.SetConnMaxLifetime(10000)

	Dbconn = db
	return isOPen, db
}

func Sele(db *sql.DB, username string, pwd string) (*UserInfo, bool) {

	//查询数据
	rows, err := db.Query("SELECT * FROM ls_p_user where fusername = ? and fpwd = ?", username, pwd)
	checkErr(err)
	//db.Exec(username, pwd)
	var fid int
	var fusername string
	var fpwd string
	var fnickname string
	var fcreatetime string

	var user UserInfo
	var flag bool
	if rows.Next() {
		for rows.Next() {
			err = rows.Scan(&fid, &fusername, &fpwd, &fnickname, &fcreatetime)
			checkErr(err)
			fmt.Println(fid)
			fmt.Println(fusername)
			fmt.Println(fpwd)
		}
		user = UserInfo{username: fusername, pwd: fpwd, nickname: fnickname}
		flag = true
	}
	return &user, flag
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
