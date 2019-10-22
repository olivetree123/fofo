# Fofo
[![Build Status](https://travis-ci.org/olivetree123/fofo.svg?branch=master)](https://travis-ci.org/olivetree123/fofo)  [![codecov](https://codecov.io/gh/olivetree123/fofo/branch/master/graph/badge.svg)](https://codecov.io/gh/olivetree123/fofo)  [![Codacy Badge](https://api.codacy.com/project/badge/Grade/265dc46712694bbdb26052f4572b30b0)](https://www.codacy.com/manual/olivetree123/fofo?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=olivetree123/fofo&amp;utm_campaign=Badge_Grade)  ![GitHub](https://img.shields.io/github/license/olivetree123/fofo)
Fofo is a service for Service Register & Discovery.

## Install
``` shell
wget -O - https://olivetree.oss-cn-hangzhou.aliyuncs.com/fofo/install.sh | bash
```

## Configuration
Fofo depends on Redis and MongoDB. Default configuration is file `/etc/fofo/config.toml`.
``` shell
# /etc/fofo/config.toml

RedisHost = "localhost"
RedisPort = 6379
RedisDB = 0

MongoDBHost = "localhost"
MongoDBPort = 27017
```
