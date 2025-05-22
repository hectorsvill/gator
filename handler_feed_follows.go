package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hectorsvill/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.GatorUser) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <name> <url>", cmd.Name)
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't get feed with url: %w", err)
	}
	
	ffrow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Println("Feed follow created: ", ffrow)
	printFeedFollow(ffrow.UserName, ffrow.FeedName)
	return nil
}

func handlerListFeedFollows(s *state, cmd command, user database.GatorUser) error {
	fFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follows: %w", err)
	}

	if len(fFollows) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	fmt.Printf("Feed follows for user %s:\n", user.Name)

	for _, ff := range fFollows {
		fmt.Printf("* %s\n", ff.FeedName)
	}
	return nil
}

func printFeedFollow(userName, feedName string) {
	fmt.Printf("* user: %s\n* feedname: %s\n", userName, feedName)
}

func handlerDeleteFeedFollowByUserAndFeed(s *state, cmd command, user database.GatorUser) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <url>", cmd.Name)
	}
	feedUrl := cmd.Args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), feedUrl)
	if err != nil {
		return err
	}

	err = s.db.DeleteFeedFollowByUserAndFeed(context.Background(), database.DeleteFeedFollowByUserAndFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		return err
	}
	return nil
}
