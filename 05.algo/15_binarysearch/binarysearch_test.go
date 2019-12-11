package main

import "testing"

func TestBinarySearch(t *testing.T) {
	a := []int{1, 3, 5, 6, 8}
	if 4 != BinarySearch(a, 8) {
		t.Fatal(BinarySearch(a, 8))
	}
	if -1 != BinarySearch(a, 4) {
		t.Fatal(BinarySearch(a, 4))
	}
}

func TestBinarySearchRecursive(t *testing.T) {
	a := []int{1, 3, 5, 6, 8}
	if 4 != BinarySearchRecursive(a, 8) {
		t.Fatal(BinarySearch(a, 8))
	}
	if -1 != BinarySearchRecursive(a, 4) {
		t.Fatal(BinarySearch(a, 4))
	}
}

func TestBinarySearchFirst(t *testing.T) {
	a := []int{1, 2, 2, 2, 3, 4}
	if 1 != BinarySearchFirst(a, 2) {
		t.Fatal(BinarySearchFirst(a, 2))
	}
	a = []int{1, 2, 2, 2, 3, 4}
	if 4 != BinarySearchFirst(a, 3) {
		t.Fatal(BinarySearchFirst(a, 3))
	}
}

func TestBinarySearchLast(t *testing.T) {
	a := []int{1, 2, 2, 2, 3, 4}
	if 3 != BinarySearchLast(a, 2) {
		t.Fatal(BinarySearchLast(a, 2))
	}
	a = []int{1, 2, 2, 2, 3, 4}
	if 4 != BinarySearchLast(a, 3) {
		t.Fatal(BinarySearchLast(a, 3))
	}
}

func TestBinarySearchFirstGT(t *testing.T) {
	a := []int{1, 2, 2, 2, 3, 4}
	if 4 != BinarySearchFirstGT(a, 2) {
		t.Fatal(BinarySearchFirstGT(a, 2))
	}
	if 1 != BinarySearchFirstGT(a, 1) {
		t.Fatal(BinarySearchFirstGT(a, 1))
	}
	if 5 != BinarySearchFirstGT(a, 3) {
		t.Fatal(BinarySearchFirstGT(a, 3))
	}
	if -1 != BinarySearchFirstGT(a, 4) {
		t.Fatal(BinarySearchFirstGT(a, 4))
	}
}

func TestBinarySearchLastLT(t *testing.T) {
	a := []int{1, 2, 2, 2, 3, 4}
	if 0 != BinarySearchLastLT(a, 2) {
		t.Fatal(BinarySearchLastLT(a, 2))
	}
	if -1 != BinarySearchLastLT(a, 1) {
		t.Fatal(BinarySearchLastLT(a, 1))
	}
	if 3 != BinarySearchLastLT(a, 3) {
		t.Fatal(BinarySearchLastLT(a, 3))
	}
	if 4 != BinarySearchLastLT(a, 4) {
		t.Fatal(BinarySearchLastLT(a, 4))
	}
}
