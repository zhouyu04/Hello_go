package main

import (
	"entity"
	"github.com/zhuxiujia/GoMybatis"
	"reflect"
)

func main() {
	var bean = entity.Terminals{} //此处只是举例，应该替换为你自己的数据库模型
	GoMybatis.OutPutXml(reflect.TypeOf(bean).Name()+"Mapper.xml", GoMybatis.CreateXml("biz_"+GoMybatis.StructToSnakeString(bean), bean))
}
