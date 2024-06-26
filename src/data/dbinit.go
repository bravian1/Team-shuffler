package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := "root:12345678@tcp(34.44.36.143:3306)/foosball_league?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// func ConnectDatabase() (*sql.DB, error) {
// 	dbUser := "root"
// 	dbPass := "12345678"
// 	dbInstance := "protean-tome-427511-p6:us-central1:team-shuffler"
// 	dbPort := "3306"
// 	dbName := "foosball_league"

// 	// Always use the TCP connection
// 	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbInstance, dbPort, dbName)

// 	db, err := sql.Open("mysql", dataSourceName)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to open database connection: %w", err)
// 	}

// 	// Test connection
// 	err = db.Ping()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to ping database: %w", err)
// 	}

// 	return db, nil
// }
