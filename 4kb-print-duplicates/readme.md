Perfect ✅ Here’s a complete **`README.md`** that captures the **question, reasoning, and final Go solution** — plus all the intuition you built step by step.

---

````markdown
# 🧮 Find Duplicates with Only 4KB of Memory

### 💡 Problem Statement
You are given an array that contains integers ranging from **1 to 32,000**.  
The array **may have duplicate entries**, but you **do not know what `n` is** (i.e., the array can be of any length).  

Your goal is to **print all duplicate elements**, using **no more than 4 KB of memory**.

---

### 🧠 Intuition Behind the Solution

- 4 KB = **4096 bytes** = **4096 × 8 = 32,768 bits**
- There are **32,000 possible numbers**, each between 1 and 32,000.  
  So, we can **assign 1 bit per number** — perfect fit!

Each bit represents whether we have **seen** a particular number before.

---

### ⚙️ Core Idea

We maintain a **bit vector** where:

- **Each bit** corresponds to a number (1 → 32,000).
- **If the bit is 1**, it means we’ve seen that number.
- **If the bit is 0**, it means we haven’t.

---

### 🧩 Key Operations

#### Mark a number as seen
```go
mem[byteIndex] |= 1 << bitIndex
````

* `1 << bitIndex` → creates a mask like `00010000`, setting just one bit.
* `|=` → flips that bit **on** in the byte.

#### Check if a number has already been seen

```go
if mem[byteIndex] & (1 << bitIndex) != 0 {
    fmt.Println("duplicate:", num)
}
```

* `&` → checks whether that bit was already set.
* If non-zero, it’s a **duplicate**.

---

### 🧮 Calculating Byte and Bit Positions

For a given number `num`:

```go
byteIndex := (num - 1) / 8  // which byte
bitIndex := (num - 1) % 8   // which bit inside the byte
```

This allows us to efficiently locate the bit that corresponds to `num`.

---

### 🧑‍💻 Go Implementation

```go
package main

import "fmt"

func main() {
	arr := []int{3, 7, 3, 15, 7, 10, 15, 1, 32000, 32000}
	mem := make([]byte, 4000) // 4000 bytes = 32,000 bits

	for _, num := range arr {
		byteIndex := (num - 1) / 8
		bitIndex := (num - 1) % 8

		if mem[byteIndex]&(1<<bitIndex) != 0 {
			fmt.Println("Duplicate:", num)
		} else {
			mem[byteIndex] |= 1 << bitIndex
		}
	}
}
```

---

### 🧩 Example Walkthrough

Let’s say `num = 6`

| Step | Expression                | Binary      | Meaning         |                       |
| ---- | ------------------------- | ----------- | --------------- | --------------------- |
| 1    | `byteIndex = (6-1)/8 = 0` | → Byte 0    | First byte      |                       |
| 2    | `bitIndex = (6-1)%8 = 5`  | → Bit 5     | Sixth bit       |                       |
| 3    | `1 << 5`                  | `00100000`  | Set the 6th bit |                       |
| 4    | `mem[0]                   | = 00100000` | `00100000`      | Mark number 6 as seen |

Next time when `6` comes again:
`mem[0] & 00100000` ≠ 0 → Duplicate found ✅

---

### 🧩 Why This Problem Is in “Sorting and Searching”

Even though it’s implemented via **bit manipulation**,
it’s conceptually a **search problem** —
we’re efficiently *searching* for duplicates while minimizing memory usage.

---

### ✅ Takeaways

* 1 bit per number = memory-efficient tracking.
* Use `|` to set bits, `&` to check bits.
* `<<` moves a single `1` into the bit position you want.
* Brilliant example of **space-time tradeoff** in algorithm design.

---

### 🗂️ File Info

| File                                 | Purpose                           |
| ------------------------------------ | --------------------------------- |
| `find-duplicates-with-bit-vector.go` | Go code implementation            |
| `README.md`                          | Problem explanation and intuition |

---

**Author:** *Adnan Abdullah*
**Concept learned:** Bit manipulation for memory optimization (from “Cracking the Coding Interview”)
**Tags:** `bitwise`, `searching`, `duplicates`, `memory-efficient`, `golang`

```

---

Would you like me to include **ASCII illustrations** of how bits flip (like `00000000 → 00100000`) in the README too?  
It makes the visualization even easier to recall later.
```
