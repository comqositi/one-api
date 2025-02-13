#!/bin/bash

# 配置线上服务器定时任务，每10分钟或者5分钟运行一次
# 定时任务配置  */10 * * * * cd /home/server/kpai_server_common && sh ./check_prod.sh
# 检查进程是否正常，如果不正常，将执行stop.sh后再拉起进程

process_name_api="kpai_server_oneapi"

# 检测进程是否在运行
is_process_running_api() {
    pgrep -f "$process_name_api" > /dev/null
}

# 拉起进程
restart_process() {
    # 停止进程
    # pkill -x "$process_name"
    sh ./stop.sh

    # 等待进程停止
    while is_process_running_api; do
        sleep 1
    done

    # 启动进程
    sh ./start_prod.sh
}

# 检测并重新拉起进程
if is_process_running_api ; then
    echo "进程 $process_name_api 在运行中."
else
    echo "进程 $process_name_api  未运行."
    restart_process
fi