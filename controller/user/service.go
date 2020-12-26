package user

import "pchkaty_web/db"

type UserService struct {
	UserRepository UserRepository
}

func ProvideUserService(u UserRepository) UserService {
	return UserService{UserRepository: u}
}

func (u *UserService) FindAll() []db.User {
	return u.UserRepository.FindAll()
}

func (u *UserService) FindByID(id uint) db.User {
	return u.UserRepository.FindByID(id)
}

func (u *UserService) Save(user db.User) db.User {
	u.UserRepository.Save(user)
	return user
}

func (u *UserService) Delete(user db.User) {
	u.UserRepository.Delete(user)
}
