package iam

import (
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cooperlutz/go-full/internal/iam/adapters/inbound"
	"github.com/cooperlutz/go-full/internal/iam/adapters/outbound"
	"github.com/cooperlutz/go-full/internal/iam/service"
)

type IamModuleConfig struct {
	JwtSecret       string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

type IamModule struct {
	iamRepo     outbound.Querier
	Service     *service.IamService
	UserRestApi http.Handler
	AuthRestApi http.Handler
}

func NewModule(pgconn *pgxpool.Pool, conf IamModuleConfig) *IamModule {
	repo := outbound.New(pgconn)
	iamSvc := service.NewIamService(
		repo,
		conf.JwtSecret,
		conf.AccessTokenTTL,
		conf.RefreshTokenTTL,
	)

	return &IamModule{
		iamRepo:     repo,
		Service:     iamSvc,
		UserRestApi: inbound.NewIamUserApiController(repo),
		AuthRestApi: inbound.NewIamAuthApiController(iamSvc),
	}
}
