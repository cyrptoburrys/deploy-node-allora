package testcommon_test

import (
	"math/rand"
	"testing"

	testcommon "github.com/allora-network/allora-chain/test/common"
)

func TestRandomKeyMap_Delete(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	rkm := testcommon.NewRandomKeyMap[int](r)

	// Insert some elements into the map
	keys := []int{1, 2, 3, 4, 5}
	for _, key := range keys {
		rkm.Insert(key)
	}

	// Delete an existing element
	keyToDelete := 3
	rkm.Delete(keyToDelete)

	// Verify that the deleted element is no longer in the map
	_, exists := rkm.Get(keyToDelete)
	if exists {
		t.Errorf("Expected key %d to be deleted, but it still exists in the map", keyToDelete)
	}

	// Verify that the length of the map has decreased by 1
	expectedLen := len(keys) - 1
	actualLen := rkm.Len()
	if actualLen != expectedLen {
		t.Errorf("Expected map length to be %d, but got %d", expectedLen, actualLen)
	}

	// Delete a non-existing element
	nonExistingKey := 6
	rkm.Delete(nonExistingKey)

	// Verify that the map remains unchanged
	actualLen = rkm.Len()
	if actualLen != expectedLen {
		t.Errorf("Expected map length to be %d, but got %d", expectedLen, actualLen)
	}
}

func TestRandomKeyMap_Get(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	rkm := testcommon.NewRandomKeyMap[int](r)
	// Insert some elements into the map
	keys := []int{1, 2, 3, 4, 5}
	for _, key := range keys {
		rkm.Insert(key)
	}
	// Get an existing element
	keyToGet := 3
	value, exists := rkm.Get(keyToGet)
	if !exists {
		t.Errorf("Expected key %d to exist in the map, but it doesn't", keyToGet)
	}
	if *value != keyToGet {
		t.Errorf("Expected value %d for key %d, but got %d", keyToGet, keyToGet, *value)
	}
	// Get a non-existing element
	nonExistingKey := 6
	_, exists = rkm.Get(nonExistingKey)
	if exists {
		t.Errorf("Expected key %d to not exist in the map, but it does", nonExistingKey)
	}
}
