package domain

type Book struct {
  Title string `json:"title"`
  Author string `json:"author"`
  Price float32 `json:"price"`
  Quantity int64 `json:"quantity"`

}

type BookRepository interface {
  SaveBook(book Book) error
  FindAll() ([]*Book, error)
}

