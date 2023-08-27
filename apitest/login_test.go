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
	//router取得 table初期化
	r := NewRouter()

	noMailLogin_400(t, r)
	successSignup_200(t, r)
	duplicateSignUp_400(t, r)
	discordancePasswordLogin_400(t, r)
	successLogin_200(t, r)
}

// 存在しないユーザーログイン
func noMailLogin_400(t *testing.T, r *gin.Engine) {
	requestJson := `{ "mail": "hoge@mail.com", "password": "hogehoge"}`
	body := bytes.NewBuffer([]byte(requestJson))

	//メールアドレスかパスワードが間違っている
	responseJson := `{ "message":"` + msg.IncorrectMailOrPasswordErr + `"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/login", body)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.JSONEq(t, w.Body.String(), responseJson)
	assert.Equal(t, w.Code, 400)
}

// 正常なサインアップ
func successSignup_200(t *testing.T, r *gin.Engine) {
	requestJson := `{ "name": "hogeman", "mail": "hoge@mail.com", "password": "hogehoge"}`
	body := bytes.NewBuffer([]byte(requestJson))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/signup", body)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
}

// 重複ユーザーサインアップ
func duplicateSignUp_400(t *testing.T, r *gin.Engine) {
	requestJson := `{ "name": "hogewoman", "mail": "hoge@mail.com", "password": "hogehoge"}`
	body := bytes.NewBuffer([]byte(requestJson))

	//メールアドレスが重複している
	responseJson := `{ "message":"` + msg.DuplicateMailErr + `"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/signup", body)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.JSONEq(t, w.Body.String(), responseJson)
	assert.Equal(t, w.Code, 400)
}

// パスワード不一致ログイン
func discordancePasswordLogin_400(t *testing.T, r *gin.Engine) {
	requestJson := `{ "mail": "hoge@mail.com", "password": "fugafuga"}`
	body := bytes.NewBuffer([]byte(requestJson))

	//メールアドレスかパスワードが間違っている
	responseJson := `{ "message":"` + msg.IncorrectMailOrPasswordErr + `"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/login", body)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.JSONEq(t, w.Body.String(), responseJson)
	assert.Equal(t, w.Code, 400)
}

// 正常なログイン
func successLogin_200(t *testing.T, r *gin.Engine) {
	requestJson := `{ "mail": "hoge@mail.com", "password": "hogehoge"}`
	body := bytes.NewBuffer([]byte(requestJson))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/login", body)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
}
