package jwt

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"time"
)

var JwtKey = []byte("k6CswdUm77WKcbM68UQUuxVsHSpTCwgK")
var RcJwtKey = []byte("k6CswdUm77WKcbM68UQUuxVsHSpTCwgA")

type RedisJwtHandler struct {
	client        redis.Cmdable
	signingMethod jwt.SigningMethod
	rcExpiration  time.Duration
}

func NewJwtHandler(client redis.Cmdable) JwtHandler {
	return &RedisJwtHandler{
		client:        client,
		signingMethod: jwt.SigningMethodHS512,
		rcExpiration:  time.Hour * 24 * 30,
	}
}

type UserClaims struct {
	jwt.RegisteredClaims
	Uid       string
	Ssid      string
	UserAgent string
}

type RefreshClaims struct {
	jwt.RegisteredClaims
	Uid  string
	Ssid string
}

func (r *RedisJwtHandler) ClearToken(ctx *gin.Context) error {
	ctx.Header("x-jwt-token", "")
	ctx.Header("x-refresh-token", "")
	uc := ctx.MustGet("user").(UserClaims)

	return r.client.Set(ctx, fmt.Sprintf("user:ssid:%s", uc.Ssid), "", r.rcExpiration).Err()
}

// ExtractToken 提取token
func (r *RedisJwtHandler) ExtractToken(ctx *gin.Context) string {
	return ctx.GetHeader("Authorization")
}

// SetLoginToken 设置用户登录token
func (r *RedisJwtHandler) SetLoginToken(ctx *gin.Context, uid string) error {
	ssid := uuid.New().String()
	err := r.setRefreshToken(ctx, uid, ssid)
	if err != nil {
		return err
	}
	return r.SetJwtToken(ctx, uid, ssid)
}

// SetJwtToken 设置token
func (r *RedisJwtHandler) SetJwtToken(ctx *gin.Context, uid string, ssid string) error {
	uc := UserClaims{
		Uid:       uid,
		Ssid:      ssid,
		UserAgent: ctx.GetHeader("User-Agent"),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(r.signingMethod, uc)
	tokenStr, err := token.SignedString(JwtKey)
	if err != nil {
		return err
	}
	ctx.Header("x-jwt-token", tokenStr)
	return nil
}

// setRefreshToken 设置刷新token
func (r *RedisJwtHandler) setRefreshToken(ctx *gin.Context, uid string, ssid string) error {
	rc := RefreshClaims{
		Uid:  uid,
		Ssid: ssid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(r.rcExpiration)),
		},
	}
	token := jwt.NewWithClaims(r.signingMethod, rc)
	tokenStr, err := token.SignedString(RcJwtKey)
	if err != nil {
		return err
	}
	ctx.Header("x-refresh-token", tokenStr)
	return nil
}

func (r *RedisJwtHandler) CheckSession(ctx *gin.Context, ssid string) error {
	cnt, err := r.client.Exists(ctx, fmt.Sprintf("user:ssid:%s", ssid)).Result()
	if err != nil {
		return err
	}

	if cnt > 0 {
		return errors.New("无效token")
	}
	return nil
}
