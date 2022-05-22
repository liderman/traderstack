package domain

import "time"

type Strategy struct {
	Id          string
	StackId     string
	AccountId   string
	RunInterval time.Duration
	Enabled     bool
}

type StrategyLogType string

const (
	StrategyLogTypeAction StrategyLogType = "action"
	StrategyLogTypeStart  StrategyLogType = "start"
	StrategyLogTypeError  StrategyLogType = "error"
)

type StrategyLog struct {
	LogType StrategyLogType
	Message string
	Time    time.Time
}
