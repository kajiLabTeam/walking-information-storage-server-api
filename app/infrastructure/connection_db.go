package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectionDB() (*sql.DB, error) {

	//環境変数の読み込み
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	// fmt.Println(host)
	// fmt.Println(port)
	// fmt.Println(name)
	// fmt.Println(user)
	fmt.Println(password)

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s", user, password, name, host, port)
	fmt.Println(connStr)

	//データベースと接続
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	//defer:処理が完了次に実行
	defer db.Close()

	//ピンを飛ばす
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("DB接続")
	return db, nil

}
