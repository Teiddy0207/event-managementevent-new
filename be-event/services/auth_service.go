package services

import (
	"be-event/models"
	"be-event/repositories"
	"be-event/utils"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type AuthService interface {
	Register(user *models.User) error
	Login(email, password string) (string, error)
	Logout(token string) error
}

type authService struct {
	repo repositories.AuthRepository
}

func NewAuthService(repo repositories.AuthRepository) AuthService {
	return &authService{
		repo: repo,
	}
}

func (s *authService) Register(user *models.User) error {
	//Kiểm tra email đã tồn tại (nếu có hàm FindByEmail)
	existing, _ := s.repo.FindByEmail(user.Email)
	if existing != nil {
		return errors.New("email already exists")
	}

	// Hash mật khẩu
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Gán default role_id nếu chưa có (VD: 2 = user)
	if user.RoleID == 0 {
		user.RoleID = 3
	}

	// Lưu user
	return s.repo.CreateUser(user)
}

func (s *authService) Login(email, password string) (string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil || user == nil {
		return "", errors.New("email hoặc mật khẩu không đúng")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("email hoặc mật khẩu không đúng")
	}

	// Sinh JWT
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	// Lưu token
	loginUser := &models.LoginUser{
		UserID:    user.ID,
		Token:     token,
		ExpiredAt: time.Now().Add(24 * time.Hour),
	}
	//_ = s.repo.SaveLoginToken(loginUser)

	if err := s.repo.SaveLoginToken(loginUser); err != nil {
		log.Println("❌ Lỗi khi lưu token vào DB:", err)
		return "", err
	}

	return token, nil
}

func (s *authService) Logout(token string) error {
	return s.repo.Logout(token)
}
