package main

// ActionTime はプレイヤーが行動可能になるまでの時間を管理する
type ActionTime struct {
	timeLeft int
}

// NewActionTime は新しいActionTimeを生成する
func NewActionTime(timeLeft int) *ActionTime {
	return &ActionTime{timeLeft: timeLeft}
}

// 行動が可能か確認する
func (at *ActionTime) CanAct() bool {
	return at.timeLeft == 0
}

// クールタイムを進める
func (at *ActionTime) Update() {
	if at.timeLeft > 0 {
		at.timeLeft--
	}
}

// 行動を実行し、クールタイムを設定する
func (at *ActionTime) SetCooldown(cooldown int) {
	at.timeLeft = cooldown
}

// 現在のActionTimeを取得する
func (at *ActionTime) Value() int {
	return at.timeLeft
}
