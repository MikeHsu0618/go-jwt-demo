package repository

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"-"`
}

type Repository interface {
	InsertUser(username string, password string) (*User, error)
	FindUserByUsername(username string) (*User, error)
	FindUserByID(id uint) (*User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) InsertUser(username string, password string) (*User, error) {
	user := User{
		Name:     username,
		Password: password,
	}
	if res := r.db.Create(&user); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func (r *repository) FindUserByUsername(username string) (*User, error) {
	var user User
	if res := r.db.Where("name = ?", username).Find(&user); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func (r *repository) FindUserByID(id uint) (*User, error) {
	var user User
	if res := r.db.Find(&user, id); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}
