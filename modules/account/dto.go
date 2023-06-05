package account

import (
	"bootcamp/dto"
	"bootcamp/entities"
	"time"
)

type RegisterAccountParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SuccessCreateAccount struct {
	dto.ResponseMeta
	Data AccountResponse `json:"data"`
}

type SuccessGetAllAccount struct {
	dto.ResponseMeta
	Data []AccountResponse `json:"data"`
}

type SuccessGetAllApprovalRequest struct {
	dto.ResponseMeta
	Data []entities.Approval `json:"data"`
}

type AccountResponse struct {
	ID        uint      `json:"id"`
	RoleID    uint      `json:"role_id"`
	Verified  bool      `json:"verified"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateAccountRequestBody struct {
	RoleID   uint `json:"role_id"`
	Verified bool `json:"verified"`
	Active   bool `json:"active"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
