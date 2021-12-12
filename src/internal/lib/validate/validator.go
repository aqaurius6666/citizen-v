package validate

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

var (
	myvalidator *validator.Validate
)

func RegisterValidator() {
	myvalidator = validator.New()
	myvalidator.RegisterValidation("myregexp", RegexTag)
	myvalidator.RegisterValidation("vietnamese", VietnameseRegexTag)

}

func Validate(s interface{}) error {
	if err := myvalidator.Struct(s); err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			fmt.Println(e.Field(), e.Error())
		}
		return err
	}
	return nil
}

func RequiredFields(s interface{}, fields ...string) bool {
	typ := reflect.TypeOf(s)
	value := reflect.ValueOf(s)
	if typ.Kind() == reflect.Ptr {
		value = value.Elem()
		typ = typ.Elem()
	}
	if typ.Kind() == reflect.Struct {
		return requiredStruct(typ, value, fields...)
	}
	return true
}
func handleField(t reflect.Type, v reflect.Value) bool {
	switch t.Kind() {
	case reflect.String:
		val, _ := v.Interface().(string)
		if val == "" {
			return false
		}
	case reflect.Int:
		val, _ := v.Interface().(int)
		if val == 0 {
			return false
		}
	case reflect.Int16:
		val, _ := v.Interface().(int16)
		if val == 0 {
			return false
		}
	case reflect.Int64:
		val, _ := v.Interface().(int64)
		if val == 0 {
			return false
		}
	}
	return true
}
func requiredStruct(t reflect.Type, v reflect.Value, fields ...string) bool {
	if len(fields) == 0 {
		numberFields := t.NumField()
		for i := 0; i < numberFields; i++ {
			if !handleField(t.Field(i).Type, v.Field(i)) {
				return false
			}
		}
		return true
	}

	for _, f := range fields {
		_f, ok := t.FieldByName(f)
		if !ok {
			fmt.Printf("missing field %s\n", f)
			return false
		}
		if !handleField(_f.Type, v.FieldByName(f)) {
			return false
		}
	}

	return true
}
