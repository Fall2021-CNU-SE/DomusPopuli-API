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
        if db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
            return nil, err
        }

        // Reset DB
        if err = resetDB(db, &User_t{}, &House_t{}); err != nil {
            return nil, err
        }
    }

    return db, nil
}

func resetDB(db *gorm.DB, schemes ...interface{}) error {
    for _, scheme := range schemes {
        if err := db.Migrator().DropTable(scheme); err != nil {
            return err
        }
        if err := db.AutoMigrate(scheme); err != nil {
            return err
        }
    }
    return nil
}
