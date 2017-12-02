package parser

import "testing"

// itemsEqualIgnoringOrder return true if the two slices contain identical elements,
// even if they may be in different orders. Duplicates matter, ["a", "a"] != ["a"]
func itemsEqualIgnoringOrder(first []string, second []string) bool {
  // if lengths are equal, one slice must have elements not in the other
  if len(first) != len(second) {
    return false
  }

  items := map[string]int{}
  // Count occurences in first
  for _, item := range first {
    items[item]++
  }
  // Check everything in second
  for _, item := range second {
    // if this wasn't in first (or we've already found a value in second for every value in first),
    // the slices aren't equal
    if items[item] == 0 {
      return false
    }
    items[item]--
  }
  // If the count isn't zero, item was incremented more than decremented (found more in first)
  // or decremented more than incremented (found more in second)
  for _, count := range items {
    if count != 0 {
      return false
    }
  }

  return true
}

func TestMentions(t *testing.T) {
  input := "@chris are you around"
  expected := []string{"chris"}
  found := Mentions(input)
  if !itemsEqualIgnoringOrder(found, expected) {
    t.Errorf("Mentions(%q): expected %q got %q", input, expected, found)
  }
}
