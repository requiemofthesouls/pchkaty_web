package user

type UserDTO struct {
	ID         uint   `json:"id,string,omitempty"`
	Username   string `form:"username" json:"username" binding:"required"`
	Email      string `form:"email" json:"email" binding:"required"`
	Password   string `form:"password" json:"password" binding:"required"`
	Name       string `form:"name" json:"name"`
	Surname    string `form:"surname" json:"surname"`
	Patronymic string `form:"patronymic" json:"patronymic"`
}
