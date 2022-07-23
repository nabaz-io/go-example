package auth

import (
	"testing"
)

func TestAuthDB(t *testing.T) {
	tt := []struct {
		name     string
		username string
		password string
	}{
		{
			name:     "DatabaseCacheHit",
			username: "admin",
			password: "admin",
		},
	}

	for _, tc := range tt {
		if !Auth(tc.username, tc.password) {
			t.Errorf("%s: Auth failed", tc.name)
		}
	}
}
