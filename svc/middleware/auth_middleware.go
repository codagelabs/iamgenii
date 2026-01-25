package middleware

import (
	"context"
	"fmt"
	"strings"

	authUtils "github.com/iamgenii/utils/auth_util"
	"github.com/iamgenii/utils/crypto_utils"
	"github.com/iamgenii/utils/http_utils"
	"github.com/dgrijalva/jwt-go"

	"net/http"

	imgnErr "github.com/iamgenii/error"
	log "github.com/iamgenii/logs"
)

type AuthMiddleware interface {
	Middleware(next http.Handler) http.Handler
}

type authMiddleware struct {
	authUtis            authUtils.AuthUtils
	jwtUtils            authUtils.JwtUtils
	cryptoUtils         crypto_utils.CryptoUtils
	httpWriter          http_utils.HTTPWriter
	protectedUrlService ProtectedUrlService
}

func (a authMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		log.Logger(ctx).Debug("AuthMiddleware: In auth middleware")
		url := req.URL.Path
		if a.protectedUrlService.IsUrlProtected(url) {
			if cookie, err := req.Cookie("id_token"); err == nil {
				ctx = context.WithValue(ctx, "id_token", cookie.Value)
				jweToken, imgnError := a.authUtis.DecodeEncryptIdToken1(ctx)
				if imgnError != nil {
					log.Logger(ctx).Error("AuthMiddleware.Middleware: Error In idtoken decode.")
				}
				jweToken = strings.ReplaceAll(jweToken, "\"", "")
				fmt.Println("jweToken ::: ", jweToken)
				key, err2 := a.cryptoUtils.GetJWTPublicKey(ctx)
				if err2 != nil {
					fmt.Println("fdsfsffdsfdsfdsffsff  fsfsfsff fsfs dfdf dfsf", err2)
					return
				}
				token, err := jwt.Parse(jweToken, func(token *jwt.Token) (interface{}, error) {
					return key, nil
				})
				fmt.Println("err ::: ", err)
				fmt.Printf("idtoken \"%+v\\n\" ::: ", fmt.Sprint(token))
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					ctx := context.WithValue(ctx, "props", claims)
					//props, _ := req.Context().Value("props").(jwt.MapClaims)
					next.ServeHTTP(w, req.WithContext(ctx))
				} else {
					fmt.Println(err)
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Unauthorized"))
				}
				if token.Valid {
					next.ServeHTTP(w, req.WithContext(ctx))
				} else if ve, ok := err.(*jwt.ValidationError); ok {
					if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
						fmt.Println("Timing is everything")
						a.httpWriter.WriteHTTPError(w, imgnErr.InvalidToTokenErrorFunc(err.Error()))
					}
					a.httpWriter.WriteHTTPError(w, imgnErr.InternalServerErrorFunc(err.Error()))
					fmt.Println("Couldn't handle this idtoken:", err)

				} else {
					fmt.Println("Couldn't handle this idtoken:", err)
					a.httpWriter.WriteHTTPError(w, imgnErr.InternalServerErrorFunc(err.Error()))
				}

			} else {
				a.httpWriter.WriteHTTPError(w, imgnErr.InternalServerErrorFunc(err.Error()))
				return
			}
		}
		next.ServeHTTP(w, req)
	})
}

func NewAuthMiddleware(authUtis authUtils.AuthUtils, jwtUtils authUtils.JwtUtils, cryptoUtils crypto_utils.CryptoUtils, httpWriter http_utils.HTTPWriter, protectedUrlService ProtectedUrlService) AuthMiddleware {
	return &authMiddleware{authUtis: authUtis, jwtUtils: jwtUtils, cryptoUtils: cryptoUtils, httpWriter: httpWriter, protectedUrlService: protectedUrlService}
}
