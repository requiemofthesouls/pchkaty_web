package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"pchkaty_web/db"
	"strconv"
)

func InitUserAPI(db *gorm.DB) UserAPI {
	userRepository := ProvideUserRepository(db)
	userService := ProvideUserService(userRepository)
	userAPI := ProvideUserAPI(userService)

	return userAPI
}

type UserAPI struct {
	UserService UserService
}

func ProvideUserAPI(u UserService) UserAPI {
	return UserAPI{UserService: u}
}

func (u *UserAPI) FindAll(c *gin.Context) {
	users := u.UserService.FindAll()
	c.JSON(http.StatusOK, gin.H{"users": ToUserDTOs(users)})
}

func (u *UserAPI) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := u.UserService.FindByID(uint(id))
	c.JSON(http.StatusOK, gin.H{"user": ToUserDTO(user)})
}

func (u *UserAPI) Create(c *gin.Context) {
	var dto UserDTO
	err := c.BindJSON(&dto)

	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Fatalln(err)
		return
	}

	createdUser := u.UserService.Save(ToUser(dto))
	c.JSON(http.StatusCreated, gin.H{"user": ToUserDTO(createdUser)})
}

func (u *UserAPI) Update(c *gin.Context) {
	var dto UserDTO
	err := c.BindJSON(&dto)

	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Fatalln(err)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	user := u.UserService.FindByID(uint(id))

	if user == (db.User{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	user.Username = dto.Username
	user.Email = dto.Email
	user.Password = dto.Password
	user.Name = dto.Name
	user.Surname = dto.Surname
	user.Patronymic = dto.Patronymic
	u.UserService.Save(user)

	c.Status(http.StatusOK)
}

func (u *UserAPI) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := u.UserService.FindByID(uint(id))

	if user == (db.User{}) {
		c.Status(http.StatusBadRequest)
	}

	u.UserService.Delete(user)

	c.Status(http.StatusOK)
}
