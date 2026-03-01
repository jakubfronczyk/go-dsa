# Go DSA

A personal study repository for data structures and algorithms implemented in Go. Each solution is written by hand to build genuine understanding of how things work under the hood.

## Testing

Each problem has its own `_test.go` file using Go's built-in testing package. No external dependencies.

```bash
# Run all tests
go test ./...

# Run tests for a specific folder
go test ./arrays/...

# Run a specific test
go test ./arrays/... -run TestReverse

# Verbose output
go test ./... -v
```

---

## Array-Based Problems

1. Find the maximum and minimum elements in an array.
2. Reverse an array.
3. Check if an array is sorted.
4. Check if two arrays are equal.
5. Count occurrences of an element in an array.
6. Remove duplicates from a sorted array.
7. Find duplicate elements in an array.
8. Find the second largest element in an array.
9. Find the missing number in an array of size n containing numbers from 1 to n+1.
10. Find all the pairs in an array that sum up to a given number.
11. Find the intersection of two arrays.
12. Find the union of two arrays.
13. Merge two sorted arrays.
14. Rotate an array by k positions.
15. Move all zeros to the end of an array.
16. Sort an array of 0s, 1s, and 2s.
17. Check if a subarray with a given sum exists.
18. Find the equilibrium index of an array.
19. Rearrange an array such that the positive and negative numbers alternate.
20. Find the contiguous subarray with the maximum sum (Kadane's Algorithm).
21. Find the majority element in an array (element that appears more than n/2 times).

---

## String-Based Problems

1. Reverse a string.
2. Check if a string is a palindrome.
3. Check if a string contains only digits.
4. Count the number of vowels and consonants in a string.
5. Count the frequency of characters in a string.
6. Convert a string to an integer.
7. Convert an integer to a string.
8. Check if two strings are anagrams of each other.
9. Remove all duplicate characters from a string.
10. Find the first non-repeating character in a string.
11. Replace all spaces in a string with %20.
12. Reverse the order of words in a string.
13. Find the longest common prefix among a set of strings.
14. Compress a string (e.g., aaabbc becomes a3b2c1).
15. Find all substrings of a string.
16. Check if a string is a rotation of another string.
17. Check if a string is a valid shuffle of two other strings.
18. Check if brackets in a string are balanced.
19. Find the longest palindrome in a string.
20. Find all permutations of a string.
21. Implement a basic pattern-matching algorithm.
22. Find the longest substring without repeating characters.
23. Find the most frequent word in a string.
