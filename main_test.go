package main

import (
	"testing"
	"time"
)

func TestWeekendTrue(t *testing.T) {
	expected := true
	result := weekend(time.Weekday(5))
	if result != expected {
		t.Fatalf("Expected %t, got %t", expected, result)
	}
}

func TestWeekendFalse(t *testing.T) {
	expected := false
	result := weekend(time.Weekday(1))
	if result != expected {
		t.Fatalf("Expected %t, got %t", expected, result)
	}
}
