package infrastructure

import (
	"fmt"
	"reflect"
	"strings"

	"duonglt.net/pkg/utils"
	"github.com/jmoiron/sqlx"
)

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

// FindById finds a model by its id
func (rep Repository[M, E]) FindById(id uint64) (E, error) {
	var model M
	sql := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", model.Table())
	if err := rep.db.Get(&model, sql, id); err != nil {
		return *new(E), err
	}
	return model.ToEntity(), nil
}

// Save saves a model
func (rep Repository[M, E]) Save(entity *E) error {
	var model M
	model = model.FromEntity(*entity).(M)
	fields := rep.getFields(model)
	namedFields := utils.Map[string, string](func(field string) string {
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
	namedFields := utils.Map[string, string](func(field string) string {
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

// Delete deletes a model by its id
func (rep Repository[M, E]) Delete(id uint64) error {
	var model M
	sql := fmt.Sprintf("DELETE FROM %s WHERE id = $1", model.Table())
	if _, err := rep.db.Exec(sql, id); err != nil {
		return err
	}
	return nil
}
