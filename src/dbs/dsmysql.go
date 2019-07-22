package dbs

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *DB

func db_connection() int64 {
	db, err := sql.Open("mysql", "root:zhouyu920414@/jdy_publish?charset=utf8")
	checkErr(err)
	db.SetMaxOpenConns(5)

	//插入数据
	stmt, err := db.Prepare("INSERT INTO ls_p_user SET fid=?,fusername=?,fpwd=?,fnickname=?")
	checkErr(err)

	res, err := stmt.Exec("1", "kingdee", "kingdee", "测试")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	return id
}

func main() {
	db_connection()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
