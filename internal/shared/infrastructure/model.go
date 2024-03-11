package infrastructure

type IModel[E any] interface {
	Table() string
	ToEntity() E
	FromEntity(entity E) any
}
