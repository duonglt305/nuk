package auth_repositories

import auth_entities "duonglt.net/internal/domain/auth/entities"

type TokenRepository interface {
	CreateToken(uid uint64) (*auth_entities.Token, error)
	Get(uid uint64) (*auth_entities.Token, error)
}
