package service

import (
	"context"

	"github.com/KXX747/wolf/getaway/kratos-getaway-servers/internal/dao"
	"github.com/bilibili/kratos/pkg/conf/paladin"
)

// Service service.
type Service struct {
	ac  *paladin.Map
	dao dao.Dao
	AppConfig *dao.Config
}

// New new a service and return.
func New() (s *Service) {

	appConf:=dao.BuildConfig()
	s = &Service{
		//ac:  ac,
		dao: dao.New(),
		AppConfig:appConf,
	}
	return s
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
	s.dao.Close()
}
