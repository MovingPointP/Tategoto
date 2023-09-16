package apitest

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"tategoto/config/msg"
	"tategoto/pkg/auth"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var token string

func TestLogin(t *testing.T) {
	//router取得 table初期化
	r := NewRouter()

	//token取得
	token, _ = auth.CreateUserJWT("1")

	noMailLogin_400(t, r)
	SuccessSignup_200(t, r)
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

	assert.JSONEq(t, responseJson, w.Body.String())
	assert.Equal(t, 400, w.Code)
}

// 正常なサインアップ
func SuccessSignup_200(t *testing.T, r *gin.Engine) {
	requestJson := `{ "name": "hogeman", "mail": "hoge@mail.com", "password": "hogehoge"}`
	body := bytes.NewBuffer([]byte(requestJson))

	responseElement := []string{`"ID":1`, `"mail":"hoge@mail.com"`, `"password":""`, `"name":"hogeman"`}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/signup", body)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	for _, v := range responseElement {
		hasElement := strings.Contains(w.Body.String(), v)
		assert.Equal(t, true, hasElement)
		if !hasElement {
			fmt.Println("expected", v)
			fmt.Println("actual", w.Body.String())
		}
	}
	assert.Equal(t, 200, w.Code)
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

	assert.JSONEq(t, responseJson, w.Body.String())
	assert.Equal(t, 400, w.Code)
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

	assert.JSONEq(t, responseJson, w.Body.String())
	assert.Equal(t, 400, w.Code)
}

// 正常なログイン
func successLogin_200(t *testing.T, r *gin.Engine) {
	requestJson := `{ "mail": "hoge@mail.com", "password": "hogehoge"}`
	body := bytes.NewBuffer([]byte(requestJson))

	responseElement := []string{`"ID":1`, `"mail":"hoge@mail.com"`, `"password":""`, `"name":"hogeman"`}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/login", body)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	for _, v := range responseElement {
		hasElement := strings.Contains(w.Body.String(), v)
		assert.Equal(t, true, hasElement)
		if !hasElement {
			fmt.Println("expected", v)
			fmt.Println("actual", w.Body.String())
		}
	}
	assert.Equal(t, 200, w.Code)
}
