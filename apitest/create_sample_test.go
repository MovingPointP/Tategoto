package apitest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"tategoto/model"
	"tategoto/pkg/auth"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var Router *gin.Engine
var TestingUser *model.User
var TestingPost *model.Post
var Token string

func TestCreateSample(t *testing.T) {
	//router取得 table初期化
	Router := NewRouter()
	successSignup_200(t, Router)
	Token, _ = auth.CreateUserJWT(TestingUser.ID)
	successCreatePost_200(t, Router)
}

// 正常なサインアップ
func successSignup_200(t *testing.T, r *gin.Engine) {
	requestJson := `{"mail": "hoge@mail.com", "password": "hogehoge", "name": "hoge"}`
	body := bytes.NewBuffer([]byte(requestJson))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/signup", body)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	data := new(resUser)
	jsonBytes := []byte(w.Body.String())
	json.Unmarshal(jsonBytes, &data)
	TestingUser = data.User

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 26, len(TestingUser.ID))
	assert.Equal(t, false, TestingUser.CreatedAt.IsZero())
	assert.Equal(t, false, TestingUser.UpdatedAt.IsZero())
	assert.Equal(t, true, TestingUser.DeletedAt.Time.IsZero())
	assert.Equal(t, "hoge@mail.com", TestingUser.Mail)
	assert.Equal(t, "", TestingUser.Password)
	assert.Equal(t, "hoge", TestingUser.Name)
}

// 正常なポストの投稿
func successCreatePost_200(t *testing.T, r *gin.Engine) {
	requestJson := `{ "content": "こんにちは", "user_id": "` + TestingUser.ID + `"}`
	body := bytes.NewBuffer([]byte(requestJson))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/posts", body)
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{
		Name:  "token",
		Value: Token,
	})
	r.ServeHTTP(w, req)

	data := new(resPost)
	jsonBytes := []byte(w.Body.String())
	json.Unmarshal(jsonBytes, &data)
	TestingPost = data.Post

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 26, len(TestingPost.ID))
	assert.Equal(t, false, TestingPost.CreatedAt.IsZero())
	assert.Equal(t, false, TestingPost.UpdatedAt.IsZero())
	assert.Equal(t, true, TestingPost.DeletedAt.Time.IsZero())
	assert.Equal(t, "こんにちは", TestingPost.Content)
	assert.Equal(t, TestingUser.ID, TestingPost.UserID)
}
