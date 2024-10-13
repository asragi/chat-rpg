package main

// Backend はゲームのバックエンドを表すインターフェース
type Backend interface {
	AddPlayer(player *Player) error          // プレイヤーを追加
	GetPlayer(id *PlayerID) (*Player, error) // プレイヤーをIDで取得
}
