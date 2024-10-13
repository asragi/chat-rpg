package main

// Cooldown はクールタイムを管理する構造体
type Cooldown struct {
	value   int
	counter int
}

// NewCooldown は新しいクールタイムを生成する
func NewCooldown(value int) *Cooldown {
	return &Cooldown{value: value, counter: 0}
}

// クールタイムの値を取得する
func (c *Cooldown) Value() int {
	return c.value
}

// クールタイムを開始する
func (c *Cooldown) Start() {
	c.counter = c.value
}

// クールタイムを更新する
func (c *Cooldown) Update() {
	if c.counter > 0 {
		c.counter--
	}
}

// クールタイムがアクティブか確認する
func (c *Cooldown) IsActive() bool {
	return c.counter > 0
}
