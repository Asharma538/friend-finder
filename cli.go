package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/Asharma538/friend-finder/controllers"
	"github.com/Asharma538/friend-finder/models"
)

func addAllPlatformsForCLI(u *models.UsernameChecker) {
	u.AddPlatform("X/Twitter", controllers.CheckIfXUserExists)
	u.AddPlatform("GitHub", controllers.CheckIfGitHubUserExists)
	u.AddPlatform("Instagram", controllers.CheckIfInstagramUserExists)
	u.AddPlatform("Reddit", controllers.CheckIfRedditUserExists)
	u.AddPlatform("Pinterest", controllers.CheckIfPinterestUserExists)
	u.AddPlatform("Threads", controllers.CheckIfThreadsUserExists)
	u.AddPlatform("LinkedIn", controllers.CheckIfLinkedInUserExists)
	// u.AddPlatform("Facebook", controllers.CheckIfFacebookUserExists)
}

func runCLI() {
	var provided_username string
	fmt.Print("Enter username: ")
	fmt.Scanln(&provided_username)

	username_checker := models.UsernameChecker{}
	addAllPlatformsForCLI(&username_checker)

	wg := &sync.WaitGroup{}

	for _, p := range username_checker.GetPlatforms() {
		wg.Add(1)
		go func(p models.Platform) {
			defer wg.Done()
			exists, err := p.UsernameChecker(provided_username)
			if err != nil {
				log.Printf("%s | Error checking username: %v", p.Name, err)
				return
			}
			if exists {
				log.Printf("%s | Username not available", p.Name)
				return
			}
			log.Printf("%s | Username available", p.Name)
		}(p)
	}
	wg.Wait()

	log.Println("\n\nUsername check completed.")
}
