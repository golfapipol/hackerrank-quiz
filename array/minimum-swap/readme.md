# Problem
You are given an unordered array consisting of consecutive integers  [1, 2, 3, ..., n] without any duplicates. You are allowed to swap any two elements. You need to find the minimum number of swaps required to sort the array in ascending order.

For example, given the array  we perform the following steps:

i   arr                         swap (indices)
0   [7, 1, 3, 2, 4, 5, 6]   swap (0,3)
1   [2, 1, 3, 7, 4, 5, 6]   swap (0,1)
2   [1, 2, 3, 7, 4, 5, 6]   swap (3,4)
3   [1, 2, 3, 4, 7, 5, 6]   swap (4,5)
4   [1, 2, 3, 4, 5, 7, 6]   swap (5,6)
5   [1, 2, 3, 4, 5, 6, 7]
It took  swaps to sort the array.

# Solve 
ใช้วิธีเช็คว่า index ปัจจุบันใช่ไหม ถ้าไม่ใช่ให้หาว่า index ปัจจุบันอยู่ตรงไหนแล้วสลับที่กัน

# Constraints

Sample Input 0
4 3 1 2
Sample Output 0
3

Sample Input 1
2 3 4 1 5
Sample Output 1
3
