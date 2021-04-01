package main

import (
	"database/sql"
	"time"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

)

type User struct {
	ID int64 `db:"id"`
	Name sql.NullString  `db:"name"`  //由于在mysql的users表中name没有设置为NOT NULL,所以name可能为null,在查询过程中会返回nil，如果是string类型则无法接收nil,但sql.NullString则可以接收nil值
	Age int `db:"age"`
}

const(
	host = "localhost"
	username="root"
	pwd="root"
	port=3306
	network="tcp"
	database="test"
)
func main(){

	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",username,pwd,network,host,port,database)
	fmt.Println(dsn)
	DB,err := sql.Open("mysql",dsn)

	fmt.Println(err)
	if err != nil {
		fmt.Println("open mysql filed,error:",err);
		return
	}

	DB.SetConnMaxLifetime(100*time.Second)  //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(100)//设置最大连接数
	DB.SetMaxIdleConns(16) //设置闲置连接数
	user := queryOne(DB)
	fmt.Println(*user)
}

func queryOne(DB *sql.DB) (*User){
	user := new(User)
	row := DB.QueryRow("select * from users where id=?",1)
	//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
	if err := row.Scan(&user.ID,&user.Name, &user.Age); err != nil{
		fmt.Printf("scan failed, err:%v",err)
	}
	return user
}


