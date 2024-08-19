package db

import (
	"fmt"
	"reflect"
	"strings"

	"duonglt.net/pkg/utils"
	"github.com/jmoiron/sqlx"
)

type IRepository[E any, M any] interface {
	FindById(id uint64) (E, error)
	Find(spec Specification) ([]E, error)
	FindOne(spec Specification) (E, error)
	Save(entity *E) error
	Update(entity *E) error
	Delete(spec Specification) error
}

type IModel[E any] interface {
	Table() string
	ToEntity() E
	FromEntity(entity E) any
}

// Repository is a generic repository
type Repository[M IModel[E], E any] struct {
	db *sqlx.DB
}

// NewRepository creates a new repository
func NewRepository[M IModel[E], E any](db *sqlx.DB) Repository[M, E] {
	return Repository[M, E]{db: db}
}

// GetFields gets the fields of a model
func (rep Repository[M, E]) getFields(model M) []string {
	val := reflect.TypeOf(model)
	fields := make([]string, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		fields[i] = val.Field(i).Tag.Get("db")
	}
	return fields
}

// FindOne finds a model by a specification
func (rep Repository[M, E]) FindOne(spec Specification) (E, error) {
	var model M
	sql := fmt.Sprintf("SELECT * FROM %s WHERE %s", model.Table(), spec.GetQuery())
	if err := rep.db.Get(&model, sql, spec.GetValues()...); err != nil {
		return *new(E), err
	}
	return model.ToEntity(), nil
}

// Find finds models by a specification
func (rep Repository[M, E]) Find(spec Specification) ([]E, error) {
	var models []M
	sql := fmt.Sprintf("SELECT * FROM %s WHERE %s", models[0].Table(), spec.GetQuery())
	if err := rep.db.Select(&models, sql, spec.GetValues()...); err != nil {
		return nil, err
	}
	entities := make([]E, len(models))
	for i, model := range models {
		entities[i] = model.ToEntity()
	}
	return entities, nil
}

// FindById finds a model by its id
func (rep Repository[M, E]) FindById(id uint64) (E, error) {
	return rep.FindOne(Eq("id", id))
}

// Save saves a model
func (rep Repository[M, E]) Save(entity *E) error {
	var model M
	model = model.FromEntity(*entity).(M)
	fields := rep.getFields(model)
	namedFields := utils.Map(func(field string) string {
		return fmt.Sprintf(":%s", field)
	}, fields)
	sql := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		model.Table(),
		strings.Join(fields, ", "),
		strings.Join(namedFields, ", "),
	)
	if _, err := rep.db.NamedExec(sql, model); err != nil {
		return err
	}
	*entity = model.ToEntity()
	return nil
}

// Update updates a model
func (rep Repository[M, E]) Update(entity *E) error {
	var model M
	model = model.FromEntity(*entity).(M)
	fields := rep.getFields(model)
	namedFields := utils.Map(func(field string) string {
		return fmt.Sprintf("%s = :%s", field, field)
	}, fields)

	sql := fmt.Sprintf(
		"UPDATE %s SET %s WHERE id = :id",
		model.Table(),
		strings.Join(namedFields, ", "),
	)
	fmt.Println(sql)
	if _, err := rep.db.NamedExec(sql, model); err != nil {
		return err
	}
	return nil
}

// Delete deletes models by a specification
func (rep Repository[M, E]) Delete(spec Specification) error {
	var model M
	sql := fmt.Sprintf("DELETE FROM %s WHERE %s", model.Table(), spec.GetQuery())
	if _, err := rep.db.Exec(sql, spec.GetValues()...); err != nil {
		return err
	}
	return nil
}
