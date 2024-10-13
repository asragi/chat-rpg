package main

// Action はプレイヤーが実行できる行動を表す構造体
type Action struct {
	name     string
	cooldown *Cooldown // 行動に必要なクールタイム
}

// NewAction は新しいアクションを生成する
func NewAction(name string, cooldownValue int) *Action {
	return &Action{
		name:     name,
		cooldown: NewCooldown(cooldownValue),
	}
}

// アクションの名前を取得する
func (a *Action) Name() string {
	return a.name
}

// クールタイムの値を取得する
func (a *Action) CooldownValue() int {
	return a.cooldown.Value()
}

// アクションを実行する
func (a *Action) Execute() {
	a.cooldown.Start() // アクション実行時にクールタイムを開始
}

// アクションのクールタイムを更新する
func (a *Action) Update() {
	a.cooldown.Update()
}

// 実行可能か確認する
func (a *Action) CanExecute() bool {
	return !a.cooldown.IsActive()
}
