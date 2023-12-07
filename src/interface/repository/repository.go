package repository

import (
  "VosLannuck/x_golang_cleanArchitecture/src/domain"
)

type DBHandler interface {
  FindAllBooks() ([]*domain.Book, error)
  SaveBook(book domain.Book)(error)
  SaveAuthor(author domain.Author)(error)
}
