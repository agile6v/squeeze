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
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type sqlite struct {
	file string
}

// NewSQLite returns an instance of sqlite
func NewSQLite(file string) Database {
	return &sqlite{
		file: file,
	}
}

func (s *sqlite) Init() (db *gorm.DB, err error) {
	db, err = gorm.Open("sqlite3", s.file)
	if err != nil {
		return nil, err
	}

	if err := db.DB().Ping(); err != nil {
		return nil, err
	}

	db.LogMode(true)

	return db, nil
}

func (s *sqlite) Name() string {
	return "SQLite"
}

func (s *sqlite) String() string {
	return fmt.Sprintf("%s file:%s", s.Name(), s.file)
}
