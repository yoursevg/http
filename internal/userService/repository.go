package userService

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]User, error)
	CreateUser(message User) (User, error)
	UpdateUserByID(id uint, message User) (User, error)
	DeleteUserByID(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

// GetAllUsers - Получение всего списка пользователей
func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

// CreateUser - Создает нового пользователя в бд
func (r *userRepository) CreateUser(newUser User) (User, error) {
	result := r.db.Create(&newUser)
	if result.Error != nil {
		return User{}, result.Error
	}
	return newUser, nil
}

// UpdateUserByID - Ищем пользователя в БД по id и обновляем его данные
func (r *userRepository) UpdateUserByID(id uint, updatedUser User) (User, error) {
	var user User
	user.ID = id
	err := r.db.First(&user).Error
	if err != nil {
		return User{}, err
	}

	if updatedUser.Email != "" {
		user.Email = updatedUser.Email
	}
	if updatedUser.Password != "" {
		user.Password = updatedUser.Password
	}

	err = r.db.Save(&user).Error
	if err != nil {
		return User{}, err
	}

	return user, err
}

// DeleteUserByID - Ищем пользователя в БД по id и удаляем его
func (r *userRepository) DeleteUserByID(id uint) error {
	var userToDelete User
	userToDelete.ID = id
	err := r.db.First(&userToDelete, id).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(&userToDelete).Error
	return err
}
