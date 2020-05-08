package AppInit

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "log"
)
var db *gorm.DB

func init()  {
    var err error
    db, err = gorm.Open(
        "mysql",
        "root:casa123@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local",
        )
    if err != nil {
        log.Fatal(err)
    }

    db.DB().SetMaxIdleConns(10)
    db.DB().SetMaxOpenConns(50)
}

func GetDb() *gorm.DB  {
    return db
}
