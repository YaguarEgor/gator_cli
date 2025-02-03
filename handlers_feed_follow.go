package main

import (
	"context"
	"fmt"
	"time"

	"github.com/YaguarEgor/gator_cli/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %v <url>", cmd.name)
	}
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("couldn't find feed to follow: %w", err)
	}
	feed_follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't follow a feed: %w", err)
	}
	printFeedFollow(feed_follow)
	return nil
}

func printFeedFollow(feed_follow database.CreateFeedFollowRow) {
	fmt.Printf(" * Name of feed: %v\n", feed_follow.FeedName)
	fmt.Printf(" * Name of user: %v\n", feed_follow.UserName)
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %v", cmd.name)
	}
	feed_follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't find feeds for current user: %w", err)
	}
	fmt.Println("All feeds for current user:")
	for _, item := range feed_follows {
		fmt.Printf(" * Name fo feed: %v\n", item.FeedName)
	}
	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %v <url>", cmd.name)
	}
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("couldn't find that feed for current user: %w", err)
	}
	err = s.db.UnfollowFeed(context.Background(), database.UnfollowFeedParams{
		FeedID: feed.ID,
		UserID: user.ID,
	})
	return err
}