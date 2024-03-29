package service

import (
	"context"
	"github.com/KXX747/wolf/getaway/kratos-getaway-servers/internal/dao"
)

// Service service.
type Service struct {
	Dao dao.Dao
	AppConfig *dao.Config
}

// New new a service and return.
func New() (s *Service) {

	appConf:=dao.BuildConfig()
	s = &Service{
		//ac:  ac,
		Dao: dao.New(),
		AppConfig:appConf,
	}
	return s
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return s.Dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
	s.Dao.Close()
}
