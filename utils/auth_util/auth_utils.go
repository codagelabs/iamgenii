package auth_util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	imgnErr "github.com/iamgenii/error"
	log "github.com/iamgenii/logs"
	"github.com/iamgenii/utils/crypto_utils"
	"github.com/square/go-jose/v3"
)

type AuthUtils interface {
	EncryptIdToken(ctx context.Context, payload IdToken) (encryptedToken string, imgnError *imgnErr.IMGNError)
	DecodeEncryptIdToken(ctx context.Context) (idToken IdToken, imgnError *imgnErr.IMGNError)
	DecodeEncryptIdToken1(ctx context.Context) (idToken string, imgnError *imgnErr.IMGNError)
}
type authUtils struct {
	jwtUtils    JwtUtils
	cryptoUtils crypto_utils.CryptoUtils
}

func (utils authUtils) EncryptIdToken(ctx context.Context, payload IdToken) (encryptedToken string, imgnError *imgnErr.IMGNError) {
	log.Logger(ctx).Debug("AuthUtils.EncryptIdToken: in id token encryption.")
	token, jwtEncodeErr := utils.jwtUtils.EncodeJwtToken(ctx, payload)
	if jwtEncodeErr != nil {
		log.Logger(ctx).Error("AuthUtils.EncryptIdToken: Error in encode jwt token.")
		return "", imgnErr.InternalServerErrorFunc(jwtEncodeErr.Error())
	}
	fmt.Println(token)
	publicKey, err := utils.cryptoUtils.GetPublicKey(ctx)
	if err != nil {
		log.Logger(ctx).Error("AuthUtils.EncryptIdToken: Error in get public key.")
		return "", imgnErr.InternalServerErrorFunc(err.Error())
	}
	encrypter, err := utils.cryptoUtils.GetEncrypter(ctx, publicKey)
	if err != nil {
		log.Logger(ctx).Error("AuthUtils.EncryptIdToken: Error in get  encrypter key.")
		return "", imgnErr.InternalServerErrorFunc(err.Error())
	}

	plaintext, err := json.Marshal(token)
	if err != nil {
		log.Logger(ctx).Error("AuthUtils.EncryptIdToken: Error in token marshaling Marshal.")
		return "", imgnErr.InternalServerErrorFunc(err.Error())
	}

	encObj, err := encrypter.Encrypt(plaintext)
	if err != nil {
		log.Logger(ctx).Error("AuthUtils.EncryptIdToken: Error in token encrypter.")
		return "", imgnErr.InternalServerErrorFunc(err.Error())
	}
	encriptedJwtTocken, err := encObj.CompactSerialize()
	if err != nil {
		log.Logger(ctx).Error("AuthUtils.EncryptIdToken: Error in token encrypter.")
		return "", imgnErr.InternalServerErrorFunc(err.Error())
	}

	return encriptedJwtTocken, nil

}

func (utils authUtils) DecodeEncryptIdToken1(ctx context.Context) (idToken string, imgnError *imgnErr.IMGNError) {
	log.Logger(ctx).Debug("AuthUtils.EncryptIdToken: in id token encryption.")
	token := ctx.Value("id_token")
	if token == nil {
		log.Logger(ctx).Error("AuthUtils.EncryptIdToken: Error in encode jwt token.")
		return "", imgnErr.InternalServerErrorFunc("invalid token")
	}
	encrypted, err := jose.ParseEncrypted(token.(string))
	if err != nil {
		log.Logger(ctx).Error("AuthUtils.EncryptIdToken: Error in get public key.")
		return "", imgnErr.InternalServerErrorFunc(err.Error())
	}
	privateKey, err := utils.cryptoUtils.GetPrivateKey(ctx)
	if err != nil {
		log.Logger(ctx).Error("AuthUtils.EncryptIdToken: Error in get public key.")
		return "", imgnErr.InternalServerErrorFunc(err.Error())
	}

	decrypted, err := encrypted.Decrypt(privateKey)
	if err != nil {
		log.Logger(ctx).Error("AuthUtils.EncryptIdToken: Error in get public key.")
		return "", imgnErr.InternalServerErrorFunc(err.Error())
	}
	return bytes.NewBuffer(decrypted).String(), nil //string(decrypted),nil
}

func (utils authUtils) DecodeEncryptIdToken(ctx context.Context) (idToken IdToken, imgnError *imgnErr.IMGNError) {
	log.Logger(ctx).Debug("AuthUtils.EncryptIdToken: in id token encryption.")
	token := ctx.Value("id_token")
	if token == nil {
		log.Logger(ctx).Error("AuthUtils.EncryptIdToken: Error in encode jwt token.")
		return IdToken{}, imgnErr.InternalServerErrorFunc("invalid token")
	}
	encrypted, err := jose.ParseEncrypted(token.(string))
	if err != nil {
		log.Logger(ctx).Error("AuthUtils.EncryptIdToken: Error in get public key.")
		return IdToken{}, imgnErr.InternalServerErrorFunc(err.Error())
	}
	privateKey, err := utils.cryptoUtils.GetPrivateKey(ctx)
	if err != nil {
		log.Logger(ctx).Error("AuthUtils.EncryptIdToken: Error in get public key.")
		return IdToken{}, imgnErr.InternalServerErrorFunc(err.Error())
	}

	decrypted, err := encrypted.Decrypt(privateKey)
	if err != nil {
		log.Logger(ctx).Error("AuthUtils.EncryptIdToken: Error in get public key.")
		return IdToken{}, imgnErr.InternalServerErrorFunc(err.Error())
	}

	idToken, err = utils.jwtUtils.DecodeJwtToken(ctx, string(decrypted))
	if err != nil {
		log.Logger(ctx).Error("AuthUtils.EncryptIdToken: Error in token marshaling Marshal.")
		return IdToken{}, imgnErr.InternalServerErrorFunc(err.Error())
	}
	return idToken, nil
}

func NewAuthUtils(jwtUtils JwtUtils, cryptoUtils crypto_utils.CryptoUtils) AuthUtils {
	return authUtils{
		jwtUtils:    jwtUtils,
		cryptoUtils: cryptoUtils,
	}
}
