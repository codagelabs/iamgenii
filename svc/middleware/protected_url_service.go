package middleware

import (
	"strings"
)

type ProtectedUrlService interface {
	IsUrlProtected(inputUrl string) bool
}

type protectedUrlService struct {
}

func NewProtectedUrlService() ProtectedUrlService {
	return &protectedUrlService{}
}

func (protectedUrlService protectedUrlService) IsUrlProtected(inputUrl string) bool {
	var protectedUrls = []string{
		"logout",
	}
	for _, url := range protectedUrls {
		if strings.Contains(inputUrl, url) {
			return true
		}
	}
	return false
}
