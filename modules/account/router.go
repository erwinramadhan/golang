package account

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AccountRouter struct {
	hndlr AccountRequestHandler
}

func NewAccountRouter(dbCrud *gorm.DB) AccountRouter {
	return AccountRouter{hndlr: NewAccountRequestHandler(
		dbCrud,
	)}
}

func (r AccountRouter) Handle(router *gin.Engine) {
	basepath := "/account"
	account := router.Group(basepath)

	account.GET("/", r.hndlr.GetAllAccount)
	account.POST("/register", r.hndlr.RegisterAccount)
	account.POST("/login", r.hndlr.LoginAccount)
	account.PATCH("/approve/:id", r.hndlr.ApproveAdmin)
	account.GET("/approval", r.hndlr.GetAllApprovalRequest)
}
