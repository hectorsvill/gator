package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hectorsvill/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	fmt.Printf("Feed: %+v\n", feed.Channel)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %v <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	id, err := s.db.GetID(context.Background(), s.cfg.UserName)
	if err != nil {
		return fmt.Errorf("couldn't get id: %w", err)
	}

	_, err = s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    id,
	})

	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	return nil
}
