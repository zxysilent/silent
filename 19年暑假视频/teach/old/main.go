package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func first(w http.ResponseWriter, r *http.Request) {
	// tcp http格式的数据 ==> http
	// 把这个数据 解析 并放到 form
	r.ParseForm()
	// r.Form//get post
	// name := r.Form["name"][0]
	// name := r.Form.Get("name")
	name1 := r.FormValue("name")
	pass := r.Form.Get("pass")
	log.Println(name1)
	w.Write([]byte(name1 + "=====>" + pass))
}

func index(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./views/index.html")
	defer f.Close()
	// 直接复制
	io.Copy(w, f)
	// f.Close()
	// 读取所有
	// buf, _ := ioutil.ReadAll(f)
	// w.Write(buf)
	// 准备容器
	// buf := make([]byte, 1024*2)
	// ln, _ := f.Read(buf)
	// w.Write(buf[:ln])
}

func indexJS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	f, _ := os.Open("./1.js")
	defer f.Close()
	// 直接复制
	io.Copy(w, f)
	// f.Close()
	// 读取所有
	// buf, _ := ioutil.ReadAll(f)
	// w.Write(buf)
	// 准备容器
	// buf := make([]byte, 1024*2)
	// ln, _ := f.Read(buf)
	// w.Write(buf[:ln])
}
func main() {
	//静态文件访问 css img js 等
	http.Handle("/res/", http.StripPrefix("/res/", http.FileServer(http.Dir("./res"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/index.js", indexJS)
	http.HandleFunc(`/first`, first)
	http.ListenAndServe(":8080", nil)
}
