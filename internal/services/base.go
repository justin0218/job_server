package services

import "job_server/store"

type baseService struct {
	Redis  store.Redis
	Config store.Config
}
