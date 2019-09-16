package main

import (
	"crypto/md5"
	"dbs"
	"entity"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func reqParameter(response http.ResponseWriter, request *http.Request) {

	request.ParseForm()

	fmt.Println(request.Form)
	fmt.Println("URL", request.URL)
	fmt.Println("path", request.URL.Path)
	fmt.Println("schem", request.URL.Scheme)
	fmt.Println("name", request.Form["name"])
	for k, v := range request.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", v)
	}
	fmt.Fprint(response, "hello,world")
}

func testParameter(response http.ResponseWriter, request *http.Request) {

	request.ParseForm()

	fmt.Println(request.Form)
	fmt.Println("URL", request.URL)
	fmt.Println("path", request.URL.Path)
	fmt.Println("schem", request.URL.Scheme)
	fmt.Println("name", request.Form["name"])
	for k, v := range request.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", v)
	}
	fmt.Fprint(response, "hello,test")

}

func toIndex(w http.ResponseWriter, r *http.Request) {

	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))

	t, _ := template.ParseFiles("src/views/login.html")
	t.Execute(w, token)
}

type ErrTipout struct {
	ErrNumber  int    //错误号
	ErrMessage string //错误信息
	ErrDescr   string //更详细的错误描述
}

func login(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	username := r.Form["username"]
	password := r.Form["password"]

	//[]string --> string
	//name := strings.Join(username, "")
	//pwd := strings.Join(password, "")

	if len(r.Form["username"][0]) == 0 {
		//为空的处理
		message := "用户名不能为空"
		t, _ := template.ParseFiles("src/views/error.html")
		t.Execute(w, message)
		return
	}
	if len(r.Form["password"][0]) == 0 {
		//为空的处理
		message := "密码不能为空"
		t, _ := template.ParseFiles("src/views/error.html")
		t.Execute(w, message)
		return
	}
	fmt.Println(username[0])
	fmt.Println(password[0])

	var userMapper entity.UserMapper

	user, e := userMapper.SelectByUserAndPwd(username[0], password[0])

	if e != nil {
		fmt.Println(e)
	}
	fmt.Println("用户:", user)

	if user.Id == 0 {
		t, _ := template.ParseFiles("src/views/error.html")
		t.Execute(w, "用户名或密码不存在")
	} else {
		t, _ := template.ParseFiles("src/views/success.html")
		t.Execute(w, "success")
	}
}

func main() {

	http.HandleFunc("/", reqParameter)
	http.HandleFunc("/index", toIndex)
	http.HandleFunc("/login", login)
	http.HandleFunc("/test", testParameter)
	dbs.InitDB()
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
