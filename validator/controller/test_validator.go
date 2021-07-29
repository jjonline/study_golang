package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jjonline/go-lib-backend/validation4gin"
	"net/http"
	"reflect"
)

type ValidRuleRequest struct {
	ID   uint   `form:"id" json:"id" binding:"required,min=10,max=1000"`
	Name string `form:"name" json:"name" binding:"required,max=3"`
}

// TestValidator controller
type TestValidator struct{}

func (t TestValidator) ValidRule(ctx *gin.Context) {
	var req ValidRuleRequest
	// binding.Validator = nil // disabled gin valid
	err := ctx.ShouldBindQuery(&req)

	fmt.Println(req)

	if err != nil {
		message := validation4gin.Message{
			"ID.*":          ":ID错误",
			"ID.type":       ":ID类型错误",
			"ID.min":        ":ID不得小于10",
			"ID.max":        ":ID不得大于100",
			"Name.required": ":Name不得为空",
			"*":             ":Name格式有误",
		}
		attribute := validation4gin.FieldMap{
			"ID":   "编号",
			"Name": "名称",
		}
		msg := validation4gin.Translate(err, message, attribute)
		fmt.Println(msg.First())

		if errors, ok := err.(validator.ValidationErrors); ok {
			for _, item := range errors {
				fmt.Println("validator.ValidationErrors")
				fmt.Println(item)
				fmt.Println(item.Tag())
				fmt.Println(item.Namespace())
				fmt.Println(item.StructNamespace())
				fmt.Println(item.Field())
				fmt.Println(item.StructField())
				fmt.Println(item.Value())
				fmt.Println(item.Param())
				fmt.Println(item.Kind())
				fmt.Println(item.Type())
				fmt.Println(item.Error())
				fmt.Println("validator.ValidationErrors")
			}
		}

		// curl --location --request GET 'http://127.0.0.1:9080/'
		// validator.ValidationErrors
		// ++++++++++++++++++++++++++++++++++++++++++++++++++++++
		// ++++++++++++++++++++++++++++++++++++++++++++++++++++++
		// curl --location --request GET 'http://127.0.0.1:9080/?id=a'
		// validator.ValidationErrors
		fmt.Println(reflect.TypeOf(err).String())

		ctx.JSON(http.StatusOK, map[string]string{"type": reflect.TypeOf(err).String(), "error": err.Error(), "msg": msg.First()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]string{"ok": "ok"})
}
