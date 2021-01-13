package curd

import (
	"github.com/liujunren93/curl_gin/file"
)

func BuildCurl(data []interface{})  {

}

// 创建文件
func BuildFile(data []interface{},dbPkg string)  {

	file.BuildRouter(data)
	file.BuildDao(data,dbPkg )
	file.BuildCtrl(data)
}
