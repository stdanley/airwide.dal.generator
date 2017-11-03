/*
 *  Copyright (c) 2017, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/jmoiron/sqlx"
	"github.com/golang/glog"
	dao2 "github.com/nebulaim/nebula-dal-generator/samples/dal/dao"
	"flag"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

func main() {
	flag.Parse()

	// 数据库
	db, err := sqlx.Connect("mysql", "root:@/nebulaim?charset=utf8")
	if err != nil {
		glog.Error("Connect database error: ", err)
		return
	}

	dao := dao2.NewAppsDAO(db)
	do, _ := dao.SelectById(1)
	glog.Info(do)
}