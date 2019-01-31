package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadProfile(t *testing.T) {
	sc := os.Getenv(envSharedCredentialsFile)
	cf := os.Getenv(envAWSConfigFile)
	os.Setenv(envSharedCredentialsFile, filepath.Join("testdata", "test.ini"))
	os.Setenv(envAWSConfigFile, filepath.Join("testdata", "config.ini"))
	defer func() {
		os.Setenv(envSharedCredentialsFile, sc)
		os.Setenv(envAWSConfigFile, cf)
	}()
	f, err := LoadProfile()
	if err != nil {
		t.Error(err)
		return
	}

	contains := func(sections []string, expected string) {
		for _, s := range sections {
			if s == expected {
				return
			}
		}
		t.Errorf("%s is not found. found: %v", expected, sections)
	}
	secs := f.SectionStrings()
	contains(secs, "default")
	contains(secs, "test")
	contains(secs, "xxx")
	contains(secs, "config")
}
