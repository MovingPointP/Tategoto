package apitest

import (
	"context"
	"tategoto/model"
	"tategoto/pkg/auth"
	"tategoto/pkg/ulid"
)

var SampleUserHoge *model.User
var SampleUserFuga *model.User
var TokenHoge string
var TokenFuga string
var SamplePostHello *model.Post
var SamplePostWorld *model.Post

func createTestSample() {
	ctx := context.Background()
	SampleUserHoge, SampleUserFuga, TokenHoge, TokenFuga = createUserSample(ctx)
	SamplePostHello, SamplePostWorld = createPostSample(ctx)
}

func createUserSample(ctx context.Context) (*model.User, *model.User, string, string) {
	name := "sample"
	//UserHoge
	password, _ := auth.EncryptPassword("hoge")
	user := &model.User{
		Mail:     "hoge@mail.com",
		Password: password,
		Name:     name,
	}
	userHoge, _ := serviceInstance.SignUp(ctx, user)
	tokenHoge, _ := auth.CreateUserJWT(userHoge.ID)

	//UserFuga
	password, _ = auth.EncryptPassword("fuga")
	user = &model.User{
		Mail:     "fuga@mail.com",
		Password: password,
		Name:     name,
	}
	userFuga, _ := serviceInstance.SignUp(ctx, user)
	tokenFuga, _ := auth.CreateUserJWT(userFuga.ID)
	return userHoge, userFuga, tokenHoge, tokenFuga
}

func createPostSample(ctx context.Context) (*model.Post, *model.Post) {
	//PostHello
	id, _ := ulid.CreateULID()
	post := &model.Post{
		ID:      id,
		Content: "hello",
		UserID:  SampleUserHoge.ID,
	}
	postHello, _ := serviceInstance.CreatePost(ctx, post)

	//PostWorld
	id, _ = ulid.CreateULID()
	post = &model.Post{
		ID:      id,
		Content: "world",
		UserID:  SampleUserHoge.ID,
	}
	postWorld, _ := serviceInstance.CreatePost(ctx, post)

	return postHello, postWorld
}
