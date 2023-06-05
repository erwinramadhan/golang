package account

import (
	"bootcamp/entities"
	"bootcamp/repositories"
	"errors"
	"time"
)

type AccountUseCase interface {
	GetAllAccount() ([]entities.Account, error)
	RegisterAccount(account RegisterAccountParam) (entities.Account, error)
	ApproveAdmin(uint64) (any, error)
	GetAllApprovalRequest() ([]entities.Approval, error)
	LoginAccount(LoginInput) error
}

type accountUseCase struct {
	accountRepo repositories.AccountRepoInterface
}

func (uc accountUseCase) GetAllAccount() ([]entities.Account, error) {
	var accounts []entities.Account
	accounts, err := uc.accountRepo.GetAllAccount()
	return accounts, err
}

func (uc accountUseCase) RegisterAccount(account RegisterAccountParam) (entities.Account, error) {
	newAccount := entities.Account{
		Username:  account.Username,
		Password:  account.Password,
		RoleID:    2,
		Verified:  false,
		Active:    false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err := uc.accountRepo.RegisterAccount(&newAccount)
	if err != nil {
		return newAccount, err
	}

	newApproval := entities.Approval{
		AdminID:      newAccount.ID,
		SuperAdminId: nil,
		Status:       nil,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	_, err = uc.accountRepo.CreateApproval(&newApproval)
	if err != nil {
		return newAccount, err
	}

	return newAccount, nil
}

func (uc accountUseCase) ApproveAdmin(adminId uint64) (any, error) {
	admin, err := uc.accountRepo.GetAdminById(adminId)
	if err != nil {
		return nil, err
	}

	_, err = uc.accountRepo.UpdateAdmin(&admin)
	if err != nil {
		return nil, err
	}

	approval, err := uc.accountRepo.GetApprovalByAdminId(adminId)
	if err != nil {
		return nil, err
	}

	_, err = uc.accountRepo.UpdateApproval(&approval)
	if err != nil {
		return nil, err
	}

	return admin, err
}

func (uc accountUseCase) GetAllApprovalRequest() ([]entities.Approval, error) {
	var entities []entities.Approval
	entities, err := uc.accountRepo.GetAllApprovalRequest()
	return entities, err
}

func (uc accountUseCase) LoginAccount(input LoginInput) error {
	admin, err := uc.accountRepo.GetAdminByUsername(input.Username)
	if err != nil {
		return err
	}

	if admin.Password != input.Password {
		return errors.New("password: didn't match")
	}

	return err
}
