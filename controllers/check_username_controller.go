package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

var redditUsername = os.Getenv("REDDIT_USERNAME")

func CheckIfInstagramUserExists(username string) (bool, error) {
	instagramURL := fmt.Sprintf("https://www.instagram.com/%s/?hl=en", username)
	response, err := http.Get(instagramURL)
	if err != nil {
		return true, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Printf("INSTAGRAM | Error: received status code %d for user %s\n", response.StatusCode, username)
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
		log.Printf("GITHUB | Error fetching user %s: %v\n", username, err)
		return true, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return false, nil
	}
	return true, nil
}

func CheckIfRedditUserExists(username string) (bool, error) {
	redditURL := fmt.Sprintf("https://www.reddit.com/api/username_available.json?user=%s", username)
	req, _ := http.NewRequest("GET", redditURL, nil)
	req.Header.Set("User-Agent", fmt.Sprintf("username-checker/1.0 (by /%s)", redditUsername))
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("REDDIT | Error making request for user %s: %v\n", username, err)
		return true, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return false, nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("REDDIT | Error reading response body for user %s: %v\n", username, err)
		return true, err
	}
	if string(body) == "false"{
		return true, nil
	} else if string(body) == "true" {
		return false, nil
	}
	log.Printf("REDDIT | Unexpected response for user %s: %s\n", username, body)
	return true, errors.New("unexpected response from Reddit API")
}

func CheckIfXUserExists(username string) (bool, error) {
	xURL := fmt.Sprintf("https://x.com/%s", username)

	// Create context with cookie handling disabled
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("disable-cookies", true),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var pageContent string

	err := chromedp.Run(ctx,
		chromedp.Navigate(xURL),
		chromedp.Sleep(2*time.Second),
		chromedp.OuterHTML("html", &pageContent),
	)

	if err != nil {
		log.Printf("X/Twitter | Error fetching user %s: %v\n", username, err)
		return true, err
	}

	if strings.Contains(pageContent, "Account suspended") {
		return true, nil
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
		log.Printf("PINTEREST | Error fetching user %s: %v\n", username, err)
		return true, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Printf("PINTEREST | Error: received status code %d for user %s\n", response.StatusCode, username)
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
		log.Printf("THREADS | Error fetching user %s: %v\n", username, err)
		return true, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Printf("THREADS | Error: received status code %d for user %s\n", response.StatusCode, username)
		return false, errors.New("something went wrong while checking that username, try again in a bit")
	}

	webpage_content, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("THREADS | Error reading response body for user %s: %v\n", username, err)
		return true, err
	}

	webpage_string := string(webpage_content)
	return !strings.Contains(webpage_string, "<title>Threads • Log in</title>"), nil
}

func CheckIfLinkedInUserExists(username string) (bool, error) {
	linkedinURL := fmt.Sprintf("https://www.googleapis.com/customsearch/v1?key=%s&cx=22f7f4f700d6e4f09&q=site:linkedin.com/in/%s", os.Getenv("SEARCH_ENGINE_KEY"), username)
	response, err := http.Get(linkedinURL)
	if err != nil {
		log.Printf("LINKEDIN | Error fetching user %s: %v\n", username, err)
		return true, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Printf("LINKEDIN | Error: received status code %d for user %s\n", response.StatusCode, username)
		return true, errors.New("something went wrong while checking that username, try again in a bit")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return true, err
	}

	
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return true, err
	}

	if items, exists := result["items"]; !exists || len(items.([]any)) == 0 {
		return false, nil
	}
	for _, item := range result["items"].([]any) {
		if itemMap, ok := item.(map[string]any); ok {
			if pageMap, exists := itemMap["pagemap"].(map[string]any); exists {
				if metaTags, exists := pageMap["metatags"].([]any); exists {
					for _, metaTag := range metaTags {
						if metaTagMap, ok := metaTag.(map[string]any); ok {
							if _, exists := metaTagMap["og:url"].(string); exists {
								// checking if og:url field exists, not always correct, but 90% accurate
								return true, nil
								// // For checking the exact url
								// if strings.HasSuffix(ogURL, "/in/"+username) {
								// 	return true, nil
								// }
							}
						}
					}
				}
			}
		}
	}
	return false, nil
}