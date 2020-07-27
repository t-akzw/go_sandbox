package lib
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)
func NewDbConnection() *gorm.DB {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	db, err := gorm.Open("postgres", "host="+host+" user="+user+" password="+password+" dbname="+dbname+" sslmode=disable timezone=Asia/Tokyo")
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	db.DB().SetMaxOpenConns(50)
	db.DB().SetMaxIdleConns(10)
	db.LogMode(true)
	return db
}





