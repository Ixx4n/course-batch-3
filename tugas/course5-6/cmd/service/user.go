package service

import (
	"course5-6/cmd/model"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserSvc struct {
	db *gorm.DB
}

func NewUserSvc(DB *gorm.DB) *UserSvc {
	return &UserSvc{DB}
}

func (userSvc UserSvc) Register(registerRequest model.RegisterRequest) (int, string, error) {

	if registerRequest.Name == "" {
		return 417, "", errors.New("name required")
	}
	if registerRequest.Email == "" {
		return 417, "", errors.New("email required")
	}
	if registerRequest.NoHP == "" {
		return 417, "", errors.New("no hp required")
	}
	if registerRequest.Password == "" {
		return 417, "", errors.New("password required")
	}
	if len(registerRequest.Password) < 6 {
		return 417, "", errors.New("password must more than 6 character")
	}

	var count int64
	err := userSvc.db.Model(&model.User{}).Where("email = ?", registerRequest.Email).Count(&count).Error
	if err != nil {
		return 500, "", err
	}
	if count > 0 {
		return 417, "", errors.New("email already used")
	}

	user := model.NewUser(registerRequest.Name, registerRequest.Email, registerRequest.Password, registerRequest.NoHP)
	err = userSvc.db.Create(user).Error
	if err != nil {
		return 500, "", err
	}

	token, err := GenerateToken(user.ID)
	if err != nil {
		return 500, "", err
	}
	return 201, token, nil
}

func (userSvc UserSvc) Login(loginRequest model.LoginRequest) (int, string, error) {

	if loginRequest.Email == "" || loginRequest.Password == "" {
		return 401, "", errors.New("wrong email/password")
	}

	var user model.User
	err := userSvc.db.Where("email = ?", loginRequest.Email).Take(&user).Error
	if err != nil {
		return 401, "", errors.New("wrong email/password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return 401, "", errors.New("wrong email/password")
	}

	token, _ := GenerateToken(user.ID)

	return 200, token, nil
}

var PrivateKey = []byte("SuperSecretKey")

func GenerateToken(id int) (string, error) {
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"iss": "admin",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(PrivateKey)
}

func DecryptJWT(token string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("auth invalid")
		}
		return PrivateKey, nil
	})
	if err != nil {
		return map[string]interface{}{}, err
	}

	if !parsedToken.Valid {
		return map[string]interface{}{}, err
	}

	return parsedToken.Claims.(jwt.MapClaims), nil
}
