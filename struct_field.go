package json

import (
	"reflect"
	"strings"
)

func getTag(field reflect.StructField) string {
	return field.Tag.Get("json")
}

func isIgnoredStructField(field reflect.StructField) bool {
	if field.PkgPath != "" && !field.Anonymous {
		// private field
		return true
	}
	tag := getTag(field)
	if tag == "-" {
		return true
	}
	return false
}

type structTag struct {
	key         string
	isOmitEmpty bool
	isString    bool
	field       reflect.StructField
}

type structTags []*structTag

func (t structTags) existsKey(key string) bool {
	for _, tt := range t {
		if tt.key == key {
			return true
		}
	}
	return false
}

func structTagFromField(field reflect.StructField) *structTag {
	keyName := field.Name
	tag := getTag(field)
	opts := strings.Split(tag, ",")
	if len(opts) > 0 {
		if opts[0] != "" {
			keyName = opts[0]
		}
	}
	st := &structTag{key: keyName, field: field}
	if len(opts) > 1 {
		st.isOmitEmpty = opts[1] == "omitempty"
		st.isString = opts[1] == "string"
	}
	return st
}
