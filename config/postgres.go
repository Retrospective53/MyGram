package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresConfig struct {
	Port              string
	Host              string
	Username          string
	Password          string
	DBName            string
	MaxIdleConnection int
	MaxOpenConnection int
	MaxIdleTime       int
}

func loadPostgresConfig() PostgresConfig {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}

	var (
		host              = os.Getenv("pgHost")
		pgport            = os.Getenv("pgPort")
		username          = os.Getenv("pgUsername")
		password          = os.Getenv("pgPassword")
		dbname            = os.Getenv("pgDBName")
		maxIdleConnection = 7
		maxOpenConnection = 5
		maxIdleTime       = int(30 * time.Minute)
	)

	return PostgresConfig{
		Host:              host,
		Port:              pgport,
		Username:          username,
		Password:          password,
		DBName:            dbname,
		MaxIdleConnection: maxIdleConnection,
		MaxOpenConnection: maxOpenConnection,
		MaxIdleTime:       maxIdleTime,
	}
}


func NewPostgresConfig() (db *sql.DB) {
	pgConfig := loadPostgresConfig()

	connString := fmt.Sprintf(`
		host=%v
		port=%v
		user=%v
		password=%v
		dbname=%v
		sslmode=disable
	`,
		pgConfig.Host,
		pgConfig.Port,
		pgConfig.Username,
		pgConfig.Password,
		pgConfig.DBName,
	)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(pgConfig.MaxOpenConnection)
	db.SetMaxIdleConns(pgConfig.MaxIdleConnection)
	db.SetConnMaxIdleTime(time.Duration(pgConfig.MaxIdleConnection))

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}

func NewPostgresGormConn() (db *gorm.DB) {
	// connect to db
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(postgresDSN()), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	dbSQL, err := db.DB()
	if err != nil {
		panic(err)
	}
	postgresPoolConf(dbSQL)

	if err := dbSQL.Ping(); err != nil {
		panic(err)
	}
	log.Println("successfully connect to Postgres")
	return db
}

func postgresDSN() string {
	pgConfig := loadPostgresConfig()
	
	return fmt.Sprintf(`
	host=%v
	port=%v
	user=%v
	password=%v
	dbname=%v
	sslmode=disable
`,
	pgConfig.Host,
	pgConfig.Port,
	pgConfig.Username,
	pgConfig.Password,
	pgConfig.DBName,
)
}

func postgresPoolConf(dbSQL *sql.DB) {
	pgConfig := loadPostgresConfig()
	
	// set extended config
	dbSQL.SetMaxIdleConns(pgConfig.MaxIdleConnection)
	dbSQL.SetMaxOpenConns(pgConfig.MaxOpenConnection)
	dbSQL.SetConnMaxIdleTime(time.Duration(pgConfig.MaxIdleTime))
}
