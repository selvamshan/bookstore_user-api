package users_db

import(
	"os"
	"fmt"
	//"log"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"	
	"github.com/selvamshan/bookstore_utils-go/logger"
)


var (
	Client *sql.DB	

	username = os.Getenv("MYSQL_USERNAME")
	password = os.Getenv("MYSQL_PASSWORD")
	host = os.Getenv("MYSQL_HOST")	
	schema = os.Getenv("MYSQL_SCHEMA")	
)



func init() {	
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, schema)
	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	
	mysql.SetLogger(logger.GetLogger())
	//log.Println("database successfully configured")
	logger.Info("database successfully configured")
}