package routes

import (
	"net/http"

	"example.com/rest-api-go/models"
	"example.com/rest-api-go/utils"
	"github.com/gin-gonic/gin"
)
func signup(context * gin.Context){
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := user.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}
func signin(context * gin.Context){
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := user.ValidateCredentials(); if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
	}
	token, err := utils.GenerateToken(user.Email,user.ID); if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Signin successful", "token": token})
}
func getUsers(context *gin.Context) {
	users, err := models.GetUsers(); if err != nil{
		 context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		 return
	}
	context.JSON(http.StatusOK, users)

}
func getUserByEmail(context *gin.Context){
	 email := context.Param("email"); if email == ""{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email"})
	 }
	 user := models.User{Email: email}
	 fetchedUser, err := user.GetUserByEmailId(); if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}
	context.JSON(http.StatusOK, fetchedUser)
}