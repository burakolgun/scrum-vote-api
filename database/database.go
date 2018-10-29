package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/go-yaml/yaml"

	_ "github.com/lib/pq"
)

type Config = struct {
	Postgres struct {
		PostgresUser     string `yaml:"POSTGRES_USER"`
		PostgresDB       string `yaml:"POSTGRES_DB"`
		PostgresPassword string `yaml:"POSTGRES_PASSWORD"`
		PostgresHost     string `yaml:"HOST"`
		PostgresPort     string `yaml:"PORT"`
	}
}

var Conf *Config

var DB *sql.DB

func Init() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	confContent, err := ioutil.ReadFile(dir + "/database/config.yml")
	if err != nil {
		panic(err)
	}

	// expand environment variables
	confContent = []byte(os.ExpandEnv(string(confContent)))
	Conf = &Config{}

	if err := yaml.Unmarshal(confContent, Conf); err != nil {
		panic(err)
	}

	log.Println("Database Configuration are Loaded")
}

func Connect() (*sql.DB, error) {

	time.Sleep(time.Second * 10)
	fmt.Println("sleep1")
	if DB != nil {
		return DB, nil
		fmt.Println("already connected")
	}

	fmt.Println("Connection Opening...")
	connStr := "user=" + Conf.Postgres.PostgresUser + " " +
		"dbname=" + Conf.Postgres.PostgresDB + " " +
		"password=" + Conf.Postgres.PostgresPassword + " " +
		"port=" + Conf.Postgres.PostgresPort + " " +
		"sslmode=disable"
	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	defer DB.Close()

	err = DB.Ping()

	if err != nil {
		return nil, err
	}

	fmt.Println("Connection Opened")

	return DB, nil
}
