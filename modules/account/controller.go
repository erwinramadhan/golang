package account

import (
	"bootcamp/dto"
	"bootcamp/entities"
)

type AccountController interface {
	GetAllAccount() ([]AccountResponse, error)
	RegisterAccount(req RegisterAccountParam) (any, error)
	ApproveAdmin(uint64) (any, error)
	GetAllApprovalRequest() ([]entities.Approval, error)
	LoginAccount(LoginInput) error
}

type accountController struct {
	accountUseCase AccountUseCase
}

func (uc accountController) GetAllAccount() ([]AccountResponse, error) {
	accounts, err := uc.accountUseCase.GetAllAccount()
	var resp []AccountResponse
	for i := 0; i < len(accounts); i++ {
		resp = append(resp, AccountResponse{
			ID:        accounts[i].ID,
			RoleID:    accounts[i].RoleID,
			Verified:  accounts[i].Verified,
			Active:    accounts[i].Active,
			CreatedAt: accounts[i].CreatedAt,
			UpdatedAt: accounts[i].UpdatedAt,
		})
	}
	return resp, err
}

func (uc accountController) RegisterAccount(req RegisterAccountParam) (any, error) {
	account, err := uc.accountUseCase.RegisterAccount(req)

	if err != nil {
		return nil, err
	}
	res := SuccessCreateAccount{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success create user",
			Message:      "Success Register",
			ResponseTime: "",
		},
		Data: AccountResponse{
			ID:        account.ID,
			RoleID:    account.RoleID,
			Verified:  account.Verified,
			Active:    account.Active,
			CreatedAt: account.CreatedAt,
			UpdatedAt: account.UpdatedAt,
		},
	}
	return res, nil
}

func (uc accountController) ApproveAdmin(adminId uint64) (any, error) {
	admin, err := uc.accountUseCase.ApproveAdmin(adminId)

	if err != nil {
		return nil, err
	}
	return admin, err
}

func (uc accountController) GetAllApprovalRequest() ([]entities.Approval, error) {
	approvals, err := uc.accountUseCase.GetAllApprovalRequest()
	if err != nil {
		return nil, err
	}
	return approvals, err
}

func (uc accountController) LoginAccount(input LoginInput) error {
	err := uc.accountUseCase.LoginAccount(input) 
	return err
}
