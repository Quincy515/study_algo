package main

import (
	"testing"
)

func TestPalindrome1(t *testing.T) {
	strs := []string{"heooeh", "hello", "heoeh", "a", ""}
	for _, str := range strs {
		l := NewLinkedList()
		for _, c := range str {
			l.InsertToTail(string(c))
		}
		l.Print()
		t.Log(isPalindrome1(l))
	}
}

func TestPalindrome2(t *testing.T) {
	strs := []string{"heooeh", "hello", "heoeh", "a", ""}
	for _, str := range strs {
		l := NewLinkedList()
		for _, c := range str {
			l.InsertToTail(string(c))
		}
		l.Print()
		t.Log(isPalindrome2(l))
		l.Print()
	}
}
