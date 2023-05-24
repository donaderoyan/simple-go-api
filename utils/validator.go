package util

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// func GoValidator(s interface{}, config []cm.ErrorMetaConfig) (interface{}, int) {
// 	var validate *validator.Validate
// 	validators := cm.NewValidator(validate)
// 	bind := cm.NewBindValidator(validators)

// 	errResponse, errCount := bind.BindValidator(s, config)
// 	return errResponse, errCount
// }

// Validation request from struct field
func Validator(s interface{}) (interface{}, error) {
	if reflect.TypeOf(s).Kind().String() != "struct" {
		return nil, fmt.Errorf("validator value not supported, because %v is not struct", reflect.TypeOf(s).Kind().String())
	} else if res, err := keyExist(s); err != nil || res == 0 {
		return nil, fmt.Errorf("validator value can't be empty struct %v", s)
	}

	val := validator.New()
	err := val.Struct(s)

	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")

	if err := en_translations.RegisterDefaultTranslations(val, trans); err != nil {
		return nil, err
	}

	if err == nil {
		return nil, err
	}

	return formatError(err, trans, s)
}

// Example Usage With cm Tags
// type Login struct {
// 	Email    string `validate:"required" cm:"required=Email tidak boleh kosong"`
// 	Password string `validate:"required" cm:"required=Password tidak boleh kosong"`
// }

// {
//   "errors": [
//     {
//       "msg": "Email tidak boleh kosong",
//       "param": "Email",
//       "tag": "required"
//     },
//     {
//       "msg": "Password tidak boleh kosong",
//       "param": "Password",
//       "tag": "required"
//     }
//   ]
// }

func formatError(err error, trans ut.Translator, customMessage interface{}) (interface{}, error) {
	errRes := make(map[string][]map[string]interface{})
	tags := []string{}

	for i, e := range err.(validator.ValidationErrors) {

		errResult := make(map[string]interface{})
		errResult["param"] = e.StructField()

		if _, ok := reflect.TypeOf(customMessage).Field(i).Tag.Lookup("cm"); !ok {
			errResult["msg"] = e.Translate(trans)
		} else {
			strucField, _ := reflect.TypeOf(customMessage).FieldByName(e.StructField())
			structTags := strucField.Tag.Get("cm")

			regexTag := regexp.MustCompile(`=+[\w].*`)
			regexVal := regexp.MustCompile(`[\w]+=`)
			strArr := strings.Split(structTags, ",")
			tags = append(tags, mergeSlice(strArr)...)

			for j, v := range tags {
				if ok := regexTag.ReplaceAllString(tags[j], ""); ok == e.ActualTag() {
					errResult["msg"] = regexVal.ReplaceAllString(v, "")
					tags = append(tags, "")
				}
			}
		}

		errResult["tag"] = e.ActualTag()
		errRes["errors"] = append(errRes["errors"], errResult)
	}

	return errRes, nil
}

func keyExist(input interface{}) (int, error) {

	if reflect.TypeOf(input).Kind().String() != "struct" {
		return -1, fmt.Errorf("validator value not supported, because %v is not struct", reflect.TypeOf(input).Kind().String())
	}

	received := make(map[string]interface{})
	mapsArr := []string{}

	stringify, err := json.Marshal(&input)

	if err != nil {
		return -1, err
	}

	if err := json.Unmarshal(stringify, &received); err != nil {
		return -1, err
	}

	for i, _ := range received {
		mapsArr = append(mapsArr, i)
	}

	return len(mapsArr), nil
}

func mergeSlice(input ...[]string) (res []string) {
	if len(input) > 0 {
		for i, _ := range input {
			for _, j := range input[i] {
				res = append(res, j)
			}
		}
	} else {
		return res
	}
	return res
}
