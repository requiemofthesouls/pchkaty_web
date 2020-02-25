package user

import (
	"github.com/jinzhu/gorm"
	"pchkaty_web/backend/db"
)

type UserRepository struct {
	DB *gorm.DB
}

func ProvideUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}

func (u *UserRepository) FindAll() []db.User {
	var users []db.User
	u.DB.Find(&users)
	return users
}

func (u *UserRepository) FindByID(id uint) db.User {
	var user db.User
	u.DB.First(&user, id)
	return user
}

func (u *UserRepository) Save(user db.User) db.User {
	u.DB.Save(&user)
	return user
}

func (u *UserRepository) Delete(user db.User) {
	u.DB.Delete(&user)
}
