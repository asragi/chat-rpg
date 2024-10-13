package main

import (
	"errors"
	"sync"
)

// InMemoryBackend はメモリ上でプレイヤーを管理するバックエンドの実装
type InMemoryBackend struct {
	players map[string]*Player
	mu      sync.RWMutex
}

// NewInMemoryBackend は新しいInMemoryBackendを生成する
func NewInMemoryBackend() *InMemoryBackend {
	return &InMemoryBackend{
		players: make(map[string]*Player),
	}
}

// プレイヤーを追加する
func (b *InMemoryBackend) AddPlayer(player *Player) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.players[player.ID.Value()] = player // PlayerIDを使用してマップに格納
	return nil
}

// プレイヤーをIDで取得する
func (b *InMemoryBackend) GetPlayer(id *PlayerID) (*Player, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	player, exists := b.players[id.Value()]
	if !exists {
		return nil, errors.New("player not found")
	}
	return player, nil
}
