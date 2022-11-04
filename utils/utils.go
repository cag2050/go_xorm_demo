package utils

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"hash/fnv"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	uuid "github.com/satori/go.uuid"
)

const (
	DefaultDateTimeFormart = "2006-01-02 15:04:05"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func StringReplaceStart(inStr string, startIndex int, oldSubStr, newSubStr string, n int) string {
	outStr := inStr[:startIndex]
	outStr += strings.Replace(inStr[startIndex:], oldSubStr, newSubStr, n)
	return outStr
}

//从字符串集合中排除指定字符串，并返回新的集合对象
func StringArrayExclude(src []string, excludeStr string) []string {
	var retStrs []string
	for _, str := range src {
		if str == excludeStr {
			continue
		}
		retStrs = append(retStrs, str)
	}
	return retStrs
}

func Container2Interfaces(container interface{}) []interface{} {
	var ret []interface{}

	containerValue := reflect.ValueOf(container)
	switch reflect.TypeOf(container).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < containerValue.Len(); i++ {
			ret = append(ret, containerValue.Index(i).Interface())
		}
	case reflect.Map:
		//todo
	}

	return ret
}

func Contains(container interface{}, obj interface{}) bool {
	containerValue := reflect.ValueOf(container)
	switch reflect.TypeOf(container).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < containerValue.Len(); i++ {
			if containerValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if containerValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}
	return false
}

//str = 100mb   ret: 100*1024*1024
func ParseByteSize(str string) (int64, error) {
	if v, err := strconv.ParseInt(str, 10, 64); err == nil {
		return v, nil
	}

	var p1 int
	var p2 string
	str = strings.ToLower(str)
	if _, err := fmt.Sscanf(str, "%d%s", &p1, &p2); err != nil {
		return 0, err
	}
	switch p2 {
	case "k", "kb":
		return int64(p1) * 1024, nil
	case "m", "mb":
		return int64(p1) * 1024 * 1024, nil
	case "g", "gb":
		return int64(p1) * 1024 * 1024 * 1024, nil
	}
	return 0, fmt.Errorf("invalid byte string:%v", str)
}

func GetUuid() string {
	v := uuid.NewV4()
	return v.String()
}

func CreateFileDirIf(filePath string) error {
	dir, _ := filepath.Split(filePath)
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return os.MkdirAll(dir, 0666)
	}
	return err
}

//深拷贝(使用Gob方式)
func GobDeepCopy(src, dst interface{}) error {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(src); err != nil {
		return err
	}

	decoder := gob.NewDecoder(&buf)
	//decoder := gob.NewDecoder(bytes.NewBuffer(buf.Bytes()))
	if err := decoder.Decode(dst); err != nil {
		return err
	}

	return nil
}

func GetStringHash64(val string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(val))
	return h.Sum64()
}
