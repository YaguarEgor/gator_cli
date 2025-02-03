package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}
	defer resp.Body.Close()
	resp.Header.Set("User-Agent", "gator")
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RSSFeed{}, err
	}
	var feed RSSFeed
	err = xml.Unmarshal(data, &feed)
	if err != nil {
		return &RSSFeed{}, err
	}
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for _, item := range feed.Channel.Item {
		item.Description = html.UnescapeString(item.Description)
		item.Title = html.UnescapeString(item.Title)
	}
	return &feed, nil
}