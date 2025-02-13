#!/bin/bash

# 进程名称
prefixName="kpai_server_oneapi"
apiName="kpai_server_oneapi"


ps -ef | grep $prefixName

echo "start apiName ..."
nohup ./${apiName} --log-dir ./logs > ./kpai_server_oneapi.log 2>&1 &


ps -ef | grep $prefixName