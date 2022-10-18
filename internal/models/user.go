package models

import "github.com/jinzhu/gorm"

type User struct {
	ID    int64
	Name  string
	Email string
}

// CreateUser create a user
func CreateUser(db *gorm.DB, User *User) (err error) {
	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUsers get users by conditions
func GetUsers(db *gorm.DB, User *[]User, conditions map[string]interface{}) (err error) {
	err = db.Where(conditions).Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUser get one user by id
func GetUser(db *gorm.DB, User *User, id int64) (err error) {
	err = db.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser update user by id
func UpdateUser(db *gorm.DB, User *User, id int64) (err error) {
	db.Model(User).Where("id = ?", id).Updates(User)
	return nil
}
