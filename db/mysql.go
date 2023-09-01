package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
	"github.com/menghua6/token-net/entity"
)

var mysqlDb *gorm.DB

func InitMysql() error {
	url := "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local"
	constr := fmt.Sprintf(url, MysqlUserName, MysqlPassword, MysqlHost, MysqlPort, MysqlDbName)
	//连接数据库
	db, err := gorm.Open(mysql.Open(constr), &gorm.Config{})
	if err != nil {
		return err
	}
	log.Println("MysqlDb初始化成功")

	db.AutoMigrate(entity.Token{})
	db.AutoMigrate(entity.Message{})

	mysqlDb = db

	return nil
}

func CreateToken(token string, expirationTime time.Time, count int) error {
	tokenInfo := entity.Token{
		Token:          token,
		ExpirationTime: expirationTime,
		Count:          count,
	}
	dbRes := mysqlDb.Create(&tokenInfo)
	if dbRes.Error != nil {
		return dbRes.Error
	}
	return nil
}

func GetToken(token string) (*entity.Token, error) {
	var tokenInfo entity.Token
	dbRes := mysqlDb.Where("token = ?", token).Last(&tokenInfo)
	if dbRes.Error != nil {
		return nil, dbRes.Error
	}
	return &tokenInfo, nil
}

func UpdateToken(token *entity.Token) error {
	values := map[string]interface{}{
		"count": token.Count,
	}
	dbRes := mysqlDb.Model(&entity.Token{}).Where("token = ?", token.Token).Updates(values)
	if dbRes.Error != nil {
		return dbRes.Error
	}
	return nil
}

func GetMessages(token string) ([]entity.Message, error) {
	messages := make([]entity.Message, 0)
	dbRes := mysqlDb.Where("token = ?", token).Find(&messages)
	if dbRes.Error != nil {
		return nil, dbRes.Error
	}
	return messages, nil
}

func CreateMessage(token string, message string) error {
	messageInfo := entity.Message{
		Token:   token,
		Message: message,
	}
	dbRes := mysqlDb.Create(&messageInfo)
	if dbRes.Error != nil {
		return dbRes.Error
	}
	return nil
}
