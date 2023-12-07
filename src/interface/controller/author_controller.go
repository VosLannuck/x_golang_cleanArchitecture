package controller

import (
	"VosLannuck/x_golang_cleanArchitecture/src/domain"
	"VosLannuck/x_golang_cleanArchitecture/src/usecases"
	"encoding/json"
	"log"
	"net/http"
)

type AuthorController struct{

  AuthorInteractor usecases.AuthorInteractor
    
}

func NewAuthorController(authorInteractor usecases.AuthorInteractor) *AuthorController{
  return &AuthorController{authorInteractor}
}

func (controller *AuthorController) Add(res http.ResponseWriter, req* http.Request) {
  res.Header().Set("Content-Type", "application/json")
  var author domain.Author 

  err_json := json.NewDecoder(req.Body).Decode(&author)
  
  if err_json != nil {
    log.Print(err_json.Error())
    var response map[string]string = map[string]string{
      "message" : "Error found in the server",
  } 
    res.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(res).Encode(response)
    return
  } 
  
  err_author := controller.AuthorInteractor.CreateAuthor(author)   

  if err_author != nil {
    log.Print(err_author.Error())
    res.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(res).Encode(map[string]string{"message" : err_author.Error()})
    return
  }

  res.WriteHeader(http.StatusOK)
  return

}




