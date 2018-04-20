package handlers

import (
	"time"

	"github.com/info344-a17/challenges-vincentmvdm/servers/gateway/models/users"
)

//SessionState contains information about the current session
type SessionState struct {
	StartTime         time.Time
	AuthenticatedUser *users.User
}
