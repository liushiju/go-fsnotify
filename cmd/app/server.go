/*
author: @liushiju
time: 2023-04-11
*/

package app

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"

	"github.com/liushiju/go-fsnotify/cmd/app/options"
	notify "github.com/liushiju/go-fsnotify/pkg/fsnotify"
)

func NewServerCommand() *cobra.Command {
	opts, err := options.NewOptions()
	if err != nil {
		klog.Fatalf("unable to initialize command options: %v", err)
	}

	cmd := &cobra.Command{
		Use:  "go-fsnotify-server",
		Long: "The go-fsnotify server controller is a daemon that embeds the core control loops.",
		Run: func(cmd *cobra.Command, args []string) {
			if err = opts.Complete(); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
			if err = Run(opts); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}

	// 绑定命令行参数
	opts.BindFlags(cmd)
	return cmd
}

func Run(opt *options.Options) error {

	// 启动优雅服务
	runServer(opt)
	return nil
}

func runServer(opt *options.Options) {
	dir := opt.ComponentConfig.Default.MonitorDir
	// klog.Infof("监控的目录是: %s ", dir)
	// klog.Infof("到这里就成功一大半了！！！！")
	watch := notify.NewNotifyFile()
	watch.WatchDir(dir)
	select {}
	return
}
