package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Player はプレイヤーのロジックを管理する
type Player struct {
	ID         *PlayerID   // プレイヤーのID
	HP         *HP         // プレイヤーのHP
	ActionTime *ActionTime // 行動に必要なクールタイム
	Actions    []*Action   // プレイヤーが実行できる行動のリスト
	Name       string      // プレイヤーの名前
}

// NewPlayer は新しいプレイヤーを生成する
func NewPlayer(id *PlayerID, name string, hpValue int, actions []*Action) *Player {
	return &Player{
		ID:         id,
		HP:         NewHP(hpValue),
		ActionTime: NewActionTime(0), // 初期クールタイムはゼロとする
		Actions:    actions,
		Name:       name,
	}
}

// プレイヤーのステータスを更新する
func (p *Player) Update() {
	// 各行動のクールタイムを更新
	for _, action := range p.Actions {
		action.Update()
	}
	p.ActionTime.Update() // プレイヤーのクールタイムを更新
}

// プレイヤーの情報を描画する
func (p *Player) Draw(screen *ebiten.Image, position int) {
	msg := fmt.Sprintf("%s HP: %d Cooldown: %d", p.Name, p.HP.Value(), p.ActionTime.Value())
	x, y := 20, 50+position*40
	ebitenutil.DebugPrintAt(screen, msg, x, y)
}

// 特定の行動を実行できるか試す
func (p *Player) TryAction(actionIdx int) bool {
	if actionIdx >= 0 && actionIdx < len(p.Actions) {
		action := p.Actions[actionIdx]
		if action.CanExecute() && p.ActionTime.CanAct() {
			action.Execute()                                 // 行動を実行し、クールタイムをリセット
			p.ActionTime.SetCooldown(action.CooldownValue()) // 行動後のクールタイムを設定
			return true
		}
	}
	return false
}
