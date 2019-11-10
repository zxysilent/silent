package control

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path"
	"time"
	"unsafe"
)

func ApiUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1 << 20 * 20) //20MB
	f, h, err := r.FormFile("upfile")
	if err != nil {
		Fail(w, "上传失败", err.Error())
	}
	// log.Println(f, h, err)
	os.MkdirAll("static/upload/", 0666)
	ext := path.Ext(h.Filename)
	name := "static/upload/" + RandStr(10) + ext
	dst, _ := os.Create(name)
	io.Copy(dst, f)
	f.Close()
	dst.Close()
	w.Header().Set("Content-Type", "application/json")
	mod := editorReply{
		Original: h.Filename,
		State:    "SUCCESS",
		Title:    h.Filename,
		Url:      "/" + name,
	}
	w.Write(mod.Json())
}

// rand.Seed(time.Now().UnixNano())
// log.Println(rand.Intn(36))
// log.Println(rand.Intn(36))
// log.Println(rand.Intn(36))
// log.Println(rand.Intn(36))

var all = "abcdefghijklmnopqrstuvwxyz0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 构建随机字符串
func RandStr(ln int) string {
	// rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	res := make([]byte, ln, ln)
	for i := 0; i < ln; i++ {
		res[i] = all[rand.Intn(36)]
	}
	// return string(res)
	return *(*string)(unsafe.Pointer(&res))
}

type editorReply struct {
	Original string `json:"original"` // len cap point  // 64bit （8 + 8 + 8 ）*4 96 B   ---8B
	State    string `json:"state"`
	Title    string `json:"title"`
	Url      string `json:"url"`
}

// 方法
func (er *editorReply) Json() []byte {
	buf, _ := json.Marshal(er)
	return buf
}
