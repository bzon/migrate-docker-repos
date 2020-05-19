package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"gopkg.in/yaml.v2"
)

type Copy struct {
	Tag        string `yaml:"tag"`
	SourceRepo string `yaml:"source_repo"`
	TargetRepo string `yaml:"target_repo"`
}

func main() {
	b, err := ioutil.ReadFile("./config.yaml")
	check(err)

	var cc []Copy

	err = yaml.Unmarshal(b, &cc)
	check(err)

	for _, c := range cc {
		// pull
		cmd := exec.Command("docker", "pull", fmt.Sprintf("%s:%s", c.SourceRepo, c.Tag))
		fmt.Println(cmd.String())
		out, err := cmd.CombinedOutput()
		check(err)
		fmt.Println(string(out))

		// tag
		cmd = exec.Command("docker", "tag", fmt.Sprintf("%s:%s", c.SourceRepo, c.Tag), fmt.Sprintf("%s:%s", c.TargetRepo, c.Tag))
		fmt.Println(cmd.String())
		out, err = cmd.CombinedOutput()
		check(err)
		fmt.Println(string(out))

		// push
		cmd = exec.Command("docker", "push", fmt.Sprintf("%s:%s", c.TargetRepo, c.Tag))
		fmt.Println(cmd.String())
		out, err = cmd.CombinedOutput()
		check(err)
		fmt.Println(string(out))
	}

}

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
