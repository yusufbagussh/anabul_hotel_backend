package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

// JWTService is a contract of what jwtService can do
type JWTService interface {
	GenerateToken(userID string, email string, userRole string, hotelID string) string
	ValidateToken(token string) (*jwt.Token, error)
	CheckExpiredToken(token string) (*jwt.Token, error, *jwtCustomClaim)
	UpdateToken(claims *jwtCustomClaim) (string, error)
}

// jwtCustomClaim untuk membuat klaim token jwtnya berdasarkan id user dan minimum standart jwt
type jwtCustomClaim struct {
	UserID  string `json:"user_id"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	HotelID string `json:"hotel_id"`
	jwt.StandardClaims
}

// jwtService is func identity of jwt user and his secret key
type jwtService struct {
	issuer    string
	secretKey string
}

// NewJWTService method is creates a new instance of JWTService
func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "ydhnwb",
		secretKey: getSecretKey(),
	}
}

// getSecretKey is func to get secret key
func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "ydhnwb"
	}
	return secretKey
}

// GenerateToken is to generate token
func (j *jwtService) GenerateToken(UserID string, Email string, Role string, HotelID string) string {
	claims := &jwtCustomClaim{
		UserID,
		Email,
		Role,
		HotelID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			//ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:   j.issuer,
			IssuedAt: time.Now().Unix(),
		},
	}
	//melakukan signing dengan algoritma HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//melakukan convert token ke dalam string
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

// ValidateToken untuk melakukan validasi token user, jika benar akan dikirimkan secret key
func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}

func (j *jwtService) CheckExpiredToken(token string) (*jwt.Token, error, *jwtCustomClaim) {
	claims := &jwtCustomClaim{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	return tkn, err, claims
}

func (j *jwtService) UpdateToken(claims *jwtCustomClaim) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24)

	claims.ExpiresAt = expirationTime.Unix()

	tokenWithClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, errToken := tokenWithClaim.SignedString([]byte(j.secretKey))

	return tokenString, errToken
}
