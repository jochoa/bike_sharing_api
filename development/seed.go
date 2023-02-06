package development

import (
	"bike_sharing_api/database"
    "log"
	"bike_sharing_api/model"
)

func Seed(){

	var count int64
	//Check if table is empty
	database.Database.Table("bicycles").Count(&count)

	if count == 0 {
		log.Print("Preparing to seed data")

		//Insert batch
		var bicycles = []model.Bicycle{
			{ Type:"folding", SerialNumber:"2343543-LIJSENR23535", Kilometers: 0 },
			{ Type:"road", SerialNumber:"2343543-232DGSRERF23", Kilometers: 0 },
			{ Type:"folding", SerialNumber:"2343543-34554GRTRTTR", Kilometers: 0 }}
		result := database.Database.Create(&bicycles)
		log.Print("Result: ", result)
	}
}