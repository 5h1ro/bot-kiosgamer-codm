package rest

import (
	"botkiosgamercodm/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Body struct {
	SessionKey string `json:"sessioin_key"`
}
type User struct {
	userService users.Service
}

func NewUser(userService users.Service) *User {
	return &User{userService}
}

func (u *User) Redeem(c *gin.Context) {
	u.userService.CreateUserService("8498880835154196154", "1fgrliutzvccvv0142rfzb66pwi3un3z")
	c.JSON(http.StatusOK, "done")
}
