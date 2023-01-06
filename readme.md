# Genius Invokation Simulator Backend

----

## Badges

| Branch | master | dev | release |
| :--: | :--: | :--: | :--: |
| drone-ci | [![Build Status](https://drone.sunist.cn/api/badges/sunist-c/genius-invokation-simulator-backend/status.svg?ref=refs/heads/master)](https://drone.sunist.cn/sunist-c/genius-invokation-simulator-backend) | [![Build Status](https://drone.sunist.cn/api/badges/sunist-c/genius-invokation-simulator-backend/status.svg?ref=refs/heads/dev)](https://drone.sunist.cn/sunist-c/genius-invokation-simulator-backend) | nil |
| github-action | [![gisb-test](https://github.com/sunist-c/genius-invokation-simulator-backend/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/sunist-c/genius-invokation-simulator-backend/actions/workflows/go.yml) | [![gisb-test](https://github.com/sunist-c/genius-invokation-simulator-backend/actions/workflows/go.yml/badge.svg?branch=dev)](https://github.com/sunist-c/genius-invokation-simulator-backend/actions/workflows/go.yml) | nil |

## Summary

这里是原神(Genshin Impact)的《七圣召唤》模拟器，是参考原神3.3版本的「七圣召唤」玩法重新实现的后端(服务端)，包括所有的原神内的游戏内容，并拓展一些米哈游没有做的内容。

**具体游戏内容的实现(卡牌、角色、技能等)将会以sub-module的形式发布，本项目仅为后端框架**

## How to use

**本环节(How to use)暂时不可用，只是先更新文档占个坑**

### From source

1. 确保您的设备上正常安装了`go 1.19`或以上版本，您可以通过在命令行执行下面的语句查看您的go版本与安装情况：

  ```shell
  go version
  ```

2. 选择一个存放代码的地方，Clone本仓库，选择一个源即可：

  ```shell
  git clone https://github.com/sunist-c/genius-invokation-simulator-backend.git # GitHub源
  git clone https://code.sunist.cn/sunist-c/genius-invokation-simulator-backend.git # 作者私有源
  ```

3. 进入本项目的目录，然后构建本项目，请根据操作系统情况与实际情况替换掉`${output_name}`：

  ```shell
  go mod tidy # 同步依赖关系
  go build -o ${output_name} # 构建可执行文件
  ```

4. 您可以在当前目录下发现新增了一个名为`${output_name}`的可执行文件，执行这个可执行文件：

  ```shell
  ./${output_name} -mode cli # 以命令行模式启动模拟器
  ```

5. 您可以使用`-h`参数来获取更多执行帮助：

  ```shell
  $ ./${output_name} -h
  > Usage of ${output_name}:
  >   -conf string
  >         setup the backend configuration file, highest priority
  >   -mode string
  >         setup the startup mode, available [backend, cli, ai] (default "backend")
  >   -port uint
  >         setup the http service port (default 8086)
  ```

### From binary

1. 转到本项目的[release]()页面，根据您的设备、平台与操作系统，下载对应的二进制可执行文件，这个文件等价于 `From source` 环节中 3. 步骤产生的可执行文件
2. 参考 `From source` 环节的 4. 和 5. 步骤进行下一步操作

### From docker

// todo

## How to modify

**本环节(How to modify)暂时不可用，只是先更新文档占个坑**

1. 进入 `data` 目录，这是模拟器自动扫描的目录
2. 找到一个您中意的mod，然后将其下载至 `data` 目录下
3. 修改**项目根目录**下的 `main.go` 文件，在imports区块内，将您获取的mod匿名引入
4. 启动模拟器，模拟器将自动加载mod

----

以下的内容供开发者参考

## Branch

- master 主要的分支，将在发生重要功能修改时与dev分支进行同步
- dev 开发中的分支，将频繁地进行更改与新功能测试
- release 稳定的分支，仅会在重大版本更新时进行功能合并

## Progress

请转到[gisb's feature development](https://github.com/users/sunist-c/projects/2)查看项目进度

## Document

+ 战斗框架： [Battle Framework of GISB](https://github.com/sunist-c/genius-invokation-simulator-backend/wiki/Battle-Framework)
+ 事件框架： Mkdir...
+ MOD制作： Mkdir...

## Announce

本模拟器侧重于提供自定义的七圣召唤对局，比如每次投十二个骰子/每回合摸四张牌/加入自定义角色、卡牌等功能，~~短期内没有~~ 项目稳定后尽快针对ai训练进行优化、适配。
考虑设计接口时兼容RLcard。

相关的[genius-invokation-gym](https://github.com/paladin1013/genius-invokation-gym)项目侧重于提供ai相关接口，请根据需求选择。

**本模拟器接口尽量和[genius-invokation-gym](https://github.com/paladin1013/genius-invokation-gym)保持一致，其项目完善后本项目也尽量拓展相应的ai接口。**

同时感谢[@Leng Yue](https://github.com/leng-yue)实现的前端项目[genius-invokation-webui](https://github.com/leng-yue/genius-invokation-webui)。

本项目的交流群(QQ)为`530924402`，欢迎讨论与PR！

## Contribute

如果您想增加一个功能或想法：

1. 加入本项目的交流群或[genius-invokation-gym](https://github.com/paladin1013/genius-invokation-gym)的交流群或在本项目的[Discussion/Ideas](https://github.com/sunist-c/genius-invokation-simulator-backend/discussions/categories/ideas)中分享您的想法
2. 在取得Contributor的广泛认可后将为您创建一个WIP的issue
3. 按照正常的流程fork->coding->pull request您的修改

如果您想为项目贡献代码，您可以转到[gisb's feature development](https://github.com/users/sunist-c/projects/2)查看这个项目目前在干什么，标有`help wanted`的内容可能需要一些帮助，处于`design`阶段的内容目前还没有进行开发，您可以直接在GitHub Projects/Project Issue页面与我们交流

本项目有意向从Jetbrains申请开源项目的All Products License，将会提供给Code Contributors

## Features

- [x] 游戏基本玩法
  - [ ] 元素反应
  - [x] 角色与技能
  - [ ] 圣遗物与武器
  - [ ] 场景和伙伴
  - [ ] 召唤物
  - [x] 卡牌与元素转化
- [ ] 游戏拓展玩法
  - [x] 多人游戏(玩家`N>=2`，队伍`N>=2`)
  - [ ] 创建对局
  - [ ] 匹配对局
  - [ ] 游戏内通信
  - [ ] 观战模式
  - [ ] 比赛模式
  - [ ] 作战记录
- [x] 自定义Mod支持
  - [x] Go/Lua支持
  - [x] 自定义角色与卡牌
  - [x] 自定义规则
- [x] 多种通信协议连接
  - [x] websocket接口
  - [x] http/https接口
  - [ ] udp/kcp接口
  - [ ] rpc接口
- [ ] 分布式支持
  - [ ] 自动化部署
  - [ ] 服务注册与服务发现
  - [ ] 服务端负载均衡
- [ ] 管理功能
  - [ ] 公告与通知
  - [ ] IP追踪/封禁
  - [ ] QPS/TPS限制器

## License

MIT License

**本项目不含任何原神(Genshin Impact)的美术素材，原神(Genshin Impact)的相关版权归米哈游(miHoYo)所有**