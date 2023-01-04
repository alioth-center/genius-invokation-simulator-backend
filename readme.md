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

## Branch

- master 主要的分支，将在发生重要功能修改时与dev分支进行同步
- dev 开发中的分支，将频繁地进行更改与新功能测试
- release 稳定的分支，仅会在重大版本更新时进行功能合并

## Progress

请转到[gisb's feature development](https://github.com/users/sunist-c/projects/2)查看项目进度

## Document

+ 战斗框架： [Battle Framework of GISB](https://github.com/sunist-c/genius-invokation-simulator-backend/wiki/Battle-Framework)
+ 事件框架： Mkdir...

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