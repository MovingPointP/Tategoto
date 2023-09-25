package apitest

import "testing"

func TestAll(t *testing.T) {
	//router取得 table初期化
	r := NewRouter()
	//sampleの作成
	createTestSample()
	//以下テスト
	authFunctions(t, r)
	userFunctions(t, r)
	postFunctions(t, r)
}
