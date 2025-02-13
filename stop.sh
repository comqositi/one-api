#!/bin/bash

prefix='kpai_server_oneapi'

# 查找包含指定前缀的进程
ps -ef | grep "$prefix" | grep -v grep | while read -r line; do
    # 提取 PID
    pid=$(echo "$line" | awk '{print $2}')

    # 检查进程是否存在
    if ps -p "$pid" > /dev/null; then
        echo "当前PID=$pid"

        # 准备结束进程
        echo "PID=$pid 准备结束"
        kill -9 "$pid"

        # 检查进程是否已经被结束
        if ! ps -p "$pid" > /dev/null; then
            echo "PID=$pid 已经结束"
        else
            echo "PID=$pid 结束失败"
        fi
    else
        echo "PID=$pid 不存在"
    fi
done

# 再次查找包含指定前缀的进程，以确认它们是否已被结束
ps -ef | grep "$prefix" | grep -v grep