package handlers

import (
	"time"

	"github.com/leedsjb/capstone2k18/servers/gateway/models/users"
)

//SessionState contains information about the current session
type SessionState struct {
	StartTime         time.Time
	AuthenticatedUser *users.User
}
