package authRepositories

import authEntities "duonglt.net/internal/auth/domain/entities"

type ITokenRepository interface {
	Create(uid uint64) (*authEntities.Token, error)
	Get(uid uint64) (*authEntities.Token, error)
}
