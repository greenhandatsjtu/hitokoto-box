package main

import (
	"context"
	"log"
	"net/url"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"
)

const (
	endpoint = "https://international.v1.hitokoto.cn"
)

// Response is response struct referring to https://developer.hitokoto.cn/sentence/#%E8%BF%94%E5%9B%9E%E6%A0%BC%E5%BC%8F
type Response struct {
	Hitokoto string `json:"hitokoto"`
	From     string `json:"from"`
	FromWho  string `json:"from_who"`
}

func main() {
	ghToken := os.Getenv("GH_TOKEN")
	gistId := os.Getenv("GIST_ID")

	hitokoto, err := getHitokoto()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	if err := updateGist(ctx, ghToken, gistId, hitokoto); err != nil {
		log.Fatal(err)
	}
}

func getHitokoto() (*Response, error) {
	client := resty.New()
	query := url.Values{}
	query.Add("c", "a")
	query.Add("encode", "json")
	query.Add("charset", "utf-8")

	var resp Response
	_, err := client.R().
		SetQueryParamsFromValues(query).
		SetResult(&resp).
		Get(endpoint)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func updateGist(ctx context.Context, token string, gistId string, hitokoto *Response) error {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// get gist
	gist, _, err := client.Gists.Get(ctx, gistId)
	if err != nil {
		return err
	}

	// set gist content to new hitokoto
	for name, f := range gist.Files {
		content := hitokoto.Hitokoto + " ---" + hitokoto.From
		f.Content = &content
		gist.Files[name] = f
		break
	}
	if _, _, err := client.Gists.Edit(ctx, gistId, gist); err != nil {
		return err
	}
	return nil
}
