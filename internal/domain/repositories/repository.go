package Repository

import "log/slog"

type Repository struct {
	tableName string
	DB        []string
	log       *slog.Logger
}

type IRepository[T any] interface {
	Find() T
	FindMany() T
	Create() T
	Update() T
	Delete() T
}

func New[T any](tableName string, db []string, log *slog.Logger) *Repository {
	return &Repository{
		tableName,
		db,
		log,
	}
}

func (r *Repository) Find() string {
	return "asd";
}
func (r *Repository) FindMany() string {
	return "asd"
}
func (r *Repository) Create() string {
	return "asd"
}
func (r *Repository) Update() string {
	return "asd"
}
func (r *Repository) Delete() string {
	return "asd"
}
