/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "context.go" LastUpdatedAt 2022/12/12 15:24:12
 */

package model

import "time"

type Context struct {
	Options           Options
	Rules             RuleSet
	Players           []Player
	ActivePlayerChain PlayerChain
	NextPlayerChain   PlayerChain
}

type Options struct {
	PlayerCount        uint              // PlayerCount 玩家人数
	PlayerTeam         map[uint]uint     // PlayerTeam 玩家阵营，key: player.uid, value: player.team
	StepTimer          time.Duration     // StepTimer 步时计时器
	RoundTimer         time.Duration     // RoundTimer 局时计时器
	RollTimes          uint              // RollTimes 投掷阶段投掷次数
	RoundElements      uint              // RoundElements 每回合投掷产生的元素个数
	RoundCards         uint              // RoundCards 每回合获得卡牌数
	ReShuffleCardStack bool              // ReShuffleCardStack 牌堆消耗后是否重新洗牌
	MaxRound           uint              // MaxRound 最大回合数
	ExtendOptions      map[string]string // ExtendOptions 拓展选项，供开发者调用，key: option.name, value: option.value
}
