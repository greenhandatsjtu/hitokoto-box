package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"
)

const (
	endpoint = "https://v1.hitokoto.cn"
)

// Response is response struct referring to https://developer.hitokoto.cn/sentence/#%E8%BF%94%E5%9B%9E%E6%A0%BC%E5%BC%8F
type Response struct {
	Hitokoto string `json:"hitokoto"`
	From     string `json:"from"`
	FromWho  string `json:"from_who"`
}

func main() {
	ghToken := strings.TrimSpace(os.Getenv("GH_TOKEN"))
	gistId := strings.TrimSpace(os.Getenv("GIST_ID"))
	category := strings.TrimSpace(os.Getenv("CATEGORY"))
	categories := strings.Split(category, "")

	if ghToken == "" {
		log.Fatal("Please add GH_TOKEN environment.")
	}

	hitokoto, err := getHitokoto(categories)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	if err := updateGist(ctx, ghToken, gistId, hitokoto); err != nil {
		log.Fatal(err)
	}
}

func getHitokoto(categories []string) (*Response, error) {
	client := resty.New()

	// set query parameters
	query := url.Values{}
	for _, v := range categories {
		if c := strings.TrimSpace(v); c == "" {
			continue
		} else {
			query.Add("c", c)
		}
	}
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

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return err
	}
	now := time.Now().In(loc).Format(time.RFC1123)

	var from string
	if hitokoto.From == "" && hitokoto.FromWho == "" {
		from = ""
	} else if hitokoto.From != "" {
		from = fmt.Sprintf("\n ---%s", hitokoto.From)
	} else {
		from = fmt.Sprintf("\n ---%s", hitokoto.FromWho)
	}

	content := fmt.Sprintf("%s%s\n\nUpdated at %s", hitokoto.Hitokoto, from, now)

	// set gist content to new hitokoto
	fileName := "🌧Hitokoto"
	f := gist.Files[github.GistFilename(fileName)]
	f.Content = &content
	gist.Files[github.GistFilename(fileName)] = f

	if _, _, err := client.Gists.Edit(ctx, gistId, gist); err != nil {
		return err
	}
	return nil
}
