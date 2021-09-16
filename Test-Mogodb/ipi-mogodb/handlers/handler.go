package handlers

import (
	"Test-Mogodb/ipi-mogodb/models"
	"Test-Mogodb/ipi-mogodb/resquest"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateUser(c *gin.Context) {

	user := resquest.User{}

	if err := c.ShouldBind(&user); err != nil {
		if ValidationError, ok := err.(validator.ValidationErrors); ok {
			MessageError := make([]string, 0)
			for _, v := range ValidationError {
				MessageError = append(MessageError, v.Error())
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  MessageError,
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err,
		})
		return
	}
	result, err := models.CreateUser(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"errors": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{

		"status":  http.StatusText(http.StatusOK),
		"payload": result,
	})
}
func UpdateUser(c *gin.Context) {

	user := resquest.Userudate{}

	if err := c.ShouldBind(&user); err != nil {
		if ValidationError, ok := err.(validator.ValidationErrors); ok {
			MessageError := make([]string, 0)
			for _, v := range ValidationError {
				MessageError = append(MessageError, v.Error())
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  MessageError,
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err,
		})
		return
	}
	result, err := models.UpdateUser(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"errors": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{

		"status":  http.StatusText(http.StatusOK),
		"payload": result,
	})
}
func FindUser(c *gin.Context) {
	User := resquest.UeserResquestID{}
	if err := c.ShouldBindUri(&User); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err,
		})
		return
	}
	user, err := models.FindUser(User.Id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": user,
	})
}
func LisAllUser(c *gin.Context) {
	list, err := models.LisAllUser()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": list,
	})
}
func LisName(c *gin.Context) {
	User := resquest.UserResquestName{}
	if err := c.ShouldBindUri(&User); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err,
		})
		return
	}
	list, err := models.LisName(User.Name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": list,
	})
}
func DeleteUser(c *gin.Context) {
	User := resquest.UeserResquestID{}
	if err := c.ShouldBindUri(&User); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err,
		})
		return
	}
	message, err := models.DeleteUser(User.Id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": message,
	})
}
