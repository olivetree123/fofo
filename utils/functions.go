package utils

import (
	"github.com/google/uuid"
	"reflect"
	"strings"
)

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

// NewUUID 生成 UUID
func NewUUID() string {
	uid := uuid.Must(uuid.NewRandom())
	var idBytes [32]byte
	copy(idBytes[:], strings.Replace(uid.String(), "-", "", -1))
	return string(idBytes[:])
}
