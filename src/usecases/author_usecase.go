package usecases

import(
  "log"
  "VosLannuck/x_golang_cleanArchitecture/src/domain"
)

type AuthorInteractor struct {
  AuthorRepository domain.AuthorRepository
}

func NewAuthorInteractor(repository domain.AuthorRepository ) AuthorInteractor {
  return AuthorInteractor{repository}
}

func (interactor *AuthorInteractor) SaveAuthor(author domain.Author) error {
  err := interactor.AuthorRepository.SaveAuthor(author)
  if(err != nil) {
    log.Print(err.Error())
    return err
  }
  return nil
}



