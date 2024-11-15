package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"

	bot "github.com/AaravSibbal/SquashWebsite/Bot"
	"github.com/AaravSibbal/SquashWebsite/server"
	"github.com/joho/godotenv"
)

func main() {
	wg := &sync.WaitGroup{}

	envFile, err := godotenv.Read(".env")
	if err != nil {
		
		fmt.Printf("%v", err)
		return	
	}	
	
	db, err := getDbConnection(envFile)
	if err != nil {
		fmt.Printf("%v",err)
		return
	}
	defer db.Close()
	
	
	botToken := envFile["BOT_TOKEN"]
	ctx := context.Background()
	botStruct := &bot.Bot{
		BotToken: botToken,
		Db: db,
		Ctx: &ctx,
		Wg: wg,
		
	}
	fmt.Println(envFile["BOT_TOKEN"])
	
	wg.Add(2)
	go botStruct.Run()	
	go server.Run()

	wg.Wait()
}

func getDbConnection(envFile map[string]string) (*sql.DB, error){

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
