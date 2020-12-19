package browser

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type Browser struct {
	// Browser name e.g. firefox or google-chrome-stable.
	Name string
	// Name of profile to create.
	Profile string
}

func (br *Browser) createDirectory() error {

	path := br.profilePath()
	// Only current user can rwx e.g. 0700.
	if err := os.MkdirAll(path, 0700); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (br *Browser) createProfile() error {

	// Example: "JoelUser c:\internet\joelusers-moz-profile"
	profileNameAndDir := fmt.Sprintf("%s %s", br.Profile, br.profilePath())

	cmd := exec.Command(br.Name, "-CreateProfile", profileNameAndDir)
	log.Print(cmd)

	if err := cmd.Run(); err != nil {
		log.Fatal("failed run")
		return err
	}

	return nil
}

func (br *Browser) profileExists() bool {

	path := br.profilePath()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func (br *Browser) profilePath() string {

	suffix := fmt.Sprintf("%s-%s", br.Name, br.Profile)
	path := filepath.Join(os.Getenv("HOME"), ".config", "browser", suffix)

	return path
}

func (br *Browser) Start() error {

	if br.profileExists() == false {

		if err := br.createDirectory(); err != nil {
			log.Fatal(err)
			return err
		}

		if err := br.createProfile(); err != nil {
			log.Fatal(err)
			return err
		}

	}

	cmd := exec.Command(br.Name, "-profile", br.profilePath())
	log.Print(cmd)

	// Use cmd.Start() instead of cmd.Run() so process does
	// not wait for browser to exit before returning.
	if err := cmd.Start(); err != nil {
		log.Fatal("failed run")
		return err
	}

	return nil
}
