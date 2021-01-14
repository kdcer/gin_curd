package file

import (
	"os"
	"strings"
)

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
	if build.Class != nil {
		_, err = f.WriteString(build.Class.String())
		if err != nil {
			panic(err)
		}
	}

	if build.Func != nil {
		_, err = f.WriteString(build.Func2String())
		if err != nil {
			panic(err)
		}
	}

	f.Close()
}

func createFile(name string) (*os.File, error) {
	fileName := name + ".go"
	return os.Create(fileName)
}

func parseTagSetting(tag string) map[string]string {
	setting := map[string]string{}
	tags := strings.Split(tag, ";")
	for _, value := range tags {
		v := strings.Split(value, ":")
		k := strings.TrimSpace(strings.ToLower(v[0]))
		if len(v) >= 2 {
			setting[k] = strings.Join(v[1:], ":")
		} else {
			setting[k] = k
		}
	}
	return setting
}

func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	lower := strings.ToLower(string(data[:]))
	return strings.Replace(lower,"i_d","id",-1)

}
