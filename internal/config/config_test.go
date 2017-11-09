package config

import "testing"

func init() {
	configPath = "./"
	configType = "yaml"
	configName = "test_config"
}

func TestParseConfigFile(t *testing.T) {
	expectedConfig := Configuration{"foo, bar, baz", "/home/user/tweetfeed.xml", 3, 827142200861609984}

	actualConfig := parseConfigFile()

	if actualConfig != expectedConfig {
		t.Errorf("Failed! Got %q != want %q", actualConfig, expectedConfig)
	}
}

func TestSetHashtags(t *testing.T) {
	config := Configuration{"foo, bar", "./", 0, 0}
	expectedHashtags := "booze, hookers"

	config.SetHashtags("booze, hookers")

	if config.Hashtags != expectedHashtags {
		t.Errorf("Failed! SetHashtags(%q) == %q, want %q", expectedHashtags, config.Hashtags, expectedHashtags)
	}
}
