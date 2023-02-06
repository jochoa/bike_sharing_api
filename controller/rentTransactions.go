package controller

import (
    "bike_sharing_api/helper"
    "bike_sharing_api/model"
    "github.com/gin-gonic/gin"
    "net/http"

	"log"
)

type bikeResponseData struct {
	ID  uint
	Type string
	Status string
	SerialNumber string
	Kilometers uint
}

type errorResponse struct {
	ErrorDescription string `json:"error_description"`
	ErrorCode int `json:"error_code"`
}
// @BasePath /api

// @Security BearerAuth
// TransactionExample godoc
// @Summary create a rental transaction
// @Schemes
// @Description create a new rental transaction
// @Tags example
// @Param json body string true "json object for a transaction" SchemaExample({"Status": "rented","requestBikeById": {"id": 8}})
// @Accept json
// @Produce json
// @Success 201 {string} AddTransaction
// @Router /api/transaction [post]
func AddTransaction(context *gin.Context) {
    var input model.RentTransactions
	// the "&" in the parameter name means that the variable maz be mutated
    if err := context.ShouldBindJSON(&input); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	var requestBikeId = input.RequestBikeById.Id

	log.Print("[AddTransaction] Request bicycle id: ", requestBikeId)

    user, err := helper.CurrentUser(context)

    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	// Checking if the user has already a rented bicycle
	checkErr := checkUserHasRented(user.ID)
	if checkErr == true {
		var t = errorResponse{"user already rented a bicycle", 1}
		context.JSON(http.StatusBadRequest, gin.H{"data":t})
		return
	}
	//TODO: additionally check if the requestBikeId is still available

	updateBicycle(requestBikeId, false)

    input.UserID = user.ID
	input.BicycleID = requestBikeId

    savedTransaction, err := input.Save()

    if err != nil {
		log.Print( "Error: ", err.Error() )
        context.JSON(http.StatusBadRequest, gin.H{"error": "Error in request"})
        return
    }

    context.JSON(http.StatusCreated, gin.H{"data": savedTransaction})
}

// @BasePath /api

// @Security BearerAuth
// GetAllTransactionsExample godoc
// @Summary get all the rental transactions
// @Schemes
// @Description get all the rental transaction (normally this is only for user with admin role)
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} GetAllTransactions
// @Router /api/transaction [get]
func GetAllTransactions(context *gin.Context) {
    user, err := helper.CurrentUser(context)

    if err != nil {
		log.Print( "Error: ", err.Error() )
        context.JSON(http.StatusBadRequest, gin.H{"error": "Error in request"})
        return
    }

	// Retrieves the bicycle id 
	var bicycleId = user.RentTransactions[0].BicycleID
	var status = user.RentTransactions[0].Status

	// we got the ids now lets find the bicycle data
	var data, resultErr = model.FindBicycleById(bicycleId) 

	if resultErr != nil {
		log.Print( "Error: ", err.Error() )
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error in request"})
		return
	}

	var t = bikeResponseData{bicycleId, data.Type, status, data.SerialNumber, data.Kilometers}
    context.JSON(http.StatusOK, gin.H{"data": t})
}

func checkUserHasRented(id uint) (result bool) {
	// returns a RentTransactions type
	var input, err = model.FindTransactionByUserId(id, helper.RENTED)
	
	if err != nil {
		log.Print( "Error checking user: ", err )
        result = true
		return
    }

	// if the result is equal to an empty struct of type RentTransactions
	// then we dont have any "rented" records for that user
	if (model.RentTransactions{}) == input {
		result = false
	} else {
		result = true
	}
	return 
}

func updateBicycle(id uint, status bool) (result bool) {
	var err = model.UpdateBicycle(id, status)
	
	if err != nil {
		log.Print( "Error checking user: ", err )
        result = true
		return
    }

	result = false
	return
}
// @BasePath /api

// @Security BearerAuth
// UpdateTransactionExample godoc
// @Summary Returns a bicycle api
// @Schemes
// @Description used for returning a bicycle so that the user is allow again to rent another bicycle
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Return a bicycle
// @Router /api/return_bicycle [patch]
func UpdateTransaction(context *gin.Context){
	//Function used when customer returns a bicycle
	user, err := helper.CurrentUser(context)

    if err != nil {
		log.Print( "Error checking user: ", err.Error() )
        context.JSON(http.StatusBadRequest, gin.H{"error": "Error in request"})
        return
    }

	// We need to update the transaction record and the bicycle
	var resultErr = model.UpdateTransactionById(user.RentTransactions) 

	if resultErr != nil {
		log.Print( "Error checking user: ", resultErr )
        context.JSON(http.StatusBadRequest, gin.H{"error": "Error in request"})
	}
	var bicycleId = user.RentTransactions[0].BicycleID

	updateBicycle(bicycleId, true)
	//TODO: handle errors
	context.JSON(http.StatusOK, gin.H{"status": "success"})

}