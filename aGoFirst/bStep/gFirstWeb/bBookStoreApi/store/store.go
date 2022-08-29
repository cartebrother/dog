package store

type Book struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Authors []string `json:"authors"`
	Press   string   `json:"Press"`
}

type Store interface {
	Create(book *Book) error
	Update(book *Book) error
	Get(id *string) (Book, error)
	GetAll() ([]Book, error)
	Delete(id string) error
}
