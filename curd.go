package curd

import (
	"github.com/liujunren93/curl_gin/file"
)


// 创建文件
//data model
//dbPkg db所在的包
func BuildFile(data []interface{},dbPkg string)  {

	file.BuildRouter(data)
	file.BuildDao(data,dbPkg )
	file.BuildCtrl(data)
}
