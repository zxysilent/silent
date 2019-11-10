package control

import (
	"io"
	"net/http"
	"os"
)

// ListView 列表页
func ListView(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./views/list.html")
	io.Copy(w, f)
	f.Close()
}

// EditView 列表页
func EditView(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./views/edit.html")
	io.Copy(w, f)
	f.Close()
}

// DetailView 详细页
func DetailView(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./views/detail.html")
	io.Copy(w, f)
	f.Close()
}

// IndexView 主页面
func IndexView(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./views/index.html")
	io.Copy(w, f)
	f.Close()
}

// ViewArticleAdd 文章添加页面
func ViewArticleAdd(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./views/add.html")
	io.Copy(w, f)
	f.Close()
}
