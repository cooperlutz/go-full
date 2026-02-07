package handlers

import (
	"github.com/go-chi/chi/v5"

	"github.com/cooperlutz/go-full/internal/iam"
	"github.com/cooperlutz/go-full/pkg/hteeteepee"
)

type IaMApiController struct {
	IaMService *iam.IamService
	IamRouter  *chi.Mux
}

func NewIAMApiController(iamSvc *iam.IamService) *IaMApiController {
	iamRouter := hteeteepee.NewRouter("iam")
	authHandler := NewAuthHandler(iamSvc)
	iamRouter.HandleFunc("/register", authHandler.Register)
	iamRouter.HandleFunc("/login", authHandler.Login)
	iamRouter.HandleFunc("/refresh", authHandler.RefreshToken)
	controller := &IaMApiController{IaMService: iamSvc, IamRouter: iamRouter}

	return controller
}
