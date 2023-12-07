package main

import (
  "fmt"
  "log"
  "net/http"
  "VosLannuck/x_golang_cleanArchitecture/src/infrastructure/db"
  "VosLannuck/x_golang_cleanArchitecture/src/infrastructure/router"
  "VosLannuck/x_golang_cleanArchitecture/src/usecases"
  "VosLannuck/x_golang_cleanArchitecture/src/interface/controller"
  "VosLannuck/x_golang_cleanArchitecture/src/interface/repository"
)

var httpRouter router.Router = router.NewMuxRouter()
var dbHandler db.DBHandler 

func getBookController() controller.BookController {
  var bookRepository repository.BookRepo = repository.NewBookRepository(dbHandler)
  var bookInteractor usecases.BookInteractor = usecases.NewBookInteractor(bookRepository)
  var bookController *controller.BookController = controller.NewBookController(bookInteractor)
  return *bookController
}

func getAuthorController() controller.AuthorController {
  var authorRepository repository.AuthorRepo = repository.NewAuthorRepo(dbHandler)
  var authorInteractor usecases.AuthorInteractor = usecases.NewAuthorInteractor(authorRepository)
  var authorController *controller.AuthorController = controller.NewAuthorController(authorInteractor)
  return *authorController
}

func main() {

  
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "App is up and running..")
	})
	var err error
	dbHandler, err = db.NewDBHandler("mongodb://localhost:27017", "bookstore")
	if err != nil {
		log.Println("Unable to connect to the DataBase")
		return
	}
	bookController := getBookController()
	authorController := getAuthorController()
	httpRouter.POST("/book/add", bookController.Add)
	httpRouter.GET("/book", bookController.FindAll)
	httpRouter.POST("/author/add", authorController.Add)
	httpRouter.SERVE(":8000")

}
