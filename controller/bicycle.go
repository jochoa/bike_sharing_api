package controller

import (
    "bike_sharing_api/model"
	//"bike_sharing_api/helper"
    "github.com/gin-gonic/gin"
    "net/http"
	"log"
)

// @BasePath /api

// @Security BearerAuth
// AddBicycleExample godoc
// @Summary add a bicycle to the stock
// @Schemes
// @Description adds a bicycle item to the stock (normally this is only for user with admin role)
// @Tags example
// @Param json body string true "json object for a transaction" SchemaExample({ "type": "<type>", "serial_number": "<serial_number>", "kilometers": <int>})
// @Accept json
// @Produce json
// @Success 201 {string} AddBicycle
// @Router /api/add_bicycle [post]
func AddBicycle(context *gin.Context) {
    var input model.Bicycle
	// the "&" in the parameter name means that the variable maz be mutated
    if err := context.ShouldBindJSON(&input); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    savedBicycle, err := input.Save()

    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusCreated, gin.H{"data": savedBicycle})
}

// @BasePath /api

// @Security BearerAuth
// GetAllBicyclesExample godoc
// @Summary get all the bicycles in the stock
// @Schemes
// @Description get all the bicycles (normally this is only for user with admin role)
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} GetAllBicycles
// @Router /api/get_all_bicycles [get]
func GetAllBicycles(context *gin.Context) {

	var input, err = model.FindAllBicycles()

	if err != nil {
        log.Print("[GetAllBicycles] Error: ", err)
        context.JSON(http.StatusBadRequest, gin.H{"error": "Error in request"})
        return
    }

    context.JSON(http.StatusOK, gin.H{"data": input})
}

// @BasePath /api

// @Security BearerAuth
// GetAllAvailableBicyclesExample godoc
// @Summary get all the available bicycles
// @Schemes
// @Description get all the available bicycles
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} GetAllAvailableBicycles
// @Router /api/get_all_available_bicycles [get]
func GetAllAvailableBicycles(context *gin.Context) {

	var input, err = model.GetAllAvailableBicycles()

	if err != nil {
        log.Print("[GetAllAvailableBicycles] Error: ", err)
        context.JSON(http.StatusBadRequest, gin.H{"error": "Error in request"})
        return
    }

    context.JSON(http.StatusOK, gin.H{"data": input})
}
