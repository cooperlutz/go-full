package iam

import (
	"crypto/rsa"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cooperlutz/go-full/internal/iam/adapters/inbound"
	"github.com/cooperlutz/go-full/internal/iam/adapters/outbound"
	"github.com/cooperlutz/go-full/internal/iam/service"
)

type IamModuleConfig struct {
	PrivateKey      *rsa.PrivateKey
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

type IamModule struct {
	Service     *service.IamService
	UserRestApi http.Handler
	AuthRestApi http.Handler
}

func NewModule(pgconn *pgxpool.Pool, conf IamModuleConfig) *IamModule {
	iamPostgres := outbound.NewPostgresAdapter(pgconn)
	iamSvc := service.NewIamService(
		iamPostgres,
		iamPostgres,
		iamPostgres,
		conf.PrivateKey,
		conf.AccessTokenTTL,
		conf.RefreshTokenTTL,
	)

	return &IamModule{
		Service:     iamSvc,
		UserRestApi: inbound.NewIamUserApiController(iamSvc),
		AuthRestApi: inbound.NewIamAuthApiController(iamSvc),
	}
}
