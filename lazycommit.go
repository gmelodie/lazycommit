package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-git/go-git/v5"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	directory := "./"

	// curl http://whatthecommit.com
	resp, err := http.Get("http://whatthecommit.com")
	checkError(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	checkError(err)

	// grep <p> | cut -d ">" -f 2
	commitMsg := strings.Split(string(body), "<p>")[1] // cut before
	commitMsg = strings.Split(commitMsg, "</p>")[0]    // cut after
	commitMsg = strings.Trim(commitMsg, "\n")          // remove newline
	fmt.Println(string(commitMsg))

	// Open the repo
	repo, err := git.PlainOpen(directory)
	checkError(err)
	worktree, err := repo.Worktree()
	checkError(err)

	// git add .
	_, err = worktree.Add(".")
	checkError(err)

	// git commit -m "$(message)"
	_, err = worktree.Commit(commitMsg, &git.CommitOptions{})
	checkError(err)

	// git push
	err = repo.Push(&git.PushOptions{})
	checkError(err)
}
