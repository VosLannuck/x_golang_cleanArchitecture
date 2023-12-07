package repository

import (
  "log"
  "VosLannuck/x_golang_cleanArchitecture/src/domain"
)

type AuthorRepo struct {
  dbHandler DBHandler
}


func NewAuthorRepo(dbHandler DBHandler) AuthorRepo {
  
  return AuthorRepo{dbHandler}
}

func (interactor AuthorRepo) SaveAuthor(author domain.Author ) error {
  err := interactor.dbHandler.SaveAuthor(author)

  if err != nil {
    log.Println(err.Error())
    return err
  }
  
  return nil
}
