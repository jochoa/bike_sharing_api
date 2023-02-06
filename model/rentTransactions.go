package model

import (
    "bike_sharing_api/database"

    "gorm.io/gorm"
)

// filter fields
type RequestedBykeId struct {
    Id uint `gorm:"-"` 
}

type RentTransactions struct { 
	gorm.Model
    Id uint `json:"Id"`
	Status string `gorm:"size:20;not null" json:"status"`
    RequestBikeById RequestedBykeId `gorm:"embedded"`
	UserID uint
	BicycleID uint
	//User User
	//Bicycle Bicycle `json:"bicycle,omitempty" gorm:"foreignKey:BicycleID;references:ID"`
}

func (rentTransactions *RentTransactions) Save() (*RentTransactions, error) {
    err := database.Database.Create(&rentTransactions).Error
    if err != nil {
        return &RentTransactions{}, err
    }
    return rentTransactions, nil
}
//Hook BeforeSave
/*
func (rentTransactions *RentTransactions) BeforeSave(*gorm.DB) error {
    //TODO: update the bicycles table to make the bicycle not available
}
*/
func FindTransactionByUserId(id uint, status string) (RentTransactions, error){

    var transaction RentTransactions
    result := database.Database.Where("user_id = ? AND status = ?", id, status).Find(&transaction)
	
    if result.Error != nil {
        return RentTransactions{}, result.Error
    }

	return transaction, nil
}

func UpdateTransactionById(transaction []RentTransactions)( error){

    result := database.Database.Model(&transaction).Update("status", "returned")

    if result.Error != nil {
        return result.Error
    }

    return nil
}