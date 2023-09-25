package apitest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"tategoto/config/msg"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// func TestAuth(t *testing.T) {
// 	//router取得 table初期化
// 	r := NewRouter()
// 	//sampleの作成
// 	createTestSample()
// 	//以下テスト
// 	userFunctions(t, r)
// }

func userFunctions(t *testing.T, r *gin.Engine) {
	beforeLoginGetUserByID_303(t, r)
	getNoUserByID_200(t, r)
	getUserByID_200(t, r)
	beforeLoginGetUsersWithQuery_303(t, r)
	getNoUsersWithQuery_200(t, r)
	getUsersWithQuery_200(t, r)
}

// tokenなしのIDによるユーザーの取得
func beforeLoginGetUserByID_303(t *testing.T, r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/users/"+SampleUserHoge.ID, nil)
	r.ServeHTTP(w, req)

	data := new(resFail)
	jsonBytes := []byte(w.Body.String())
	json.Unmarshal(jsonBytes, &data)
	message := data.Message

	assert.Equal(t, 303, w.Code)
	assert.Equal(t, msg.ShouldLoginErr, message)
}

// IDによる存在しないユーザーの取得
func getNoUserByID_200(t *testing.T, r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/users/hoge", nil)
	req.AddCookie(&http.Cookie{
		Name:  "token",
		Value: TokenHoge,
	})
	r.ServeHTTP(w, req)

	data := new(resUser)
	jsonBytes := []byte(w.Body.String())
	json.Unmarshal(jsonBytes, &data)
	testUser := data.User

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "", testUser.ID)
	assert.Equal(t, true, testUser.CreatedAt.IsZero())
	assert.Equal(t, true, testUser.UpdatedAt.IsZero())
	assert.Equal(t, true, testUser.DeletedAt.Time.IsZero())
	assert.Equal(t, "", testUser.Mail)
	assert.Equal(t, "", testUser.Password)
	assert.Equal(t, "", testUser.Name)
}

// 正常なIDによるユーザーの取得
func getUserByID_200(t *testing.T, r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/users/"+SampleUserHoge.ID, nil)
	req.AddCookie(&http.Cookie{
		Name:  "token",
		Value: TokenHoge,
	})
	r.ServeHTTP(w, req)

	data := new(resUser)
	jsonBytes := []byte(w.Body.String())
	json.Unmarshal(jsonBytes, &data)
	testUser := data.User

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, SampleUserHoge.ID, testUser.ID)
	assert.Equal(t, true, SampleUserHoge.CreatedAt.Equal(testUser.CreatedAt))
	assert.Equal(t, true, SampleUserHoge.UpdatedAt.Equal(testUser.UpdatedAt))
	assert.Equal(t, SampleUserHoge.DeletedAt, testUser.DeletedAt)
	assert.Equal(t, "", testUser.Mail)
	assert.Equal(t, "", testUser.Password)
	assert.Equal(t, SampleUserHoge.Name, testUser.Name)
}

// tokenなしのクエリによるユーザーの取得
func beforeLoginGetUsersWithQuery_303(t *testing.T, r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/users?name="+SampleUserHoge.Name, nil)
	r.ServeHTTP(w, req)

	data := new(resFail)
	jsonBytes := []byte(w.Body.String())
	json.Unmarshal(jsonBytes, &data)
	message := data.Message

	assert.Equal(t, 303, w.Code)
	assert.Equal(t, msg.ShouldLoginErr, message)
}

// クエリによる存在しないユーザーの取得
func getNoUsersWithQuery_200(t *testing.T, r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/users?name=noting", nil)
	req.AddCookie(&http.Cookie{
		Name:  "token",
		Value: TokenHoge,
	})
	r.ServeHTTP(w, req)

	data := new(resUsers)
	jsonBytes := []byte(w.Body.String())
	json.Unmarshal(jsonBytes, &data)
	testUsers := data.Users

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 0, len(testUsers))
}

// 正常なクエリによるユーザーの取得
func getUsersWithQuery_200(t *testing.T, r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/users?name="+SampleUserHoge.Name, nil)
	req.AddCookie(&http.Cookie{
		Name:  "token",
		Value: TokenHoge,
	})
	r.ServeHTTP(w, req)

	data := new(resUsers)
	jsonBytes := []byte(w.Body.String())
	json.Unmarshal(jsonBytes, &data)
	testUsers := data.Users

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, SampleUserHoge.ID, testUsers[0].ID)
	assert.Equal(t, true, SampleUserHoge.CreatedAt.Equal(testUsers[0].CreatedAt))
	assert.Equal(t, true, SampleUserHoge.UpdatedAt.Equal(testUsers[0].UpdatedAt))
	assert.Equal(t, SampleUserHoge.DeletedAt, testUsers[0].DeletedAt)
	assert.Equal(t, "", testUsers[0].Mail)
	assert.Equal(t, "", testUsers[0].Password)
	assert.Equal(t, SampleUserHoge.Name, testUsers[0].Name)
	assert.Equal(t, SampleUserFuga.ID, testUsers[1].ID)
	assert.Equal(t, true, SampleUserFuga.CreatedAt.Equal(testUsers[1].CreatedAt))
	assert.Equal(t, true, SampleUserFuga.UpdatedAt.Equal(testUsers[1].UpdatedAt))
	assert.Equal(t, SampleUserFuga.DeletedAt, testUsers[1].DeletedAt)
	assert.Equal(t, "", testUsers[1].Mail)
	assert.Equal(t, "", testUsers[1].Password)
	assert.Equal(t, SampleUserFuga.Name, testUsers[1].Name)
}
