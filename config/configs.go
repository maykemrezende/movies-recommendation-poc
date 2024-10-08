package config

import (
	"github.com/go-chi/jwtauth"
	// "github.com/maykemrezende/expenses-manager-golang/internal/models"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type conf struct {
	DBDriver     string `mapstructure:"dbDriver" env:"DB_DRIVER" env-default:"mysql"`
	DBHost       string `mapstructure:"dbHost" env:"DB_HOST" env-default:"localhost"`
	DBPort       string `mapstructure:"dbPort" env:"DB_PORT" env-default:"3306"`
	DBUser       string `mapstructure:"dbUser" env:"DB_USER" env-default:"root"`
	DBPassword   string `mapstructure:"dbPassword" env:"DB_PASSWORD" env-default:"root"`
	DBName       string `mapstructure:"dbName" env:"DB_NAME" env-default:"expenses_manager"`
	JwtSecret    string `mapstructure:"jwtSecret" env:"JWT_SECRET" env-default:"secret"`
	JwtExpiresIn int    `mapstructure:"jwtExpiresIn" env:"JWT_EXPIRESIN" env-default:"3600"`
	RootUrl      string `mapstructure:"rootUrl" env:"ROOT_URL" env-default:"http://localhost:8080"`
	ApiPort      string `mapstructure:"apiPort" env:"API_PORT" env-default:"8080"`
	TokenAuth    *jwtauth.JWTAuth
}

type application struct {
	Db      *gorm.DB
	RootUrl string
	ApiPort string
}

func LoadApplication() (*application, error) {
	config, err := loadConfig()
	if err != nil {
		return nil, err
	}

	// db, err := initDatabase(config)
	// if err != nil {
	// 	return nil, err
	// }

	application := &application{
		// Db:      db,
		RootUrl: config.RootUrl,
		ApiPort: config.ApiPort,
	}

	return application, nil
}

func loadConfig() (*conf, error) {
	var cfg *conf

	viper.AddConfigPath("./config")
	viper.SetConfigName("app_config")
	viper.SetConfigType("json")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)

	if err != nil {
		return nil, err
	}

	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JwtSecret), nil)

	return cfg, nil
}

// func initDatabase(config *conf) (*gorm.DB, error) {
// 	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
// 	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

// 	if err != nil {
// 		return nil, err
// 	}
// 	err = db.AutoMigrate(&models.User{}, &models.Bill{}, &models.Tag{})

// 	if err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }
