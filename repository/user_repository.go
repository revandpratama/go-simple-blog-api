package repository

import (
	"github.com/revandpratama/go-simple-blog-api/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(email string) (*entity.User, error)
	GetUserByUsername(username string) (*entity.User, error)
	UsernameExists(username string) bool
	EmailExists(email string) bool
	Register(userData *entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) UsernameExists(username string) bool {
	var user entity.User

	err := r.db.First(&user, "username = ?", username).Error

	return err == nil
}
func (r *userRepository) EmailExists(email string) bool {
	var user entity.User

	err := r.db.First(&user, "email = ?", email).Error

	return err == nil
}

func (r *userRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User

	err := r.db.First(&user, "email = ?", email).Error

	return &user, err
}

func (r *userRepository) GetUserByUsername(username string) (*entity.User, error) {
	var user entity.User

	err := r.db.First(&user, "username = ?", username).Error

	return &user, err
}

func (r *userRepository) Register(userData *entity.User) error {

	err := r.db.Create(&userData).Error

	return err
}
