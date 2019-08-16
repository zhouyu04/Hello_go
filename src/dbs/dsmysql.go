package dbs

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zhuxiujia/GoMybatis"
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

func InitDB() {
	_, err := Engine.Open("mysql", DB_Driver2)
	if err != nil {
		panic(err)
	}
	fmt.Println("初始化DB完成.............")
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
