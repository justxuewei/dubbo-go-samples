package pkg

import (
	"context"
	"strconv"
	"time"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
)

type UserProviderTriple struct{}

func (u *UserProviderTriple) GetUser1(_ context.Context, userID string) (*User, error) {
	logger.Infof("req:%#v", userID)
	rsp := User{userID, "Joe", 48, time.Now()}
	logger.Infof("rsp:%#v", rsp)
	return &rsp, nil
}

func (u *UserProviderTriple) GetUser2(_ context.Context, userID string, name string) (*User, error) {
	logger.Infof("req:%#v, %#v", userID, name)
	rsp := User{userID, name, 48, time.Now()}
	logger.Infof("rsp:%#v", rsp)
	return &rsp, nil
}

func (u *UserProviderTriple) GetUser3(_ context.Context, userCode int) (*User, error) {
	logger.Infof("req:%#v", userCode)
	rsp := User{strconv.Itoa(userCode), "Alex Stocks", 18, time.Now()}
	logger.Infof("rsp:%#v", rsp)
	return &rsp, nil
}

func (u *UserProviderTriple) GetUser4(_ context.Context, userCode int, name string) (*User, error) {
	logger.Infof("req:%#v, %#v", userCode, name)
	rsp := User{strconv.Itoa(userCode), name, 18, time.Now()}
	logger.Infof("rsp:%#v", rsp)
	return &rsp, nil
}

func (u *UserProviderTriple) GetOneUser(_ context.Context) (*User, error) {
	return &User{
		ID:   "1000",
		Name: "xavierniu",
		Age:  24,
		Time: time.Now(),
	}, nil
}

func (u *UserProviderTriple) GetUsers(_ context.Context, userIdList []string) (*UserResponse, error) {
	logger.Infof("req:%#v", userIdList)
	var users []*User
	for _, i := range userIdList {
		users = append(users, userMap[i])
	}
	return &UserResponse{
		Users: users,
	}, nil
}

func (u *UserProviderTriple) GetUsersMap(_ context.Context, userIdList []string) (map[string]*User, error) {
	logger.Infof("req:%#v", userIdList)
	var users = make(map[string]*User)
	for _, i := range userIdList {
		users[i] = userMap[i]
	}
	return users, nil
}

func (u *UserProviderTriple) QueryUser(_ context.Context, user *User) (*User, error) {
	logger.Infof("req1:%#v", user)
	rsp := User{user.ID, user.Name, user.Age, time.Now()}
	logger.Infof("rsp1:%#v", rsp)
	return &rsp, nil
}

func (u *UserProviderTriple) QueryUsers(_ context.Context, users []*User) (*UserResponse, error) {
	return &UserResponse{
		Users: users,
	}, nil
}

func (u *UserProviderTriple) QueryAll(_ context.Context) (*UserResponse, error) {
	users := []*User{
		{
			ID:   "001",
			Name: "Joe",
			Age:  18,
			Time: time.Now(),
		},
		{
			ID:   "002",
			Name: "Wen",
			Age:  20,
			Time: time.Now(),
		},
	}

	return &UserResponse{
		Users: users,
	}, nil
}

func (u *UserProviderTriple) MethodMapper() map[string]string {
	return map[string]string{
		"QueryUser":  "queryUser",
		"QueryUsers": "queryUsers",
		"QueryAll":   "queryAll",
	}
}

func (u *UserProviderTriple) Reference() string {
	return "UserProviderTriple"
}
