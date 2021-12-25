package validate

import (
	"fmt"
	"regexp"
	"strconv"

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
	regexString := `^([0-9a-zA-ZÀÁÂÃÈÉÊÌÍÒÓÔÕÙÚĂĐĨŨƠàáâãèéêếìíòóôõùúăđĩũơƯĂẠẢẤẦẨẪẬẮẰẲẴẶẸẺẼỀỀỂưăạảấầẩẫậắằẳẵặẹẻẽềềểỄỆỈỊỌỎỐỒỔỖỘỚỜỞỠỢỤỦỨỪễệỉịọỏốồổỗộớờởỡợụủứừỬỮỰỲỴÝỶỸửữựỳỵỷỹ\s]+)$`
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

func CodeRegexTag(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	if field == "" {
		return true
	}
	_, err := strconv.Atoi(field)
	if err != nil {
		return false
	}
	if len(field)%2 == 0 || len(field) > 0 {
		return true
	}
	return false
}
