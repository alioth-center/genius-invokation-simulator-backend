# Genius Invokation Simulator Backend

## 1. 综述

### 1.1 概况

这里是原神(Genshin Impact)的《七圣召唤》模拟器，是参考原神3.3版本的「七圣召唤」玩法重新实现的后端(服务端)，包括所有的原神内的游戏内容，并拓展一些米哈游没有做的内容。

**本自述文件供开发者参考，若您只是使用本项目，您可以转到我们的[文档](https://sunist-c.github.io/genius-invokation-simulator-backend/)**

[English Version](https://sunist-c.github.io/genius-invokation-simulator-backend/#/en/)

### 1.2 声明

本项目的交流群(QQ)为`530924402`，欢迎讨论与PR

本模拟器侧重于提供自定义的七圣召唤对局，比如每次投十二个骰子/每回合摸四张牌/加入自定义角色、卡牌等功能，~~短期内没有~~ 项目稳定后尽快针对ai训练进行优化、适配。
考虑设计接口时兼容RLcard

相关的[genius-invokation-gym](https://github.com/paladin1013/genius-invokation-gym)项目侧重于提供ai相关接口，请根据需求选择

**本模拟器接口尽量和[genius-invokation-gym](https://github.com/paladin1013/genius-invokation-gym)保持一致，其项目完善后本项目也尽量拓展相应的ai接口**

同时感谢[@Leng Yue](https://github.com/leng-yue)实现的前端项目[genius-invokation-webui](https://github.com/leng-yue/genius-invokation-webui)

**本项目不含任何原神(Genshin Impact)的美术素材，原神(Genshin Impact)的相关版权归米哈游(miHoYo)所有**

### 1.3 技术栈与组件

+ golang 1.19： 这是本项目的编码语言与运行环境
+ mingw-gcc/clang/gcc： 因为内嵌sqlite3产生的cgo编译需求
+ sqlite3： 这是本项目用于存储玩家信息、玩家卡组信息的持久化组件，已内嵌，无需下载

## 2. 徽章

### 2.1 构建情况

| Branch | master | dev | release |
| :--: | :--: | :--: | :--: |
| drone-ci | [![Build Status](https://drone.sunist.cn/api/badges/sunist-c/genius-invokation-simulator-backend/status.svg?ref=refs/heads/master)](https://drone.sunist.cn/sunist-c/genius-invokation-simulator-backend) | [![Build Status](https://drone.sunist.cn/api/badges/sunist-c/genius-invokation-simulator-backend/status.svg?ref=refs/heads/dev)](https://drone.sunist.cn/sunist-c/genius-invokation-simulator-backend) | nil |
| github-action | [![gisb-test](https://github.com/sunist-c/genius-invokation-simulator-backend/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/sunist-c/genius-invokation-simulator-backend/actions/workflows/go.yml) | [![gisb-test](https://github.com/sunist-c/genius-invokation-simulator-backend/actions/workflows/go.yml/badge.svg?branch=dev)](https://github.com/sunist-c/genius-invokation-simulator-backend/actions/workflows/go.yml) | nil |

### 2.2 代码质量

[![Go Report Card](https://goreportcard.com/badge/github.com/sunist-c/genius-invokation-simulator-backend)](https://goreportcard.com/report/github.com/sunist-c/genius-invokation-simulator-backend)

## 3. 分支说明

- master 主要的分支，将在发生重要功能修改时与dev分支进行同步
- dev 开发中的分支，将频繁地进行更改与新功能测试
- release 稳定的分支，仅会在重大版本更新时进行功能合并
- cli 脚手架分支，将提供一个代码生成脚手架，用于快速实现mod

## 4. 开发进度

请转到[gisb's feature development](https://github.com/users/sunist-c/projects/2)查看项目进度

## 5. 开发文档

您可以转到[genius-invokation-simulator-mod-template](https://github.com/sunist-c/genius-invokation-simulator-mod-template)查看我们的示例mod，并从该模板创建您的mod

+ 战斗框架： [Battle Framework of GISB](https://github.com/sunist-c/genius-invokation-simulator-backend/wiki/Battle-Framework)
+ 事件框架： Mkdir...
+ MOD制作： Mkdir...

## 6. 参与项目

如果您想增加一个功能或想法：

1. 加入本项目的交流群或[genius-invokation-gym](https://github.com/paladin1013/genius-invokation-gym)的交流群或在本项目的[Discussion/Ideas](https://github.com/sunist-c/genius-invokation-simulator-backend/discussions/categories/ideas)中分享您的想法
2. 在取得Contributor的广泛认可后将为您创建一个WIP的issue
3. 按照正常的流程fork->coding->pull request您的修改

如果您想为项目贡献代码，您可以转到[gisb's feature development](https://github.com/users/sunist-c/projects/2)查看这个项目目前在干什么，标有`help wanted`的内容可能需要一些帮助，处于`design`阶段的内容目前还没有进行开发，您可以直接在GitHub Projects/Project Issue页面与我们交流

## 7. 功能特性

转到我们的[文档](https://sunist-c.github.io/genius-invokation-simulator-backend/)查看最新的功能特性

## 8. 开源许可

[MIT License](license)

## 9. 鸣谢

感谢 [JetBrains](https://www.jetbrains.com) 为本项目提供的 [Open Source development license(s)](https://www.jetbrains.com/community/opensource/#support)

<img width="75px" src="https://resources.jetbrains.com.cn/storage/products/company/brand/logos/GoLand_icon.svg" alt="GoLand logo."/><img width="50px" src="https://resources.jetbrains.com.cn/storage/products/company/brand/logos/IntelliJ_IDEA_icon.svg" alt="IntelliJ_IDEA logo."/><img width="50px" src="https://resources.jetbrains.com.cn/storage/products/company/brand/logos/WebStorm_icon.svg" alt="WebStorm logo."/><img width="50px" src="https://resources.jetbrains.com.cn/storage/products/company/brand/logos/PyCharm_icon.svg" alt="PyCharm logo."/>
