package models

type Platform struct {
	Name            string
	UsernameChecker func(string) (bool, error)
}