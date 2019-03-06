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
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type mysql struct {
    dsn      string
}

// NewMySQL returns an instance of mysql
func NewMySQL(dsn string) Database {
    return &mysql{
        dsn: dsn,
    }
}

func (m *mysql) Init() (db *gorm.DB, err error) {
    if db, err = gorm.Open("mysql", m.dsn); err != nil {
        return nil, err
    }

    if err := db.DB().Ping(); err != nil {
        return nil, err
    }

    db.LogMode(true)
    db.SingularTable(true)

    return db, nil
}

func (m *mysql) Name() string {
    return "MySQL"
}

func (m *mysql) String() string {
    return fmt.Sprintf("%s %s",
        m.Name(), m.dsn)
}
