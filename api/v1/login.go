package v1

import (
	"ginblog/middleware"
	"ginblog/model"
	"ginblog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)
	var token string
	var code int

	code = model.CheckLogin(formData.Username, formData.Password)

	if code == errmsg.SUCCSE {
		token, code = middleware.SetToken(formData.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    formData.Username,
		"id":      formData.ID,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
