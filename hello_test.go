package main

import "testing"

func Test(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		hello := Hello("hahaha")
		if hello != "hello, hahaha" {
			t.Errorf("xxxxxxxx")
		}
	})
}