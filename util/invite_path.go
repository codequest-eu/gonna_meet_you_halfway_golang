package util

import "os"

var host = os.Getenv("HOST")

func InvitePath(hash string) string {
	return "http://" + host + "/accept_meeting/" + hash
}
