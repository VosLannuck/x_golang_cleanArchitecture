package usecases

import(
  "log"
  "VosLannuck/x_golang_cleanArchitecture/src/domain"
)

type BookInteractor struct {
  BookRepository domain.BookRepository
}

func NewBookInteractor(repository domain.BookRepository) BookInteractor {
  return BookInteractor{repository}
}

func (interactor *BookInteractor) CreateBook(book domain.Book) error {
  err := interactor.BookRepository.SaveBook(book)
  if err != nil {
    log.Print(err.Error())
    return err
  }
  return nil
} 

func (interactor *BookInteractor) FindAll() ([]*domain.Book,error) {
  books, err := interactor.BookRepository.FindAll()
  
  if err != nil {
    log.Print(err.Error())
    return nil, err
  }

  return books, nil
}

