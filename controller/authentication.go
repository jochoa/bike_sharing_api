package controller

import (
	"bike_sharing_api/helper"
    "bike_sharing_api/model"
    "github.com/gin-gonic/gin"
    "net/http"
    //"github.com/swaggo/gin-swagger" // gin-swagger middleware
	//"github.com/swaggo/files" // swagger embed files
    //docs "bike_sharing_api/docs"
)

// @BasePath /auth

// RegisterExample godoc
// @Summary register a user example
// @Schemes
// @Description register a user
// @Tags example
// @Param json body string true "register data for user" SchemaExample({"username":"<username>", "password":"<password>"})
// @Accept json
// @Produce json
// @Success 201 {string} Register
// @Router /auth/register [post]
// Validates, creates a new user, returns the saved user as JSON
func Register(context *gin.Context){
	var input model.AuthenticationInput

    if err := context.ShouldBindJSON(&input); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user := model.User{
        Username: input.Username,
        Password: input.Password,
        Role: input.Role,
    }

    savedUser, err := user.Save()

    if err != nil {
        //TODO: change return message to something generic
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusCreated, gin.H{"user": savedUser})	
}

// @BasePath /auth

// RLoginExample godoc
// @Summary user login
// @Schemes
// @Description user login
// @Tags example
// @Param json body string true "login data for user" SchemaExample({"username":"<username>", "password":"<password>"})
// @Accept json
// @Produce json
// @Success 200 {string} Login
// @Router /auth/login [post]
func Login(context *gin.Context) {
    var input model.AuthenticationInput

    if err := context.ShouldBindJSON(&input); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := model.FindUserByUsername(input.Username)

    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err = user.ValidatePassword(input.Password)

    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    jwt, err := helper.GenerateJWT(user)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusOK, gin.H{"jwt": jwt})
} 