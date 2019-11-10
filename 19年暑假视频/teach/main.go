package main

import (
	"net/http"
	"teach/control"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/", control.IndexView)                      //页面
	http.HandleFunc("/list", control.ListView)                   //页面
	http.HandleFunc("/edit", control.EditView)                   //页面
	http.HandleFunc("/detail", control.DetailView)               //详细页面
	http.HandleFunc("/add", control.ViewArticleAdd)              //
	http.HandleFunc("/api/index/data", control.IndexData)        //数据
	http.HandleFunc("/api/list/data", control.ListData)          //数据
	http.HandleFunc("/api/article/page", control.ApiArticlePage) //分页数据
	http.HandleFunc("/api/list/del", control.ListDel)            //操作
	http.HandleFunc("/api/article/add", control.ApiArticleAdd)   //添加
	http.HandleFunc("/api/article/edit", control.ApiArticleEdit) //修改
	http.HandleFunc("/api/upload", control.ApiUpload)
	http.ListenAndServe(":8080", nil)
}
