# Go-fsnotify Overview

Go-fsnotify is used to monitor the change of directory files, add new files, determine whether the attributes are qualified, and the permissions are unchanged if they are qualified, otherwise the modification permission is 0600.


## Installation

- 下载源码

```bash
git clone https://github.com/liushiju/go-fsnotify.git
cd go-fsnotify
```

- 修改配置文件`config.yaml`
    - example:

    ```bash
    default:
    # 日志类型，支持 stdout, stderr 和 file, 默认为 stdout
    log_type: file
    # 日志级别，支持 INFO, WARN 和 ERROR, 默认为 INFO
    log_level: INFO
    # 日志路径，在日志类型为 file 的时候生效
    log_dir: /home/liushiju/go-fsnotify/logs
    # 配置监控目录
    monitor_dir: /home/liushiju/test_file
    ```

- 编译运行

```bash
make run
```

Copyright 2023 liushiju (lsj_tedu@163.com)