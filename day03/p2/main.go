package main

/*
37  36  35  34  33  32  31
38  17  16  15  14  13  30
39  18   5   4   3  12  29
40  19   6   1   2  11  28
41  20   7   8   9  10  27
42  21  22  23  24  25  26
43  44  45  46  47  48  49

33333X
X222X3
2X1X23
210123
2X1X23
X222X3 <-- always starts to right of last level corner
33333X

ending corners per level:
0,0
1,1 (l, l)
2,2

number of squares at each level:
l = level,  0, 1, 2, 3, 4...

level 0: 1
level 1: 8    4 + 1*4    (2,... 9)
level 2: 16   4 + 3*4    (10,.. 25)
level 3: 24   4 + 5*4    (26,.. 49)
level 4: 32   4 + 7*4    (50,...81)
level 5: 40   4 + 9*4    (82,...

we always add 4 new corners...
( 2 x (l-1) + 1) * 4  + 4 = 1, 8, 16, 24, 32, 40

length of side at level:
2(l-1)+1  + 2 = 0, 3, 5, 7, 9

C = count at level
C/4 number per side

LEVEL 1  (8: 2-9,  side length 3)
(x-prev max -1) % (C/4)
2, 3,   4, 5,   6, 7,   8, 9    (numbers in level)
0, 1,   0, 1,   0, 1,   0, 1
^ sub l-1

LEVEL 2  (16: 10-25, side length 5)
10, 11, 12, 13,   14, 15, 16, 17,   18, 19, 20, 21,   22, 23, 24, 25
0   1   2   3     0   1   2   3     0   1   2   3     0   1   2   3
    ^ sub l-1

LEVEL 3  (24: 26-49, side length 7)
(26-25-1) % 6
26, 27, 28, 29, 30, 31    32, 33, 34, 35, 36, 37   38, 39, 40, 41, 42, 43    44, 45, 46, 47, 48, 49
0   1   2   3   4   5     0   1   2   3   4   5    0   1   2   3   4   5     0   1   2   3   4   5
        ^  sub l-2

how to tell which level a number is in?
if 1, it's level 1, and distance == 0
count up levels getting max number, if <= max, it's in that level, then break
numbers = prev max + 1 + number squares in level

*/

func main() {

}
