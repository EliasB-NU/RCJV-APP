package database

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	gorm.Model
	ID uint64 `gorm:"primaryKey;autoIncrement" json:"id"`

	Username string `json:"username"`
	Email    string `gorm:"unique" json:"email"`
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

type Leagues struct {
	gorm.Model
	ID uint64 `gorm:"primaryKey"`

	RescueLineEntry bool `json:"rescueLineEntry"`
	RescueLine      bool `json:"rescueLine"`
	RescueMazeEntry bool `json:"rescueMazeEntry"`
	RescueMaze      bool `json:"rescueMaze"`

	SoccerEntry            bool `json:"soccerEntry"`
	SoccerLightWeightEntry bool `json:"soccerLightWeightEntry"`
	SoccerLightWeight      bool `json:"soccerLightWeight"`
	SoccerOpen             bool `json:"soccerOpen"`

	OnStageEntry bool `json:"onStageEntry"`
	OnStage      bool `json:"onStage"`
}

type Config struct {
	gorm.Model
	ID uint64 `gorm:"primaryKey"`

	AppEnabled bool   `json:"appEnabled"`
	EventName  string `json:"eventName"`

	SoccerURL          string `json:"soccerUrl"`
	SoccerAbbreviation string `json:"soccerAbbreviation"`

	RescueURL string `json:"rescueUrl"`
}

type Institution struct {
	gorm.Model
	ID uint64 `gorm:"primaryKey"`

	Name string `json:"name"`

	Matches []Match

	Teams []Team
}

type Team struct {
	gorm.Model
	ID uint64 `gorm:"primaryKey"`

	Name   string
	League string

	Matches []Match

	InstitutionID uint64 `gorm:"index"`
	Institution   Institution
}

type Field struct {
	gorm.Model
	ID uint64 `gorm:"primaryKey" json:"id"`

	Name   string `json:"name"`
	League string `json:"league"`
}

type Match struct {
	gorm.Model
	ID uint64 `gorm:"primaryKey" json:"id"`

	Name      string
	StartTime time.Time
	Duration  time.Duration
	Field     string

	League string

	Institution   Institution
	InstitutionID uint64 `gorm:"index"`

	Team   Team
	TeamID uint64 `gorm:"index"`
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

	err = db.AutoMigrate(&Leagues{})
	if err != nil {
		return errors.New("failed to auto migrate leagues table: " + err.Error())
	}

	err = db.AutoMigrate(&Config{})
	if err != nil {
		return errors.New("failed to auto migrate configs table: " + err.Error())
	}

	err = db.AutoMigrate(&Institution{})
	if err != nil {
		return errors.New("failed to auto migrate institution table: " + err.Error())
	}

	err = db.AutoMigrate(&Team{})
	if err != nil {
		return errors.New("failed to auto migrate team table: " + err.Error())
	}

	err = db.AutoMigrate(&Field{})
	if err != nil {
		return errors.New("failed to auto migrate field table: " + err.Error())
	}

	// I am lazy, so all game tables are initiated at once here
	err = db.AutoMigrate(&Match{})
	if err != nil {
		return errors.New("failed to auto migrate matches tables: " + err.Error())
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

	// Create initial league entry
	// Check if league entry already exists
	result = db.First(&Leagues{})
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// Create initial league
		var leagues = Leagues{
			RescueLineEntry:        false,
			RescueLine:             false,
			RescueMazeEntry:        false,
			RescueMaze:             false,
			SoccerEntry:            false,
			SoccerLightWeightEntry: false,
			SoccerLightWeight:      false,
			SoccerOpen:             false,
			OnStageEntry:           false,
			OnStage:                false,
		}
		err = db.Create(&leagues).Error
		if err != nil {
			return errors.New("failed to create leagues: " + err.Error())
		}
	} else if result.Error != nil {
		return errors.New("failed to check for leagues: " + result.Error.Error())
	} else {
		log.Println("Leagues already exists")
	}

	// Create config entry
	result = db.First(&Config{})
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		var config = Config{
			AppEnabled:         false,
			EventName:          "RCJ - Test",
			SoccerURL:          "https://catigoal.com",
			SoccerAbbreviation: "VO25",
			RescueURL:          "https://google.com",
		}
		err = db.Create(&config).Error
		if err != nil {
			return errors.New("failed to create config: " + err.Error())
		}
	} else if result.Error != nil {
		return errors.New("failed to check for config: " + result.Error.Error())
	} else {
		log.Println("Config already exists")
	}

	log.Println("Database initialized successfully")
	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
