package api

import (
	"net/http"

	"github.com/SimulatedSakura/go-gin-example/models"
	"github.com/SimulatedSakura/go-gin-example/pkg/e"
	"github.com/SimulatedSakura/go-gin-example/pkg/logging"
	"github.com/SimulatedSakura/go-gin-example/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required;MaxSize(50)"`
	Password string `valid:"Required;MaxSize(50)"`
}

// @Summary Get Auth
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500 {string} string "{"code":500,"data":{},"msg":"err"}"
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{
		Username: username,
		Password: password,
	}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			for _, err := range valid.Errors {
				logging.Info(err.Key, err.Value)
			}
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Value)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
