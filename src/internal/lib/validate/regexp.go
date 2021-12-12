package validate

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func RegexTag(fl validator.FieldLevel) bool {
	fmt.Println(123)
	field := fl.Field().String()
	regexString := fl.Param()
	regex := regexp.MustCompile(regexString)
	return field == "" || regex.Match([]byte(field))
}
func VietnameseRegexTag(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	if field == "" {
		return true
	}
	regexString := `^([a-zA-ZÀÁÂÃÈÉÊÌÍÒÓÔÕÙÚĂĐĨŨƠàáâãèéêìíòóôõùúăđĩũơƯĂẠẢẤẦẨẪẬẮẰẲẴẶẸẺẼỀỀỂưăạảấầẩẫậắằẳẵặẹẻẽềềểỄỆỈỊỌỎỐỒỔỖỘỚỜỞỠỢỤỦỨỪễệỉịọỏốồổỗộớờởỡợụủứừỬỮỰỲỴÝỶỸửữựỳỵỷỹ\s]+)$`
	regex := regexp.MustCompile(regexString)
	return regex.Match([]byte(field))
}

func PidRegexTag(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	if field == "" {
		return true
	}
	regexString := `^[0-9]{12}$`
	regex := regexp.MustCompile(regexString)
	return regex.Match([]byte(field))
}
