/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package pkg

import (
	"context"
	"math/rand"
	"time"
)

import (
	hessian "github.com/apache/dubbo-go-hessian2"
	"github.com/apache/dubbo-go/config"

	"github.com/dubbogo/gost/log"
)

var userProviderB = new(UserProviderB)

func init() {
	config.SetProviderService(new(UserProvider))
	config.SetConsumerService(userProviderB)
	hessian.RegisterPOJO(&User{})
}

type User struct {
	ID   string
	Name string
	Age  int32
	Time time.Time
}

type UserProvider struct {
}

func (u *UserProvider) GetUser(ctx context.Context, req []interface{}) (*User, error) {
	gxlog.CInfo("req:%#v", req)
	time.Sleep(time.Duration(rand.Intn(977)+300) * time.Millisecond)
	rsp := User{"A001", "Alex Stocks In Group A", 18, time.Now()}
	gxlog.CInfo("rsp:%#v", rsp)

	gxlog.CInfo("request provider B , req:%#v", req)
	getUserB(ctx)

	return &rsp, nil
}

func getUserB(ctx context.Context) {

	time.Sleep(time.Duration(rand.Intn(977)+300) * time.Millisecond)
	user := &User{}
	err := userProviderB.GetUser(ctx, []interface{}{"A001"}, user)
	if err != nil {
		gxlog.CError("error: %v\n", err)
		return
	}
	gxlog.CInfo("response result: %v\n", user)
}

func (u *UserProvider) Reference() string {
	return "UserProviderA"
}

func (u User) JavaClassName() string {
	return "org.apache.dubbo.User"
}

type UserProviderB struct {
	GetUser func(ctx context.Context, req []interface{}, rsp *User) error
}

func (u *UserProviderB) Reference() string {
	return "UserProviderB"
}
