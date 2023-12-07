package controller

import(
  "encoding/json"
  "net/http"
  "VosLannuck/x_golang_cleanArchitecture/src/domain"
  "VosLannuck/x_golang_cleanArchitecture/src/usecases"
)

type BookController struct {
  bookInteractor usecases.BookInteractor
}

func NewBookController(bookInteractor usecases.BookInteractor) *BookController {
  
  return &BookController{bookInteractor}
}


func (controller *BookController) Add(res http.ResponseWriter, req *http.Request) {
  res.Header().Set("Content-Type", "application/json")
  var book domain.Book
  err := json.NewDecoder(req.Body).Decode(&book)
  
  if err != nil {
    var response map[string]string = make(map[string]string)
    response["message"] = "Error found in the server"
    res.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(res).Encode(response)
    return
  }


  err1 := controller.bookInteractor.CreateBook(book)

  if err1 != nil {
    var response map[string]string = map[string]string {
      "message" : "Error found in the server",
    }
    res.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(res).Encode(response)
    return
  }
    res.WriteHeader(http.StatusOK)
    return
}

func (controller* BookController) FindAll(res http.ResponseWriter, req *http.Request){
  res.Header().Set("Content-Type", "application/json")
  books, err := controller.bookInteractor.FindAll()
  if err != nil {
    var response map[string]string = map[string]string{
      "message": "Could not find your files",
    }
    res.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(res).Encode(response)
    return 
  }
  res.WriteHeader(http.StatusOK)
  json.NewEncoder(res).Encode(books)
  return 
}


