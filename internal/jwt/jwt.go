package jwt

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwtlib "github.com/golang-jwt/jwt/v5"

	"my-auction-market-api/internal/config"
)

type Claims struct {
	UserID   uint64 `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwtlib.RegisteredClaims
}

type UserInfo struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var cfg config.JWTConfig

func Init(jwtCfg config.JWTConfig) {
	cfg = jwtCfg
}

func getSecret() string {
	if cfg.Secret == "" {
		panic("JWT secret is not configured")
	}
	return cfg.Secret
}

func getExpiration() time.Duration {
	if cfg.Expiration == 0 {
		return 24 * time.Hour
	}
	return cfg.Expiration
}

func GenerateToken(userID uint64, username, email string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		Email:    email,
		RegisteredClaims: jwtlib.RegisteredClaims{
			ExpiresAt: jwtlib.NewNumericDate(time.Now().Add(getExpiration())),
			IssuedAt:  jwtlib.NewNumericDate(time.Now()),
			NotBefore: jwtlib.NewNumericDate(time.Now()),
		},
	}

	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims)
	return token.SignedString([]byte(getSecret()))
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwtlib.ParseWithClaims(tokenString, &Claims{}, func(token *jwtlib.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtlib.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(getSecret()), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func ExtractUser(tokenString string) (*UserInfo, error) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return nil, err
	}

	return &UserInfo{
		ID:       claims.UserID,
		Username: claims.Username,
		Email:    claims.Email,
	}, nil
}

func ExtractUserFromHeader(authHeader string) (*UserInfo, error) {
	if authHeader == "" {
		return nil, errors.New("authorization header missing")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return nil, errors.New("invalid authorization header format")
	}

	return ExtractUser(parts[1])
}

func ExtractUserFromContext(c *gin.Context) (*UserInfo, error) {
	return ExtractUserFromHeader(c.GetHeader("Authorization"))
}
