package main

import "fmt"

/*
MAPS NOTES:
1. Syntax: map[KeyType]ValueType - e.g., map[string]int
2. Declaration creates nil map - CANNOT add items to nil map (will panic)
3. Use make() to initialize: map1 := make(map[string]string)
4. Or initialize with values: map2 := map[string]int{"age": 30}
5. Access: value := map1["key"]
6. Check existence: value, exists := map1["key"]
7. Delete: delete(map1, "key")
8. Iterate: for key, value := range map1 { }
9. Maps are UNORDERED - iteration order is random
10. Cannot use numeric index like arrays - use range to iterate
11. len(map) gives number of key-value pairs

WHY NIL MAPS PANIC:
Maps are reference types - declaration without initialization creates nil pointer:
   var m map[string]int    // m is nil (no memory allocated)
   m["key"] = 42           // ❌ PANIC: assignment to entry in nil map

Why? Nil map has NO underlying data structure (hash table, buckets, metadata).
It's just an empty pointer pointing to nothing!

Solution - use make() to allocate memory:
   m := make(map[string]int)  // ✅ Allocates hash table, ready to use
   m["key"] = 42              // Works!

Or initialize with literal:
   m := map[string]int{"key": 42}  // ✅ Also allocates memory

NIL BEHAVIOR COMPARISON:
┌──────────┬─────────────────┬──────────┬───────────┬─────────────────┐
│ Type     │ Nil Declaration │ Can Read?│ Can Write?│ Solution        │
├──────────┼─────────────────┼──────────┼───────────┼─────────────────┤
│ Map      │ var m map[K]V   │ ✅ (zero)│ ❌ PANIC  │ make(map[K]V)   │
│ Channel  │ var ch chan T   │ ❌ Blocks│ ❌ PANIC  │ make(chan T)    │
│ Slice    │ var s []T       │ ❌ PANIC │ ✅ append │ make([]T, len)  │
└──────────┴─────────────────┴──────────┴───────────┴─────────────────┘

Reading from nil map is safe (returns zero value), but writing panics!
   var m map[string]int  // nil
   value := m["key"]     // ✅ Returns 0 (zero value), no panic
   m["key"] = 1          // ❌ PANIC!

CHECKING IF KEY EXISTS IN MAP:
Method 1 (Two-value assignment):
   value, exists := myMap[key]
   if exists {
       // Key exists, use value
   } else {
       // Key doesn't exist
   }

Method 2 (Check only):
   if _, exists := myMap[key]; exists {
       // Key exists
   }

Method 3 (For bool maps):
   if myMap[key] {
       // Key exists and value is true
   }

IMPORTANT:
- Accessing non-existent key returns ZERO VALUE (not error)
  Example: myMap["missing"] returns "" for string, 0 for int, false for bool
- Use two-value assignment to distinguish between "key doesn't exist" vs "key exists with zero value"

EXAMPLE:
   ages := map[string]int{"Alice": 30, "Bob": 0}
   
   age1 := ages["Alice"]  // 30
   age2 := ages["Bob"]    // 0 (exists, but value is 0)
   age3 := ages["Charlie"] // 0 (doesn't exist, returns zero value)
   
   // To distinguish Bob (exists with 0) from Charlie (doesn't exist):
   if age, exists := ages["Bob"]; exists {
       fmt.Println("Bob's age:", age)  // Prints: Bob's age: 0
   }
   
   if age, exists := ages["Charlie"]; exists {
       fmt.Println("Charlie's age:", age)
   } else {
       fmt.Println("Charlie not found")  // This executes
   }
*/

func main() {
	/*
		var map1 map[string]string
		above creates nil map - which creates panic while trying to initialise - always use make or initialise with values while declaration

			map1["name"] = "Rajat"
			map1["age"] = "31"
			map1["gender"] = "Male"
			fmt.Println(map1)
	*/
	map2 := make(map[string]string)
	map2["name"] = "Rajat"
	map2["age"] = "31"
	map2["gender"] = "Male"
	fmt.Println(map2)
	fmt.Println("Length of map2 is: ", len(map2))
	fmt.Println("Value of map2 for key name is: ", map2["name"])
	/* traverse a map
	for i:= 0; i<len(map2); i++ {
		fmt.Println(map2[i])
	}
	*/
	for key, value := range map2 {
		fmt.Printf("Map2 key is %s : and value is %s:\n", key, value)
	}
	/*
		a := 10
		for i := range a {
			fmt.Println(i)
		}
	*/

}
