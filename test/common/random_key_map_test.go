package testcommon_test

import (
	"math/rand"
	"testing"

	testcommon "github.com/allora-network/allora-chain/test/common"
)

func TestRandomKeyMap_Delete(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	rkm := testcommon.NewRandomKeyMap[int, string](r)

	// Insert some elements into the map
	keys := []int{1, 2, 3, 4, 5}
	values := []string{"one", "two", "three", "four", "five"}
	for i, key := range keys {
		rkm.Insert(key, values[i])
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
	rkm := testcommon.NewRandomKeyMap[int, string](r)
	// Insert some elements into the map
	keys := []int{1, 2, 3, 4, 5}
	values := []string{"one", "two", "three", "four", "five"}
	for i, key := range keys {
		rkm.Insert(key, values[i])
	}
	// Get an existing element
	keyToGet := 3
	valueToGet := "three"
	value, exists := rkm.Get(keyToGet)
	if !exists {
		t.Errorf("Expected key %d to exist in the map, but it doesn't", keyToGet)
	}
	if value != valueToGet {
		t.Errorf("Expected value %s for key %d, but got %s", valueToGet, keyToGet, value)
	}
	// Get a non-existing element
	nonExistingKey := 6
	_, exists = rkm.Get(nonExistingKey)
	if exists {
		t.Errorf("Expected key %d to not exist in the map, but it does", nonExistingKey)
	}
}

func TestRandomKeyMap_RandomKey(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	rkm := testcommon.NewRandomKeyMap[int, string](r)
	// Insert some elements into the map
	keys := []int{1, 2, 3, 4, 5}
	values := []string{"one", "two", "three", "four", "five"}
	for i, key := range keys {
		rkm.Insert(key, values[i])
	}
	// Get a random key from the map
	randomKey := rkm.RandomKey()
	// Verify that the random key is one of the keys in the map
	found := false
	for _, key := range keys {
		if key == randomKey {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected random key to be one of %v, but got %v", keys, randomKey)
	}
}

func TestRandomKeyMap_Len(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	rkm := testcommon.NewRandomKeyMap[int, string](r)
	// Insert some elements into the map
	keys := []int{1, 2, 3, 4, 5}
	values := []string{"one", "two", "three", "four", "five"}
	for i, key := range keys {
		rkm.Insert(key, values[i])
	}
	// Verify the initial length of the map
	expectedLen := len(keys)
	actualLen := rkm.Len()
	if actualLen != expectedLen {
		t.Errorf("Expected map length to be %d, but got %d", expectedLen, actualLen)
	}
	// Delete an element and verify the length decreases by 1
	keyToDelete := 3
	rkm.Delete(keyToDelete)
	expectedLen--
	actualLen = rkm.Len()
	if actualLen != expectedLen {
		t.Errorf("Expected map length to be %d, but got %d", expectedLen, actualLen)
	}
	// Insert a new element and verify the length increases by 1
	newKey := 6
	newValue := "six"
	rkm.Insert(newKey, newValue)
	expectedLen++
	actualLen = rkm.Len()
	if actualLen != expectedLen {
		t.Errorf("Expected map length to be %d, but got %d", expectedLen, actualLen)
	}
}
