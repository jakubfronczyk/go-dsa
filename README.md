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
4. Rotate an array by k positions.
5. Find the second largest element in an array.
6. Find all the pairs in an array that sum up to a given number.
7. Find duplicate elements in an array.
8. Remove duplicates from a sorted array.
9. Merge two sorted arrays.
10. Find the intersection of two arrays.
11. Find the union of two arrays.
12. Find the contiguous subarray with the maximum sum (Kadane's Algorithm).
13. Count occurrences of an element in an array.
14. Find the missing number in an array of size n containing numbers from 1 to n+1.
15. Check if a subarray with a given sum exists.
16. Sort an array of 0s, 1s, and 2s.
17. Rearrange an array such that the positive and negative numbers alternate.
18. Find the equilibrium index of an array.
19. Find the majority element in an array (element that appears more than n/2 times).
20. Check if two arrays are equal.

---

## String-Based Problems

21. Reverse a string.
22. Check if a string is a palindrome.
23. Find the first non-repeating character in a string.
24. Count the number of vowels and consonants in a string.
25. Remove all duplicate characters from a string.
26. Check if two strings are anagrams of each other.
27. Find all substrings of a string.
28. Find the longest common prefix among a set of strings.
29. Count the frequency of characters in a string.
30. Replace all spaces in a string with %20.
31. Check if a string is a valid shuffle of two other strings.
32. Compress a string (e.g., aaabbc becomes a3b2c1).
33. Find the longest palindrome in a string.
34. Implement a basic pattern-matching algorithm.
35. Convert a string to an integer.
36. Convert an integer to a string.
37. Check if a string contains only digits.
38. Reverse the order of words in a string.
39. Find the longest substring without repeating characters.
40. Find the most frequent word in a string.
