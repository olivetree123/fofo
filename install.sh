#! /bin/sh

mkdir /etc/fofo
wget https://olivetree.oss-cn-hangzhou.aliyuncs.com/fofo -O /usr/local/bin/fofo
wget https://olivetree.oss-cn-hangzhou.aliyuncs.com/config.toml -O /etc/fofo/config.toml
chmod +x /usr/local/bin/fofo

