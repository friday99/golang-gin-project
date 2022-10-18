package setting

import (
	"fmt"
	"github.com/go-ini/ini"
)

type DatabaseSetting struct {
	Type     string
	User     string
	Password string
	Host     string
	Name     string
}

var DatabaseSettingS = &DatabaseSetting{}

// New initialize the configuration instance
func New() error {
	// read the configuration file
	err := ini.MapTo(DatabaseSettingS, "config/database.ini")
	fmt.Println("setting", DatabaseSettingS)
	if err != nil {
		return err
	}
	return nil
}
