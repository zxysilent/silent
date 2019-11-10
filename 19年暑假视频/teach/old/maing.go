package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql" // init
	"github.com/jmoiron/sqlx"          //sql x
)

type Class struct {
	Id   int64
	Name string
	Desc string
}

func view(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./views/index.html")
	io.Copy(w, f)
	f.Close()
}

func data(w http.ResponseWriter, r *http.Request) {
	db, err := sqlx.Open(`mysql`, `root:root@tcp(127.0.0.1:3306)/news?charset=utf8&parseTime=true`)
	log.Println(db, err)
	// 查询
	// Get 查询单个(count())/一条  struct
	// Select 一个集合 []
	// 非查询
	// db.Exec()//执行 insert update delete
	r.ParseForm()
	idStr := r.Form.Get("id")
	log.Println("id===> ", idStr)
	// string convert  atoi 字符转 数字
	// int64
	id, _ := strconv.ParseInt(idStr, 10, 64)
	mod := Class{}
	log.Println("----", mod)
	db.Get(&mod, "select * from class where id = ?", id)
	log.Println("--------", mod)
	// 序列化  struct --> json 格式的字符串
	jsonbytes, _ := json.Marshal(mod)
	w.Write(jsonbytes)
}

func main() {
	// 路由
	// 访问前面的路径 / ----->执行后面的方法
	http.HandleFunc("/", view)
	http.HandleFunc("/data", data)
	// http.ListenAndServe(":8080", nil)
	db, err := sqlx.Open(`mysql`, `root:root@tcp(127.0.0.1:3306)/news?charset=utf8&parseTime=true`)
	log.Println(db, err)
	//db.Ping()//检测是否能够正常执行
	// ping
	// 查询
	// Get 查询单个(count())/一条  struct
	// Select 一个集合 []
	// 非查询
	// db.Exec()//执行 insert update delete
	// string convert  atoi 字符转 数字
	// int64
	mod := Class{}
	log.Println("----", mod)
	db.Get(&mod, "select * from class where id = ?", 1)
	mods := make([]Class, 0, 4)
	//  1 2 3 4 5 6
	// 2 ,2==> 3 4
	// pi,ps pageindex ,pagesize
	// 2,2// 数据量小的分页--->200万
	pi, ps := 2, 1
	db.Select(&mods, "select * from class  limit ?,?", (pi-1)*ps, ps)
	log.Println(mods)
	// 查询单个 或者一条数据
	// db.Get(1,2,3...)
	// 1 查询的结构放到
	// 2 sql语句
	// 3 给sql里面的参数 ?,?,? ,100,1000,10000
	// db.Get(ptr,"select * from class limit ?,?" ,2,4)

	// 查询数据集合 [] ,make([]T,0,预估大小)
	// db.Select(1,2,3...)
	// get、select 返回值 err

	// 非查询- 执行
	// 1 sql 语句
	// 2 参数 参数 ?,?,? ,100,1000,10000
	// db.Exec(1，2...)
	// 返回值
	// err 执行过程遇到的错误 语法？
	// result, err := db.Exec("update class set name = ? where id= ?", "update", 1)
	// result
	// LastInsertId() (int64, error) // insert pk-id 自增 每次插入一条数据 id ++
	// result, err := db.Exec("insert into class (name) values(?)", "insert12")
	// result, err := db.Exec("delete from class where id >3")
	result, err := db.Exec("update class set `desc`=?", "描述")
	log.Println("error ==> ", err)
	lastid, _ := result.LastInsertId()
	log.Println("lastid ==> ", lastid)
	// RowsAffected() (int64, error) // 返回 执行操作  操作了多少条数据
	rows, _ := result.RowsAffected()
	log.Println("rows ==> ", rows)

}
