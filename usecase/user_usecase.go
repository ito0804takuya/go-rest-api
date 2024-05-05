package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repositoty"
	"go-rest-api/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
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
	uv validator.IUserValidator
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return model.UserResponse{}, err
	}

	// パスワードを平文→ハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}

	// 登録するユーザーオブジェクト
	newUser := model.User{Email: user.Email, Password: string(hash)}
	// レポジトリにユーザー登録を依頼
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}

	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return "", err
	}

	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}

	// 入力されたパスワードが、登録されているパスワードと一致しているか検証
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		return "", err
	}

	// 認証したため、jwtトークンオブジェクトを生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	// 環境変数SECRETで、トークンに署名してトークン（文字列）を生成
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", nil
	}
	return tokenString, nil

}

// コンストラクタ（レポジトリのインターフェースをDI（注入））
func NewUserUsecase(ur repositoty.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}
