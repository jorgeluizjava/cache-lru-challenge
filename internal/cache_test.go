package internal

import (
	"testing"
)

func TestInitializeCache(t *testing.T) {
	c := New(10)

	if c == nil {
		t.Error("cache is nil")
	}
}

func TestGet_ReturnValueForAKey(t *testing.T) {
	c := New(3)
	if c == nil {
		t.Error("cache is nil")
	}
	c.Set("key1", "test")

	key1 := c.Get("key1")
	if key1 == -1 {
		t.Error("key1 not found")
	}
	if key1 == nil {
		t.Error("key1 is nil")
	}

	if key1 != "test" {
		t.Error("value of key1 is not test")
	}
}

func TestSet_WithEnoughCapacity(t *testing.T) {
	c := New(3)
	if c == nil {
		t.Error("cache is nil")
	}

	key1 := "key"
	value1 := "value"
	c.Set(key1, value1)

	if c.Get(key1) == -1 {
		t.Error("key was not added")
	}
}

func TestSet_ReplaceOldestItemWhenThereIsNoCapacity(t *testing.T) {
	c := New(3)
	if c == nil {
		t.Error("cache is nil")
	}
	key1 := "key1"
	value1 := "value1"
	c.Set(key1, value1)

	key2 := "key2"
	value2 := "value2"
	c.Set(key2, value2)

	key3 := "key3"
	value3 := "value3"
	c.Set(key3, value3)

	key4 := "key4"
	value4 := "value4"
	c.Set(key4, value4)

	result := c.Get(key1)
	if result != -1 {
		t.Errorf("the value of key1 should be -1, current value is %d", result)
	}
}

func TestSet_ReplaceOldestItemWhenThereIsNoCapacityAfterCallGet(t *testing.T) {
	c := New(3)
	if c == nil {
		t.Error("cache is nil")
	}
	key1 := "key1"
	value1 := "value1"
	c.Set(key1, value1)

	key2 := "key2"
	value2 := "value2"
	c.Set(key2, value2)

	key3 := "key3"
	value3 := "value3"
	c.Set(key3, value3)

	c.Get(key1)

	key4 := "key4"
	value4 := "value4"
	c.Set(key4, value4)

	result2 := c.Get(key2)
	if result2 != -1 {
		t.Errorf("the value of key2 should be -1, current value is %d", result2)
	}
}
