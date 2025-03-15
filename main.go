package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/webhook", handleGithubWebhook)
	port := "9090"
	fmt.Println("Listening on port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

type GithubPushEvent struct {
	Repository struct {
		Name string `json:"name"`
	} `json:"repository"`
	Commits []struct {
		Message string `json:"message"`
		Author  struct {
			Name string `json:"name"`
		} `json:"author"`
	} `json:"commits"`
	Sender struct {
		Login string `json:"login"`
	} `json:"sender"`
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("GoCi"))
}

func handleGithubWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Request", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Internal Server Error. Please read log", http.StatusBadRequest)
		fmt.Println(err.Error())
	}
	defer r.Body.Close()

	fmt.Println(string(body))

	var event GithubPushEvent
	if err := json.Unmarshal(body, &event); err != nil {
		http.Error(w, "Internal Server Error. Please read log", http.StatusBadRequest)
		fmt.Println(err.Error())
	}

	eventType := r.Header.Get("X-Github-Event")
	if eventType != "push" {
		w.WriteHeader(http.StatusOK)
		//ignore
		return
	}

	var sb strings.Builder
	sb.WriteString("Received PUSH event.")
	sb.WriteString(fmt.Sprintf("%s uploaded %d commits to Repository %s\n", event.Sender.Login, len(event.Commits), event.Repository.Name))

	for _, commit := range event.Commits {
		author := commit.Author.Name
		if author == "" {
			author = "unknown"
		}
		sb.WriteString(fmt.Sprintf("Author：%s\nDesc: %s\n-----------------------------------\n", author, commit.Message))
	}

	go executeShellScript("runGradle.sh")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success!"))
}
