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

package integration

import (
	"context"
	"testing"
	"time"
)

import (
	"dubbo.apache.org/dubbo-go/v3/config/generic"

	hessian "github.com/apache/dubbo-go-hessian2"

	"github.com/stretchr/testify/assert"
)

func TestGetUser1(t *testing.T) {
	o, err := dubboRefConf.GetRPCService().(*generic.GenericService).Invoke(
		context.TODO(),
		[]interface{}{
			"GetUser1",
			[]string{"java.lang.String"},
			[]hessian.Object{"A003"},
		},
	)
	assert.Nil(t, err)
	assert.IsType(t, make(map[interface{}]interface{}, 0), o)
	resp := o.(map[interface{}]interface{})
	assert.Equal(t, "Joe", resp["name"])
	assert.Equal(t, int32(48), resp["age"])
	assert.Equal(t, "A003", resp["iD"])
}

func TestGetUser2(t *testing.T) {
	o, err := dubboRefConf.GetRPCService().(*generic.GenericService).Invoke(
		context.TODO(),
		[]interface{}{
			"GetUser2",
			[]string{"java.lang.String", "java.lang.String"},
			[]hessian.Object{"A003", "lily"},
		},
	)
	assert.Nil(t, err)
	assert.IsType(t, make(map[interface{}]interface{}, 0), o)
	resp := o.(map[interface{}]interface{})
	assert.Equal(t, "lily", resp["name"])
	assert.Equal(t, int32(48), resp["age"])
	assert.Equal(t, "A003", resp["iD"])
}

func TestGetUser3(t *testing.T) {
	o, err := dubboRefConf.GetRPCService().(*generic.GenericService).Invoke(
		context.TODO(),
		[]interface{}{
			"GetUser3",
			[]string{"int"},
			[]hessian.Object{1},
		},
	)
	assert.Nil(t, err)
	assert.IsType(t, make(map[interface{}]interface{}, 0), o)
	resp := o.(map[interface{}]interface{})
	assert.Equal(t, "Alex Stocks", resp["name"])
	assert.Equal(t, int32(18), resp["age"])
	assert.Equal(t, "1", resp["iD"])
}

func TestGetUser4(t *testing.T) {
	o, err := dubboRefConf.GetRPCService().(*generic.GenericService).Invoke(
		context.TODO(),
		[]interface{}{
			"GetUser4",
			[]string{"int", "java.lang.String"},
			[]hessian.Object{1, "zhangsan"},
		},
	)
	assert.Nil(t, err)
	assert.IsType(t, make(map[interface{}]interface{}, 0), o)
	resp := o.(map[interface{}]interface{})
	assert.Equal(t, "zhangsan", resp["name"])
	assert.Equal(t, int32(18), resp["age"])
	assert.Equal(t, "1", resp["iD"])
}

func TestGetOneUser(t *testing.T) {
	o, err := dubboRefConf.GetRPCService().(*generic.GenericService).Invoke(
		context.TODO(),
		[]interface{}{
			"GetOneUser",
			[]hessian.Object{},
			[]hessian.Object{},
		},
	)
	assert.Nil(t, err)
	assert.IsType(t, make(map[interface{}]interface{}, 0), o)
	resp := o.(map[interface{}]interface{})
	assert.Equal(t, "xavierniu", resp["name"])
	assert.Equal(t, int32(24), resp["age"])
	assert.Equal(t, "1000", resp["iD"])
}

func TestGetUsers(t *testing.T) {
	o, err := dubboRefConf.GetRPCService().(*generic.GenericService).Invoke(
		context.TODO(),
		[]interface{}{
			"GetUsers",
			[]string{"java.util.List"},
			[]hessian.Object{
				[]hessian.Object{
					"001", "002", "003", "004",
				},
			},
		},
	)
	assert.Nil(t, err)
	assert.IsType(t, make(map[interface{}]interface{}, 0), o)
	//resp := o.(map[interface{}]interface{})
	//assert.Equal(t, "other-zhangsan", resp[0].(*pkg.User).Name)
	//assert.Equal(t, "other-lisi", resp[1].(*pkg.User).Name)
	//assert.Equal(t, "other-lily", resp[2].(*pkg.User).Name)
	//assert.Equal(t, "other-lisa", resp[3].(*pkg.User).Name)
}

func TestQueryUser(t *testing.T) {
	o, err := dubboRefConf.GetRPCService().(*generic.GenericService).Invoke(
		context.TODO(),
		[]interface{}{
			"queryUser",
			[]string{"org.apache.dubbo.User"},
			[]hessian.Object{map[string]hessian.Object{
				"iD":   "3213",
				"name": "panty",
				"age":  25,
				"time": time.Now(),
			}},
		},
	)

	assert.Nil(t, err)
	assert.IsType(t, make(map[interface{}]interface{}, 0), o)
	resp := o.(map[interface{}]interface{})
	assert.Equal(t, "panty", resp["name"])
	assert.Equal(t, int32(25), resp["age"])
	assert.Equal(t, "3213", resp["iD"])
}

//
//func TestQueryUsers(t *testing.T) {
//	o, err := referenceConfig.GetRPCService().(*generic.GenericService).Invoke(
//		context.TODO(),
//		[]interface{}{
//			"queryUsers",
//			[]string{"org.apache.dubbo.User"},
//			[]hessian.Object{
//				map[string]hessian.Object{
//					"id":    "3212",
//					"name":  "XavierNiu",
//					"age":   24,
//					"time":  time.Now().Add(4),
//					"class": "org.apache.dubbo.User",
//				},
//				map[string]hessian.Object{
//					"iD":    "3213",
//					"name":  "zhangsan",
//					"age":   21,
//					"time":  time.Now().Add(4),
//					"class": "org.apache.dubbo.User",
//				},
//			},
//		},
//	)
//
//	assert.Nil(t, err)
//	assert.IsType(t, make(map[interface{}]interface{}, 0), o)
//	resp := o.(map[interface{}]interface{})
//	assert.Equal(t, "XavierNiu", resp[0].(*pkg.User).Name)
//	assert.Equal(t, "zhangsan", resp[1].(*pkg.User).Name)
//}
//
//func TestQueryAll(t *testing.T) {
//	o, err := referenceConfig.GetRPCService().(*generic.GenericService).Invoke(
//		context.TODO(),
//		[]interface{}{
//			"queryAll",
//			[]hessian.Object{},
//			[]hessian.Object{},
//		},
//	)
//
//	assert.Nil(t, err)
//	assert.IsType(t, make(map[interface{}]interface{}, 0), o)
//	resp := o.(map[interface{}]interface{})
//	assert.Equal(t, "Joe", resp[0].(*pkg.User).Name)
//	assert.Equal(t, "Wen", resp[1].(*pkg.User).Name)
//}