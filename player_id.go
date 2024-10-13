package main

// PlayerID はプレイヤーの識別子を表すValueObject
type PlayerID struct {
	value string
}

// NewPlayerID は新しいプレイヤーIDを生成する
func NewPlayerID(value string) *PlayerID {
	return &PlayerID{value: value}
}

// 値を取得する
func (p *PlayerID) Value() string {
	return p.value
}
