package test

import (
	"hxbk_resp/utils"
	"math/rand"
	"testing"
	"time"
)

func TestRandom(t *testing.T) {
	//rand.Seed(time.Now().Unix())
	//
	//t.Log(rand.Int())
	//t.Log(rand.Int())
	//t.Log(rand.Int())
	//t.Log(rand.Int())
	//t.Log(rand.Int())
	//t.Log(rand.Int())
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		t.Log(rand.Int())
	}

	t.Log("------------")
	//t.Log(rand.Intn(100))
	//t.Log(rand.Int31())
	//t.Log(rand.Uint64())
	//t.Log(rand.Float32())
	//t.Log(rand.Float64())
	//
	//t.Log("------------")
	//
	//r := rand.New(rand.NewSource(time.Now().Unix()))
	//t.Log(r.Int())
	//t.Log(r.Intn(100))
	//t.Log(r.Int31())
	//t.Log(r.Float64())
	//t.Log(r.Float32())

	for i := 0; i < 100; i++ {
		t.Log(utils.GeneraterIntIdStr())

	}

}

func TestRandomStr(t *testing.T) {

	for i := 0; i < 10; i++ {
		t.Log(utils.GenID())
		t.Log(utils.GenIDStr())
	}

}
