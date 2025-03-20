package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDB(dsn string) (*gorm.DB, error) {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Perform migrations
	/*
		err = db.AutoMigrate(&Orders{}, &Transaction{})
		if err != nil {
			panic("failed to migrate database")
		}

		fmt.Println("Migrations completed successfully!")

		err = SeedProducts(db)
		if err != nil {
			fmt.Println("couldnt seed db")
		}
	*/
	return db, nil

}

/*
func SeedProducts(db *gorm.DB) error {
	products := []Product{
		{
			ID:             1,
			Slug:           "not-avg-black-joggers",
			Name:           "Not Average Black Joggers",
			Description:    "100% Cotton Joggers designed for style and comfort",
			InventoryLevel: 20,
			Price:          60000,
			CreatedAt:      time.Now(),
		},
		{
			ID:             2,
			Slug:           "not-avg-grey-joggers",
			Name:           "Not Average Grey Joggers",
			Description:    "100% Cotton Joggers designed for style and comfort",
			InventoryLevel: 20,
			Price:          60000,
			CreatedAt:      time.Now(),
		},
		{
			ID:             3,
			Slug:           "not-avg-skull-cap",
			Name:           "Not Average New Religion Skull Cap",
			Description:    "Silk Skull caps designed for perfect fit",
			InventoryLevel: 20,
			Price:          20000,
			CreatedAt:      time.Now(),
		},
		{
			ID:             4,
			Slug:           "not-average-nr-shirt",
			Name:           "Not Average New Religion Shirt",
			Description:    "100% Cotton Joggers designed for style and comfort",
			InventoryLevel: 20,
			Price:          50000,
			CreatedAt:      time.Now(),
		},
		{
			ID:             5,
			Slug:           "not-average-NR-sleevless-shirt",
			Name:           "Not Average New Religion Sleeveless Shirt",
			Description:    "100% Cotton Joggers designed for style and comfort",
			InventoryLevel: 20,
			Price:          45000,
			CreatedAt:      time.Now(),
		},
	}

	for _, product := range products {
		db.Create(&product)
	}

	return nil
}
*/
