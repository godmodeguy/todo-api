package models

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type User struct {
	Id       int    `json:"id"       gorm:"primary key"`
	Name     string `json:"name" 	 binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
	TypeId int
}

type LinkType struct {
	Id   int
	Type string
}

type Task struct {
	Id          int    `json:"id"`
	ListId      int    `json:"list_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	Priority    int    `json:"priority"`
}
