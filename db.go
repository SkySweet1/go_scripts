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

func Init() []Client {
	clients := []Client{}

	return clients
}

func CheckUser(log string, C []Client) bool {
	for _, n := range C {
		if n.Login == log {
			return true
		}
	}

	return false
}

func AddUser(login, password string, id int, conn net.Conn, clients *[]Client) {
	for _, n := range *clients {
		if n.Login == login {
			println("this login exist!")
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

func SearchUser(log string, C []Client) *Client {
	for _, n := range C {
		if n.Login == log {
			return &n
		}
	}
	return nil
}

func SearchUserByID(id int, C []Client) *Client {
	for _, n := range C {
		if n.ID == id {
			return &n
		}
	}
	return nil
}

func AuthUser(log string, pass string, C []Client) bool {
	for i, n := range C {
		if n.Login == log && n.Password == pass {
			C[i].Authorized = true
			return true
		}
	}
	return false
}

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

func Connect() net.Conn {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err.Error())
		return nil
	}
	return conn
}

func Disconnect(conn net.Conn) {
	if conn != nil {
		conn.Close()
	}
}
