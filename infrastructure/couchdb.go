package infrastructure

import (
	"log"
	"os"

	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/joho/godotenv"
	"github.com/timjacobi/go-couchdb"
)

var client *couchdb.Client
var dbName = "mydb"

func GetDB() *couchdb.DB {
	return client.DB(dbName)
}
func InitDB() error {
	//When running locally, get credentials from .env file.
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file does not exist")
	}
	cloudantUrl := os.Getenv("CLOUDANT_URL")

	appEnv, _ := cfenv.Current()
	if appEnv != nil {
		cloudantService, _ := appEnv.Services.WithLabel("cloudantNoSQLDB")
		if len(cloudantService) > 0 {
			cloudantUrl = cloudantService[0].Credentials["url"].(string)
		}
	}

	client, err = couchdb.NewClient(cloudantUrl, nil)
	if err != nil {
		log.Println("Can not connect to Cloudant database")
	}
	_, err = client.CreateDB(dbName)
	if err != nil {
		log.Print(err)
		return nil
	}
	return nil
}
