// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package commands

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sonnnnnnp/reverie/server/adapter/http/controller"
	"github.com/sonnnnnnp/reverie/server/pkg/line"
	"github.com/sonnnnnnp/reverie/server/usecase/authorize"
	"github.com/sonnnnnnp/reverie/server/usecase/call"
	"github.com/sonnnnnnp/reverie/server/usecase/call_timeline"
	"github.com/sonnnnnnp/reverie/server/usecase/gateway"
	"github.com/sonnnnnnp/reverie/server/usecase/post"
	"github.com/sonnnnnnp/reverie/server/usecase/post_timeline"
	"github.com/sonnnnnnp/reverie/server/usecase/user"
)

// Injectors from wire.go:

func Wire(pool *pgxpool.Pool) *controller.Controller {
	client := line.New()
	authorizeUsecase := authorize.New(pool, client)
	callUsecase := call.New(pool)
	callTimelineUsecase := call_timeline.New(pool)
	postUsecase := post.New(pool)
	postTimelineUsecase := post_timeline.New(pool)
	gatewayUsecase := gateway.New()
	userUsecase := user.New(pool)
	controllerController := controller.New(authorizeUsecase, callUsecase, callTimelineUsecase, postUsecase, postTimelineUsecase, gatewayUsecase, userUsecase)
	return controllerController
}
