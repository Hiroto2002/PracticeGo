package infrastructure

import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "practiceGo/app/src/domain"
)

type DB struct {
    Host string
    Username string
    Password string
    DBName string
    Port string
    Connection *gorm.DB
}

func NewDB()*DB{
    c := NewConfig()
    return newDB(&DB{
        Host: c.DB.Production.Host,
        Username: c.DB.Production.Username,
        Password: c.DB.Production.Password,
        DBName: c.DB.Production.DBName,
        Port: c.DB.Production.Port,
    })
}

// func NewTestDB() *DB {
//     c := NewConfig()
//     return newDB(&DB{
//         Host: c.DB.Test.Host,
//         Username: c.DB.Test.Username,
//         Password: c.DB.Test.Password,
//         DBName: c.DB.Test.DBName,
//     })
// }

func newDB(d *DB) *DB {
    dsn := "host=" + d.Host + " user=" + d.Username + " password=" + d.Password + " dbname=" + d.DBName + " port=" + d.Port + " sslmode=disable TimeZone=Asia/Tokyo"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }
    d.Connection = db
    return d
}

// Begin begins a transaction
func (db *DB) Begin() *gorm.DB {
    return db.Connection.Begin()
}

func (db *DB) Connect() *gorm.DB {
    return db.Connection
}

func (db *DB) AutoMigrate() {
    db.Connection.AutoMigrate(&domain.Todos{}) // Todoモデルに基づくマイグレーション
}

func (db *DB) ResetDatabase() {
    // テーブルを削除
    db.Connection.Migrator().DropTable(&domain.Todos{})
    // 必要に応じて再マイグレーション
    db.AutoMigrate()
}

func (db *DB) InsertSeedData() {
    // シードデータの挿入例
    db.Connection.Create(&domain.Todos{Title: "Seed Data", Description: "This is a seed data.", IsCompleted: false})
}


