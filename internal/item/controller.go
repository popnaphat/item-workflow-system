package item

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"task-api/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Controller struct {
	Service Service
}

func NewController(db *gorm.DB) Controller {
	return Controller{
		Service: NewService(db),
	}
}

type ApiError struct {
	Field  string
	Reason string
}

func msgForTag(tag, param string) string {
	switch tag {
	case "required":
		return "จำเป็นต้องกรอกข้อมูลนี้"
	case "email":
		return "Invalid email"
	case "gt":
		return fmt.Sprintf("Number must be greater than %v", param)
	case "gte":
		return fmt.Sprintf("Number must be greater than or equal to %v", param)
	}
	return ""
}

func getValidationErrors(err error) []ApiError {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ApiError, len(ve))
		for i, fe := range ve {
			out[i] = ApiError{fe.Field(), msgForTag(fe.Tag(), fe.Param())}
		}
		return out
	}
	return nil
}

func (controller Controller) CreateItem(ctx *gin.Context) {
	// อ่าน userID จากคุกกี้
	userIDStr, err := ctx.Cookie("userID")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization required"})
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32) // แปลงเป็น uint
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid user ID"})
		return
	}

	// Bind the request body to the model.RequestItem struct
	var request model.RequestItem
	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": getValidationErrors(err),
		})
		return
	}

	// Create the item with the userID as the OwnerID
	item, err := controller.Service.Create(request, uint(userID)) // Cast userID to uint
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Respond with the created item
	ctx.JSON(http.StatusCreated, item)
}

func (controller Controller) FindItems(ctx *gin.Context) {
	// Bind query parameters
	var request model.RequestFindItem

	log.Println("find items")
	if err := ctx.BindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	// Find
	items, err := controller.Service.Find(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"data": items,
	// })
	ctx.JSON(http.StatusOK, items)
}
func (controller Controller) FindItem(ctx *gin.Context) {
	// Bind query parameters
	var (
		request model.RequestFindItem
	)
	if err := ctx.BindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Find

	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	items, err := controller.Service.FindByID(uint(id), request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"data": items,
	// })
	ctx.JSON(http.StatusOK, items)
}
func (controller Controller) UpdateItemStatus(ctx *gin.Context) {
	// Bind the request body to the model.RequestUpdateItem struct
	var request model.RequestUpdateItem

	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	// Path param
	id, _ := strconv.Atoi(ctx.Param("id"))

	// Update status
	item, err := controller.Service.UpdateStatus(uint(id), request.Status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"data": item,
	// })
	ctx.JSON(http.StatusOK, item)
}
func (controller Controller) Update(ctx *gin.Context) {
	// Bind
	var (
		request model.RequestItem
	)
	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Path param
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	// Update
	item, err := controller.Service.UpdateItem(uint(id), request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"data": item,
	// })
	ctx.JSON(http.StatusOK, item)
}
func (controller Controller) DeleteItem(ctx *gin.Context) {

	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	items, err := controller.Service.DeleteItem(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data that delete Success": items,
	})
}
