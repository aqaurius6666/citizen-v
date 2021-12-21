package validate

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

var (
	myvalidator *validator.Validate
)

func RegisterValidator() error {
	var err error
	myvalidator = validator.New()
	err = myvalidator.RegisterValidation("regexp", RegexTag)
	if err != nil {
		return err
	}
	err = myvalidator.RegisterValidation("vietnamese", VietnameseRegexTag)
	if err != nil {
		return err
	}
	err = myvalidator.RegisterValidation("pid", PidRegexTag)
	if err != nil {
		return err
	}
	err = myvalidator.RegisterValidation("code", CodeRegexTag)
	if err != nil {
		return err
	}
	return nil
}

func Validate(s interface{}) error {
	if myvalidator == nil {
		panic("missing validator")
	}
	if err := myvalidator.Struct(s); err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			fmt.Println(e.Field(), e.Error())
		}
		return err
	}
	return nil
}

func RequiredFields(s interface{}, fields ...string) (string, bool) {
	typ := reflect.TypeOf(s)
	value := reflect.ValueOf(s)
	if typ.Kind() == reflect.Ptr {
		value = value.Elem()
		typ = typ.Elem()
	}
	if typ.Kind() == reflect.Struct {
		return requiredStruct(typ, value, fields...)
	}
	return "", true
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
	case reflect.Int32:
		val, _ := v.Interface().(int32)
		if val == 0 {
			return false
		}
	case reflect.Int64:
		val, _ := v.Interface().(int64)
		if val == 0 {
			return false
		}
	case reflect.Slice:
		size := v.Len()
		if size == 0 {
			return false
		}
		for i := 0; i < size; i++ {
			val := v.Index(i)
			newT := reflect.TypeOf(val.Interface())
			if !handleField(newT, val) {
				return false
			}
		}
	case reflect.Ptr:
		if !handleField(t.Elem(), v.Elem()) {
			return false
		}
	default:
		fmt.Println("unsupported kind %w", t.Kind())
	}
	return true
}
func requiredStruct(t reflect.Type, v reflect.Value, fields ...string) (string, bool) {
	if len(fields) == 0 {
		numberFields := t.NumField()
		for i := 0; i < numberFields; i++ {
			if !handleField(t.Field(i).Type, v.Field(i)) {
				return t.Field(i).Name, false
			}
		}
		return "", true
	}

	for _, f := range fields {
		_f, ok := t.FieldByName(f)
		if !ok {
			return f, false
		}
		if !handleField(_f.Type, v.FieldByName(f)) {
			return f, false
		}
	}

	return "", true
}
