package db

// gorm db connection
import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/oranzh/cls/pkg/setting"
	"log"
	"time"
)

var db *gorm.DB

type Model struct {
	ID        int       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Setup() {
	var err error
	db, err = gorm.Open(setting.DatabaseSettingS.Type,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			setting.DatabaseSettingS.User,
			setting.DatabaseSettingS.Password,
			setting.DatabaseSettingS.Host,
			setting.DatabaseSettingS.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}

func GetDB() *gorm.DB {
	return db
}
