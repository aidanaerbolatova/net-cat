package client

import (
	"net"
	"sync"
)

type Client struct {
	Nick     string
	ChatRoom *ChatRoom
	Conn     net.Conn
	Msg      chan<- Message
}

type ChatRoom struct {
	Members      map[net.Addr]*Client
	HistoryStore *fileStore
	MaxSize      int
}

type fileStore struct {
	Mu    sync.Mutex
	Store []string
}

type Message struct {
	From *Client
	Msg  string
}
