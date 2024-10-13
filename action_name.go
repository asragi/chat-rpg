package main

// ActionName は行動の名前を表すクラス
type ActionName struct {
	value string
}

// NewActionName は新しい行動名を生成する
func NewActionName(value string) *ActionName {
	return &ActionName{value: value}
}

// 名前を取得する
func (a *ActionName) Value() string {
	return a.value
}
