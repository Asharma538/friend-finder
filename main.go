package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/Asharma538/friend-finder/models"
	"github.com/Asharma538/friend-finder/controllers"
)

func main() {
	fmt.Print("Welcome to the Free Username Finder! \n You can check the availability of usernames across multiple platforms at once.\n\n")

	var provided_username string
	fmt.Print("Enter username: ")
	fmt.Scanln(&provided_username)

	username_checker := models.UsernameChecker{}

	username_checker.AddPlatform(	"X/Twitter",		controllers.CheckIfXUserExists			)
	username_checker.AddPlatform(	"GitHub",			controllers.CheckIfGitHubUserExists		)
	username_checker.AddPlatform(	"Instagram", 		controllers.CheckIfInstagramUserExists	)
	username_checker.AddPlatform(	"Reddit", 			controllers.CheckIfRedditUserExists		)
	username_checker.AddPlatform(	"Pinterest", 		controllers.CheckIfPinterestUserExists	)
	username_checker.AddPlatform(	"Threads", 			controllers.CheckIfThreadsUserExists	)

	wg := &sync.WaitGroup{}

	for _, p := range username_checker.GetPlatforms() {
		wg.Add(1)
		go func(p models.Platform) {
			defer wg.Done()
			exists, err := p.UsernameChecker(provided_username)
			if err != nil {
				log.Fatalf("%s | Error checking username: %v", p.Name, err)
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
