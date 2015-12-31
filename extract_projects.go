package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Project struct {
	Name        string `yaml:"project"`
	Description string
	WebsiteLink string `yaml:"website-link"`
	GithubLink  string `yaml:"github-link"`
	Tags        []string
	Creator     string
	CreatorLink string `yaml:"creator-link"`
}

func getProjects() []Project {
	projects := []Project{}
	data := readProjectsFile()
	err := yaml.Unmarshal(data, &projects)
	if err != nil {
		return []Project{}
	}

	// Reverse `projects` to display newest first
	l := len(projects)
	for i := 0; i < l/2; i++ {
		j := l - i - 1
		projects[i], projects[j] = projects[j], projects[i]
	}

	return projects
}

func readProjectsFile() []byte {
	filename := "projects.yaml"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte{}
	}

	return content
}
