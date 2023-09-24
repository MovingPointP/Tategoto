package apitest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"tategoto/config/msg"
	"tategoto/pkg/funk"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// func TestAuth(t *testing.T) {
// 	//router取得 table初期化
// 	r := NewRouter()
// 	//sampleの作成
// 	createTestSample()
// 	//以下テスト
// 	authFunctions(t, r)
// }

func authFunctions(t *testing.T, r *gin.Engine) {
	duplicateSignUp_400(t, r)
	signup_200(t, r)
	noMailLogin_400(t, r)
	discordancePasswordLogin_400(t, r)
}

// 重複ユーザーサインアップ
func duplicateSignUp_400(t *testing.T, r *gin.Engine) {
	requestJson := `{"mail": "hoge@mail.com", "password": "hogehoge", "name": "hoge"}`
	body := bytes.NewBuffer([]byte(requestJson))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/signup", body)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	data := new(resFail)
	jsonBytes := []byte(w.Body.String())
	json.Unmarshal(jsonBytes, &data)
	message := data.Message

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, msg.DuplicateMailErr, message)
}

// 正常なサインアップ
func signup_200(t *testing.T, r *gin.Engine) {
	requestJson := `{"mail": "piyo@mail.com", "password": "piyo", "name": "piyo"}`
	body := bytes.NewBuffer([]byte(requestJson))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/signup", body)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	data := new(resUser)
	jsonBytes := []byte(w.Body.String())
	json.Unmarshal(jsonBytes, &data)
	testUser := data.User
	now := time.Now()

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 26, len(testUser.ID))
	assert.Equal(t, true, funk.CompareAboutTime(now, testUser.CreatedAt))
	assert.Equal(t, true, funk.CompareAboutTime(now, testUser.UpdatedAt))
	assert.Equal(t, true, testUser.DeletedAt.Time.IsZero())
	assert.Equal(t, "piyo@mail.com", testUser.Mail)
	assert.Equal(t, "", testUser.Password)
	assert.Equal(t, "piyo", testUser.Name)
}

// 存在しないユーザーでログイン
func noMailLogin_400(t *testing.T, r *gin.Engine) {
	requestJson := `{ "mail": "noting@mail.com", "password": "noting"}`
	body := bytes.NewBuffer([]byte(requestJson))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/login", body)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	data := new(resFail)
	jsonBytes := []byte(w.Body.String())
	json.Unmarshal(jsonBytes, &data)
	message := data.Message

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, msg.IncorrectMailOrPasswordErr, message)
}

// パスワード不一致ログイン
func discordancePasswordLogin_400(t *testing.T, r *gin.Engine) {
	requestJson := `{ "mail": "hoge@mail.com", "password": "hogehoge"}`
	body := bytes.NewBuffer([]byte(requestJson))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/login", body)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	data := new(resFail)
	jsonBytes := []byte(w.Body.String())
	json.Unmarshal(jsonBytes, &data)
	message := data.Message

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, msg.IncorrectMailOrPasswordErr, message)
}
