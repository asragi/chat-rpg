package main

// Enemy は敵キャラクターを表す構造体
type Enemy struct {
	name string
	hp   *HP
}

// NewEnemy は新しい敵を生成する
func NewEnemy(name string, hpValue int) *Enemy {
	return &Enemy{
		name: name,
		hp:   NewHP(hpValue),
	}
}

// HP を取得する
func (e *Enemy) HP() *HP {
	return e.hp
}

// 敵の名前を取得する
func (e *Enemy) Name() string {
	return e.name
}

// 敵にダメージを与える
func (e *Enemy) TakeDamage(amount int) {
	if amount > e.hp.Value() {
		amount = e.hp.Value() // HPを超えるダメージは無視
	}
	e.hp = NewHP(e.hp.Value() - amount)
}
