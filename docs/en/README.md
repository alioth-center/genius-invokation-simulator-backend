# 1. Summary

## 1.1 Overview

Here is a simulator of the game Genshin Impact's `Genius Invokation TCG`, which is a backend implementation of the `Genius Invokation TCG`, playing methods referring to the version 3.3 of Genshin Impact, including all the game mechanisms in Genshin Impact, and expanding some content that miHoYo has not released.

> **Attention: This readme file is for developers' reference. If you only want to use this project, you can explore the guide page of this website**

## 1.2 Announcement

The community channel(QQ) of this project is `530924402`, you can also use the [GitHub Discussions](https://github.com/sunist-c/genius-invokation-simulator-backend/discussions) for communication. Welcome to discuss and PR :)

This simulator focuses on providing customized `Genius Invokation TCG` game, such as:

+ Roll 12 dice each round
+ Obtain four cards each round
+ Play with customized roles, cards and rules

As soon as the project is stable, we will optimize and adapt for AI training at first time, we consider to Compatible with `RLcard` when designing interfaces

> **This project does not contain any art materials of Genshin Impact, and the relevant copyright of Genshin Impact belongs to miHoYo**

## 1.3 Technological Stacks and Components

+ golang 1.19 or higher: this is the coding language and running environment of the project
+ gcc: compile cgo requirements caused by embedded sqlite3
+ sqlite3: the persistent component used to store player information and player card set information in this project, it is embedded and does not need to be downloaded

# 2. Badges

## 2.1 Build Status

| Branch | master | dev | release |
| :--: | :--: | :--: | :--: |
| drone-ci | [![Build Status](https://drone.sunist.cn/api/badges/sunist-c/genius-invokation-simulator-backend/status.svg?ref=refs/heads/master)](https://drone.sunist.cn/sunist-c/genius-invokation-simulator-backend) | [![Build Status](https://drone.sunist.cn/api/badges/sunist-c/genius-invokation-simulator-backend/status.svg?ref=refs/heads/dev)](https://drone.sunist.cn/sunist-c/genius-invokation-simulator-backend) | nil |
| github-action | [![gisb-test](https://github.com/sunist-c/genius-invokation-simulator-backend/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/sunist-c/genius-invokation-simulator-backend/actions/workflows/go.yml) | [![gisb-test](https://github.com/sunist-c/genius-invokation-simulator-backend/actions/workflows/go.yml/badge.svg?branch=dev)](https://github.com/sunist-c/genius-invokation-simulator-backend/actions/workflows/go.yml) | nil |

## 2.2 Code Analysis

[![Go Report Card](https://goreportcard.com/badge/github.com/sunist-c/genius-invokation-simulator-backend)](https://goreportcard.com/report/github.com/sunist-c/genius-invokation-simulator-backend)

# 3. Branches

- master: the main branch, will synchronize with the dev branch when important function changes occur
- dev: the branch under development, will undergo frequent changes and new function tests
- release: the stable branch, will only merge functions when major version is updated
- document: the document branch, the source file of this website

# 4. Progress

Please turn to [gisb's feature development](https://github.com/users/sunist-c/projects/2) to explore the progress of this project

# 5. Contribution

> // todo

# 6. Features


- [ ] Basic game playing method
    - [ ] Elemental reaction
    - [x] Roles and skills
    - [ ] Artifacts and weapons
    - [ ] Scenarios and partners
    - [ ] Summons
    - [ ] Cards
    - [x] Element transformation
- [ ] Game expansion
  - [ ] Multiplayer game(players more than 2)
  - [ ] Match game
  - [ ] Creat room
  - [ ] Join room
  - [ ] In-game communication
  - [ ] Viewer mode
  - [ ] Competition mode
  - [ ] Operational record
- [x] Custom Mod support
  - [ ] Go/Lua/Python support
  - [x] Custom roles and cards
  - [x] Custom rules
- [x] Multiple communication protocol connections
  - [x] websocket interface
  - [x] http/https interface
  - [ ] udp/kcp interface
  - [ ] rpc interface
- [ ] Distributed support
  - [ ] Automated deployment
  - [ ] Service registration and service discovery 
  - [ ] Server load balancing
- [ ] Management function
  - [ ] Announcements and notices
  - [x] IP tracking/blocking
  - [x] QPS/TPS limiter