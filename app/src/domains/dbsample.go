package domains

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/t-akzw/go_sandbox/models"
)

func main() {
	fmt.Println("hogehoge")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	db, err := gorm.Open("postgres", "host="+host+" user="+user+" password="+password+" dbname="+dbname+" sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// スキーマのマイグレーション
	var model models.Model
	db.AutoMigrate(model)

	// Create
	db.Create(model{Code: "L1212", Price: 1000})

	db.First(&model, 1)
	fmt.Printf("model %#v\n", model.id)
}
