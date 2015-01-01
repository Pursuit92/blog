package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Pursuit92/github"
)

func main() {
	ch, err := github.ReceiveHooks(":8080")
	if err != nil {
		log.Fatal(err)
	}
	for out := range ch {
		fmt.Print(json.MarshalIndent(out, "", "  "))
		/*
			repo := out.Repository.Name
			if repo == "" {
				return
			}
			cwd := "./" + repo
			gitPull := exec.Command("/usr/bin/git", "pull")
			hugo := exec.Command("/usr/local/bin/hugo", "-s", cwd)
			gitPull.Dir = cwd
			hugo.Dir = cwd
			err := gitPull.Run()
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Git ran successfully!")
			err = hugo.Run()
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Hugo ran successfully!")
		*/
	}
}
