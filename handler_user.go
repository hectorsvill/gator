package main

import (
	"context"
	"fmt"
	"os/user"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/hectorsvill/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	err := s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("ussage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]

	err := s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	ctx := context.Background()
	uuidStr := uuid.NewString()


	return nil
}
