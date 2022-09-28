package server

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"

	"net-cat-v0.1/client"
)

const alphaNumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 "

type Server struct {
	ChatRoom *client.ChatRoom
	Command  chan client.Message
	Mutex    sync.Mutex
}

func NewServer() *Server {
	return &Server{
		ChatRoom: client.NewChatRoom(),
		Command:  make(chan client.Message),
	}
}

func (s *Server) Run() {
	for msg := range s.Command {
		s.processMsg(msg)
	}
}

func (s *Server) processMsg(input client.Message) {
	msg := "[" + time.Now().Format("2006-01-02 15:04:05") + "]" + "[" + input.From.Nick + "]" + ": " + input.Msg
	input.From.ChatRoom.Broadcast(input.From, msg)
	input.From.ChatRoom.HistoryStore.Write(msg)
	input.From.Conn.Write([]byte("[" + time.Now().Format("2006-01-02 15:04:05") + "]" + "[" + input.From.Nick + "]: "))
}

func (s *Server) handleRequest(conn net.Conn) {
	s.Mutex.Lock()
	lenMembers := len(s.ChatRoom.Members)
	s.Mutex.Unlock()

	if lenMembers == s.ChatRoom.MaxSize {
		conn.Write([]byte("The chat room already is full. Try later"))
		if err := conn.Close(); err != nil {
			log.Printf("The error in closing connection: %s", conn.RemoteAddr().String())
		}
		return
	}
	log.Printf("new client has connected: %s", conn.RemoteAddr().String())

	c := &client.Client{
		Conn: conn,
		Nick: "anonymous",
		Msg:  s.Command,
	}

	c.Conn.Write([]byte(LinuxLogoMsg() + "\n"))

	err := s.ProcessNaming(c)

	for err != nil {
		if errors.Is(err, ErrReadInput) {
			log.Printf("%s was disconnected from the server\n", c.Conn.RemoteAddr().String())
			return
		} else if errors.Is(err, ErrInvalidNick) {
			c.Conn.Write([]byte("The client with the same nick exists, try again\n"))
			err = s.ProcessNaming(c)
		} else if errors.Is(err, ErrEmptyNick) {
			c.Conn.Write([]byte("The client nick cannot be empty, try again\n"))
			err = s.ProcessNaming(c)
		} else if errors.Is(err, ErrWrongNick) {
			c.Conn.Write([]byte("The client nick is not correct, try again\n"))
			err = s.ProcessNaming(c)
		}
	}

	for {
		msg, err := bufio.NewReader(c.Conn).ReadString('\n')
		if err != nil {

			s.Mutex.Lock()
			delete(c.ChatRoom.Members, c.Conn.RemoteAddr())
			s.Mutex.Unlock()

			c.ChatRoom.Broadcast(c, fmt.Sprintf("%s has left our chat ...", c.Nick))
			c.ChatRoom.HistoryStore.Write(c.Nick + " has left our room ...")
			log.Printf("%s was disconnected from the server ", c.Conn.RemoteAddr().String())
			return
		}

		msg = strings.Trim(strings.TrimSpace(msg), "\r\n")

		if msg == "" {
			c.Conn.Write([]byte("[" + time.Now().Format("2006-01-02 15:04:05") + "]" + "[" + c.Nick + "]: "))
			continue
		}

		c.Msg <- client.Message{
			From: c,
			Msg:  msg,
		}

	}
}

func (s *Server) ProcessNaming(c *client.Client) error {
	c.Conn.Write([]byte("[ENTER YOUR NAME]: "))

	msgNick, err := bufio.NewReader(c.Conn).ReadString('\n')
	if err != nil {
		return fmt.Errorf("error in reading nick: %w", ErrReadInput)
	}

	if strings.TrimSpace(msgNick) == "" {
		return fmt.Errorf("invalid nick: %w", ErrEmptyNick)
	}

	for _, letter := range msgNick {
		if !strings.ContainsAny(alphaNumeric, string(letter)) && letter != 10 {
			return fmt.Errorf("invalid nick: %w", ErrWrongNick)
		}
	}

	c.Nick = strings.TrimSpace(strings.Trim(msgNick, "\r\n"))

	for _, m := range s.ChatRoom.Members {
		if m.Nick == c.Nick {
			return fmt.Errorf("invalid nick: %w", ErrInvalidNick)
		}
	}

	s.Mutex.Lock()
	s.ChatRoom.Members[c.Conn.RemoteAddr()] = c
	s.Mutex.Unlock()

	c.ChatRoom = s.ChatRoom

	s.ChatRoom.Broadcast(c, fmt.Sprintf("%s has joined our room ...", c.Nick))
	c.ChatRoom.HistoryStore.Read(c)
	c.ChatRoom.HistoryStore.Write(c.Nick + " has joined  our room ...")
	c.Conn.Write([]byte("[" + time.Now().Format("2006-01-02 15:04:05") + "]" + "[" + c.Nick + "]: "))

	return nil
}
