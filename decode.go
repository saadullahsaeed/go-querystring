//Package querystring provides encoding & decoding of querystring to & from structs
//
// Example:
//
// 	type someStruct struct {
// 		Field1 string `url:"f1"`
// 		Field2 int `url:"f2"`
// 	}
//
// ts := &someStruct{}
// querystring.Decode(httpReq.URL.Query(), ts)
package querystring

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

const tagName = "url"

var (
	// ErrInvalidDestination ...
	ErrInvalidDestination = errors.New("destination is nil")
)

// Decode will accept url.Values and decode the querystring params to the struct passed in
func Decode(values url.Values, dst interface{}) error {
	val := reflect.ValueOf(dst)
	for val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return ErrInvalidDestination
		}
		val = val.Elem()
	}

	if dst == nil {
		return ErrInvalidDestination
	}

	if val.Kind() != reflect.Struct {
		return fmt.Errorf("query: Values() expects struct input. Got %v", val.Kind())
	}

	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		sf := typ.Field(i)
		if sf.PkgPath != "" && !sf.Anonymous { // unexported
			continue
		}

		sv := val.Field(i)
		tag := sf.Tag.Get(tagName)
		if tag == "-" {
			continue
		}

		if sv.CanSet() {
			tagv := values.Get(tag)
			switch sv.Kind() {
			case reflect.String:
				sv.SetString(tagv)
				break
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				uv, err := strconv.Atoi(tagv)
				if err != nil {
					continue
				}
				sv.SetInt(int64(uv))
				break
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				uv, err := strconv.Atoi(tagv)
				if err != nil {
					continue
				}
				sv.SetUint(uint64(uv))
				break
			case reflect.Float32, reflect.Float64:
				uv, err := strconv.ParseFloat(tagv, 64)
				if err != nil {
					continue
				}
				sv.SetFloat(uv)
				break
			}
		}
	}
	return nil
}
