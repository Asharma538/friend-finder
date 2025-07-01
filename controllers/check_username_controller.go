package controllers

import (
	"fmt"
	"net/http"
	"errors"
	"log"
	"io"
	"strings"
	"context"
	"time"
	"github.com/chromedp/chromedp"
)

func CheckIfInstagramUserExists(username string) (bool, error) {
	instagramURL := fmt.Sprintf("https://www.instagram.com/%s/?hl=en", username)
	response, err := http.Get(instagramURL)
	if err != nil {
		return true, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Fatalf("Error: received status code %d for user %s\n", response.StatusCode, username)
		return false, errors.New("something went wrong while checking that username, try again in a bit")
	}

	webpage_content, err := io.ReadAll(response.Body)
	if err != nil {
		return true, err
	}

	webpage_string := string(webpage_content)
	return strings.Contains(webpage_string, "og:title"), nil
}

func CheckIfGitHubUserExists(username string) (bool, error) {
	githubURL := fmt.Sprintf("https://www.github.com/%s", username)
	response, err := http.Get(githubURL)
	if err != nil {
		return true, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return false, nil
	}
	return true, nil
}

func CheckIfRedditUserExists(username string) (bool, error) {
	redditURL := fmt.Sprintf("https://www.reddit.com/user/%s", username)
	response, err := http.Get(redditURL)
	if err != nil {
		return true, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Fatalf("Error: received status code %d for user %s\n", response.StatusCode, username)
		return false, errors.New("something went wrong while checking that username, try again in a bit")
	}

	webpage_content, err := io.ReadAll(response.Body)
	if err != nil {
		return true, err
	}
	webpage_string := string(webpage_content)
	return !strings.Contains(webpage_string, "Sorry"), nil
}

func CheckIfXUserExists(username string) (bool, error) {
	xURL := fmt.Sprintf("https://twitter.com/%s", username)

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var pageContent string

	err := chromedp.Run(ctx,
		chromedp.Navigate(xURL),
		chromedp.Sleep(2*time.Second),
		chromedp.OuterHTML("html", &pageContent),
	)

	if err != nil {
		return true, err
	}

	if strings.Contains(pageContent, "<title>Profile / X</title>") {
		return false, nil
	}

	return true, nil
}

func CheckIfPinterestUserExists(username string) (bool, error) {
	pinterestURL := fmt.Sprintf("https://pinterest.com/%s", username)
	response, err := http.Get(pinterestURL)
	if err != nil {
		return true, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Fatalf("Error: received status code %d for user %s\n", response.StatusCode, username)
		return false, errors.New("something went wrong while checking that username, try again in a bit")
	}

	webpage_content, err := io.ReadAll(response.Body)
	if err != nil {
		return true, err
	}

	webpage_string := string(webpage_content)
	return !strings.Contains(webpage_string, "<title></title>"), nil
}

func CheckIfThreadsUserExists(username string) (bool, error) {
	threadsURL := fmt.Sprintf("https://threads.com/@%s", username)
	response, err := http.Get(threadsURL)
	if err != nil {
		return true, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Fatalf("Error: received status code %d for user %s\n", response.StatusCode, username)
		return false, errors.New("something went wrong while checking that username, try again in a bit")
	}

	webpage_content, err := io.ReadAll(response.Body)
	if err != nil {
		return true, err
	}

	webpage_string := string(webpage_content)
	return !strings.Contains(webpage_string, "<title>Threads • Log in</title>"), nil
}
