package server

import (
	"fmt"
	"github.com/Xhofe/alist/drivers"
	"github.com/Xhofe/alist/model"
	"github.com/gin-gonic/gin"
	"time"
)

func GetAccounts(c *gin.Context) {
	accounts, err := model.GetAccounts()
	if err != nil {
		ErrorResp(c, err, 500)
		return
	}
	SuccessResp(c, accounts)
}

func SaveAccount(c *gin.Context) {
	var req model.Account
	if err := c.ShouldBind(&req); err != nil {
		ErrorResp(c, err, 400)
		return
	}
	driver, ok := drivers.GetDriver(req.Type)
	if !ok {
		ErrorResp(c, fmt.Errorf("no [%s] driver", req.Type), 400)
		return
	}
	old, ok := model.GetAccount(req.Name)
	now := time.Now()
	req.UpdatedAt = &now
	if err := model.SaveAccount(req); err != nil {
		ErrorResp(c, err, 500)
	} else {
		if ok {
			err = driver.Save(&req, &old)
		} else {
			err = driver.Save(&req, nil)
		}
		if err != nil {
			ErrorResp(c, err, 500)
			return
		}
		SuccessResp(c)
	}
}

func DeleteAccount(c *gin.Context) {
	name := c.Query("name")
	if err := model.DeleteAccount(name); err != nil {
		ErrorResp(c, err, 500)
		return
	}
	SuccessResp(c)
}