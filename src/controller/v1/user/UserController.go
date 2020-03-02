package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	httpEntity "github.com/sofyan48/ggwp/src/entity/v1/http"
	services "github.com/sofyan48/ggwp/src/service/v1/user"
	"github.com/sofyan48/ggwp/src/util/helper/rest"
)

// V1UserController types
type V1UserController struct {
	UserService services.V1UserServiceInterface
}

// UpdateUsersByID params
// @context: gin context
func (service *V1UserController) UpdateUsersByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, "Bad Request")
		return
	}
	payload := httpEntity.UserRequest{}
	if err := context.ShouldBind(&payload); err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, "Bad Request")
		return
	}
	service.UserService.UpdateUserByID(id, payload)
	rest.ResponseMessages(context, http.StatusNoContent, "Edited")
	return
}

// InsertUsers params
// @context: gin context
func (service *V1UserController) InsertUsers(context *gin.Context) {
	payload := httpEntity.UserRequest{}
	if err := context.ShouldBind(&payload); err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, "Bad Request")
		return
	}
	fmt.Println(payload)
	results, err := service.UserService.InsertUsers(payload)
	if err != nil {
		rest.ResponseMessages(context, http.StatusUnauthorized, fmt.Sprintf("%e", err))
		return
	}
	rest.ResponseData(context, http.StatusCreated, results)
	return
}
