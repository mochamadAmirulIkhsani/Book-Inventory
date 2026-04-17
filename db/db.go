package db

import (
	"book_inventory/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error load env")
	}

	conn := os.Getenv("POSTGRES_URL")
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	Migrate(db)
	return db
}

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&models.Books{}); err != nil {
		log.Fatal(err)
	}

	var count int64
	db.Model(&models.Books{}).Count(&count)
	if count == 0 {
		seederBook(db)
	}

	data := models.Books{}
	if db.Find(&data).Error != nil {
		seederBook(db)
	}
}

func seederBook(db *gorm.DB) {
	data := []models.Books{{
		Title:       "The Great Gatsby",
		Author:      "F. Scott Fitzgerald",
		Description: "A novel set in the Roaring Twenties, exploring themes of wealth, love, and the American Dream.",
		Stock:       10,
	}, {
		Title:       "To Kill a Mockingbird",
		Author:      "Harper Lee",
		Description: "A powerful story of racial injustice and moral growth in the American South.",
		Stock:       15,
	}, {
		Title:       "1984",
		Author:      "George Orwell",
		Description: "A dystopian novel that delves into themes of totalitarianism, surveillance, and individuality.",
		Stock:       20,
	}, {
		Title:       "Pride and Prejudice",
		Author:      "Jane Austen",
		Description: "A classic romance novel that explores themes of love, class, and societal expectations.",
		Stock:       12,
	},
	}

	for _, v := range data {
		db.Create(&v)
	}
}
