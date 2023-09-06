package apitest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"tategoto/config/msg"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	//router取得 table初期化
	r := NewRouter()

	SuccessSignup_200(t, r)
	beforeLoginGetUserByID_303(t, r)
	GetNoUserByID_200(t, r)
	successGetUserByID_200(t, r)
	beforeLoginGetUsersWithQuery_303(t, r)
	GetNoUsersWithQuery_200(t, r)
	successGetUsersWithQuery_200(t, r)
}

// tokenなしのIDによるユーザーの取得
func beforeLoginGetUserByID_303(t *testing.T, r *gin.Engine) {

	//ログインが必要な処理
	responseJson := `{ "message":"` + msg.ShouldLoginErr + `", "path":"/api/users/1"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/users/1", nil)
	r.ServeHTTP(w, req)

	assert.JSONEq(t, responseJson, w.Body.String())
	assert.Equal(t, 303, w.Code)
}

// IDによる存在しないユーザーの取得
func GetNoUserByID_200(t *testing.T, r *gin.Engine) {

	responseElement := []string{`"ID":0`}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/users/2", nil)
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

// 正常なIDによるユーザーの取得
func successGetUserByID_200(t *testing.T, r *gin.Engine) {

	responseElement := []string{`"ID":1`, `"mail":""`, `"password":""`, `"name":"hogeman"`}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/users/1", nil)
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

// tokenなしのクエリによるユーザーの取得
func beforeLoginGetUsersWithQuery_303(t *testing.T, r *gin.Engine) {

	//ログインが必要な処理
	responseJson := `{ "message":"` + msg.ShouldLoginErr + `", "path":"/api/users"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/users?name=hogeman", nil)
	r.ServeHTTP(w, req)

	assert.JSONEq(t, responseJson, w.Body.String())
	assert.Equal(t, 303, w.Code)
}

// クエリによる存在しないユーザーの取得
func GetNoUsersWithQuery_200(t *testing.T, r *gin.Engine) {

	responseElement := []string{`"users":[]`}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/users?name=hogewoman", nil)
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

// 正常なクエリによるユーザーの取得
func successGetUsersWithQuery_200(t *testing.T, r *gin.Engine) {

	responseElement := []string{`"ID":1`, `"mail":""`, `"password":""`, `"name":"hogeman"`}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/users?name=hogeman", nil)
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
