/*
author: @liushiju
time: 2023-04-11
*/

package fsnotify

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/liushiju/go-fsnotify/pkg/log"

	notify "github.com/fsnotify/fsnotify"
	"github.com/liushiju/go-fsnotify/pkg/modify"
)

type NotifyFile struct {
	watch *notify.Watcher
}

func NewNotifyFile() *NotifyFile {
	w := new(NotifyFile)
	w.watch, _ = notify.NewWatcher()
	return w
}

// 监控目录
func (this *NotifyFile) WatchDir(dir string) {
	//通过Walk来遍历目录下的所有子目录
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		//判断是否为目录，监控目录,目录下文件也在监控范围内，不需要加
		if info.IsDir() {
			path, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			err = this.watch.Add(path)
			if err != nil {
				return err
			}
			log.Logger.Infof("monitor %v ", path)
		}
		return nil
	})

	go this.WatchEvent() //协程
}

func (this *NotifyFile) WatchEvent() {
	for {
		select {
		case ev := <-this.watch.Events:
			{
				if ev.Op&notify.Create == notify.Create {
					log.Logger.Infof("create directory or file %v ", ev.Name)
					//获取新创建文件的信息，如果是目录，则加入监控中
					file, err := os.Stat(ev.Name)
					if err == nil && file.IsDir() {
						this.watch.Add(ev.Name)
						log.Logger.Infof("add monitor %v ", ev.Name)
					}
				}
				if ev.Op&notify.Write == notify.Write {
					modify.ChmodFile(ev.Name)

				}
				if ev.Op&notify.Remove == notify.Remove {
					log.Logger.Warnf("delete file %v ", ev.Name)
					//如果删除文件是目录，则移除监控
					fi, err := os.Stat(ev.Name)
					if err == nil && fi.IsDir() {
						this.watch.Remove(ev.Name)
						log.Logger.Warnf("delete monitor %v", ev.Name)
					}
				}

				// if ev.Op&fsnotify.Rename == fsnotify.Rename {
				// 	//如果重命名文件是目录，则移除监控 ,注意这里无法使用os.Stat来判断是否是目录了
				// 	//因为重命名后，go已经无法找到原文件来获取信息了,所以简单粗爆直接remove
				// 	// fmt.Println("重命名文件 : ", ev.Name)
				// 	log.Logger.Infof("rename file %v", ev.Name)
				// 	this.watch.Remove(ev.Name)
				// }
				// if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
				// 	file_info, err := os.Stat(ev.Name)
				// 	if err != nil {
				// 		log.Logger.Info("os.Stat failed:", err)
				// 	}
				// 	file_mode := file_info.Mode()
				// 	log.Logger.Infof("file primitive attribute %v: %v", ev.Name, file_mode)
				// 	os.Chmod(ev.Name, 0600)
				// 	log.Logger.Warnf("modify %v attribute : %v", ev.Name, file_mode)
				// 	// fmt.Println(                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     "修改权限 : ", ev.Name)
				// }
			}
			// fileinfo, _ := os.Stat(ev.Name)
			// log.Logger.Infof("修改文件%v权限完成，文件类型为%v", ev.Name, fileinfo.Mode())
		case err := <-this.watch.Errors:
			{
				fmt.Println("error : ", err)
				return
			}
		}
	}

}
