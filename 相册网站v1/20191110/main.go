package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func IndexView(w http.ResponseWriter, r *http.Request) {
	html := loadHtml("./views/index.html")
	w.Write(html)
}

// Upload 上传
func Upload(w http.ResponseWriter, r *http.Request) {
	//GET
	// fmt.Println(r.Method)
	if r.Method == "GET" {
		w.Write(loadHtml("./views/upload.html"))
		return
	}
	//POST
	if r.Method == "POST" {
		f, h, err := r.FormFile("file")
		if err != nil {
			//有错误
			w.Write([]byte("文件上传有误 :" + err.Error()))
			return
		}
		t := h.Header.Get("Content-Type")
		log.Println(t)
		if !strings.Contains(t, "image") {
			io.WriteString(w, "<html> <a href=\"/upload\">请上传图片</a><html>")
			return
		}
		os.Mkdir("./static", 0666) //以 main.go 作为相对路径
		out, err := os.Create("./static/" + h.Filename)
		if err != nil {
			io.WriteString(w, "文件创建失败："+err.Error())
			return
		}
		_, err = io.Copy(out, f)
		if err != nil {
			io.WriteString(w, "文件保存失败："+err.Error())
			return
		}
		//io.WriteString(w, "/static/"+h.Filename) //web服务器 / 以 main.go 作为绝对路径的开始
		http.Redirect(w, r, "/detail?name="+h.Filename, 302)
		return
	}
}

// ImageView 返回指定的图片
func ImageView(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //把url 或者 form 表单的数据解析到对应的容器 r.From
	name := r.Form.Get("name")
	// fmt.Println(name)
	f, err := os.Open("./static/" + name)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	defer f.Close()
	w.Header().Set("Content-Type", "image")
	io.Copy(w, f)
}

//
func DetailView(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	html := loadHtml("./views/detail.html")
	html = bytes.Replace(html, []byte("@src"), []byte("/image?name="+name), 1)
	w.Write(html)
}

func ListView(w http.ResponseWriter, r *http.Request) {
	html := loadHtml("./views/list.html")
	names, err := ioutil.ReadDir("./static")
	if err != nil {
		w.Write([]byte("errors"))
		return
	}
	temp := ""
	for i := 0; i < len(names); i++ {
		temp += `<li><a href="/detail?name=` + names[i].Name() + `"><img src="/image?name=` + names[i].Name() + `" alt="未发现"></a></li>`
	}
	html = bytes.Replace(html, []byte("@html"), []byte(temp), 1)
	w.Write(html)
}

// loadHtml 通用加载html
func loadHtml(name string) []byte {
	f, err := os.Open(name)
	if err != nil {
		return []byte("<html><head></head><body><h3>Errors </h3></body></html>")
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return []byte("<html><head></head><body><h3>Errors </h3></body></html>")
	}
	return buf
}

func main() {
	fmt.Println("hello")
	http.HandleFunc("/upload", Upload)
	http.HandleFunc("/", IndexView)
	http.HandleFunc("/index", IndexView)
	http.HandleFunc("/image", ImageView) ///iamge --> ImageView 函数 ==>路由
	http.HandleFunc("/detail", DetailView)
	http.HandleFunc("/list", ListView)
	log.Println("run .. ")
	http.ListenAndServe(":8080", nil)
}
