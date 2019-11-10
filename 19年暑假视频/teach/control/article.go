package control

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"teach/model"
	"time"
)

// func ApiArticleAdd(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm() //解析数据
// 	mod := &model.Article{}
// 	mod.Title = r.Form.Get("title")
// 	mod.Author = r.Form["author"][0]
// 	mod.Content = r.FormValue("content")
// 	mod.Hits, _ = strconv.Atoi(r.Form.Get("hits"))
// 	mod.Utime = time.Now()
// 	log.Println(mod)
// 	err := model.ArticleAdd(mod)
// 	if err == nil {
// 		Succ(w, "添加成功")
// 		return
// 	}
// 	Fail(w, "添加失败", err.Error())
// 	return
// }
func ApiArticleAdd(w http.ResponseWriter, r *http.Request) {
	mod := &model.Article{}
	// err := json.NewDecoder(r.Body).Decode(mod)
	buf, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(buf, mod)
	if err != nil {
		Fail(w, "输入数据有误", err.Error())
		return
	}
	mod.Utime = time.Now()
	err = model.ArticleAdd(mod)
	if err != nil {
		Fail(w, "添加失败", err.Error())
		return
	}
	Succ(w, "添加成功")
	return
}
func ApiArticleEdit(w http.ResponseWriter, r *http.Request) {
	mod := &model.Article{}
	// err := json.NewDecoder(r.Body).Decode(mod)
	buf, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(buf, mod)
	if err != nil {
		Fail(w, "输入数据有误", err.Error())
		return
	}
	err = model.ArticleEdit(mod)
	if err != nil {
		Fail(w, "修改失败", err.Error())
		return
	}
	Succ(w, "修改成功")
	return
}

// IndexData 首页数据
func IndexData(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Header().Set("Content-Type", "application/json")
	idStr := r.Form.Get("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	mod, err := model.ArticleGet(id)
	if err != nil {
		Fail(w, err.Error())
		return
	}
	buf, _ := json.Marshal(mod)
	w.Write(buf)
}

// ListData 列表页数据
func ListData(w http.ResponseWriter, r *http.Request) {
	mods, err := model.ArticleList()
	if err != nil {
		Fail(w, err.Error())
		return
	}
	Succ(w, "列表", mods)
}

type Ext struct {
	Count int         `json:"count"`
	Items interface{} `json:"items"`
}

func ApiArticlePage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	pi, _ := strconv.Atoi(r.Form.Get("pi"))
	ps, _ := strconv.Atoi(r.Form.Get("ps"))
	count := model.ArticlePageCount()
	if count < 1 {
		Fail(w, "未查询到数据", "count < 1")
		return
	}
	mods, _ := model.ArticlePage(pi, ps)
	if len(mods) < 1 {
		Fail(w, "未查询到数据", "len(mods) < 1")
		return
	}
	ext := Ext{
		Count: count,
		Items: mods,
	}
	Succ(w, "分页数据", ext)
}

func ListDel(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idStr := r.Form.Get("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	if model.ArticleDel(id) {
		Succ(w, "删除成功")
		return
	}
	Fail(w, "删除失败")
	return
}

func Fail(w http.ResponseWriter, msg string, data ...interface{}) {
	mod := Reply{
		Code: 300, //失败
		Msg:  msg,
	}
	if len(data) > 0 {
		mod.Data = data[0]
	}
	buf, _ := json.Marshal(mod)
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf)
}

func Succ(w http.ResponseWriter, msg string, data ...interface{}) {
	mod := Reply{
		Code: 200, //成功
		Msg:  msg,
	}
	if len(data) > 0 {
		mod.Data = data[0]
	}
	buf, _ := json.Marshal(mod)
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf)
}

type Reply struct {
	Code int         `json:"code"` // 作为状态标识  200 成功，300 失败 ，310 输入有误 ，320 输出有误
	Msg  string      `json:"msg"`  // 给用户提示的
	Data interface{} `json:"data"` //返回数据/返回开发者查看的错误信息 void* object 可以存放任意数据类型
}

// tag json code msg data
//
