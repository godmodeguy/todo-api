package postgres

import (
	"learn/todoapi/pkg/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)


type PgList struct {
	db	*gorm.DB
}

func NewPgList(db *gorm.DB) PgList {
	return PgList{db: db}
}

func (p PgList) GetAll() ([]models.TodoList, error) {
	users := make([]models.TodoList, 0, 100)
	tx := p.db.Find(&users)
	return users, tx.Error
}

func (p PgList) CreateList(list models.TodoList) (int, error) {
	tx := p.db.Create(&list)
	logrus.Debugf("%+v", tx)
	return list.Id, tx.Error
}

func (p PgList) GetById(id int) (models.TodoList, error) {
	list := models.TodoList{Id: id}
	tx := p.db.Find(&list)
	return list, tx.Error
}

func (p PgList) UpdateList(list models.TodoList) (models.TodoList, error) {
	tx := p.db.Updates(&list)
	return list, tx.Error
}

func (p PgList) DeleteList(id int) error {
	list := models.TodoList{Id: id}
	tx := p.db.Delete(&list)
	return tx.Error
}
