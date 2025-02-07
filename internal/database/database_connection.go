package database

import (
	"database/sql"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/sohWenMing/finance/internal/loadenv"
)

type dbEnvVars struct {
	host     string
	user     string
	password string
	dbname   string
	port     int
}

func InitDB(envPath string) (db *sql.DB, err error) {
	// envVars, err := getDBEnvVars(envPath)
	// if err != nil {
	// 	return nil, err
	// }
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	envVars.host, envVars.port, envVars.user,
	// 	envVars.password, envVars.dbname,
	// )

	db, dbErr := sql.Open("postgres", os.Getenv("DB_STRING"))
	if dbErr != nil {
		return nil, dbErr
	}
	return db, nil
}

func getDBEnvVars(envPath string) (envVarsStruct dbEnvVars, err error) {
	loadEnvErr := loadenv.LoadEnv(envPath)
	if loadEnvErr != nil {
		return envVarsStruct, loadEnvErr
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return envVarsStruct, err
	}
	loadedVars := dbEnvVars{
		host:     os.Getenv("HOST"),
		user:     os.Getenv("USER"),
		password: os.Getenv("PASSWORD"),
		dbname:   os.Getenv("DB_NAME"),
		port:     port,
	}

	return loadedVars, nil
}
