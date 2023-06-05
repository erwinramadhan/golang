package account

import (
	"bootcamp/dto"
	"bootcamp/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AccountRequestHandler struct {
	ctr AccountController
}

func NewAccountRequestHandler(
	dbCrud *gorm.DB,
) AccountRequestHandler {
	return AccountRequestHandler{
		ctr: accountController{
			accountUseCase: accountUseCase{
				accountRepo: repositories.NewAccount(dbCrud),
			},
		}}
}

func (h AccountRequestHandler) GetAllAccount(c *gin.Context) {
	res, err := h.ctr.GetAllAccount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, SuccessGetAllAccount{Data: res})
}

func (h AccountRequestHandler) RegisterAccount(c *gin.Context) {
	request := RegisterAccountParam{}
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	res, err := h.ctr.RegisterAccount(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h AccountRequestHandler) ApproveAdmin(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
	}
	res, err := h.ctr.ApproveAdmin(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
	}
	c.JSON(http.StatusOK, res)
}

func (h AccountRequestHandler) GetAllApprovalRequest(c *gin.Context) {
	res, err := h.ctr.GetAllApprovalRequest()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, SuccessGetAllApprovalRequest{
		ResponseMeta: dto.ResponseMeta{Success: true, MessageTitle: "Success", Message: "et all approval request", ResponseTime: gin.Mode()},
		Data:         res,
	})
}

func (h AccountRequestHandler) LoginAccount(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.ctr.LoginAccount(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "login success", "token": ""})
}
