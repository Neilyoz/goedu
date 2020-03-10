package model

import (
	"database/sql/driver"
	"fmt"
	"goedu/system"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库对象
var DB *gorm.DB

func init() {
	// 加载配置文件
	connectStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		system.GetConfigure().Mysql.User,
		system.GetConfigure().Mysql.Pass,
		system.GetConfigure().Mysql.Host,
		system.GetConfigure().Mysql.Port,
		system.GetConfigure().Mysql.Database,
		system.GetConfigure().Mysql.Charset)

	db, err := gorm.Open("mysql", connectStr)
	if err != nil {
		panic(err)
	}

	DB = db
	db.AutoMigrate(&User{}, &Article{})
}

// GetDB 返回数据库对象
func GetDB() *gorm.DB {
	if system.GetConfigure().Debug == true {
		return DB.Debug()
	} else {
		return DB
	}
}

type MyTime struct {
	time.Time
}

func (t MyTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t MyTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *MyTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = MyTime{Time: value}
		return nil
	}

	return fmt.Errorf("can not convert %v to timestamp", v)
}
