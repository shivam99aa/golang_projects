package models

import (
	"golang_projects/photogallery_app/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open(postgres.Open(connectionInfo), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, err
	}
	userService := UserService{
		db: db,
	}
	return &userService, nil
}

type UserService struct {
	db *gorm.DB
}

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;uniqueIndex"`
}

// ByID will look up the user by id provided.
// Return will be of 3 types
// 1 - user, nil -> user is found
// 2 - nil, ErrNotFound -> user is not found
// 3 - nil, otherError -> something went wrong
func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	err := us.db.Where("id = ?", id).First(&user).Error
	switch err {
	case nil:
		return &user, nil
	case gorm.ErrRecordNotFound:
		return nil, utils.ErrNotFound
	default:
		return nil, err
	}
}

// create the provided user.
func (us *UserService) Create(user *User) error {
	return us.db.Create(user).Error
}

func (us *UserService) Close() error {
	db, err := us.db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func (us *UserService) DestructiveReset() {
	us.db.Migrator().DropTable(&User{})
	us.db.Migrator().AutoMigrate(&User{})
}
