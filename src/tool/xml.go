package main

import (
	"entity"
	"fmt"
	"github.com/zhuxiujia/GoMybatis"
	"os"
	"path/filepath"
	"reflect"
)

func main() {
	var bean = entity.User{} //此处只是举例，应该替换为你自己的数据库模型

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	mapperPath := dir + "\\src\\mapper\\"
	fmt.Println("mapper文件夹路径：", mapperPath)

	GoMybatis.OutPutXml(mapperPath+reflect.TypeOf(bean).Name()+"Mapper.xml", GoMybatis.CreateXml("biz_"+GoMybatis.StructToSnakeString(bean), bean))
}
