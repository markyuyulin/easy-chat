#!/bin/bash
reso_addr='crpi-zonjje4rqqueswq8.cn-hangzhou.personal.cr.aliyuncs.com/easy-chat-ns/user-api-dev'
tag='latest'

container_name="easy-chat-user-api-test"

docker stop ${container_name}

docker rm ${container_name}

docker rmi ${reso_addr}:${tag}

docker pull ${reso_addr}:${tag}


# 如果需要指定配置文件的
# docker run -p 10001:8080 --network imooc_easy-chat -v /easy-chat/config/user-rpc:/user/conf/ --name=${container_name} -d ${reso_addr}:${tag}
docker run -p 8888:8888  --name=${container_name} -d ${reso_addr}:${tag}