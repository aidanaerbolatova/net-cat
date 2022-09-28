package client

import (
	"net"
	"strings"
	"time"
)

func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		Members:      make(map[net.Addr]*Client),
		HistoryStore: &fileStore{},
		MaxSize:      10,
	}
}

func (r *ChatRoom) Broadcast(from *Client, msg string) {
	for addr, member := range r.Members {
		if addr != from.Conn.RemoteAddr() {
			member.Conn.Write([]byte(clean(time.Now().Format("2006-01-02 15:04:05")+"]["+member.Nick+"]: ") + msg + "\n[" + time.Now().Format("2006-01-02 15:04:05") + "][" + member.Nick + "]: "))
		}
	}
}

func clean(line string) string {
	return "\r" + strings.Repeat(" ", len(line)) + "\r"
}
