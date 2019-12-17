package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"go_exp_atomic_mutex/lib"
	"testing"
)

func TestAtomicMutex(t *testing.T) {
	Convey("Atomic mutex", t, func() {
		res := lib.CountAtomicMutex()
		So(res, ShouldEqual, lib.CountInc*lib.CountGorotine)
	})
}

func TestAtomic(t *testing.T) {
	Convey("Atomic", t, func() {
		res := lib.CountAtomic()
		So(res, ShouldEqual, lib.CountInc*lib.CountGorotine)
	})
}

func TestMutex(t *testing.T) {
	Convey("Mutex", t, func() {
		res := lib.CountMutex()
		So(res, ShouldEqual, lib.CountInc*lib.CountGorotine)
	})
}

func BenchmarkAtomicMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lib.CountAtomicMutex()
	}
}

func BenchmarkAtomic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lib.CountAtomic()
	}
}

func BenchmarkMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lib.CountMutex()
	}
}
