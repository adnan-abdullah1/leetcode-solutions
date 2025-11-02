### 🧩 Problem: Find Duplicates with Limited Memory

You have an array containing integers in the range **1 to n**, where **n = 32,000**.
The array **may contain duplicate entries**, and **you don’t know n** in advance.

However, you have access to **only 4 KB of memory**.
Your task is to **print all duplicate elements** in the array.

---

### 💡 Key Constraints

* Each integer ranges from **1 to 32,000**.
* You have **only 4 KB (4096 bytes)** of memory available.
* Each integer typically takes **4 bytes**, so you cannot store the entire array in memory.
* You may assume the input can be read sequentially (e.g., from a stream or file).

---

### 🎯 Intuition

We cannot store all numbers directly, but we can **track which numbers have been seen** using **bits** instead of full integers.

* 1 byte = 8 bits
* 4 KB = 4096 bytes = **4096 × 8 = 32,768 bits**
* That’s **enough bits to represent all numbers from 1 to 32,000**

So we create a **bit vector** — an array of bytes, where each bit represents whether a number has been seen.

---

### 🧠 Bitwise Idea

For each number `v` in the array:

1. Compute which **byte** it belongs to:
   `byteIndex = (v - 1) / 8`
2. Compute which **bit** inside that byte:
   `bitIndex = (v - 1) % 8`
3. To **check if it’s seen**:

   ```go
   if seenBits[byteIndex] & (1 << bitIndex) != 0 {
       fmt.Println("duplicate:", v)
   }
   ```
4. Otherwise, **mark it as seen**:

   ```go
   seenBits[byteIndex] |= 1 << bitIndex
   ```

---

### 🧮 Example

For `v = 6`:

* `byteIndex = 0`
* `bitIndex = 5`
* To set bit 5 → `seenBits[0] |= 1 << 5` (marks 6 as seen)
* If 6 appears again → `seenBits[0] & (1 << 5)` will be nonzero, so we detect a duplicate.

---

### ✅ Summary

| Concept         | Operation                 | Meaning                               |                     |
| --------------- | ------------------------- | ------------------------------------- | ------------------- |
| `1 << bitIndex` | Left shift                | Set a single bit mask                 |                     |
| `               | =`                        | Bitwise OR                            | Mark number as seen |
| `&`             | Bitwise AND               | Check if bit already set              |                     |
| `/` and `%`     | Integer division, modulus | Locate byte and bit within the vector |                     |

---