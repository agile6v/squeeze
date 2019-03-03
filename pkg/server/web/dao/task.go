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

package dao

import (
    "time"
    "github.com/agile6v/squeeze/pkg/server/web/db"
)

type Task struct {
    Id          int       `sql:"AUTO_INCREMENT"`
    Status      int       `sql:"type:tinyint"`
    Result      int       `sql:"type:tinyint"`
    Request     string    `sql:"type:varchar(2048)"`
    Response    string    `sql:"type:varchar(2048)"`
    CreatedAt   time.Time `sql:"timestamp"`
    UpdatedAt   time.Time `sql:"timestamp"`
}

func Init() error {
    orm := db.GetOrm()
    orm.AutoMigrate(&Task{})
    return nil
}

func CreateTask(reqData string) error {
    orm := db.GetOrm()
    err := orm.Create(&Task{Request: reqData}).Error
    if err != nil {
        return err
    }

    return nil
}

func UpdateTaskResponse(id int, data string) error {
    orm := db.GetOrm()
    err := orm.Model(Task{}).Where("id = ?", id).Updates(Task{Response: data}).Error
    if err != nil {
        return err
    }

    return nil
}

func DeleteTask(id int) error {
    orm := db.GetOrm()
    err := orm.Delete(&Task{Id: id}).Error
    if err != nil {
        return err
    }

    return nil
}

func ListTask() ([]Task, error) {
    var tasks []Task
    orm := db.GetOrm()
    err := orm.Find(&tasks).Error
    if err != nil {
        return nil, err
    }

    return tasks, nil
}

func SearchTask(id int) (*Task, error) {
    var task Task
    orm := db.GetOrm()
    err := orm.Where("id = ?", id).Find(&task).Error
    if err != nil {
        return nil, err
    }

    return &task, nil
}







