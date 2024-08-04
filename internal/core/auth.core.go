package core

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/needl3/goreact-template/internal/domain"
	"github.com/needl3/goreact-template/internal/repository"
)

type UserInfo struct {
	Id     string
	Fname  string `json:"fname"`
	Lname  string `json:"lname"`
	Avatar string `json:"avatar"`
	Email  string `json:"email"`
}

type GoogleResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
	IdToken     string `json:"id_token"`
}

type GoogleUserResponse struct {
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	Picture       string `json:"picture"`
	EmailVerified bool   `json:"email_verified"`
}

type AuthCore struct {
	clientId     string
	clientSecret string
	redirectUri  string
	db           domain.Database
}

func NewAuthCore(conn *pgxpool.Pool) *AuthCore {
	return &AuthCore{
		db:           repository.NewPostgresRepository(conn),
		clientId:     os.Getenv("GOOGLE_CLIENT_ID"),
		clientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		redirectUri:  os.Getenv("GOOGLE_REDIRECT_URI"),
	}
}

func (a *AuthCore) GenerateAuthUrl() (string, error) {
	if a.clientId == "" || a.clientSecret == "" || a.redirectUri == "" {
		return "", errors.New("Auth credentials not in environment")
	}
	scopes := strings.Join([]string{
		"https://www.googleapis.com/auth/userinfo.email",
	}, " ")
	return fmt.Sprintf("https://accounts.google.com/o/oauth2/v2/auth?redirect_uri=%s&prompt=consent&response_type=code&client_id=%s&scope=%s", a.redirectUri, a.clientId, scopes), nil
}

func (a *AuthCore) Signup(ctx context.Context, code string) (string, error) {
	userInfo, err := a.fetchUserInfoFromGoogle(code)
	if err != nil || userInfo == nil {
		return "", err
	}

	userFromSystem := a.createOrUpsertUser(ctx, userInfo)
	if userFromSystem == nil {
		return "", errors.New("Failed to fetch user from system")
	}

	accessToken, err := a.createAccessToken(userFromSystem)

	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (a *AuthCore) fetchUserInfoFromGoogle(code string) (*GoogleUserResponse, error) {
	oauthUrl := "https://oauth2.googleapis.com/token"
	data := url.Values{}
	data.Set("client_id", a.clientId)
	data.Set("client_secret", a.clientSecret)
	data.Set("code", code)
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", a.redirectUri)

	client := http.Client{}
	tokenRequest, _ := http.NewRequest(http.MethodPost, oauthUrl, strings.NewReader(data.Encode()))
	tokenRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rawResp, err := client.Do(tokenRequest)
	if rawResp.StatusCode != 200 {
		return nil, errors.New("Failed to fetch user token from google")
	}
	if err != nil {
		return nil, err
	}

	resp, err := io.ReadAll(rawResp.Body)
	if err != nil {
		return nil, err
	}

	var userInfo GoogleResponse
	err = json.Unmarshal(resp, &userInfo)
	if err != nil {
		return nil, err
	}

	userRequest, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://www.googleapis.com/oauth2/v3/userinfo?alt=json&access_token=%s", userInfo.AccessToken), nil)
	userRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", userInfo.IdToken))
	userResponse, err := client.Do(userRequest)
	if err != nil {
		return nil, err
	}
	if userResponse.StatusCode != 200 {
		return nil, errors.New("Failed to fetch user info from google")
	}

	resp, err = io.ReadAll(userResponse.Body)
	if err != nil {
		return nil, err
	}

	var user GoogleUserResponse
	err = json.Unmarshal(resp, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *AuthCore) createOrUpsertUser(ctx context.Context, userInfo *GoogleUserResponse) *domain.User {
	user, err := a.db.FindUserByEmail(ctx, userInfo.Email)
	if err == nil {
		return user
	}
	fmt.Println(err)

	user, err = a.db.CreateUser(ctx, "John", "Doe", userInfo.Picture, userInfo.Email)
	if err == nil {
		return user
	}
	fmt.Println(err)
	return nil
}

type JwtType struct {
	jwt.StandardClaims
	AuthenticatedUserRequestValues
}
type AuthenticatedUserRequestValues struct {
	Email string `json:"email"`
	Id    string `json:"id"`
}

func (a *AuthCore) createAccessToken(user *domain.User) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, JwtType{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		AuthenticatedUserRequestValues: AuthenticatedUserRequestValues{
			Id:    user.Id,
			Email: user.Email,
		},
	}).SignedString([]byte(os.Getenv("APP_SECRET")))
}
