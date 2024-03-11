package domain

type IRepository[E any, M any] interface {
	FindById(id uint64) (E, error)
	Save(entity *E) error
	Update(entity *E) error
	Delete(id uint64) error
}
