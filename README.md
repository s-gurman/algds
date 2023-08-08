# Algorithms & Data Structures in Golang

Basic implementation of standard algorithms and data structures in Golang.
This repository is intended for educational purposes only.

## Algorithms

### Search

| Name          | Data structure | Average        | Worse         |
|:--------------|:--------------:|:--------------:|:-------------:|
| Linear search | Array          | $$O(n)$$       | $$O(n)$$      |
| Binary search | Sorted array   | $$O(\log n)$$  | $$O(\log n)$$ |

### Sort

| Name           | Data structure | Average             | Worse               |
|:---------------|:--------------:|:-------------------:|:-------------------:|
| Selection sort | Array          | $$O(n^2)$$          | $$O(n^2)$$          |
| Quick sort     | Array          | $$O(n\cdot\log n)$$ | $$O(n^2)$$          |
| Merge sort     | Array          | $$O(n\cdot\log n)$$ | $$O(n\cdot\log n)$$ |

## Data structures

| Name        | Search        | Insert        | Delete        | Space    |
|:------------|:-------------:|:-------------:|:-------------:|:--------:|
| Linked List | $$O(n)$$      | $$O(1)$$      | $$O(1)$$      | $$O(n)$$ |
| Stack       | -             | $$O(1)$$      | $$O(1)$$      | $$O(n)$$ |
| Queue       | -             | $$O(1)$$      | $$O(1)$$      | $$O(n)$$ |
| Hash Table  | $$O(1)$$      | $$O(1)$$      | $$O(1)$$      | $$O(n)$$ |
| Binary Tree | $$O(\log n)$$ | $$O(\log n)$$ | $$O(\log n)$$ | $$O(n)$$ |
| Binary Heap | $$O(n)$$      | $$O(\log n)$$ | $$O(\log n)$$ | $$O(n)$$ |