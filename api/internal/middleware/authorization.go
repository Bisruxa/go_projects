package middleware

import (
	"errors"
	"net/http"
	//"github.com/Bisruxa/go_projects/go"
	"github.com/Bisruxa/go_projects/internal/tools"
	"github.com/Bisruxa/go_projects/api"

	log "github.com/sirupsen/logrus"
)

var ErrFoo = errors.New("invalid username or token")
func Authorization(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		var username string= r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error 
		if username == "" || token == ""{
			log.Error(ErrFoo)
			api.RequestErrorHandler(w,ErrFoo)
			return 
		}
		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return 
		}
		var loginDetails *tools.LoginDetails
		if database == nil {
    log.Error("Database connection failed")
    api.InternalErrorHandler(w)
    return
}

		loginDetails = (*database).GetUserLoginDetails(username)
		if(loginDetails == nil || (token != (*loginDetails).AuthToken)){
			log.Error(ErrFoo)
			api.RequestErrorHandler(w,ErrFoo)
			return 
		}
		next.ServeHTTP(w,r)
	})
}