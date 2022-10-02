package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func main() {
	flInit := flag.Bool("init", false, "init gh-mirror")
	flAdd := flag.String("add", "", "add repo to gh-mirror (username/repo)")
	flList := flag.Bool("list", false, "list repos in gh-mirror")
	flag.Parse()
	log.SetFlags(0)

	if *flInit {
		initCmd()
		return
	}

	if *flAdd != "" {
		addCmd(*flAdd)
		return
	}

	if *flList {
		listCmd()
		return
	}

	baseCmd()
}

func initCmd() {
	createRoot()
	createConfig()
}

func addCmd(repo string) {
	if !isValidRepoPath(repo) {
		log.Printf("usage: gh-mirror -add username/repo")
		log.Fatalf("repo path is invalid: %v", repo)
	}
	addRepoToConfig(repo)
}

func listCmd() {
	config := readConfig()
	for _, repo := range config.Repos {
		fmt.Println(repo)
	}
}

func baseCmd() {
	config := readConfig()
	for i, repo := range config.Repos {
		log.SetPrefix(fmt.Sprintf("[%s]: ", repo))

		if !isValidRepoPath(repo) {
			log.Fatalf("repo path is invalid: %v", repo)
		}

		if i > 0 {
			log.Printf("sleeping for %d seconds...\n", config.SleepDuration)
			time.Sleep(time.Duration(config.SleepDuration) * time.Second)
		}

		if notExistRepo(repo) {
			cloneRepo(repo)
		} else {
			updateRepo(repo)
		}
	}
}

func cloneRepo(repo string) {
	changeWorkDirToRoot()
	cloneCmd := "git"
	cloneArgs := []string{"clone", "--mirror", fmt.Sprintf("https://github.com/%s", repo), repo}
	runCommand(cloneCmd, cloneArgs)
}

func updateRepo(repo string) {
	changeWorkDirToRepo(repo)
	updateCmd := "git"
	updateArgs := []string{"remote", "-v", "update"}
	runCommand(updateCmd, updateArgs)
}

func changeWorkDirToRoot() {
	dir := getRootPath()
	err := os.Chdir(dir)
	if err != nil {
		log.Fatalf("could not change work directory: %v", err)
	}
	log.Printf("change work directory to %v\n", dir)
}

func changeWorkDirToRepo(repo string) {
	path := filepath.Join(getRootPath(), repo)
	err := os.Chdir(path)
	if err != nil {
		log.Fatalf("could not change work directory: %v", err)
	}
	log.Printf("change work directory to %v\n", path)
}

func runCommand(cmd string, args []string) {
	log.Printf("%s %s\n", cmd, strings.Join(args, " "))
	output, err := exec.Command(cmd, args...).CombinedOutput()
	if err != nil {
		log.Printf("%s\v", string(output))
		log.Fatalf("fatal: could not run command: %v", err)
	}
	log.Println(string(output))
}

func notExistRepo(repo string) bool {
	path := filepath.Join(getRootPath(), repo)
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

var reValidRepoPath = regexp.MustCompile(`^[\w-]+/[\w-]+$`)

func isValidRepoPath(repo string) bool {
	return reValidRepoPath.MatchString(repo)
}
