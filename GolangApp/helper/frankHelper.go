package helper

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func PrintUserInfo(userName string, email string, userTickets uint) {
	data := map[string]string{
		"Name":    userName,
		"Email":   email,
		"tickets": strconv.FormatUint(uint64(userTickets), 10),
	}

	PrintJson(data)
}

func PrintJson(mapData map[string]string) {
	jsonStr, err := json.Marshal(mapData)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println("---Json Data---")
		fmt.Println(string(jsonStr))
		fmt.Println("---Json Data---")
	}
}
