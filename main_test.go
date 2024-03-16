package main

import (
	"testing"
)

func Test_addTask(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "test",
		},
		{
			name: "test2",
		},
		{
			name: "test3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addTask()
		})
	}
}

func Test_deleteTask(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "test",
		},
		{
			name: "test2",
		},
		{
			name: "test3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteTask()
		})
	}
}
