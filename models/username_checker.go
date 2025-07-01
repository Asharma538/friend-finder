package models

type UsernameChecker struct {
	platforms []Platform
}

func (u *UsernameChecker) AddPlatform(name string, checker func(string) (bool, error)) {
	u.platforms = append(u.platforms, Platform{Name: name, UsernameChecker: checker})
}

func (u *UsernameChecker) GetPlatforms() []Platform {
	return u.platforms
}