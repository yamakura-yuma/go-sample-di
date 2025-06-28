package handler

import (
	"clean-architecture-sample/internal/api"
	"clean-architecture-sample/internal/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

// CreateUserHandler ユーザー作成API
// @Summary ユーザー作成
// @Description 新しいユーザーを登録します
// @Tags users
// @Accept json
// @Produce json
// @Param user body api.CreateUserRequest true "ユーザー情報"
// @Success 201 {object} api.UserResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	var req api.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}
	if err := h.userUseCase.CreateUser(req.ID, req.Name, req.Email); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, api.UserResponse(req))
}

// GetUserHandler ユーザー取得API
// @Summary ユーザー取得
// @Description 指定IDのユーザー情報を取得します
// @Tags users
// @Produce json
// @Param id path string true "ユーザーID"
// @Success 200 {object} api.UserResponse
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
func (h *UserHandler) GetUserHandler(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userUseCase.GetUser(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}
	c.JSON(200, api.UserResponse{ID: user.ID, Name: user.Name, Email: user.Email})
}

// UpdateUserHandler ユーザー更新API
// @Summary ユーザー更新
// @Description 指定IDのユーザー情報を更新します
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "ユーザーID"
// @Param user body api.UpdateUserRequest true "ユーザー情報"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUserHandler(c *gin.Context) {
	id := c.Param("id")
	var req api.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}
	if err := h.userUseCase.UpdateUser(id, req.Name, req.Email); err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, api.UpdateUserRequest(req))
}

// DeleteUserHandler ユーザー削除API
// @Summary ユーザー削除
// @Description 指定IDのユーザーを削除します
// @Tags users
// @Param id path string true "ユーザーID"
// @Success 204 {string} string "No Content"
// @Failure 404 {object} map[string]string
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUserHandler(c *gin.Context) {
	id := c.Param("id")
	if err := h.userUseCase.DeleteUser(id); err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, nil)
}
