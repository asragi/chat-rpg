package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"log"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

// Game 構造体に敵を追加
type Game struct {
	players            []*Player
	enemies            []*Enemy // 敵キャラクターのリスト
	backend            Backend
	currentPlayerIndex int
}

func NewGame(backend Backend) *Game {
	players := []*Player{
		NewPlayer(NewPlayerID("Player1"), "Player 1", 100, []*Action{NewAction("Attack", 5), NewAction("Defend", 30)}),
		NewPlayer(NewPlayerID("Player2"), "Player 2", 100, []*Action{NewAction("Attack", 5), NewAction("Heal", 100)}),
		NewPlayer(NewPlayerID("Player3"), "Player 3", 100, []*Action{NewAction("Attack", 6), NewAction("Magic", 80)}),
		NewPlayer(NewPlayerID("Player4"), "Player 4", 100, []*Action{NewAction("Attack", 5), NewAction("Steal", 70)}),
	}

	enemies := []*Enemy{
		NewEnemy("Goblin", 50),
		NewEnemy("Orc", 80),
	}

	for _, player := range players {
		backend.AddPlayer(player)
	}

	return &Game{players: players, enemies: enemies, backend: backend, currentPlayerIndex: 0}
}

func (g *Game) Update() error {
	for _, player := range g.players {
		player.Update()
	}

	// クールタイムの更新
	for _, player := range g.players {
		player.ActionTime.Update() // クールタイムを1減らす
	}

	b := inpututil.IsKeyJustPressed(ebiten.KeySpace)
	if b {
		actionIndex := g.currentPlayerIndex % len(g.players[g.currentPlayerIndex].Actions)
		if g.players[g.currentPlayerIndex].TryAction(actionIndex) {
			// 敵にダメージを与える処理
			if len(g.enemies) > 0 {
				damage := 10 // ダメージを固定で設定（今後アクションに基づいて変更）
				g.enemies[0].TakeDamage(damage)
				if g.enemies[0].HP().Value() <= 0 {
					g.enemies = g.enemies[1:] // 敵が倒れたらリストから削除
				}
			}
			g.currentPlayerIndex = (g.currentPlayerIndex + 1) % len(g.players)
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// プレイヤー情報の表示
	for i, player := range g.players {
		playerInfo := fmt.Sprintf("Player %d: %s HP: %d", i+1, player.Name, player.HP.Value())
		// プレイヤー情報をy座標に基づいて異なる位置に表示
		ebitenutil.DebugPrintAt(screen, playerInfo, 10, 10+i*20)
		// クールタイムの表示
		cooldownInfo := fmt.Sprintf("Cooldown: %d", player.ActionTime.Value())
		ebitenutil.DebugPrintAt(screen, cooldownInfo, 200, 10+i*20)
	}

	// 敵情報の表示
	for i, enemy := range g.enemies {
		enemyInfo := fmt.Sprintf("Enemy %d: %s HP: %d", i+1, enemy.Name(), enemy.HP().Value())
		// 敵情報をy座標に基づいて異なる位置に表示、プレイヤー情報の下に表示
		ebitenutil.DebugPrintAt(screen, enemyInfo, 10, 100+i*20)
	}

	// 現在のプレイヤーのターン情報を表示
	currentPlayerMsg := "Current Player: " + g.players[g.currentPlayerIndex].Name
	ebitenutil.DebugPrintAt(screen, currentPlayerMsg, 10, 300)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	backend := NewInMemoryBackend()
	game := NewGame(backend)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("RPG Prototype")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
