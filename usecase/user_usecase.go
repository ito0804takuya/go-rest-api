package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repositoty"
)

// ユーザがシステムを使うときに行う行動（ユースケース）
// インターフェース
type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

// 実装
type userUsecase struct {
	// （レポジトリでなく）レポジトリのインターフェース に依存させる
	ur repositoty.IUserRepository
}

// TODO: インターフェースを満たすメソッドの実装はまだ

// コンストラクタ（レポジトリのインターフェースをDI（注入））
func NewUserUsecase(ur repositoty.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}
