package repositories

import (
	"be-event/models"
	"gorm.io/gorm"
	"time"
)



type AuthRepository interface {
	CreateUser(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	SaveLoginToken(token *models.LoginUser) error
	Logout(token string) error
}

type authRepository struct {
	masterDB  *gorm.DB
	replicaDB *gorm.DB
}

func NewAuthRepository(masterDB, replicaDB *gorm.DB) AuthRepository {
	return &authRepository{

		masterDB:  masterDB,
		replicaDB: replicaDB,
	}
}

func (r *authRepository) CreateUser(user *models.User) error {
	return r.masterDB.Create(user).Error
}

func (r *authRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.replicaDB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Không tìm thấy người dùng
		}
		return nil, err // Lỗi khác
	}
	return &user, nil // Trả về người dùng tìm thấy
}

func (r *authRepository) SaveLoginToken(token *models.LoginUser) error {
	return r.masterDB.Create(token).Error
}

func (r *authRepository) Logout(token string) error {
	return r.masterDB.Model(&models.LoginUser{}).
		Where("token = ?", token).
		Update("expired_at", time.Now()).Error
}

