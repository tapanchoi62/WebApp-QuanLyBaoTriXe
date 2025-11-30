package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/config"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/models"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/utils"
	"golang.org/x/crypto/bcrypt"
)

func generateRandomPassword(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	// Encode base64 và lấy length ký tự cần thiết
	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}

func CreateUser(c *gin.Context) {
	var input models.UserInputCreate

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var password = "123456"

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể hash password"})
		return
	}

	// Tạo user
	user := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
		Role:     input.Role,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể tạo user"})
		return
	}

	// Trả về user + password tạm để frontend gửi reset
	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
	})
}

func GetUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	search := c.Query("search")

	// Paginate User
	data, paging, err := utils.Paginate[models.User](config.DB.Model(&models.User{}), page, pageSize, search, []string{"username"})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Lấy tất cả files của users trong page này 1 lần
	userIDs := make([]uint, 0, len(data))
	for _, u := range data {
		userIDs = append(userIDs, u.ID)
	}

	var files []models.File
	if len(userIDs) > 0 {
		config.DB.Where("entity_type = ? AND entity_id IN ?", "users", userIDs).Find(&files)
	}

	// Map files theo userID
	fileMap := make(map[uint][]models.File)
	for _, f := range files {
		id, _ := strconv.Atoi(f.EntityID)
		fileMap[uint(id)] = append(fileMap[uint(id)], f)
	}

	// Gắn files vào UserResponse
	var result []models.UserResponse
	for _, u := range data {
		resp := models.UserResponse{
			ID:       u.ID,
			Username: u.Username,
			Role:     u.Role,
			Files:    fileMap[u.ID],
		}
		result = append(result, resp)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       result,
		"pagination": paging,
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	var u models.User
	if err := config.DB.First(&u, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy user"})
		return
	}
	var files []models.File
	config.DB.Where("entity_id = ? AND entity_type = ?", id, "users").Find(&files)
	resp := models.UserResponse{
		ID:       u.ID,
		Username: u.Username,
		Role:     u.Role,
		Files:    files,
	}
	c.JSON(http.StatusOK, resp)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User không tồn tại"})
		return
	}

	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Username = input.Username
	user.Password = input.Password
	user.Role = input.Role

	config.DB.Save(&user)

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể xóa user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Xóa thành công"})
}
