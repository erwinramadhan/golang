package routes

import (
	"bootcamp/modules/account"
	"bootcamp/utils/db"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.New()
	dbCrud := db.GormMysql()

	accountHandler := account.NewAccountRouter(dbCrud)
	accountHandler.Handle(router)

	return router
}
