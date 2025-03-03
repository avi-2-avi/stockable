package repositories

type BaseRepository[T any] interface {
	Create(entity *T) error
	GetByID(id uint) (*T, error)
}
