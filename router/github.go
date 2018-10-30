package router

import (
	"context"
	"sync"
	"time"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

var myRepos = map[string]struct{}{}
var mu sync.Mutex

type query struct {
	User struct {
		Repositories struct {
			PageInfo struct {
				HasNextPage bool
				EndCursor   string
			}
			TotalCount int
			Nodes      []struct {
				Name            string
				PrimaryLanguage struct {
					Name string
				}
			}
		} `graphql:"repositories(first: 100, after: $cursor, ownerAffiliations: OWNER)"`
	} `graphql:"user(login: $login)"`
}

func exists(repo string) bool {
	mu.Lock()
	_, ok := myRepos[repo]
	mu.Unlock()
	return ok
}

func getRepos(tok string) map[string]struct{} {
	resp := map[string]struct{}{}
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: tok},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	c := githubv4.NewClient(httpClient)

	vars := map[string]interface{}{
		"login":  githubv4.String("marwan-at-work"),
		"cursor": (*githubv4.String)(nil),
	}

	var q query

	for {
		var qry query
		err := c.Query(context.Background(), &qry, vars)
		if err != nil {
			panic(err)
		}

		q.User.Repositories.Nodes = append(q.User.Repositories.Nodes, qry.User.Repositories.Nodes...)
		if !qry.User.Repositories.PageInfo.HasNextPage {
			break
		}
		vars["cursor"] = githubv4.String(qry.User.Repositories.PageInfo.EndCursor)
	}

	for _, repo := range q.User.Repositories.Nodes {
		if repo.PrimaryLanguage.Name != "Go" {
			continue
		}

		resp[repo.Name] = struct{}{}
	}

	return resp
}

func runVanityUpdater(tok string) {
	ticker := time.NewTicker(10 * time.Minute)
	for {
		mu.Lock()
		myRepos = getRepos(tok)
		mu.Unlock()
		<-ticker.C
	}
}
