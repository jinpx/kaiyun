package tidb

import (
	"database/sql"
	"fmt"
	"log"
)

type Database struct {
	db *sql.DB
}

func initConfig() {

	// MySQL 数据库连接信息
	// mysql -h127.0.0.1 -umariadb -pD8Sb3V6789
	dsn := "timog:D8Sb3V6789@tcp(192.168.60.149:3306)/ebingo_test"

	// 打开数据库连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// 尝试连接数据库
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database!")
}

func (ti *Database) BingoQuery(db *sql.DB, index int) {

	// 在此处可以执行数据库查询、插入等操作
	rows, err := ti.db.Query("SELECT vid, gmtype, video_url FROM bingo_video")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	// 处理查询结果
	for rows.Next() {
		var vid string
		var gmtype string
		var video_url string

		err := rows.Scan(&vid, &gmtype, &video_url)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("vid: %s, gmtype: %s, video_url: %s, index: %d \n",
			vid, gmtype, video_url, index)
	}

	// 检查遍历过程中是否有错误
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
