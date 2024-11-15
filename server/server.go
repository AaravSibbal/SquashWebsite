package server

import (
	"context"
	"fmt"
	"database/sql"
	// "fmt"
	"log"
	"net/http"
	"os"
	"time"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

type application struct {
	ctx            *context.Context
	db             *sql.DB
	infoLog        *log.Logger
	errorLog       *log.Logger
}

func Run() {
	envFile, err := godotenv.Read(".env")
	if err != nil {
		log.Fatalf("Couldn't read the .env file %v", err)
	}

	db, err := getDBConnection(envFile)
	if err != nil {
		log.Fatalf("could not connect to the DB: %v", err)
	}

	infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	errLogger := log.New(os.Stdout, "ERROR:\t", log.Ltime|log.Ldate|log.Lshortfile)

	ctx := context.Background()

	app := &application{
		ctx: &ctx,
		db: db,
		infoLog: infoLog,
		errorLog: errLogger,

	}

	addr := fmt.Sprintf("%s:%s",envFile["ADDRESS"], envFile["PORT"])

	srv := &http.Server{
		Addr:         addr,
		ErrorLog:     errLogger,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      app.Routes(),
	}

	err = srv.ListenAndServe()

	log.Fatal(err)
	
}

func getDBConnection(envFile map[string]string) (*sql.DB, error){
	host := envFile["POSTGRES_HOST"]
	postPort := envFile["POSTGRES_PORT"]
	user := envFile["POSTGRES_USER"]
	password := envFile["POSTGRES_PASSWORD"]
	dbName := envFile["POSTGRES_NAME"]

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, postPort, user, password, dbName)

	fmt.Println(psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalln("there was an error getting the db connection", err)
		return nil, err
	}


	if err = db.Ping(); err != nil {
		log.Fatalln("we couldn't ping the db for some reason", err)
	}

	fmt.Println("db was connected successfuly")

	return db, nil
}
