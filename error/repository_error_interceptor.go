package error

import (
	"context"
	"errors"

	log "github.com/iamgenii/logs"
	"github.com/jinzhu/gorm"
)

type RepoErrorInterceptor interface {
	ErrorMapper(ctx context.Context, err error) *IMGNError
}
type repoErrorInterceptor struct {
}

var ErrDuplicateEntry = errors.New("DUPLICATE ENTRY ERROR")

func NewRepoErrorInterceptor() RepoErrorInterceptor {
	return &repoErrorInterceptor{}
}

func (r repoErrorInterceptor) ErrorMapper(ctx context.Context, err error) *IMGNError {
	log.Logger(ctx).Debug("RepoErrorInterceptor.ErrorMapper: In Error Mapper function")
	switch err {
	case gorm.ErrRecordNotFound:
		return ErrorRecordNotFound
	case ErrDuplicateEntry:
		return ErrorDuplicateEntry
	default:
		return InternalServerErrorFunc(err.Error())
	}
}
