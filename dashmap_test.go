package dashmap_test

import (
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/projekt-go/dashmap"
)

func TestDashMap(t *testing.T) {
	mp := dashmap.New[string, int]()

	mp.Put("fish", 3)
	if v, ok := mp.Get("fish"); !ok || v != 3 {
		t.Errorf("invalid read or write. expected %d got %d.", 3, v)
	}

	mp.Put("orange", 10)
	if v, ok := mp.Get("orange"); !ok || v != 10 {
		t.Errorf("invalid read or write. expected %d got %d.", 10, v)
	}

	expectedEntries := []dashmap.Entry[string, int]{
		{Key: "fish", Value: 3},
		{Key: "orange", Value: 10},
	}

	entries := mp.Entries()

	sort.Slice(entries, func(i, j int) bool {
		cmp := strings.Compare(entries[i].Key, entries[j].Key)
		return (cmp < 0) || (cmp == 0 && (entries[i].Value < entries[j].Value))
	})

	if !reflect.DeepEqual(entries, expectedEntries) {
		t.Errorf("invalid entries. expected %v got %v.", expectedEntries, entries)
	}
}

func TestRace(t *testing.T) {
	mp := dashmap.New[int, int]()

	for i := 0; i < 256; i++ {
		go mp.Put(i, i + 1)
	}

	for i := 0; i < 256; i++ {
		go func(t *testing.T, idx int) {
			if val, ok := mp.Get(idx); !ok || val != idx + 1 {
				t.Errorf("invalid concurrent reads. expected %d got %d.", idx + 1, val)
			}
		}(t, i)
	}
}
