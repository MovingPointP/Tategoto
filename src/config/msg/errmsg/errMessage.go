package errmsg

var (
	DuplicateMailErr           = "既に同じメールアドレスが使われています。"
	EncryptionErr              = "パスワードの暗号化に失敗しました。"
	IncorrectMailOrPasswordErr = "メールアドレスまたはパスワードが一致しませんでした。"
	VerifyTokenErr             = "トークンの検証に失敗しました。"
	ExpiredTokenErr            = "トークンの期限が切れています。"
	ShouldLoginErr             = "ログインが必要な処理です。"
	IncorrectUserIDErr         = "ログインしたユーザーIDと異なります。"
	PostBindErr                = "postの割り当てに失敗しました。"
	GenerateIDErr              = "IDの生成に失敗しました。"
	NoDataErr                  = "データが存在しません。"
)
