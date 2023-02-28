package config

import (
<<<<<<< HEAD
	"fmt"
=======
<<<<<<< HEAD
	"log"
>>>>>>> development
	"os"
_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

<<<<<<< HEAD
=======
var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

=======
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

>>>>>>> development
type Config struct {
	DataSourceName string
}

func (c *Config) initDb() {
	err := godotenv.Load()
<<<<<<< HEAD
	if err != nil {
		panic(err.Error())
	}
	// tampung nilai ENV dari terminal
=======
	if err != nil{
		panic(err.Error())
	}// 


	// Tampung nilai ENV dari terminal
>>>>>>> development
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
<<<<<<< HEAD
=======
	//dbDriver := os.Getenv("DB_DRIVER")
>>>>>>> development
	dbName := os.Getenv("DB_NAME")

	// data source name
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
<<<<<<< HEAD
=======

>>>>>>> development
	c.DataSourceName = dsn
}

func NewConfig() Config {
	config := Config{}
<<<<<<< HEAD
	config.initDb() // untuk menjalankan method yang didalamnya membuka koneksi ke DB
	return config   // Mengirimkan Object Config yang didalamnya terdapat attribute bertipe data Koneksi, namun koneksi tidak dapat diakses secara
	// langsung karena attribute bertipe private
=======
	config.initDb() // untuk menjalankan method yang didalmnya membuka koneksi ke DB
	return config   // Mengirimkan Object Config yang dalamnya terdapat attribute bertipe data koneksi, namun koneksi tidak dapat diakses secara
	// langsung karena attribute bertipe private
>>>>>>> 70dbb88e8dcdea7190f7104e570f79f11b616a32
>>>>>>> development
}
