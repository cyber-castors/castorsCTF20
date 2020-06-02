package controllers

import (
	"fmt"
	"os/exec"
	"time"

	"../models"
)

func CleanSession() {
	for k, v := range models.DBSessions {
		if time.Now().Sub(v.LastActivity) > (time.Minute * 10) {
			delete(models.DBSessions, k)
			name := "shortcuts" + v.UserID
			cmd := exec.Command("rm", "-rf", name)
			err := cmd.Run()
			if err != nil {
				fmt.Print(err)
			}
		}
	}
}
