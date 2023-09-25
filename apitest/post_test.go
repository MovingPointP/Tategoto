package apitest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"tategoto/config/msg"
	"tategoto/pkg/funk"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// func TestPost(t *testing.T) {
// 	//router取得 table初期化
// 	r := NewRouter()
// 	//sampleの作成
// 	createTestSample()
// 	//以下テスト
// 	postFunctions(t, r)
// }

func postFunctions(t *testing.T, r *gin.Engine) {
	beforeLoginPost_303(t, r)
	differentUserIDPost_303(t, r)
	postPost_200(t, r)
	beforeLoginGetPostByID_303(t, r)
	getNoPostByID_200(t, r)
	getPostByID_200(t, r)
	beforeLoginGetPostsWithQuery_303(t, r)
	getNoPostsWithQuery_200(t, r)
	getPostsWithQuery_200(t, r)
}

// tokenなしのポストの投稿
func beforeLoginPost_303(t *testing.T, r *gin.Engine) {
	requestJson := `{ "content": "こんにちは", "user_id": "` + SampleUserHoge.ID + `"}`
	body := bytes.NewBuffer([]byte(requestJson))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/posts", body)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	data := new(resFail)
	jsonBytes := []byte(w.Body.String())
	json.Unmarshal(jsonBytes, &data)
	message := data.Message

	assert.Equal(t, 303, w.Code)
	assert.Equal(t, msg.ShouldLoginErr, message)
}

// tokenとは異なるuserIDでのポストの投稿
func differentUserIDPost_303(t *testing.T, r *gin.Engine) {
	requestJson := `{ "content": "こんにちは", "user_id": "` + SampleUserFuga.ID + `"}`
	body := bytes.NewBuffer([]byte(requestJson))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/posts", body)
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{
		Name:  "token",
		Value: TokenHoge,
	})
	r.ServeHTTP(w, req)

	data := new(resFail)
	jsonBytes := []byte(w.Body.String())
	json.Unmarshal(jsonBytes, &data)
	message := data.Message

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, msg.IncorrectUserIDErr, message)
}

// 正常なポストの投稿
func postPost_200(t *testing.T, r *gin.Engine) {
	requestJson := `{ "content": "こんにちは", "user_id": "` + SampleUserFuga.ID + `"}`
	body := bytes.NewBuffer([]byte(requestJson))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/posts", body)
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{
		Name:  "token",
		Value: TokenFuga,
	})
	r.ServeHTTP(w, req)

	data := new(resPost)
	jsonBytes := []byte(w.Body.String())
	json.Unmarshal(jsonBytes, &data)
	testPost := data.Post
	now := time.Now()

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 26, len(testPost.ID))
	assert.Equal(t, true, funk.CompareAboutTime(now, testPost.CreatedAt))
	assert.Equal(t, true, funk.CompareAboutTime(now, testPost.UpdatedAt))
	assert.Equal(t, true, testPost.DeletedAt.Time.IsZero())
	assert.Equal(t, "こんにちは", testPost.Content)
	assert.Equal(t, SampleUserFuga.ID, testPost.UserID)
}

// tokenなしのIDによるポストの取得
func beforeLoginGetPostByID_303(t *testing.T, r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/posts/"+SamplePostHello.ID, nil)
	r.ServeHTTP(w, req)

	data := new(resFail)
	jsonBytes := []byte(w.Body.String())
	json.Unmarshal(jsonBytes, &data)
	message := data.Message

	assert.Equal(t, 303, w.Code)
	assert.Equal(t, msg.ShouldLoginErr, message)
}

// IDによる存在しないポストの取得
func getNoPostByID_200(t *testing.T, r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/posts/noting", nil)
	req.AddCookie(&http.Cookie{
		Name:  "token",
		Value: TokenHoge,
	})
	r.ServeHTTP(w, req)

	data := new(resPost)
	jsonBytes := []byte(w.Body.String())
	json.Unmarshal(jsonBytes, &data)
	testPost := data.Post

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "", testPost.ID)
	assert.Equal(t, true, testPost.CreatedAt.IsZero())
	assert.Equal(t, true, testPost.UpdatedAt.IsZero())
	assert.Equal(t, true, testPost.DeletedAt.Time.IsZero())
	assert.Equal(t, "", testPost.Content)
	assert.Equal(t, "", testPost.UserID)
}

// 正常なIDによるポストの取得
func getPostByID_200(t *testing.T, r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/posts/"+SamplePostHello.ID, nil)
	req.AddCookie(&http.Cookie{
		Name:  "token",
		Value: TokenHoge,
	})
	r.ServeHTTP(w, req)

	data := new(resPost)
	jsonBytes := []byte(w.Body.String())
	json.Unmarshal(jsonBytes, &data)
	testPost := data.Post

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, SamplePostHello.ID, testPost.ID)
	assert.Equal(t, true, SamplePostHello.CreatedAt.Equal(testPost.CreatedAt))
	assert.Equal(t, true, SamplePostHello.UpdatedAt.Equal(testPost.UpdatedAt))
	assert.Equal(t, SamplePostHello.DeletedAt, testPost.DeletedAt)
	assert.Equal(t, SamplePostHello.Content, testPost.Content)
	assert.Equal(t, SamplePostHello.UserID, testPost.UserID)
}

// tokenなしのクエリによるポストの取得
func beforeLoginGetPostsWithQuery_303(t *testing.T, r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/posts?uid="+SampleUserHoge.ID, nil)
	r.ServeHTTP(w, req)

	data := new(resFail)
	jsonBytes := []byte(w.Body.String())
	json.Unmarshal(jsonBytes, &data)
	message := data.Message

	assert.Equal(t, 303, w.Code)
	assert.Equal(t, msg.ShouldLoginErr, message)
}

// クエリによる存在しないポストの取得
func getNoPostsWithQuery_200(t *testing.T, r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/posts?uid=noting", nil)
	req.AddCookie(&http.Cookie{
		Name:  "token",
		Value: TokenHoge,
	})
	r.ServeHTTP(w, req)

	data := new(resPosts)
	jsonBytes := []byte(w.Body.String())
	json.Unmarshal(jsonBytes, &data)
	testPosts := data.Posts

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 0, len(testPosts))
}

// 正常なクエリによるポストの取得
func getPostsWithQuery_200(t *testing.T, r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/posts?uid="+SampleUserHoge.ID, nil)
	req.AddCookie(&http.Cookie{
		Name:  "token",
		Value: TokenHoge,
	})
	r.ServeHTTP(w, req)

	data := new(resPosts)
	jsonBytes := []byte(w.Body.String())
	json.Unmarshal(jsonBytes, &data)
	testPosts := data.Posts

	fmt.Println(SamplePostHello)
	fmt.Println(SamplePostWorld)
	fmt.Println(testPosts)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, SamplePostHello.ID, testPosts[0].ID)
	assert.Equal(t, true, SamplePostHello.CreatedAt.Equal(testPosts[0].CreatedAt))
	assert.Equal(t, true, SamplePostHello.UpdatedAt.Equal(testPosts[0].UpdatedAt))
	assert.Equal(t, SamplePostHello.DeletedAt, testPosts[0].DeletedAt)
	assert.Equal(t, SamplePostHello.Content, testPosts[0].Content)
	assert.Equal(t, SamplePostHello.UserID, testPosts[0].UserID)
	assert.Equal(t, SamplePostWorld.ID, testPosts[1].ID)
	assert.Equal(t, true, SamplePostWorld.CreatedAt.Equal(testPosts[1].CreatedAt))
	assert.Equal(t, true, SamplePostWorld.UpdatedAt.Equal(testPosts[1].UpdatedAt))
	assert.Equal(t, SamplePostWorld.DeletedAt, testPosts[1].DeletedAt)
	assert.Equal(t, SamplePostWorld.Content, testPosts[1].Content)
	assert.Equal(t, SamplePostWorld.UserID, testPosts[1].UserID)
}
