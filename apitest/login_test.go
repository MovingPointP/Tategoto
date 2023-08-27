package apitest

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"tategoto/config/msg"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	//r取得 table初期化
	r := NewRouter()

	noMailLogin(t, r)
	successSignup(t, r)
	duplicateSignUp(t, r)
	discordancePasswordLogin(t, r)
	successLogin(t, r)
}

// 存在しないユーザーログイン
func noMailLogin(t *testing.T, r *gin.Engine) {
	requestJson := `{ "mail": "hoge@mail.com", "password": "hogehoge"}`
	body := bytes.NewBuffer([]byte(requestJson))

	//メールアドレスかパスワードが間違っている
	responseJson := `{ "message":"` + msg.IncorrectMailOrPasswordErr + `"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/login", body)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.JSONEq(t, w.Body.String(), responseJson)
	assert.Equal(t, w.Code, 500)
}

// 正常なサインアップ
func successSignup(t *testing.T, r *gin.Engine) {
	requestJson := `{ "name": "hogeman", "mail": "hoge@mail.com", "password": "hogehoge"}`
	body := bytes.NewBuffer([]byte(requestJson))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/signup", body)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
}

// 重複ユーザーサインアップ
func duplicateSignUp(t *testing.T, r *gin.Engine) {
	requestJson := `{ "name": "hogewoman", "mail": "hoge@mail.com", "password": "hogehoge"}`
	body := bytes.NewBuffer([]byte(requestJson))

	//メールアドレスが重複している
	responseJson := `{ "message":"` + msg.DuplicateMailErr + `"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/signup", body)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.JSONEq(t, w.Body.String(), responseJson)
	assert.Equal(t, w.Code, 500)
}

// パスワード不一致ログイン
func discordancePasswordLogin(t *testing.T, r *gin.Engine) {
	requestJson := `{ "mail": "hoge@mail.com", "password": "fugafuga"}`
	body := bytes.NewBuffer([]byte(requestJson))

	//メールアドレスかパスワードが間違っている
	responseJson := `{ "message":"` + msg.IncorrectMailOrPasswordErr + `"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/login", body)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.JSONEq(t, w.Body.String(), responseJson)
	assert.Equal(t, w.Code, 500)
}

// 正常なログイン
func successLogin(t *testing.T, r *gin.Engine) {
	requestJson := `{ "mail": "hoge@mail.com", "password": "hogehoge"}`
	body := bytes.NewBuffer([]byte(requestJson))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/login", body)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
}
