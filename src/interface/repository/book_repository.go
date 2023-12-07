package repository

import (
  "log"
  "VosLannuck/x_golang_cleanArchitecture/src/domain"
)

type BookRepo struct {
  handler DBHandler
}

func NewBookRepository(dbHandler DBHandler) BookRepo {
  return BookRepo{dbHandler}
}

func (bookRepo BookRepo) FindAll()([]*domain.Book, error) {
  books, err := bookRepo.FindAll()
  if err != nil {
    log.Print(err.Error())
    return nil, err
  }
  return books, nil
}

func (bookRepo BookRepo) SaveBook(book domain.Book)(error) {
  err := bookRepo.SaveBook(book)
  if err != nil {
    log.Print(err.Error())
    return err
  }
  return nil
}


