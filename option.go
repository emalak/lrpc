package lrpc

type optionCode int

type Option func(args ...any) optionCode

const (
	withFeedClient = iota
)
