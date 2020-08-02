# confsyncer

[![](https://img.shields.io/badge/language-English-333333.svg?longCache=true&style=flat-square&colorA=E62B1E)](README.md)

A little sync files tool in the **Linux**.

## What's 

`confsyncer` is a little tool about push and pull files in git repo,  基于此 你可以 很方便的将一些配置文件分发到多台机器上.

## Features

- 手动 Push 监视中的 文件到指定的 git 仓库
- 自动/手动 从指定的 git 仓库 拉取文件到本地的指定位置
- 基于 docker-compose 的一键式部署

## Install
```shell
# use confsyncer with bin
$ wget https://github.com/Kuri-su/confSyncer/eleases/download/v0.0.1/confsyncer-0.0.1-amd64
```

#### or
```shell
# use confsyncer with docker-compose and confsyncer-gen
$ wget https://github.com/Kuri-su/confSyncer/eleases/download/v0.0.1/confsyncer-0.0.1-amd64 
$ wget https://github.com/Kuri-su/confSyncer/eleases/download/v0.0.1/confsyncer-gen-0.0.1-amd64
$ chmod +x confsyncer-gen-0.0.1-amd64 confsyncer-0.0.1-amd64
# install
$ sudo cp confsyncer-gen-0.0.1-amd64 /usr/local/bin/confsyncer-gen 
$ sudo cp confsyncer-0.0.1-amd64     /usr/local/bin/confsyncer
```

## How to Use

### In Host

1. init confsyncer's config

    ```shell
    # init confsyncer's config
    
    $ confsyncer init
    Success! 
    
    Create config file in ~/.confsyncer/config.yaml 
    You should modify it before use.  
    
    This is your config: 
    {
        "gitpulltimeinternal": 600,
        "maps": [
            {
                "gitRepoPath": "/.confsyncer/config.yaml",
                "local": "~/.confsyncer/config.yaml"
            }
        ],
        "gitrepo": ""
    } 
    ```

2. modify config file `~/.confsyncer/config.yaml `

    ```shell
    $ vim ~/.confsyncer/config.yaml
    ```

3. create the git repo and push your config files to repo

   ```shell
   # show confsyncer's config
   $ confsyncer config
   This is your config: 
   {
       "gitrepo": "git@github.com:Kurisu-public/ktx1.git",
       "gitpulltimeinternal": 600,
       "maps": [
           {
               "gitRepoPath": "/.confsyncer/config.yaml",
               "local": "~/.confsyncer/config.yaml"
           }
       ]
   } 
   
   # commit && push files to git repo
   $ confsyncer push
   copy '~/.confsyncer/config.yaml' to '/tmp/confsyncer-20200713/.confsyncer/config.yaml' success
   ```

4. add crontab task to pull config every 1 minute

    ```shell
    # set crontab task
    * * * * * bash -c "confsyncer pull"
    ```

5. 或者你也可以手动的拉取

    ```shell
    # pull config files
    $ confsyncer pull
    copy '~/.confsyncer/config.yaml' to '/tmp/confsyncer-20200713/.confsyncer/config.yaml' success
    ```

### In Container 

1. init confsyncer's config 

   ```shell
   # init confsyncer's config
   $ confsyncer init
   Success! 
   
   Create config file in ~/.confsyncer/config.yaml 
   You should modify it before use.  
   
   This is your config: 
   {
       "gitpulltimeinternal": 600,
       "maps": [
           {
               "gitRepoPath": "/.confsyncer/config.yaml",
               "local": "~/.confsyncer/config.yaml"
           }
       ],
       "gitrepo": ""
   } 
   ```

2. modify config file `~/.confsyncer/config.yaml `
    ```shell
    $ vim ~/.confsyncer/config.yaml
    ```

3. run confsyncer-gen with you confsyncer config

   ```shell
   # run (然后使用 docker-compose 启动服务即可)
   $ confsyncer-gen && docker-compose up -f ~/.confsyncer/docker-compose.yaml -d 
   ```


## docker-compose.yaml Gen

-f config path

-o output path

// TODO docs

## LICENSE

GPL-3.0
