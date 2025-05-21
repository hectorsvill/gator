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

	user, err := s.db.GetUser(context.Background(), s.cfg.UserName)
	if err != nil {
		return err
	}

	id, err := s.db.GetID(context.Background(), s.cfg.UserName)
	if err != nil {
		return fmt.Errorf("couldn't get id: %w", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
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

	ff, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: feed.UserID,
		FeedID: feed.ID,

	})

	if err != nil {
		return fmt.Errorf("couldn't create feed follow %w",  err)
	}

	fmt.Println("Feed created succesfully:")
	printFeed(feed, user)
	fmt.Println("Feed followed successfully:")
	printFeedFollow(ff.UserName, ff.FeedName)
	fmt.Println("=============================")

	return nil
}

func handlerGetFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeedNameUrlUser(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get users: %w", err)
	}
	for _, feed := range feeds {
		user, err := s.db.GetUserWithID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("couldnt find user when getting feeds")
		}

		fmt.Printf("* %s\n", feed.Name)
		fmt.Printf("* %s\n", feed.Url)
		fmt.Printf("* %s\n ", user.Name)
	}
	return nil
}

func printFeed(feed database.GatorFeed, user database.GatorUser) {
	fmt.Printf("* ID: %s\n", feed.ID)
	fmt.Printf("* Created: %v\n", feed.CreatedAt)
	fmt.Printf("* Updated: %v\n", feed.UpdatedAt)
	fmt.Printf("* Name: %s\n", feed.Name)
	fmt.Printf("* URL: %s\n", feed.Url)
	fmt.Printf("* User: %s\n", feed.Name)
	
}