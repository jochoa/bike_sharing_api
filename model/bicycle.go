package model

import (
	"bike_sharing_api/database"
	"gorm.io/gorm"
)

var logString string = "[bicycle.model] "

type Bicycle struct {
	gorm.Model
	Type string `gorm:"size:255;not null" json:"type"`
	SerialNumber string `gorm:"size:255;not null;" json:"serialNumber"`
	Kilometers uint `json:"kilometers"`
	IsAvailable bool `gorm:"not null;default:true;" json:"isAvailable"`
	RentTransactions []RentTransactions
}

func (bicycle *Bicycle) Save() (*Bicycle, error) {
    err := database.Database.Create(&bicycle).Error
    if err != nil {
        return &Bicycle{}, err
    }
    return bicycle, nil
}

func FindAllBicycles() ([]Bicycle, error) {
	var bikes []Bicycle
	result := database.Database.Find(&bikes).Scan(&bikes)
	
    if result.Error != nil {
        return nil, result.Error
    }

    return bikes, nil
}

func FindBicycleById(id uint) (Bicycle, error){
	var bike Bicycle
	result := database.Database.First(&bike).Scan(&bike)
	if result.Error != nil {
        return Bicycle{}, result.Error
    }
	return bike, nil
}

func GetAllAvailableBicycles() ([]Bicycle, error) {
	
	var bikes []Bicycle
	result := database.Database.Where("is_rented = ?", false).Find(&bikes)
	
    if result.Error != nil {
        return nil, result.Error
    }

    return bikes, nil
}

func UpdateBicycle(id uint, status bool )(error){
	var bikes []Bicycle

	result := database.Database.Model(&bikes).Where("id = ?", id).Update("is_available", status)

	if result.Error != nil {
        return result.Error
    }
	return nil
}