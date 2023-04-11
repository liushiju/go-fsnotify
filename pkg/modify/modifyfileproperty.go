/*
author: @liushiju
time: 2023-04-11
*/

package modify

import (
	"os"
	"syscall"

	"github.com/liushiju/go-fsnotify/pkg/log"

	"github.com/gabriel-vasile/mimetype"
)

func ChmodFile(filename string) (string, error) {
	allowed := []string{"image/png", "image/jpeg", "application/zip", "application/pdf", "video/mp4", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", "application/vnd.openxmlformats-officedocument.wordprocessingml.document"}
	mtype, _ := mimetype.DetectFile(filename)

	if mimetype.EqualsAny(mtype.String(), allowed...) {
		log.Logger.Infof("allowed file: %v, %v", filename, mtype)
		return mtype.String(), nil
	} else {
		file_info, err := os.Stat(filename)
		if err != nil {
			log.Logger.Info("os.Stat failed:", err)
		}
		file_mode := file_info.Mode()
		log.Logger.Warnf("not allowed file: %v, %v, [ %v ]", filename, mtype.String(), file_mode)
		syscall.Umask(0)
		os.Chmod(filename, 0600)
		file_info, err = os.Stat(filename)
		ch_file_mode := file_info.Mode()
		log.Logger.Infof("The file permission has been changed to [ %v ]", ch_file_mode)
		return mtype.String(), nil
	}

}
