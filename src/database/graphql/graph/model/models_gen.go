// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type GithubLicense struct {
	Key     string  `json:"key"`
	Name    string  `json:"name"`
	SpdxID  *string `json:"spdxId"`
	HTMLURL string  `json:"htmlURL"`
}

type Link struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Timestamp struct {
	Time time.Time `json:"Time"`
}

type VersionInfo struct {
	Name      string  `json:"name"`
	Version   string  `json:"version"`
	Commit    string  `json:"commit"`
	Date      string  `json:"date"`
	Command   string  `json:"command"`
	GoVersion string  `json:"goVersion"`
	Os        string  `json:"os"`
	Arch      string  `json:"arch"`
	Links     []*Link `json:"links"`
}
