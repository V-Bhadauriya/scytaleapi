package auto

import (
"fmt"
"github.com/V-Bhadauriya/scytaleapi/api/database"
"github.com/V-Bhadauriya/scytaleapi/api/models"
"github.com/V-Bhadauriya/scytaleapi/api/utils/console"
"log"

)

// Load autogenerates the tables and records
func Load() {
	fmt.Println("Ignore")

		db, err := database.Connect()
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		err = db.Debug().DropTableIfExists(&models.Workload{}).Error
		if err != nil {
			log.Fatal(err)
		}
		err = db.Debug().AutoMigrate(&models.Workload{}).Error
		if err != nil {
			log.Fatal(err)
		}

		for i, _ := range workload {
			err = db.Debug().Model(&models.workload{}).Create(&workload[i]).Error
			if err != nil {
				log.Fatal(err)
			}
			console.Pretty(workload[i])
		}


}
