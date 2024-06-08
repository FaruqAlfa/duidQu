package repository

import (
    "duidQu/db"
    "duidQu/model"
    "gorm.io/gorm"
)

type UserRepository struct {
    DB *gorm.DB
}

func NewUserRepository() *UserRepository {
    return &UserRepository{
        DB: db.DB,
    }
}

func (r *UserRepository) CreateUser(user *model.User) error {
    return r.DB.Create(user).Error
}

func (r *UserRepository) GetUserByID(id uint) (*model.User, error) {
    var user model.User
    if err := r.DB.First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) UpdateUser(user *model.User) error {
    return r.DB.Save(user).Error
}

func (r *UserRepository) DeleteUser(id uint) error {
    return r.DB.Delete(&model.User{}, id).Error
}
