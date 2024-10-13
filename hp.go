package main

// HP はプレイヤーのHPを表す
type HP struct {
	value int
}

// NewHP は新しいHPを生成する
func NewHP(value int) *HP {
	return &HP{value: value}
}

// 減少させる（例えばダメージを受ける）
func (hp *HP) Decrease(amount int) {
	hp.value -= amount
	if hp.value < 0 {
		hp.value = 0
	}
}

// 現在のHPを取得する
func (hp *HP) Value() int {
	return hp.value
}
