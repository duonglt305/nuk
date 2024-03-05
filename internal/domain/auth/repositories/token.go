package auth_repositories

import authEntities "duonglt.net/internal/domain/auth/entities"

type TokenRepository interface {
	Create(uid uint64) (*authEntities.Token, error)
	Get(uid uint64) (*authEntities.Token, error)
}
