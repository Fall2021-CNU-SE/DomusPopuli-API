package repository

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"

    . "github.com/NamSoGong/DomusPopuli-API/domain/db"
)

var db *gorm.DB = nil

func getDB() (*gorm.DB, error) {
    if db == nil {
        var err error

        // Get connection
        dsn := "host=localhost user=hrk password=1q2w3e4r dbname=domus port=5432 sslmode=disable TimeZone=Asia/Seoul"
        db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err != nil { return nil, err }

        // Clear Table
        err = db.Migrator().DropTable(&User_t{})
        if err != nil { return nil, err }
        err = db.Migrator().DropTable(&House_t{})
        if err != nil { return nil, err }

        // Init Table
        err = db.AutoMigrate(&User_t{})
        if err != nil { return nil, err }
        err = db.AutoMigrate(&House_t{})
        if err != nil { return nil, err }
    }

    return db, nil
}
