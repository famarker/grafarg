package models

import (
	"os"
)

type InstalledPlugin struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`

	Info         PluginInfo   `json:"info"`
	Dependencies Dependencies `json:"dependencies"`
}

type Dependencies struct {
	GrafargVersion string   `json:"grafargVersion"`
	Plugins        []Plugin `json:"plugins"`
}

type PluginInfo struct {
	Version string `json:"version"`
	Updated string `json:"updated"`
}

type Plugin struct {
	ID       string    `json:"id"`
	Category string    `json:"category"`
	Versions []Version `json:"versions"`
}

type Version struct {
	Commit  string `json:"commit"`
	URL     string `json:"url"`
	Version string `json:"version"`
	// Arch contains architecture metadata.
	Arch map[string]ArchMeta `json:"arch"`
}

type ArchMeta struct {
	SHA256 string `json:"sha256"`
}

type PluginRepo struct {
	Plugins []Plugin `json:"plugins"`
	Version string   `json:"version"`
}

type IoUtil interface {
	Stat(path string) (os.FileInfo, error)
	RemoveAll(path string) error
	ReadDir(path string) ([]os.FileInfo, error)
	ReadFile(filename string) ([]byte, error)
}
