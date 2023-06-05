package repositories

import (
	"bootcamp/entities"
	"fmt"

	"gorm.io/gorm"
)

type Account struct {
	db *gorm.DB
}

func NewAccount(dbCrud *gorm.DB) Account {
	return Account{
		db: dbCrud,
	}
}

type AccountRepoInterface interface {
	GetAllAccount() ([]entities.Account, error)
	RegisterAccount(account *entities.Account) (*entities.Account, error)
	CreateApproval(*entities.Approval) (*entities.Approval, error)
	GetAdminById(uint64) (entities.Account, error)
	GetApprovalByAdminId(uint64) (entities.Approval, error)
	UpdateApproval(*entities.Approval) (*entities.Approval, error)
	UpdateAdmin(*entities.Account) (*entities.Account, error)
	GetAllApprovalRequest() ([]entities.Approval, error)
	GetAdminByUsername(username string) (entities.Account, error)
}

func (repo Account) GetAllAccount() ([]entities.Account, error) {
	var accounts []entities.Account
	result := repo.db.Table("actors").Find(&accounts)

	return accounts, result.Error
}

func (repo Account) RegisterAccount(account *entities.Account) (*entities.Account, error) {
	err := repo.db.Table("actors").Model(&entities.Account{}).Create(account).Error
	return account, err
}

func (repo Account) CreateApproval(approval *entities.Approval) (*entities.Approval, error) {
	err := repo.db.Table("register_approvals").Model(&entities.Approval{}).Create(approval).Error
	return approval, err
}

func (repo Account) GetAdminById(adminId uint64) (entities.Account, error) {
	admin := entities.Account{}
	err := repo.db.Table("actors").Take(&admin, adminId).Error
	return admin, err
}

func (repo Account) GetApprovalByAdminId(adminId uint64) (entities.Approval, error) {
	approval := entities.Approval{}
	err := repo.db.Table("register_approvals").Where("admin_id = ?", adminId).Take(&approval).Error
	return approval, err
}

func (repo Account) UpdateApproval(approval *entities.Approval) (*entities.Approval, error) {
	var superAdminId *uint
	newSuperAdminId := uint(1)
	superAdminId = &newSuperAdminId

	var status *bool
	newStatus := true
	status = &newStatus

	err := repo.db.Table("register_approvals").Model(&approval).Updates(entities.Approval{SuperAdminId: superAdminId, Status: status}).Error
	return approval, err
}

func (repo Account) UpdateAdmin(account *entities.Account) (*entities.Account, error) {
	err := repo.db.Table("actors").Model(&account).Updates(entities.Account{Active: true, Verified: true}).Error
	return account, err
}

func (repo Account) GetAllApprovalRequest() ([]entities.Approval, error) {
	approvals := []entities.Approval{}
	err := repo.db.Table("register_approvals").Where("super_admin_id IS NULL AND status IS NULL").Find(&approvals).Error
	fmt.Println(approvals)
	return approvals, err
}

func (repo Account) GetAdminByUsername(username string) (entities.Account, error) {
	admin := entities.Account{}
	err := repo.db.Table("actors").Where("username = ?", username).First(&admin).Error
	return admin, err
}
