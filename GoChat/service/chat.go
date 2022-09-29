package serviceChat

import (
	"encoding/json"

	"github.com/olahol/melody"
)

type Message struct {
	Event   string `json:"event"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type UserInfo struct {
	ID   int
	Name string
}

var index = 1

func (m *Message) GetByteMessage() []byte {
	result, _ := json.Marshal(m)
	return result
}

// 廣播使用
func GetBroadcastMsg(name string, msg string) []byte {
	res := &Message{
		Event:   "other",
		Name:    name,
		Content: msg,
	}

	return res.GetByteMessage()
}

// 加入設定個人資料
func JoinIn(name string, s *melody.Session) {
	s.Set("Info", &UserInfo{
		ID:   index,
		Name: name,
	})

	index++
}

func GetUserInfo(s *melody.Session) *UserInfo {
	info := s.MustGet("Info").(*UserInfo)
	return info
}
