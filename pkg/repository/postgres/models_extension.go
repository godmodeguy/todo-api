package postgres

type userDB struct {
	Id       	int     `json:"id"       	  gorm:"primary key"`
	Name     	string  `json:"name" 	 	  binding:"required"`
	Username 	string  `json:"username" 	  binding:"required"`
	PasswordHash string `json:"password_hash" binding:"required"`
} 