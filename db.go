package main

import (
	"fmt"
	"net"
)

type Client struct {
	Login      string
	Password   string
	ID         int
	Conn       net.Conn
	Authorized bool
}

/*
Login      string		имя(уникальный)
Password   string		пароль
ID         int			айди(уникальный)
Conn       net.Conn		соединение
Authorized bool			флаг авторизации(true/false)
*/

func Init() []Client {
	clients := []Client{}

	return clients
}

/*
создает и возвращает пустой лист клиентов
инициализирует базу данных пользователей
*/

func CheckUser(log string, C []Client) bool {
	for _, n := range C {
		if n.Login == log {
			return true
		}
	}

	return false
}

/*
проверка на существование пользователя с уникальным логином
возвращает true или false
*/

func AddUser(login, password string, id int, conn net.Conn, clients *[]Client) {
	for _, n := range *clients {
		if n.Login == login {
			fmt.Println("this login exist!")
			return
		}
	}

	*clients = append(*clients, Client{
		Login:      login,
		Password:   password,
		ID:         id,
		Conn:       conn,
		Authorized: false,
	})
}

/*
добавляет нового пользователя в лист клиентов
перед этим проверяет, не занят ли логин
*/

func SearchUser(log string, C []Client) *Client {
	for _, n := range C {
		if n.Login == log {
			return &n
		}
	}
	return nil
}

/*
ищет пользователя по логину и возвращает указатель на найденного пользователя
если не найден то возвращает nil (NULL)
*/

func SearchUserByID(id int, C []Client) *Client {
	for _, n := range C {
		if n.ID == id {
			return &n
		}
	}
	return nil
}

/*
ищет пользователя по уникальному айди
также возвращает указатель на найденного пользователя или nil
*/

func AuthUser(log string, pass string, C []Client) bool {
	for i, n := range C {
		if n.Login == log && n.Password == pass {
			C[i].Authorized = true
			return true
		}
	}
	return false
}

/*
проверяет логин и пароль пользователя
если данные верны то устанавливает флаг auth на true,
а так же возвращает true при успешной авторизации
*/

// func Registration_of_conn(Conn *net.Conn, log string, pass string, C []Client) {
// 	if CheckUser(log, C) == true {
// 		if AuthUser(log, pass, C) == true {
// 			x := &SearchUser(log, C)						тут было неправильно изза сcылки
// 			x.Conn = *Conn
// 		}
// 	}

// }

func Registration_of_conn(Conn *net.Conn, log string, pass string, C []Client) {
	if CheckUser(log, C) == true {
		if AuthUser(log, pass, C) == true {
			x := SearchUser(log, C)
			x.Conn = *Conn
		} else {
			(*Conn).Close()
		}
	} else {
		AddUser(log, pass, 0, *Conn, &C)
		AuthUser(log, pass, C)
	}
}

/*
одна из основных функций

CheckUser(log, C) == true					существование пользователя
AuthUser(log, pass, C) == true				верный ли пароль

x := SearchUser(log, C)						находим пользователя
x.Conn = *Conn								устанавливаем соединение

else {(*Conn).Close()}						если пароль не верный, то закрываем соединение

а так же если пользователя не существует -> AuthUser(log, pass, C) == false
AddUser(log, pass, 0, *Conn, &C)			создание нового пользователя
AuthUser(log, pass, C)						его авторизация
*/

func Connect() net.Conn {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err.Error())
		return nil
	}
	return conn
}

/*
устанавливает tcp соединение с сервером localhost:8080
при ошибке выводит сообщение и возвращает nil
при успехе возвращает обьект соединение
*/

func Disconnect(conn net.Conn) {
	if conn != nil {
		conn.Close()
	}
}

/*
отключение пользователя
*/
