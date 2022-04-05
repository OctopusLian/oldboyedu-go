package main

// 手动引个mysql的驱动包
import (
	"fmt"
	"log"
	"math/rand"
	"oldboyedu-go/my-micro/demo/src/share/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 建表语句
// 因为是id，所以是自增无符号的int
var schema = `
CREATE TABLE IF NOT EXISTS user (
    id INT UNSIGNED AUTO_INCREMENT,
    name VARCHAR(20),
 	address VARCHAR(20),
    phone VARCHAR(15),
	PRIMARY KEY (id)
)`

// 对应表的结构体
type User struct {
	Id      int32  `db:"id"`
	Name    string `db:"name"`
	Address string `db:"address"`
	Phone   string `db:"phone"`
}

func main() {
	// 打开并连接数据库，返回错误信息
	// mysql是驱动名称，上面需要引入驱动的包
	// 第二个参数是指定连接到哪个数据库，在 share/config中配置
	db, err := sqlx.Connect("mysql", config.MysqlDSN)
	if err != nil {
		log.Fatalln(err)
	}
	// 执行建表语句
	db.MustExec(schema)
	// 开启事务
	tx := db.MustBegin()
	// 设置随机数种子，可以保证每次随机都是随机的
	rand.Seed(time.Now().UnixNano())
	// 事务执行SQL插入
	// 创建 GetRandomString() 方法，按指定个数生成随机字符
	tx.MustExec("INSERT INTO user (id, name, address,phone) VALUES (?, ?, ? ,?)",
		nil, GetRandomString(10), "beijing "+GetRandomString(10),
		"1591"+GetRandomString(7))
	// 提交事务
	err = tx.Commit()
	if err != nil {
		// 回滚
		_ = tx.Rollback()
	}
	fmt.Println("执行完毕!")
}

// 按指定个数生成随机数
func GetRandomString(leng int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	// 用系统时间生成随机数
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 按指定个数添加到返回结果中
	for i := 0; i < leng; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
