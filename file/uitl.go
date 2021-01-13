package file

import "os"

func Build(f *os.File, build *File) {
	_, err := f.WriteString(build.Pkg2String())
	if err != nil {
		panic(err)
	}
	// import
	if len(build.Import) != 0 {
		_, err = f.WriteString(build.Import2String())
		if err != nil {
			panic(err)
		}
	}
	if  build.Class!=nil{
		_, err = f.WriteString(build.Class.String())
		if err != nil {
			panic(err)
		}
	}

	if build.Func!=nil {
		_, err = f.WriteString(build.Func2String())
		if err != nil {
			panic(err)
		}
	}

	f.Close()
}

func createFile(name string) (*os.File,error) {
		fileName := name + ".go"
		return os.Create(fileName)


}