package msg

var (
	DuplicateMailErr           = "既に同じメールアドレスが使われています。"
	EncryptionErr              = "パスワードの暗号化に失敗しました。"
	IncorrectMailOrPasswordErr = "メールアドレスまたはパスワードが一致しませんでした。"
	VerifyTokenErr             = "トークンの検証に失敗しました。"
	ExpiredTokenErr            = "トークンの期限が切れています。"
)
