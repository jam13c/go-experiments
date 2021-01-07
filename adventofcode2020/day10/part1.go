package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := strings.Split(RawInput, "\n")

	data := make([]int, len(input))
	for i, v := range input {
		val, _ := strconv.Atoi(v)
		data[i] = val
	}
	sort.Ints(data)
	data = append(data, data[len(data)-1]+3)
	result := map[int]int{
		1: 0,
		2: 0,
		3: 0,
	}

	prev := 0
	for i := 0; i < len(data); i++ {
		diff := data[i] - prev
		result[diff]++
		prev = data[i]
	}

	fmt.Println(result)
	fmt.Println(result[1] * result[3])
}

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

const DemoInput = `28
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
