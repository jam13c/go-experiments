package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := strings.Split(DemoInput, "\n")

	data := make([]int, len(input))
	for i, v := range input {
		val, _ := strconv.Atoi(v)
		data[i] = val
	}
	sort.Ints(data)
	fmt.Println(data)

	data = append([]int{0}, append(data, data[len(data)-1]+3)...)
	fmt.Println(data)

	memo := map[int]int{0: 1}
	for _, v := range data[1:] {
		fmt.Printf("Before %v: %v\n", v, memo)
		memo[v] = memo[v-1] + memo[v-2] + memo[v-3]
		fmt.Printf("After %v: %v\n\n", v, memo)
	}
	fmt.Println(memo)
	fmt.Println(memo[data[len(data)-1]])
	return

	result := findWays(data, []int{0})

	fmt.Println(result)
}

const DemoInput = `16
10
15
5
1
11
7
19
6
12
4`

const DemoInput2 = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

const RawInput = `99
151
61
134
112
70
75
41
119
137
158
50
167
60
116
117
62
82
31
3
72
88
165
34
8
14
27
108
166
71
51
42
135
122
140
109
1
101
2
77
85
76
143
100
127
7
107
13
148
118
56
159
133
21
154
152
130
78
54
104
160
153
95
49
19
69
142
63
11
12
29
98
84
28
17
146
161
115
4
94
24
126
136
91
57
30
155
79
66
141
48
125
162
37
40
147
18
20
45
55
83`
