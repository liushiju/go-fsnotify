/*
author: @liushiju
time: 2023-04-11
*/

package config

import (
	"fmt"
	"strings"
)

type Config struct {
	Default DefaultOptions `yaml:"default"`
}

type DefaultOptions struct {
	LogType    string `yaml:"log_type"`
	LogDir     string `yaml:"log_dir"`
	LogLevel   string `yaml:"log_level"`
	MonitorDir string `yaml:"monitor_dir"`
}

func (c *Config) Valid() error {
	if strings.ToLower(c.Default.LogType) == "file" {
		if len(c.Default.LogDir) == 0 {
			return fmt.Errorf("log_dir should be config when log type is file")
		}
	}
	return nil
}
