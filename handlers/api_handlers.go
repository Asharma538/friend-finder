package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/Asharma538/friend-finder/controllers"
	"github.com/Asharma538/friend-finder/models"
)

type CheckUsernameRequest struct {
	Username string `json:"username"`
}

type PlatformResult struct {
	Platform  string `json:"platform"`
	Available bool   `json:"available"`
	Error     string `json:"error,omitempty"`
}

type CheckUsernameResponse struct {
	Username string           `json:"username"`
	Results  []PlatformResult `json:"results"`
}

func CheckUsernameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CheckUsernameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	// Initialize username checker
	usernameChecker := models.UsernameChecker{}
	addAllPlatforms(&usernameChecker)

	// Check all platforms concurrently
	var wg sync.WaitGroup
	var mutex sync.Mutex
	results := make([]PlatformResult, 0, len(usernameChecker.GetPlatforms()))

	for _, platform := range usernameChecker.GetPlatforms() {
		wg.Add(1)
		go func(p models.Platform) {
			defer wg.Done()

			available, err := p.UsernameChecker(req.Username)

			result := PlatformResult{
				Platform:  p.Name,
				Available: !available, // Invert because the function returns true if user exists
			}

			if err != nil {
				result.Error = err.Error()
			}

			mutex.Lock()
			results = append(results, result)
			mutex.Unlock()
		}(platform)
	}

	wg.Wait()

	response := CheckUsernameResponse{
		Username: req.Username,
		Results:  results,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func addAllPlatforms(u *models.UsernameChecker) {
	u.AddPlatform("X/Twitter", controllers.CheckIfXUserExists)
	u.AddPlatform("GitHub", controllers.CheckIfGitHubUserExists)
	u.AddPlatform("Instagram", controllers.CheckIfInstagramUserExists)
	u.AddPlatform("Reddit", controllers.CheckIfRedditUserExists)
	u.AddPlatform("Pinterest", controllers.CheckIfPinterestUserExists)
	u.AddPlatform("Threads", controllers.CheckIfThreadsUserExists)
	u.AddPlatform("LinkedIn", controllers.CheckIfLinkedInUserExists)
	// u.AddPlatform("Facebook", controllers.CheckIfFacebookUserExists)
}