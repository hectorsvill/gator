package main

import (
	"context"

	"github.com/hectorsvill/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.GatorUser) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.UserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}
