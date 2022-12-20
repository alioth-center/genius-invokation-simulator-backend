# Genius Invokation Simulator Backend

----

| Branch | master | dev | release |
| :--: | :--: | :--: |
| status | [![Build Status](https://drone.sunist.cn/api/badges/sunist-c/genius-invokation-simulator-backend/status.svg?ref=refs/heads/master)](https://drone.sunist.cn/sunist-c/genius-invokation-simulator-backend) | [![Build Status](https://drone.sunist.cn/api/badges/sunist-c/genius-invokation-simulator-backend/status.svg?ref=refs/heads/dev)](https://drone.sunist.cn/sunist-c/genius-invokation-simulator-backend) | nil |
 

这里是原神(Genshin Impact)的《七圣召唤》模拟器，是参考原神3.3版本的「七圣召唤」玩法重新实现的后端(服务端)，包括所有的原神内的游戏内容，并拓展一些米哈游没有做的内容。

## Branch

- dev分支是重构时硬怼的，是新项目，和master没有同源根，正在分批次同步到master里

## Progress

- 22.12.9 创建新文件夹
- 22.12.11 设计、定义相关enum与接口
- 22.12.13 重构对局生命周期
- 22.12.15 战斗框架基本完成
- 22.12.16 开始编写文档0.0.1版本
- 22.12.16 (试验性)命令行操作对局
- 22.12.19 对部份结构/实体进行了封装与优化，由于改动较大目前还未push。重新设计了HandlerChain的结构，HandlerChain目前的执行性能达到了每秒70万次，每个Chain含有128个HandlerFunc
- 22.12.19 将ModifierChain分类为Local(仅对当前角色生效)和Global(对玩家的被激活角色生效)，由此产生的逻辑更改目前正在调试
- 22.12.20 完成了大部份的ModifierContext和Event，重构、优化了部份entity和model

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

## Features

- [x] 游戏基本玩法
  - [x] 元素反应
  - [x] 角色与技能
  - [ ] 圣遗物与武器
  - [ ] 场景和伙伴
  - [ ] 召唤物
  - [ ] 卡牌与元素转化
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
- [ ] 多种通信协议连接
  - [ ] websocket接口
  - [ ] http/https接口
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