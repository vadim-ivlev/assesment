package main

import (
	"reflect"
	"testing"
)

func TestAnalyzeGrades_Example1(t *testing.T) {
	input := "15 25 35 45 55 65 75 85 95"
	want := []string{
		"0–20: 1 оценок, 11.1%",
		"21–40: 2 оценок, 22.2%",
		"41–60: 2 оценок, 22.2%",
		"61–80: 2 оценок, 22.2%",
		"81–100: 2 оценок, 22.2%",
	}
	got := AnalyzeGrades(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Example1 failed.\nGot: %v\nWant: %v", got, want)
	}
}

func TestAnalyzeGrades_Example2(t *testing.T) {
	input := "10 20 30 40 50 60 70 80 90 100"
	want := []string{
		"0–20: 2 оценок, 20.0%",
		"21–40: 2 оценок, 20.0%",
		"41–60: 2 оценок, 20.0%",
		"61–80: 2 оценок, 20.0%",
		"81–100: 2 оценок, 20.0%",
	}
	got := AnalyzeGrades(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Example2 failed.\nGot: %v\nWant: %v", got, want)
	}
}

func TestAnalyzeGrades_NoValid(t *testing.T) {
	input := "-1 101 200 0 150"
	want := []string{"Нет корректных оценок"}
	got := AnalyzeGrades(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("NoValid failed.\nGot: %v\nWant: %v", got, want)
	}
}
