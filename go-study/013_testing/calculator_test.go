package main

import "testing"

// === 基本のテスト ===
func TestAddBasic(t *testing.T) {
	got := Add(2, 3)
	want := 5
	if got != want {
		t.Errorf("Add(2, 3) = %d, want %d", got, want)
	}
}

// === テーブル駆動テスト（Go で最も重要なパターン） ===
func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"正の数同士", 2, 3, 5},
		{"ゼロ同士", 0, 0, 0},
		{"負の数同士", -1, -2, -3},
		{"正と負", 5, -3, 2},
		{"大きな数", 1000000, 2000000, 3000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"正の数", 10, 4, 6},
		{"結果が負", 3, 5, -2},
		{"同じ数", 7, 7, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Subtract(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Subtract(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// === エラーケースのテスト ===
func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a, b    float64
		want    float64
		wantErr bool
	}{
		{"正常な割り算", 10, 3, 3.3333333333333335, false},
		{"割り切れる", 10, 2, 5, false},
		{"0除算", 10, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)

			if tt.wantErr {
				if err == nil {
					t.Error("エラーを期待したが、nil が返された")
				}
				return
			}

			if err != nil {
				t.Errorf("予期しないエラー: %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("Divide(%f, %f) = %f, want %f", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// === テーブル駆動テスト + bool ===
func TestIsPrime(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{7, true},
		{10, false},
		{13, true},
		{100, false},
		{-1, false},
		{0, false},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := IsPrime(tt.input)
			if got != tt.want {
				t.Errorf("IsPrime(%d) = %t, want %t", tt.input, got, tt.want)
			}
		})
	}
}
