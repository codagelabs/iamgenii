package auth_util

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/iamgenii/configs"
	customError "github.com/iamgenii/error"
	log "github.com/iamgenii/logs"
	"github.com/iamgenii/utils/crypto_utils"
	"github.com/dgrijalva/jwt-go"
)

type JwtUtils interface {
	DecodeJwtToken(ctx context.Context, jwtToken string) (IdToken, error)
	EncodeJwtToken(ctx context.Context, payload IdToken) (string, error)
}

type jwtUtils struct {
	jwtConfig   configs.JwtConfig
	cryptoUtils crypto_utils.CryptoUtils
}

func NewJwtUtils(jwtConfig configs.JwtConfig,
	cryptoUtils crypto_utils.CryptoUtils,
) JwtUtils {
	return &jwtUtils{
		jwtConfig:   jwtConfig,
		cryptoUtils: cryptoUtils,
	}
}
func (jwtUtils jwtUtils) DecodeJwtToken(ctx context.Context, jwtToken string) (IdToken, error) {
	tokenParts := strings.Split(jwtToken, ".")
	if len(tokenParts) != 3 {
		log.Logger(ctx).Error("JwtUtils.DecodeJwtToken: Token is in invalid format. ")
		return IdToken{}, errors.New(" Token is in invalid format")
	}
	payload := tokenParts[1]
	bytes, decodeErr := jwt.DecodeSegment(payload)
	if decodeErr != nil {
		return IdToken{}, decodeErr
	}
	var idToken IdToken
	unmarshalErr := json.Unmarshal(bytes, &idToken)
	if unmarshalErr != nil {
		return IdToken{}, unmarshalErr
	}
	return idToken, nil

}

func (jwtUtils jwtUtils) EncodeJwtToken(ctx context.Context, token IdToken) (string, error) {
	claims := NewJWTClaims(token)
	fmt.Println("claims: ", claims)
	expirationTime := time.Now().Add(24 * time.Hour)
	claims.ExpiresAt = expirationTime.Unix()
	tokenDetails := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	key, err := jwtUtils.cryptoUtils.GetJWTPrivateKey(ctx)
	if err != nil {
		log.Logger(ctx).Error("JwtUtils.EncodeJwtToken: Error in getting public key. ", err)
		return "", err
	}
	idToken, err := tokenDetails.SignedString(key)
	if err != nil {
		log.Logger(ctx).Error("JwtUtils.EncodeJwtToken:Error in token signed in. ", err)
		return "IdToken{}", customError.ErrInvalidLoginDetails
	}
	return idToken, nil

}
