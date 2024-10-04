package TokenRepo

import (
	"log/slog"

	models "github.com/MoreWiktor/auth.sso.back/internal/domain"
	repository "github.com/MoreWiktor/auth.sso.back/internal/domain/repositories"
	"github.com/MoreWiktor/auth.sso.back/internal/entity"
	TableNameEnum "github.com/MoreWiktor/auth.sso.back/internal/enums/table-name"
)

type Repository[T any] interface {
	Find() T
	FindMany() T
	Create() T
	Update() T
	Delete() T
}

// type TokenRepo = Repository[models.Token]

type TokenRepo struct {
	log *slog.Logger
	db []string
}

func New(log *slog.Logger, db []string) Repository[entity.Token] {
	tokenRepo := repository.New[entity.Token](TableNameEnum.TOKENS, db, log)
	var IRepository Repository[entity.Token]

	IRepository = tokenRepo

	return IRepository
}
