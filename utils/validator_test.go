package util

import (
	"fmt"
	"reflect"
	"testing"
)

type Login struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,gt=7"`
}

type LoginSingleCM struct {
	Email    string `validate:"required,email" cm:"required=email tidak boleh kosong"`
	Password string `validate:"required,gt=7" cm:"required=password tidak boleh kosong"`
}

type LoginMultiCM struct {
	Email    string `validate:"required,email" cm:"required=email tidak boleh kosong,email=email format tidak valid"`
	Password string `validate:"required,gt=7" cm:"required=password tidak boleh kosong,gt=password harus lebih besar dari 7"`
}

func TestValidator(action *testing.T) {
	action.Run("Should be TestValidator - email is not valid", func(t *testing.T) {
		payload := Login{Email: "johndoe@#gmail.com", Password: "qwerty12"}
		res, _ := Validator(payload)

		for _, v := range res.(map[string][]map[string]interface{})["errors"] {
			switch v["msg"] {
			case "Email must be a valid email address":
				assert(t, v["msg"], "Email must be a valid email address")

			default:
				t.FailNow()
			}
		}
	})

	action.Run("Should be TestValidator - password must be greater than 7", func(t *testing.T) {
		payload := Login{Email: "johndoe@gmail.com", Password: "qwerty"}
		res, _ := Validator(payload)

		for _, v := range res.(map[string][]map[string]interface{})["errors"] {
			switch v["msg"] {
			case "Password must be greater than 7 characters in length":
				assert(t, v["msg"], "Password must be greater than 7 characters in length")

			default:
				t.FailNow()
			}
		}
	})

	action.Run("Should be TestValidator with single cm - email and password is empty", func(t *testing.T) {
		payload := LoginSingleCM{Email: "", Password: ""}
		res, _ := Validator(payload)

		if count := len(res.(map[string][]map[string]interface{})["errors"]); count == 2 {
			assert(t, 2, count)
		} else {
			t.FailNow()
		}
	})

	action.Run("Should be TestValidator with single cm - email and password not valid", func(t *testing.T) {
		payload := LoginSingleCM{Email: "johndoe#gmail.com", Password: "qwert12"}
		res, _ := Validator(payload)

		if count := len(res.(map[string][]map[string]interface{})["errors"]); count == 2 {
			assert(t, 2, count)
		} else {
			t.FailNow()
		}
	})

	action.Run("Should be TestValidator with single cm - email is not empty", func(t *testing.T) {
		payload := LoginSingleCM{Email: "", Password: "qwerty12"}
		res, _ := Validator(payload)

		for _, v := range res.(map[string][]map[string]interface{})["errors"] {
			switch v["msg"] {
			case "email tidak boleh kosong":
				assert(t, v["msg"], "email tidak boleh kosong")

			default:
				t.FailNow()
			}
		}
	})

	action.Run("Should be TestValidator with single cm - password is not empty", func(t *testing.T) {
		payload := LoginSingleCM{Email: "johndoe@gmail.com", Password: ""}
		res, _ := Validator(payload)

		for _, v := range res.(map[string][]map[string]interface{})["errors"] {
			switch v["msg"] {
			case "password tidak boleh kosong":
				assert(t, v["msg"], "password tidak boleh kosong")

			default:
				t.FailNow()
			}
		}
	})

	action.Run("Should be TestValidator with multi cm - email and password is empty", func(t *testing.T) {
		payload := LoginMultiCM{Email: "", Password: ""}
		res, _ := Validator(payload)

		if count := len(res.(map[string][]map[string]interface{})["errors"]); count == 2 {
			assert(t, 2, count)
		} else {
			t.FailNow()
		}
	})

	action.Run("Should be TestValidator with multi cm - email and password not valid", func(t *testing.T) {
		payload := LoginMultiCM{Email: "johndoe#gmail.com", Password: "qwert12"}
		res, _ := Validator(payload)

		if count := len(res.(map[string][]map[string]interface{})["errors"]); count == 2 {
			assert(t, 2, count)
		} else {
			t.FailNow()
		}
	})

	action.Run("Should be TestValidator with multi cm - email not valid", func(t *testing.T) {
		payload := LoginMultiCM{Email: "johndoe#gmail.com", Password: "qwerty12"}
		res, _ := Validator(payload)

		for _, v := range res.(map[string][]map[string]interface{})["errors"] {
			switch v["msg"] {
			case "email format tidak valid":
				assert(t, v["msg"], "email format tidak valid")

			default:
				t.FailNow()
			}
		}
	})

	action.Run("Should be TestValidator with multi cm - password must be greater than 7", func(t *testing.T) {
		payload := LoginMultiCM{Email: "johndoe@gmail.com", Password: "qwerty"}
		res, _ := Validator(payload)

		for _, v := range res.(map[string][]map[string]interface{})["errors"] {
			switch v["msg"] {
			case "password harus lebih besar dari 7":
				assert(t, v["msg"], "password harus lebih besar dari 7")

			default:
				t.FailNow()
			}
		}
	})
}

func assert(t *testing.T, val1, val2 interface{}) {
	defer t.Cleanup(func() {
		val1 = nil
		val2 = nil
	})

	if val1 != val2 {
		t.FailNow()
	}
}

func TestMergeSlice(action *testing.T) {
	action.Run("Should be TestMergeSlice - is empty response", func(t *testing.T) {
		data := []string{}

		res := mergeSlice(data)
		if len(res) == 0 {
			assert(t, len(res), 0)
		}
	})

	action.Run("Should be TestMergeSlice - not empty response", func(t *testing.T) {
		data := []string{"a", "b", "c"}

		res := mergeSlice(data)
		if len(res) == 0 {
			assert(t, len(res), 0)
		}
	})
}

func TestKeyExist(action *testing.T) {
	action.Run("Should be TestKeyExist - key not exist", func(t *testing.T) {
		data := struct{}{}
		res, err := keyExist(data)
		if err != nil {
			assert(t, res, -1)
		}
	})

	action.Run("Should be TestKeyExist - key is exist", func(t *testing.T) {
		data := struct {
			Username string `json:"username"`
			Email    string `json:"email"`
		}{}

		res, err := keyExist(data)
		if err != nil {
			t.FailNow()
		}

		assert(t, res, 2)
	})
}

func TestFormatError(action *testing.T) {
	action.Run("Should be TestFormatError - value is not struct", func(t *testing.T) {
		payload := "hello world"

		_, err := Validator(payload)
		if err != nil {
			assert(t, err.Error(), fmt.Sprintf("validator value not supported, because %v is not struct", reflect.TypeOf(payload).Kind().String()))
		}
	})

	action.Run("Should be TestFormatError - is empty struct", func(t *testing.T) {
		payload := struct{}{}
		_, err := Validator(payload)

		if err != nil {
			assert(t, err.Error(), fmt.Sprintf("validator value can't be empty struct %v", payload))
		}
	})

	action.Run("Should be TestFormatError - not empty value", func(t *testing.T) {
		payload := Login{Email: "", Password: ""}
		res, _ := Validator(payload)

		for _, v := range res.(map[string][]map[string]interface{})["errors"] {
			switch v["msg"] {
			case "Email is a required field":
				assert(t, v["msg"], "Email is a required field")

			case "Password is a required field":
				assert(t, v["msg"], "Password is a required field")

			default:
				t.FailNow()
			}
		}
	})

	action.Run("Should be TestFormatError - success", func(t *testing.T) {
		payload := Login{Email: "johndoe@gmail.com", Password: "qwerty12"}
		res, _ := Validator(payload)
		assert(t, res, nil)
	})
}
