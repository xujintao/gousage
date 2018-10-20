package main

import "testing"

func Test1(t *testing.T) {
	t.Log("testlog111")
}

func Test2(t *testing.T) {
	t.Fatal("testlog222")
}
