package model

import "time"

// Article 新闻
type Article struct {
	Id      int64     `json:"id"` //tag
	Title   string    `json:"title"`
	Author  string    `json:"author"`
	Content string    `json:"content"`
	Hits    int       `json:"hits"`
	Utime   time.Time `json:"utime"`
}

// ArticleGet 查询一条数据
func ArticleGet(id int64) (Article, error) {
	mod := Article{}
	err := Db.Unsafe().Get(&mod, "select * from Article where id =? limit 1", id)
	return mod, err
}

//  ArticleList 返回数据列表
func ArticleList() ([]Article, error) {
	mods := make([]Article, 0, 10)
	err := Db.Unsafe().Select(&mods, "select * from Article order by id desc limit 10")
	return mods, err
}

// ArticlePage 分页
func ArticlePage(pi int, ps int) ([]Article, error) {
	mods := make([]Article, 0, 10)
	err := Db.Unsafe().Select(&mods, "select * from Article order by id desc limit ?,?", (pi-1)*ps, ps)
	return mods, err
}

// ArticlePageCount 数据总数
func ArticlePageCount() int {
	count := 0
	Db.Get(&count, "select count(id) as count from article")
	return count
}

// ArticleDel 删除数据
func ArticleDel(id int64) bool {
	res, _ := Db.Exec("delete from Article where id = ?", id)
	if res == nil {
		return false
	}
	rows, _ := res.RowsAffected()
	if rows >= 1 {
		return true
	}
	return false
}

func ArticleAdd(mod *Article) error {
	_, err := Db.Exec("insert into Article (title,author,content,hits,utime) values(?,?,?,?,?)", mod.Title, mod.Author, mod.Content, mod.Hits, mod.Utime)
	return err
}

func ArticleEdit(mod *Article) error {
	_, err := Db.Exec("update article set title=?,author=?,content=?,hits=? where id=?", mod.Title, mod.Author, mod.Content, mod.Hits, mod.Id)
	return err
}
