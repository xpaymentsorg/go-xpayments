package xpsdb_test

import (
	"testing"

	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/xpsdb"
)

func TestStaticPartitioner_Partition(t *testing.T) {
	p := xpsdb.StaticPartitioner{Name: "TEST"}
	if v := p.Partition([]byte("foo")); v != "TEST" {
		t.Fatalf("unexpected partition: %v", v)
	} else if v := p.Partition([]byte("bar")); v != "TEST" {
		t.Fatalf("unexpected partition: %v", v)
	}
}

func TestBlockNumberPartitioner_Partition(t *testing.T) {
	p := xpsdb.NewBlockNumberPartitioner(100)
	if v := p.Partition(numHashKey('t', 1234, common.Hash{})); v != `00000000000004b0` {
		t.Fatalf("unexpected partition: %v", v)
	}
}
