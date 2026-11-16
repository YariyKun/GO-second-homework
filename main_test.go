package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name        string
		size        int
		expectedLen int
	}{
		{"Положительный размер", 100, 100},
		{"Нулевой размер", 0, 0},
		{"Отрицательный размер", -5, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := generateRandomElements(tt.size)
			if len(res) != tt.expectedLen {
				t.Errorf("Ожидалась длина %d, получена %d", tt.expectedLen, len(res))
			}
			// Проверка на положительность чисел
			for _, v := range res {
				if v <= 0 {
					t.Errorf("Ожидалось положительное число, получено %d", v)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Пустой слайс", []int{}, 0},
		{"Один элемент", []int{42}, 42},
		{"Максимум в конце", []int{1, 2, 3, 100}, 100},
		{"Максимум в начале", []int{100, 2, 3, 1}, 100},
		{"Отрицательные числа", []int{-10, -5, -20}, -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maximum(tt.input)
			if res != tt.expected {
				t.Errorf("Тест '%s' не пройден: ожидалось %d, получено %d", tt.name, tt.expected, res)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Пустой слайс", []int{}, 0},
		{"Меньше CHUNKS элементов", []int{3, 1, 4, 2}, 4},
		{"Обычный слайс (100 элементов)", func() []int {
			data := make([]int, 100)
			for i := 0; i < 100; i++ {
				data[i] = i
			}
			return data
		}(), 99},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maxChunks(tt.input)
			if res != tt.expected {
				t.Errorf("Тест '%s' не пройден: ожидалось %d, получено %d", tt.name, tt.expected, res)
			}
		})
	}
}
