package apitest

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"tategoto/config/msg"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPost(t *testing.T) {
	//router取得 table初期化
	r := NewRouter()

	SuccessSignup_200(t, r)
	beforeLoginPost_303(t, r)
	successPostPost_200(t, r)
	beforeLoginGetPostByID_303(t, r)
	GetNoPostByID_200(t, r)
	successGetPostByID_200(t, r)
	beforeLoginGetPostsWithQuery_303(t, r)
	GetNoPostsWithQuery_200(t, r)
	successGetPostsWithQuery_200(t, r)
}

// tokenなしのポストの投稿
func beforeLoginPost_303(t *testing.T, r *gin.Engine) {
	requestJson := `{ "content": "hello", "user_id": "1"}`
	body := bytes.NewBuffer([]byte(requestJson))

	//ログインが必要な処理
	responseJson := `{ "message":"` + msg.ShouldLoginErr + `", "path":"/api/posts"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/posts", body)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.JSONEq(t, responseJson, w.Body.String())
	assert.Equal(t, 303, w.Code)
}

// 正常なポストの投稿
func successPostPost_200(t *testing.T, r *gin.Engine) {
	requestJson := `{ "content": "hello", "user_id": 1}`
	body := bytes.NewBuffer([]byte(requestJson))

	responseElement := []string{`"ID":1`, `"content":"hello"`, `"user_id":1`}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/posts", body)
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{
		Name:  "token",
		Value: token,
	})
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

// tokenなしのIDによるポストの取得
func beforeLoginGetPostByID_303(t *testing.T, r *gin.Engine) {

	//ログインが必要な処理
	responseJson := `{ "message":"` + msg.ShouldLoginErr + `", "path":"/api/posts/1"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/posts/1", nil)
	r.ServeHTTP(w, req)

	assert.JSONEq(t, responseJson, w.Body.String())
	assert.Equal(t, 303, w.Code)
}

// IDによる存在しないポストの取得
func GetNoPostByID_200(t *testing.T, r *gin.Engine) {

	responseElement := []string{`"ID":0`}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/posts/2", nil)
	req.AddCookie(&http.Cookie{
		Name:  "token",
		Value: token,
	})
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

// 正常なIDによるポストの取得
func successGetPostByID_200(t *testing.T, r *gin.Engine) {

	responseElement := []string{`"ID":1`, `"content":"hello"`, `"user_id":1`}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/posts/1", nil)
	req.AddCookie(&http.Cookie{
		Name:  "token",
		Value: token,
	})
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

// tokenなしのクエリによるポストの取得
func beforeLoginGetPostsWithQuery_303(t *testing.T, r *gin.Engine) {

	//ログインが必要な処理
	responseJson := `{ "message":"` + msg.ShouldLoginErr + `", "path":"/api/posts"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/posts?uid=1", nil)
	r.ServeHTTP(w, req)

	assert.JSONEq(t, responseJson, w.Body.String())
	assert.Equal(t, 303, w.Code)
}

// クエリによる存在しないポストの取得
func GetNoPostsWithQuery_200(t *testing.T, r *gin.Engine) {

	responseElement := []string{`"posts":[]`}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/posts?uid=2", nil)
	req.AddCookie(&http.Cookie{
		Name:  "token",
		Value: token,
	})
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

// 正常なクエリによるポストの取得
func successGetPostsWithQuery_200(t *testing.T, r *gin.Engine) {

	responseElement := []string{`"ID":1`, `"content":"hello"`, `"user_id":1`}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/posts?uid=1", nil)
	req.AddCookie(&http.Cookie{
		Name:  "token",
		Value: token,
	})
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
