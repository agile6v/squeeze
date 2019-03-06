// Copyright 2019 Squeeze Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db

import (
    "github.com/jinzhu/gorm"
    log "github.com/golang/glog"
    "fmt"
)

const (
    DB_MYSQL   = "mysql"
    DB_SQLITE  = "sqlite"
)

type Database interface {
    Name() string
    String() string
    Init() (*gorm.DB, error)
}

var database Database
var orm *gorm.DB

func Init(dbType, dsn, file string) (err error) {
    log.Infof("Database type: %s", dbType)

    switch dbType {
    case DB_MYSQL:
        database = NewMySQL(dsn)
    case DB_SQLITE:
        database = NewSQLite(file)
    default:
        log.Error("Invalid database type: ", dbType)
        return fmt.Errorf("Invalid database type: ", dbType)
    }

    log.Info(database.String())

    log.Info("Initializing database...")
    orm, err = database.Init()
    if  err != nil {
        return err
    }
    log.Info("Initialize database completed")

    return nil
}

func GetOrm() *gorm.DB {
    return orm
}

