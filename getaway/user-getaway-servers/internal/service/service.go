package service

import (
	"context"

	"github.com/KXX747/wolf/getaway/user-getaway-servers/internal/dao"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	account "github.com/KXX747/wolf/public/user-account-server/internal/model/user"

)

// Service service.
type Service struct {
	ac  *paladin.Map
	dao dao.Dao
	accountRPC account.UsersClient


}

// New new a service and return.
func New() (s *Service) {
	var ac = new(paladin.TOML)
	if err := paladin.Watch("application.toml", ac); err != nil {
		panic(err)
	}
	s = &Service{
		ac:  ac,
		dao: dao.New(),
	}
	return s
}



//AddUserReq function is  add user
func (s *Service) AddUser(ctx context.Context,name string,mobile string)( reply *account.UserReply, err error) {
	u:=&account.AddUserReq{Name:name,Mobile:mobile}
	reply ,err=s.accountRPC.AddUser(ctx,u)


	return
}

//update user of id_no
func (s *Service) UpdateUser(ctx context.Context,id_no string, name string,mobile string,address string) (reply *account.UserReply,err error) {

	return
}

//function delete user of id_no and name
func(s *Service)DeleteUser(ctx context.Context,id_no string,content string)(eply *account.UserReply,err error){

	return
}


// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
	s.dao.Close()
}
