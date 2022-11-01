package v1

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"homework/specs"
)

func (a apiServer) GetAuthUser(w http.ResponseWriter, r *http.Request) {
	user := &specs.UserProfile{
		AvatarUrl: "https://example.com/avatar.jpg",
		Id:        uuid.New().String(),
		Login:     "admin",
	}

	response, err := json.Marshal(user)
	if err != nil {
		return
	}

	w.Write(response)
}
