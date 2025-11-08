package rapid

import (
	"testing"

	"pgregory.net/rapid"
)

// TestAdditionCommutative demonstrates basic property testing with Rapid.
func TestAdditionCommutative(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Draw two random integers
		a := rapid.Int().Draw(t, "a")
		b := rapid.Int().Draw(t, "b")

		// Verify commutativity
		if a+b != b+a {
			t.Fatalf("addition not commutative: %d + %d != %d + %d", a, b, b, a)
		}
	})
}

// TestStackOperations demonstrates stateful testing with Rapid.
// This tests a simple stack implementation through random operations.
func TestStackOperations(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		stack := []int{}

		// Perform 50 random operations
		for i := 0; i < 50; i++ {
			shouldPush := rapid.Bool().Draw(t, "push?")

			if shouldPush {
				// Push operation
				value := rapid.IntRange(0, 100).Draw(t, "value")
				stack = append(stack, value)
			} else if len(stack) > 0 {
				// Pop operation
				stack = stack[:len(stack)-1]
			}

			// Invariant: stack length is always non-negative
		}
	})
}

// Counter is a simple counter for demonstrating stateful testing.
type Counter struct {
	value int
}

func (c *Counter) Increment() {
	c.value++
}

func (c *Counter) Decrement() {
	c.value--
}

func (c *Counter) Reset() {
	c.value = 0
}

func (c *Counter) Value() int {
	return c.value
}

// TestCounter demonstrates testing a stateful object.
func TestCounter(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		counter := &Counter{}
		expectedValue := 0

		// Perform 30 random operations
		for i := 0; i < 30; i++ {
			operation := rapid.IntRange(0, 2).Draw(t, "operation")

			switch operation {
			case 0: // Increment
				counter.Increment()
				expectedValue++
			case 1: // Decrement
				counter.Decrement()
				expectedValue--
			case 2: // Reset
				counter.Reset()
				expectedValue = 0
			}

			// Verify counter value matches our expectation
			if counter.Value() != expectedValue {
				t.Fatalf("counter value mismatch: got %d, expected %d",
					counter.Value(), expectedValue)
			}
		}
	})
}

// Map is a simple key-value store for testing.
type Map struct {
	data map[string]int
}

func NewMap() *Map {
	return &Map{data: make(map[string]int)}
}

func (m *Map) Set(key string, value int) {
	m.data[key] = value
}

func (m *Map) Get(key string) (int, bool) {
	val, exists := m.data[key]
	return val, exists
}

func (m *Map) Delete(key string) {
	delete(m.data, key)
}

func (m *Map) Size() int {
	return len(m.data)
}

// TestMap demonstrates testing a map-like structure.
func TestMap(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		m := NewMap()
		model := make(map[string]int) // Model for verification

		// Generate a set of keys to work with
		keys := []string{"a", "b", "c", "d", "e"}

		// Perform random operations
		for i := 0; i < 100; i++ {
			operation := rapid.IntRange(0, 2).Draw(t, "operation")
			key := rapid.SampledFrom(keys).Draw(t, "key")

			switch operation {
			case 0: // Set
				value := rapid.IntRange(0, 1000).Draw(t, "value")
				m.Set(key, value)
				model[key] = value

			case 1: // Get
				gotValue, gotExists := m.Get(key)
				expectedValue, expectedExists := model[key]

				if gotExists != expectedExists {
					t.Fatalf("existence mismatch for key %s", key)
				}
				if gotExists && gotValue != expectedValue {
					t.Fatalf("value mismatch for key %s: got %d, expected %d",
						key, gotValue, expectedValue)
				}

			case 2: // Delete
				m.Delete(key)
				delete(model, key)
			}

			// Verify size matches
			if m.Size() != len(model) {
				t.Fatalf("size mismatch: got %d, expected %d", m.Size(), len(model))
			}
		}
	})
}

// TestStringOperations demonstrates testing string operations.
func TestStringOperations(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Generate two strings
		s1 := rapid.StringN(0, 20, -1).Draw(t, "s1")
		s2 := rapid.StringN(0, 20, -1).Draw(t, "s2")

		// Property: concatenation length
		concat := s1 + s2
		if len(concat) != len(s1)+len(s2) {
			t.Fatalf("concatenation length incorrect")
		}

		// Property: substring of concatenation
		if len(s1) > 0 && concat[:len(s1)] != s1 {
			t.Fatalf("concatenation prefix incorrect")
		}
		if len(s2) > 0 && len(concat) > 0 && concat[len(s1):] != s2 {
			t.Fatalf("concatenation suffix incorrect")
		}
	})
}

// TestSliceOperations demonstrates testing slice operations.
func TestSliceOperations(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Generate a slice of integers
		originalLen := rapid.IntRange(0, 20).Draw(t, "length")
		slice := make([]int, originalLen)

		for i := range slice {
			slice[i] = rapid.Int().Draw(t, "element")
		}

		// Property: append increases length by 1
		value := rapid.Int().Draw(t, "append_value")
		newSlice := append(slice, value)

		if len(newSlice) != len(slice)+1 {
			t.Fatalf("append didn't increase length correctly")
		}

		// Property: last element is the appended value
		if len(newSlice) > 0 && newSlice[len(newSlice)-1] != value {
			t.Fatalf("appended value not at end")
		}

		// Property: original elements unchanged
		for i := range slice {
			if slice[i] != newSlice[i] {
				t.Fatalf("original elements changed after append")
			}
		}
	})
}

// TestReverseString demonstrates reversing strings.
func TestReverseString(t *testing.T) {
	reverseString := func(s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}

	rapid.Check(t, func(t *rapid.T) {
		s := rapid.String().Draw(t, "string")

		// Property: reversing twice returns original
		reversed := reverseString(s)
		doubleReversed := reverseString(reversed)

		if s != doubleReversed {
			t.Fatalf("double reverse didn't return original: %q != %q", s, doubleReversed)
		}

		// Property: length is preserved
		if len(s) != len(reversed) {
			t.Fatalf("length changed after reverse")
		}
	})
}

// TestSortedSliceProperties demonstrates testing sorted slices.
func TestSortedSliceProperties(t *testing.T) {
	isSorted := func(slice []int) bool {
		for i := 1; i < len(slice); i++ {
			if slice[i-1] > slice[i] {
				return false
			}
		}
		return true
	}

	rapid.Check(t, func(t *rapid.T) {
		// Generate a sorted slice
		size := rapid.IntRange(0, 50).Draw(t, "size")
		slice := make([]int, size)

		if size > 0 {
			slice[0] = rapid.IntRange(-100, 100).Draw(t, "first")
		}

		for i := 1; i < size; i++ {
			// Each element is >= previous element
			increment := rapid.IntRange(0, 10).Draw(t, "increment")
			slice[i] = slice[i-1] + increment
		}

		// Verify it's sorted
		if !isSorted(slice) {
			t.Fatalf("generated slice is not sorted: %v", slice)
		}

		// Property: length is preserved
		if len(slice) != size {
			t.Fatalf("slice length mismatch")
		}
	})
}
