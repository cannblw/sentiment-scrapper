package repository

type NewsItemRepository interface {
	Save(item *entity.NewsItem) (*entity.NewsItem, error)
	FindAll() ([]entity.NewsItem, error)
	FindById(id string) (*entity.NewsItem, error)
	Delete(item *entity.NewsItem) error
}
