FROM alpine:3.19

# Define the project name | 定义项目名称
ARG PROJECT=openned8-rpc
# Define the config file name | 定义配置文件名
ARG CONFIG_FILE=openned8.yaml

WORKDIR /app
ENV PROJECT=${PROJECT}
ENV CONFIG_FILE=${CONFIG_FILE}

COPY ./openned8_rpc ./
COPY ./etc/${CONFIG_FILE} ./etc/

EXPOSE 9190

ENTRYPOINT ./openned8_rpc -f etc/${CONFIG_FILE}