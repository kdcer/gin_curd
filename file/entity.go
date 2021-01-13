package file

const (
	entityDir = "handler/entity"

)

//func init()  {
//	router = new(File)
//	_ = os.Mkdir(entityDir, 0777)
//}
//
//func InitEntity(data []interface{})  {
//	for _, da := range data {
//		doFile := new(File)
//		doFile.Pkg = "entity"
//		of := reflect.TypeOf(da)
//		name := of.Elem().Name()
//		file, err := createCtrlFile(name)
//		if err != nil {
//			panic(err)
//		}
//
//		Build(file, doFile)
//	}
//}
