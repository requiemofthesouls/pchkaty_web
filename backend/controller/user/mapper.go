package user

import (
	"pchkaty_web/backend/db"
)

func ToUser(dto UserDTO) db.User {
	return db.User{
		Username:   dto.Username,
		Email:      dto.Email,
		Password:   dto.Password,
		Name:       dto.Name,
		Surname:    dto.Surname,
		Patronymic: dto.Patronymic,
	}
}

func ToUserDTO(user db.User) UserDTO {
	return UserDTO{
		ID:         user.ID,
		Username:   user.Username,
		Email:      user.Email,
		Password:   user.Password,
		Name:       user.Name,
		Surname:    user.Surname,
		Patronymic: user.Patronymic,
	}
}

func ToUserDTOs(users []db.User) []UserDTO {
	userdtos := make([]UserDTO, len(users))

	for index, user := range users {
		userdtos[index] = ToUserDTO(user)
	}
	return userdtos
}
