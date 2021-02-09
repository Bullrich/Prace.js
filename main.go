package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Bullrich/prace.go/githubApi"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	eventPath := os.Getenv("GITHUB_EVENT_PATH")

	githubToken := os.Getenv("GITHUB_TOKEN")
	client, err := githubApi.BuildClient(context.Background(), githubToken)

	r := strings.NewReader("my request")
	request, err := http.NewRequest(
		"POST",
		fmt.Sprintf("htpps://github.com/repos/%s/%s/actions/workflows/%s/dispatches", client.Owner, client.Repo, client.WorkflowId),
		r)

	//byteArray := make([]byte, 1024)
	resDo, errDo := client.Client.Do(context.Background(), request, os.Stdout)
	if errDo != nil {
		log.Fatal(errDo)
	}
	_, ioEr := io.Copy(os.Stdout, resDo.Body)
	if ioEr != nil {
		log.Fatal(ioEr)
	}

	repositoryData := strings.Split(os.Getenv("GITHUB_REPOSITORY"), "/")
	owner, repo := repositoryData[0], repositoryData[1]

	fmt.Println("github token length:", len(githubToken))

	fmt.Println(fmt.Sprintf("Owner: %v, repo: %v", owner, repo))

	//github.WebHookPayload{}
	fmt.Println("GITHUB_EVENT_PATH", eventPath)

	content, err := ioutil.ReadFile(eventPath)

	if err != nil {
		fmt.Println("Couldn't read ", eventPath)
		log.Fatal(err)
	}

	//fmt.Println(string(content))

	prData := githubApi.GitHubEvent{}
	er := json.Unmarshal(content, &prData)
	if er != nil {
		log.Fatal(er)
	}

	fmt.Println("PR DATA!", fmt.Sprintf("%+v", prData))
}
