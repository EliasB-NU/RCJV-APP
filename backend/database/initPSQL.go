package database

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	ID uint64 `gorm:"primaryKey;autoIncrement"`

	Username string
	Email    string `gorm:"unique"`
	Password string

	Tokens *[]BrowserToken
}

type BrowserToken struct {
	gorm.Model
	ID uint64 `gorm:"primaryKey"`

	DeviceId string
	Token    string `gorm:"unique"`
	UserID   uint64 `gorm:"index"`
	User     User
}

type Institution struct {
	gorm.Model
	ID uint64 `gorm:"primaryKey"`

	Name string

	Teams []Team
}

type League struct {
	gorm.Model
	ID uint64 `gorm:"primaryKey"`

	Name string

	Teams []Team
}

type Team struct {
	gorm.Model
	ID uint64 `gorm:"primaryKey"`

	Name string

	InstitutionID uint64 `gorm:"index"`
	Institution   Institution

	LeagueID uint64 `gorm:"index"`
	League   League
}

// InitPSQLDatabase Creates all the necessary tables for the app to work
func InitPSQLDatabase(db *gorm.DB) error {
	var err error

	err = db.AutoMigrate(&User{})
	if err != nil {
		return errors.New("failed to auto migrate users table: " + err.Error())
	}

	err = db.AutoMigrate(&BrowserToken{})
	if err != nil {
		return errors.New("failed to auto migrate browser token table: " + err.Error())
	}

	err = db.AutoMigrate(&Institution{})
	if err != nil {
		return errors.New("failed to auto migrate institution table: " + err.Error())
	}

	err = db.AutoMigrate(&League{})
	if err != nil {
		return errors.New("failed to auto migrate league table: " + err.Error())
	}

	err = db.AutoMigrate(&Team{})
	if err != nil {
		return errors.New("failed to auto migrate team table: " + err.Error())
	}

	// Create initial admin user, if not exists (email: admin@example.com, username: admin, password: admin)
	result := db.Where("username = ?", "admin").First(&User{})
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		var user User
		user.Username = "admin"
		user.Password, _ = hashPassword("admin")
		user.Email = "admin@example.com"
		err = db.Create(&user).Error
		if err != nil {
			return errors.New("failed to create initial admin user: " + err.Error())
		}
	} else if result.Error != nil {
		return errors.New("failed to check for initial admin user: " + result.Error.Error())
	} else {
		log.Println("Initial admin user already exists")
	}

	log.Println("Database initialized successfully")
	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
