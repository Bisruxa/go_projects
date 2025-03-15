package handlers

import (
"github.com/go-chi/chi"
chimiddle"github.com/go-chi/chi/middleware"
"github.com/Bisruxa/go_projects/internal/middleware"
)

func Handler(r *chi.Mux){
// if handler tarts with small letter it means the function is private and ca not be imported
r.Use(chimiddle.StripSlashes)
// ignores the tripple slashes
r.Route("/account",func(router chi.Router){
	//chi router allows as to define a get request 
router.Use(middleware.Authorization)
router.Get("/coins",GetCoinBalance)
})


}
