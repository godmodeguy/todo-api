package postgres

import (
	"learn/todoapi/pkg/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)


type PgAuth struct {
	db	*gorm.DB
}

func NewPgAuth(db *gorm.DB) PgAuth {
	return PgAuth{db: db}
}

func (p PgAuth) CreateUser(user models.User) (int, error) {
	userdb := userDB{
		Id: user.Id,
		Name: user.Name,
		Username: user.Username,
		PasswordHash: hash(user.Password),
	}
	tx := p.db.Create(&userdb)
	return userdb.Id, tx.Error
}

func (p PgAuth) GetUser(username, password string) (models.User, error) {
	var userdb userDB
	tx := p.db.Where("username = ? AND password_hash = ?", username, hash(password)).First(&userdb)

	logrus.Debugf("found: %+v", userdb)

	return models.User{
		Id: userdb.Id,
		Name: userdb.Name,
		Username: userdb.Username,
	}, tx.Error
}

func (p PgAuth) GetUserById(id int) (models.User, error) {
	user := models.User{Id: id}
	tx := p.db.First(&user)
	return user, tx.Error
}