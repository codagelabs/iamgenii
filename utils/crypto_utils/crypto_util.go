package crypto_utils

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/iamgenii/configs"
	imgnErr "github.com/iamgenii/error"
	log "github.com/iamgenii/logs"
	"github.com/dgrijalva/jwt-go"
	"github.com/square/go-jose/v3"
	"golang.org/x/crypto/bcrypt"
)

type CryptoUtils interface {
	Encrypt(ctx context.Context, plainText string) (string, *imgnErr.IMGNError)
	Decrypt(ctx context.Context, cipherText string) (string, *imgnErr.IMGNError)
	GetEncrypter(ctx context.Context, publicKey *rsa.PublicKey) (jose.Encrypter, error)
	GetPublicKey(ctx context.Context) (*rsa.PublicKey, error)
	GetJWTPublicKey(ctx context.Context) (*rsa.PublicKey, error)
	GetJWTPrivateKey(ctx context.Context) (*rsa.PrivateKey, error)
	GetPrivateKey(ctx context.Context) (*rsa.PrivateKey, error)
}
type cryptoUtils struct {
	authConfig configs.AuthConfig
}

func (crypto cryptoUtils) GetJWTPublicKey(ctx context.Context) (*rsa.PublicKey, error) {
	keyData, e := ioutil.ReadFile("utils/crypto_utils/jwtRS256.key.pub")
	if e != nil {
		panic(e.Error())
	}
	key, e := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if e != nil {
		panic(e.Error())
	}
	return key, nil
}

func (crypto cryptoUtils) GetJWTPrivateKey(ctx context.Context) (*rsa.PrivateKey, error) {
	keyData, e := ioutil.ReadFile("utils/crypto_utils/jwtRS256.key")
	if e != nil {
		panic(e.Error())
	}
	key, e := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if e != nil {
		panic(e.Error())
	}
	return key, nil
}

func (crypto cryptoUtils) Encrypt(ctx context.Context, plainText string) (string, *imgnErr.IMGNError) {
	log.Logger(ctx).Debug("CryptoUtils.Encrypt: In encryption plain.")
	cipherText, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.MinCost)
	if err != nil {
		log.Logger(ctx).Error("CryptoUtils.Encrypt: Error in create admin: ", err)
		return "", imgnErr.InternalServerErrorFunc(err.Error())
	}
	return string(cipherText), nil
}

func (crypto cryptoUtils) Decrypt(ctx context.Context, cipherText string) (string, *imgnErr.IMGNError) {
	panic("implement me")
}

func (crypto cryptoUtils) GetEncrypter(ctx context.Context, publicKey *rsa.PublicKey) (jose.Encrypter, error) {
	encrypter, err := jose.NewEncrypter(jose.A128GCM, jose.Recipient{Algorithm: jose.RSA_OAEP, Key: publicKey}, nil)
	if err != nil {
		return nil, err
	}
	return encrypter, err
}

var publicKeyData = "-----BEGIN RSA PUBLIC KEY-----\nMIIBCgKCAQEAh7BcU+ZMHqXoYvNFfX1WsuTvaEJzgq31N4brjl09gLAPW6hbr/Jf\nLcJU+KU9tglr9MT5prMXMdSDAM4DXUOusV5C2EJoh5EiSpCrWQXGAfCPV5YYiauu\nISh6KStyZ/jL1fWA2PhuEBkOdYLeQCKRdhjORjqT9GftjuTXRLf70ji/XPal+qeZ\n9TFyFWVP8UZH1U+5AL1qq1aGRrPwoIVjSMMIP20+ONpfFGOCTITrImpL4eq0LLZl\n7n/+N99ijsA5Idr+c2Rwh6tfJIz9FvZ08TjQOiXp7gA+KEYtvFjmBxw54X6ipiP7\n7bbIy3B6EWaR16UceIhxExsi0vFcdt/5JwIDAQAB\n-----END RSA PUBLIC KEY-----"

func (crypto cryptoUtils) GetPublicKey(ctx context.Context) (*rsa.PublicKey, error) {
	log.Logger(ctx).Debug("CryptoUtils.GetPrivateKey: In get public key function.")
	publicKeyData := crypto.authConfig.GetPublicKey()
	fmt.Println(publicKeyData)
	data, _ := pem.Decode([]byte(publicKeyData))
	if data == nil {
		log.Logger(ctx).Error("CryptoUtils.Encrypt: Error In pem.Decode() : Public key not found.")
		return nil, errors.New("public key not found")
	}
	fmt.Println(data)
	publicKey, parsingError := x509.ParsePKCS1PublicKey(data.Bytes)
	if parsingError != nil {
		log.Logger(ctx).Error("CryptoUtils.Encrypt: Error in public key not parsing: ", parsingError)
		return nil, parsingError
	}
	return publicKey, nil
}
func (crypto cryptoUtils) GetPrivateKey(ctx context.Context) (*rsa.PrivateKey, error) {
	log.Logger(ctx).Debug("CryptoUtils.GetPrivateKey: In get private key function.")
	privateKeyData := crypto.authConfig.GetPrivateKey()
	fmt.Println("Private Key : ", privateKeyData)
	data, _ := pem.Decode([]byte(privateKeyData))
	if data == nil {
		log.Logger(ctx).Error("CryptoUtils.Encrypt: Error In pem.Decode() : Private key not found.")
		return nil, errors.New("private key not found")
	}
	privateKey, parsingError := x509.ParsePKCS1PrivateKey(data.Bytes)
	if parsingError != nil {
		return nil, parsingError
	}
	return privateKey, nil

}

func NewCryptoUtils(authConfig configs.AuthConfig) CryptoUtils {
	return &cryptoUtils{authConfig: authConfig}
}
