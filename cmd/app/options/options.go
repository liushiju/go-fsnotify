/*
author: @liushiju
time: 2023-04-11
*/

package options

import (
	"os"
	"strings"

	pixiuConfig "github.com/caoyingjunz/pixiulib/config"
	"github.com/spf13/cobra"

	"github.com/liushiju/go-fsnotify/cmd/app/config"
	"github.com/liushiju/go-fsnotify/pkg/log"
	"github.com/liushiju/go-fsnotify/pkg/util"
)

const (
	defaultConfigFile = "/etc/go-fsnotify/config.yaml"
)

// Options has all the params needed to run a go-fsnotify .
type Options struct {
	// The default values.
	ComponentConfig config.Config

	// ConfigFile is the location of the go-fsnotify server's configuration file.
	ConfigFile string
}

func NewOptions() (*Options, error) {
	return &Options{
		ConfigFile: defaultConfigFile,
	}, nil
}

// Complete completes all the required options
func (o *Options) Complete() error {
	// 配置文件优先级: 默认配置，环境变量，命令行
	if len(o.ConfigFile) == 0 {
		// Try to read config file path from env.
		if cfgFile := os.Getenv("ConfigFile"); cfgFile != "" {
			o.ConfigFile = cfgFile
		} else {
			o.ConfigFile = defaultConfigFile
		}
	}

	c := pixiuConfig.New()
	c.SetConfigFile(o.ConfigFile)
	c.SetConfigType("yaml")
	if err := c.Binding(&o.ComponentConfig); err != nil {
		return err
	}

	// 注册依赖组件
	if err := o.register(); err != nil {
		return err
	}
	return nil
}

// BindFlags binds the pixiu Configuration struct fields
func (o *Options) BindFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.ConfigFile, "configfile", "", "The location of the gopixiu configuration file")
}

func (o *Options) register() error {
	if err := o.registerLogger(); err != nil { // 注册日志
		return err
	}

	return nil
}

func (o *Options) registerLogger() error {
	logType := strings.ToLower(o.ComponentConfig.Default.LogType)
	if logType == "file" {
		// 判断文件夹是否存在，不存在则创建
		if err := util.EnsureDirectoryExists(o.ComponentConfig.Default.LogDir); err != nil {
			return err
		}
	}
	// 注册日志
	log.Register(logType, o.ComponentConfig.Default.LogDir, o.ComponentConfig.Default.LogLevel)

	return nil
}
