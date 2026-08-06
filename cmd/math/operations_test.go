package math

import (
	"testing"
	// go get github.com/stretchr/testify
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	got := Sum(2, 3)
	want := 5

	if got != want {
		t.Errorf("Sum(2, 3) = %d; хотим %d", got, want)
	}
}

/*
Самый узнаваемый идиом Go — табличные тесты (table‑driven).
Идея: описываем набор случаев как срез структур, затем в цикле прогоняем один и тот же
код проверки. Так мы не копируем логику, а лишь добавляем строки в таблицу.
*/

func TestSum_Table(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{name: "положительные", a: 2, b: 3, want: 5},
		{name: "с нулём", a: 0, b: 7, want: 7},
		{name: "отрицательные", a: -4, b: -6, want: -10},
		{name: "разные знаки", a: -5, b: 8, want: 3},
	}

	// t.Run(name, func) создаёт подтест со своим именем. Это даёт аккуратный вывод (каждый случай отдельной строкой)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Sum(%d, %d) = %d; хотим %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	got, err := Divide(10, 2)
	if err != nil {
		t.Fatalf("неожиданная ошибка: %v", err)
	}

	if got != 5 {
		t.Errorf("Divide(10,2) = %d; хотим 5", got)
	}
}

/*
Это первое, что должны усвоить студенты. Оба пакета содержат одинаковый набор
функций с одинаковыми именами. Разница — в поведении при провале:
• assert — помечает тест проваленным, но продолжает выполнение (аналог t.Error). Увидите все несовпадения сразу.
• require — помечает провал и немедленно останавливает тест (аналог t.Fatal). Дальше код не идёт.
*/
func TestDivideTestify(t *testing.T) {
	got, err := Divide(10, 2)

	require.NoError(t, err) // если ошибка — стоп с понятным сообщением
	assert.Equal(t, 5, got) // сравнение; при провале покажет want/got
}
