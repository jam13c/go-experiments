package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	part2(strings.Split(RawInput, "\n"))
}

func part1(input string) {
	sections := strings.Split(input, "\n\n")
	rules := parseRules(strings.Split(sections[0], "\n"))
	pattern := rules.createPatterns()
	fmt.Println(pattern)
	re := regexp.MustCompile(pattern)
	count := 0
	for _, row := range strings.Split(sections[1], "\n") {
		if re.MatchString(row) {
			count++
		}
	}
	fmt.Println(count)
}
func part2(input []string) {
	type rule struct {
		val    string
		groups [][]string
	}

	rules := map[string]rule{}

	i := 0

	for ; ; i++ {
		line := input[i]

		if line == "" {
			i++
			break
		}

		tokens := strings.Split(line, ": ")
		k := tokens[0]

		if k == "8" {
			tokens[1] = "42 | 42 8"
		}

		if k == "11" {
			tokens[1] = "42 31 | 42 11 31"
		}

		if strings.HasPrefix(tokens[1], "\"") {
			rules[k] = rule{
				val: strings.Trim(tokens[1], "\""),
			}

			continue
		}

		r := rule{
			groups: [][]string{},
		}

		for _, str := range strings.Split(tokens[1], " | ") {
			r.groups = append(r.groups, strings.Split(str, " "))
		}

		rules[k] = r
	}

	seen := map[string]bool{}

	var iter func(key string, then []string, search string, index int)
	iter = func(key string, then []string, search string, index int) {
		rule := rules[key]

		if rule.val != "" {
			if index >= len(search) {
				return
			}

			if rule.val == search[index:index+1] {
				if len(then) > 0 {
					iter(then[0], then[1:], search, index+1)
				} else {
					if index == len(search)-1 {
						seen[search] = true
					}
				}
			}

			return
		}

		g1 := rule.groups[0]
		iter(g1[0], append(g1[1:], then...), search, index)

		if len(rule.groups) > 1 {
			g2 := rule.groups[1]
			iter(g2[0], append(g2[1:], then...), search, index)
		}
	}

	count := 0

	for ; i < len(input); i++ {
		iter(rules["0"].groups[0][0], rules["0"].groups[0][1:], input[i], 0)
		_, ok := seen[input[i]]
		if ok {
			count++
		}
	}

	fmt.Print(count)
}
func parseRules(input []string) rules {

	reChar := regexp.MustCompile(`^(\d+):\s"(\w)"$`)
	reOpt := regexp.MustCompile(`^(\d+):\s([\d\s*]+)$`)
	reChoice := regexp.MustCompile(`^(\d+):\s([\d+\s*]+)\s\|\s([\d+\s*]+)$`)

	r := make(rules)

	parseRuleList := func(s string) []int {
		rules := strings.Split(s, " ")
		idxs := make([]int, len(rules))
		for i, rule := range rules {
			r, _ := strconv.Atoi(rule)
			idxs[i] = r
		}
		return idxs
	}

	for _, line := range input {
		if m := reChar.FindStringSubmatch(line); m != nil {
			idx, _ := strconv.Atoi(m[1])
			c := rune(m[2][0])
			r[idx] = charRule{c, &r}
		} else if m := reOpt.FindStringSubmatch(line); m != nil {
			idx, _ := strconv.Atoi(m[1])
			rules := parseRuleList(m[2])
			r[idx] = optionRule{rules, &r}
		} else if m := reChoice.FindStringSubmatch(line); m != nil {
			idx, _ := strconv.Atoi(m[1])
			rules1 := parseRuleList(m[2])
			rules2 := parseRuleList(m[3])
			r[idx] = choiceRule{optionRule{rules1, &r}, optionRule{rules2, &r}}
		}
	}
	return r
}

type rules map[int]rule

func (r rules) createPatterns() string {

	regex := r[0].getRegexString()
	return fmt.Sprintf("^%s$", regex)
}

type rule interface {
	getRegexString() string
}

type optionRule struct {
	options []int
	rules   *rules
}

func (r optionRule) getRegexString() string {
	rtn := ""
	for _, opt := range r.options {
		rtn += (*r.rules)[opt].getRegexString()
	}
	return rtn
}

type charRule struct {
	c     rune
	rules *rules
}

func (r charRule) getRegexString() string {
	return string(r.c)
}

type choiceRule struct {
	o1 optionRule
	o2 optionRule
}

func (r choiceRule) getRegexString() string {
	return fmt.Sprintf("(%s|%s)", r.o1.getRegexString(), r.o2.getRegexString())
}

const DemoInput = `0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb`

const DemoInput2 = `42: 9 14 | 10 1
9: 14 27 | 1 26
10: 23 14 | 28 1
1: "a"
11: 42 31
5: 1 14 | 15 1
19: 14 1 | 14 14
12: 24 14 | 19 1
16: 15 1 | 14 14
31: 14 17 | 1 13
6: 14 14 | 1 14
2: 1 24 | 14 4
0: 8 11
13: 14 3 | 1 12
15: 1 | 14
17: 14 2 | 1 7
23: 25 1 | 22 14
28: 16 1
4: 1 1
20: 14 14 | 1 15
3: 5 14 | 16 1
27: 1 6 | 14 18
14: "b"
21: 14 1 | 1 14
25: 1 1 | 1 14
22: 14 14
8: 42
26: 14 22 | 1 20
18: 15 15
7: 14 5 | 1 21
24: 14 1

abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
bbabbbbaabaabba
babbbbaabbbbbabbbbbbaabaaabaaa
aaabbbbbbaaaabaababaabababbabaaabbababababaaa
bbbbbbbaaaabbbbaaabbabaaa
bbbababbbbaaaaaaaabbababaaababaabab
ababaaaaaabaaab
ababaaaaabbbaba
baabbaaaabbaaaababbaababb
abbbbabbbbaaaababbbbbbaaaababb
aaaaabbaabaaaaababaa
aaaabbaaaabbaaa
aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
babaaabbbaaabaababbaabababaaab
aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba`

const RawInput = `59: 48 127
37: 68 121 | 56 125
51: 48 125 | 125 121
31: 121 120 | 125 126
118: 121 127 | 125 96
127: 125 125 | 121 48
0: 8 11
42: 125 34 | 121 103
89: 86 125 | 39 121
44: 125 125 | 121 125
60: 124 121 | 54 125
38: 24 125 | 7 121
102: 23 121 | 65 125
33: 74 125 | 68 121
119: 97 121 | 128 125
92: 121 41 | 125 100
32: 82 121 | 29 125
12: 125 51 | 121 44
14: 83 125 | 130 121
10: 125 102 | 121 88
129: 125 56 | 121 56
46: 121 121 | 121 125
40: 125 92 | 121 6
6: 48 64
4: 96 125 | 124 121
91: 108 125 | 39 121
53: 121 22 | 125 127
41: 125 113 | 121 124
13: 70 121 | 47 125
58: 121 17 | 125 77
15: 104 121 | 95 125
110: 33 125 | 60 121
116: 125 117 | 121 54
96: 125 121
121: "a"
124: 121 125 | 48 121
20: 35 121 | 27 125
19: 125 90 | 121 3
3: 109 121 | 12 125
88: 125 94 | 121 30
48: 121 | 125
73: 125 72 | 121 18
115: 121 4 | 125 81
95: 32 125 | 55 121
66: 121 68 | 125 85
94: 37 125 | 61 121
67: 121 54 | 125 22
71: 122 121 | 73 125
112: 9 121 | 77 125
63: 121 61 | 125 53
50: 125 93 | 121 14
2: 117 125 | 51 121
43: 121 49 | 125 114
55: 118 121 | 69 125
72: 125 9 | 121 107
101: 79 121 | 35 125
87: 121 51 | 125 96
23: 121 84
80: 91 125 | 112 121
84: 121 46
77: 96 121 | 46 125
1: 113 125 | 68 121
122: 63 125 | 26 121
8: 42
62: 132 125 | 75 121
45: 127 125 | 74 121
22: 125 125 | 48 121
21: 125 44 | 121 96
120: 125 76 | 121 119
103: 125 15 | 121 71
123: 113 121 | 68 125
70: 121 87 | 125 1
82: 48 44
65: 125 45 | 121 25
9: 121 124 | 125 46
57: 22 121 | 51 125
52: 121 124
27: 44 121 | 127 125
18: 125 111 | 121 66
131: 68 121 | 117 125
78: 54 121 | 74 125
17: 48 113
100: 96 125 | 54 121
104: 20 125 | 115 121
98: 56 121 | 54 125
132: 121 68 | 125 113
109: 127 48
76: 121 13 | 125 38
130: 106 121 | 57 125
126: 121 10 | 125 36
83: 125 2 | 121 81
24: 125 78 | 121 5
99: 96 125 | 96 121
49: 82 121 | 21 125
54: 121 125 | 125 121
107: 121 113 | 125 22
105: 125 80 | 121 40
128: 101 121 | 110 125
34: 50 121 | 105 125
47: 59 125 | 69 121
64: 121 117 | 125 85
113: 121 121 | 125 121
30: 125 116
28: 124 121 | 74 125
61: 46 125 | 85 121
75: 121 127 | 125 117
97: 89 125 | 62 121
111: 48 68
90: 28 121 | 41 125
79: 46 48
106: 127 121 | 117 125
7: 125 52 | 121 131
39: 46 121 | 44 125
117: 48 48
93: 121 58 | 125 16
11: 42 31
114: 121 123 | 125 67
68: 125 125 | 125 121
69: 22 125 | 96 121
16: 121 99 | 125 98
5: 125 74 | 121 127
74: 125 125 | 121 121
81: 121 46 | 125 68
125: "b"
26: 121 77 | 125 129
29: 56 121
86: 85 125 | 54 121
56: 125 125
36: 125 19 | 121 43
108: 74 121 | 56 125
35: 121 85 | 125 56
85: 121 125
25: 121 56 | 125 44

abbbbabaaababbbbbbabbbbb
ababbabbbbabbaaaabbabbabababaaba
baabbaaaaaaaabbaaaababaa
abababbaaababbaaaabaaabb
babaababbaaabaaababaabaabbbabbaabbbbabab
aaabbbbbaaaabbaaaabbaabb
abbbbabaaaaababababbbbbb
bbbbabbabbbbbabbaaaaaaab
bababbbaabaabaabbaabbaab
aababaabaaabbbbbbbaabbabbabaabbabbbaabbb
aabbbabbabababbaabaaaaba
babbbbbaababaaaaabbbbbaabaaaaabbbbaaabab
abbbabbbbaabaaaaaaabbaaaaabaabbb
abbaaaaababbbbbaaabbbaba
aababbaaabaabaaaababaabb
babaabaaababbbbabbbababaabbabbabaabbbaaabbabbbbabbbbaaba
bbabaabababbababaaabbaba
bbbbbabababbabaaaaaaabaaaaaaabbabababbbabbbaaaabababaabb
aababbbabbabaaabbbababaaaaaaaabbbabbbabaaabaabab
baababbbabbabaaaaaabaabbbaaaaabbbbbbababbaaabbaa
abbbbabaaabbbbbaaaaaaabb
baaaababbbaaaabbababbbbaaaaababaaaaaabaabbaaabbabbbbbaaa
abaabaabbbaaabbbaabbaabaaaabbbbb
abbbbaaabababaabaaabaaaaabbbbbbaabbbbaab
abbaabaaabbbbbababbaaabb
aaaaabbaabbbabbbaababbbbbabbbaababbabaaa
abbaaaaabaabaaabbbaabbbb
aaaabbbabbabaaabbbaababa
baaaabbaababaaaabaabbabbabbaababaaababaa
bababbababaaaaaaababbaaa
aabbaababaabbababbbabbba
aabbabbaabaaaaaabbabbabbbaaaabaababbbbbaaabababb
bbaaabaababbaababaabaaaaabbabbbaabbbabaa
bbabbabbbababababaabbababbbababbbbaaaaba
ababbaabbbabbbabbababbbb
baabaaababababaaaaaaabbabbaaaaabaabbaabaaaaaaaab
ababbbbababaaaabbabaaaaa
aaaaababaababbababbbabab
baabbababaababbbaabbabaa
bbaabbaabbaabaabababbbab
bababaaaaabbaabaaaabaaba
babbaababbbababaaaaabaaa
bbbabaaababaabaabababbbb
baaaaabaabbaaaaaaaaabbbb
aabbbabbaaaabbaabbbabaabaabbbaba
aababababaaaaaaabbbaaabb
babbababbbabbbabbaababba
bababababaaaaaaabaaaaabababbaabbbbbababb
aababaabbbaabbaabbbababb
bbaaaaabbbbbaaabbaabababaabbaababbbaababbabbbabbaabbbaabbaaabbaaabbbaaaa
aabaababbabbabaabbaaaaba
bbaabaaababbbaababbbbbbabaaaaaaaabbabaabbababbbb
aaaaabaababbbbaaabbaaabb
bbaabbaaabbaababbababbbaabbaabbb
bbbbbbbbbabbababbbabaaabaabbbbaaabbaaaba
bbabaabbbbaabaaaaabaabaa
baaaabaababaabbaaabaabababaaaabbabbaaaba
abaababababaabaaababbaba
abbabbabbbaabaabbabbaabb
babababaabbaabaaabbbaaaa
babbaababaaaaababaaabbaaabbaaaba
baaaaabbabbaababaaaaabaaababbbbbaabbbaba
babbbbaabaaaabbaabbaaaab
babbbaaababbbaaaabbaaabb
abaababaaabbaababbbbbaab
babbbaabaaaababaabaababbaabbaababbabaabbbbbbbaabaabbaaaabaababaa
aaaabababaaababbabbbbabaaabbbaba
aaabbbbbaababaabababbbaa
aabababaaabaaaabbaaaaaaabbabaabbabbbaabaaaaabaabbbabbbaa
abababaabbababbabbbbabab
bbaabbabaabaababaaaaabbbaaabbaab
bbabaaabbababaabbaaaabaabaaababa
aabbabbaabababbaaaabaaaaababbbaa
aaaaabbbbabbbababaaabbbb
baabbabbababaaaabbbbaaaa
babbbbbaaaaaabababbbabab
bbbaababbaaababbaaabbbaa
abbabbabbaaaabbababbaaab
aabaabababbabbbbabaababaabababbaaababbabbbbaabaabbbbabbabbabbababbbabbab
abaabbabaabbababbbabaaababaaabab
baaabaaabbbabaaaabaaaaba
babaabaababaabbabbabaaabbabbabbaabaaabba
abbbbbaabbbaababbaababba
ababbbbaaaababbbbabbabbb
baaaaaaababbbaababaaabab
aabaaaabaabbaaabaababaaa
ababababaaababaababbbabaaabbabbabbbaabbb
abbbaaabbabbbabbaabaaabb
aaabbbabbbbaaaababbabbba
abbbbaaabaaaabbbaaaabaaa
abaabbabbbbbbababbbbbbaabaabbbab
bbabaabaabbbabbabbbaaaba
abbbaaabaabbababaaabaabb
aaaabbaabbaabbaababaababbbababbb
aababbbaaaaababbabaaabba
bbababbabbbbabbabaababba
bbbababaaaaabbaaabaabaaabababbaaaaabaaba
babaaaabbbaaaabbaaabbaaa
bbaaabbbbabbbabaaaaababababababaaabbbaba
aaaaabbababbbaabbbaaaaaaabbabbbb
baaaabababbbbababbbaaabb
aaaababbabaabaabbaaaababbbbabaaababbaaaaabaaabbbaaabaabbaabaaabb
abaaaaabbabbaabaaaaaabbaaabaababbbabaaabaabababb
aabbaabababaabbabbababbabaaabbab
abbbbbaabbbbbabaaabbbaaa
bbaabbabababbbbabbabbaaabbaaaaaaabaaaaaaabbbbbaabbabaaaa
abbaabababbbbabbbbbababaabaabaaaabaabbabbbaabaabaaaabaab
baabbaaabaabbaaaabbaaaaaaabbbababaabaabaabbbabaaabbabbaaaaabaaabbbaaabba
bbaaabaaabababbaaaababababbaaaba
baaaababbbababbaaaabbabb
bbbbbabbabaabbabbbbaabba
abbababbbbbbabbbaaaaabaa
aababbabbaabbabaaabaabbb
babbbabbbbbbbbbbbbababbaababbbbb
abababbabaaaaabababbbbab
babbaabaaabaababbaaabbaa
baabbbaaaababbbbaabbaaabbabbaababbbabbba
baaaabbaabaabaabaababbabbbbbaaaa
babbaabababbbbaaabbababa
abaabbbaabbbbbbababbbabaabbbabaa
baabababaabbbbaabbabbaaa
babbaabaaabbbaababababbb
babaabbaaababbbbaaababbbabbbaaba
baaaaababbbbbbababaabaabaabbaaababbaabbaabbbaaaa
aaaaababbbbbbabbbaaaabaaaaaabbaa
bbaabaababbbaaabaababbbbabababbababaababbbabbababbabbbba
babaababbaaaaaaaabbabbaa
baabababbbbbbababaabbbab
abbbbbabababaaaababbbbab
abbaababbbbababaabbabbaa
bbbaaabbabaaabaaaaababababbbbbabababbaabbbaaabbababbbbabbaaaababaaaaabbbbabbbbbbbaabbaba
baabaaabbabbbbbaabbbbaab
baabaaababaaaaaabaabaaba
bbaaabbbbabbbbbabbbabbab
aabbbbaaaabbbaabababbbaa
bbabbbbbaaaaabbbababbbbaaababbbaabbaabab
bbbbbbaaaabaababaababaaa
bababbaabaaaaababbbaaaaa
aabaabaabaabbabbbbbaaaabbaabbbba
bbbabbbbbbabbbbaaaaabaaaaaaabbab
baabbbabbbbabbbbbbabbbbaabbababa
abababaaabbaaaaaaabbababbabaaaabaabaaaabbbbaabbbaaabbabaaaababbaabbabbbb
bbaabaaaabbbbabbbbbabbab
abaaaaaaabaabaabaabbabaa
babbabaaaababaabbbaababa
aabbaaabaabbbabbabbbbaab
abbaaaaabbabbaaabababbabbbaaabbaababbaba
babbababbbbbbbbbabaabbbb
bbbaaaababbbabbaabaaaaaabababbbabaaaabbababaabaabaababaa
ababbabbbbabbaaaaaababba
babbaaaabababbabababbabbbbabaabbaabaabbbabbaabbbbbbaaaaa
bbabbaaaaababbaababbaaababbabbaa
aabbabbabababaababbbbbbabaabbbbbbbbbbaaa
bbaaaaabbaabaabbabababbb
aabbaabbabbbbbaabbbabbaa
aaaababbbbabaaabaabaabaa
aabbbabbbbaaaaaaabbbaabb
babbbabaaabbaababbabbbaa
abbbbababbbabaaabbbaaabb
baabbbaabbaabaaaabbabbbb
abbbbaaabbbbbabaaaabbabb
aaaababbbaaababbbbabaaaa
bbaaaabbbabbbabbabaaaabb
babbabaababbbbbaaabaaaabaaababbbabaaababaaababbabaaabbbb
bababbababaabaaaababbbab
bbbbbbbbabbbabbaababaabb
abbbbabababaaabbaabbbbbb
bbbbabbbbbbabaaaabbaaabb
bbabbbabbbaabaaaabaabaaaaabaababbabaaababbabbaababbaabba
baabbaaaaaababbbabaabbbb
aaaaaaabbbbbbbaabbbbabbabbbabaabaaababbbbbbbbbaabbaabaaa
babbbaabbaaaabbbbabbaaab
bbaabbbbabbbbaababababab
aaabbbabaaaaabbbbaaabbaa
baabaaabaabbabbaababbaaa
abbbbbabbaaaaabaaaaababbbbbbbbbbaababbbaabbabababaaabbbb
abbbbbaababaaabbabbaaaaaabbbbaabbbbbbaaa
baababababbbbabaabaaabbb
baaaabbaaababbaaabababbb
bbaabbababbbbbbaabbbabab
bbaaaaabaababbabbabbbaaabbaabbaaababbaaabbbbbaaa
babbbabbbaabbbaaaaaaaaaa
baabababbaababbbaabaaabb
bababaababaabbabbaabbbbb
bbaabaabbbabbabbbbaaabba
bbaabbababbaabaabaabbbab
babbbbaabaabbabbaabbabaa
aabbbbbabaaaabbbbbaaaabbbaaabaaabbbaaaabbbbabbaabbbbbaaaaaabaabbbaabbbab
babbbbaaaaababbbbbaabaaaaababaab
aaaaababababbabaabbabababaaababbabaaabba
bbaaabaabaaaaaababbbbbbb
aabababaabbbaabaaaabbabbbaabbbbb
bbbbbbabababaabbaaababaabaaaaababbbbbaabababaaaababbbaabaaaaabbabbabbbbbbbabbbbababaaaaababaaabb
bbbbabbbaabbbbaaaabbbbaaabbaabaabbaabbba
bbbabaaabbbbbabbbbbabbab
bbbbabbbbbbbbabbbbaaabaaaabbabbb
bbbabaaabbbbabbababbaabb
aabbabbabbabbabbabaaabab
baaaabbbbaabaaaaababbaaa
bbbababbbababaaaaabaaaaaabaaaaababbbbbaabbababababbbaaababaababbaaaaaaba
babbbbaaaaaabababbbbabbbaaabbbabaabbabaabbaababb
abbbbaaaababaabbababbbbbbaabbabbbbababbbbbbbbaaaaabaaaaaaaabaabaabaabbabbbbaabaaaabaaaab
bbaabaabaababbbaabbbabbaaaaabaab
abbaababababbbbaabbaaabb
baabababbbaaaaabaaabbaaa
baaaabababbbbbabaaabaaab
abbbbbabbaabbabaabaaabaa
baabbbbabbaaaaaabbababaa
babbbbbaaaabbbabaababbaabababaabbbaabbbb
bbaaaabbbababbaaabababaaababbabb
bbbbbababababbbaaaaabbaabbabaababbabbabbaabaabababbbbaabbbabbbba
abbbbbabaaabababbaabbbab
bbbbbabbaabaababababbbaa
babbabaaaaaababbbaababaa
aabbbbaaabbbbbaaabaababbbbbabaaabbabbbba
bbaaabbbaabaababbbaabbbb
abbbbbabbaaaabbbaaabbabb
baabaabbbababbbaaabbbaaa
bbaabaaaaaaaabababbbabbbbaaabbaa
bbababbababbbaabbbababbaabbbaaabababbbab
babbbaaababbaaaaaaabbabb
abaabaaabababaaabbbaaaabaabbbababbbbaaaabababaaabaaabaab
baabaaaaaaabaaaabbaaabab
abaabaabbabbbababaaaabbaabbaaaab
baabaaaabbaabaaabbbbbbba
aaabbbbbbaaaaababbbbaaba
aabbababbbaabbababbbbababaaaabbaabaaaaababbbaaaa
abbabbbbbbbaaabbbbbaabaaabaaabbbbbbaabaabbbabbbb
babbbbbabababbbabbbbbaaa
bababaaaaaaabababaaaaaaaaaabbabaaabbbaaabaaabbaababaaaaa
abaabaaaaaaaabbabbbbbbaaaabbaababbabbbba
abababbaabbbbabaabaabbbb
bbabbbabaaaaabbbbbbaabbabababbaabbaaaaabaababbaaaaaaababaaaabbababaaabba
baaaaabbaaaabbbaaaabaabb
baabaaaabaaaabbbaabbbbbb
babaaabbababbabbaaaaaaaa
abbabbabbabbababaaaaaaaa
baabaaaababbbabababbaabb
abbbabbbbaabaabbbabbababaabbaabaaaabbbaa
bbabbbabbabbbbbabbabbbaabababbba
baaaaabbabaaaaababbababa
aaabbbbbbaaabaaabbbaabaa
ababbabbaaababbbbabbbbbb
bababbabbabbabbaaabbabbbbbabaababbbabbaabbbabaabbabbaaaabbaaaaabbbbabaabaaaabbbabaababbb
aaabaaaaabbbbbaabbaabaababbabbabbbababbb
bbbbabbabaaaaabbaabbabbb
abbaaaaabaaabaaaaabaabba
aaaabbaaabbbbabbbbbbbbab
ababaaaaaaaabbbaaabbabbabbbababaabbababbaaaabaabababaaabbaaabbba
bbbaaaabbabbbbaaabaabaaaabaabbabaaabaaaaaaaabaab
aabbbaabbaabaabbbbaabbba
babbbbaaabbabbabaabbbaabbaaaaaabbbaaabbbabbabbaa
baabaabbbbbabaabababbbaa
bbabaaabbabaaabbaaaabaab
baabbbaabababbbabbbaaaaaaabbbaba
abbbbabaabaabaababbbbbaaabaababbbbaaaaaabaaabbaa
babbabbaabbaababaabaabba
baaaabbabbbbabbbbbababab
baaaabbaababbbbaaaaabbaaaababbaaaabbaabbababbbaaaabbabbb
aaaabababaabaaaabaabbbbb
aaabababbbbbbbbbabbabbaa
bbaabbaaaabbaababbaaabaaababbbaabbabbaba
abaaaaabaaabababbaaaabbaabbbbaba
bbbbbabbbaaaaaababababababbbabaa
bbbabababbaaaaabbbbabbba
bbabbbababbaaababbabaababbabbabbabbbababaaabbbaaabbaaaaaabbaaaababaabbab
abbbbabaabbbabbaabbaababbabaaaba
baaaaabbaaabbbababbabbba
bbbababaabbbbaaabbaabbba
aabbabbababbbaaaaaaaaaaa
abaabbbaaaaaabbaabbbbbbb
babbbabbaabaababaaaaaaaa
bbabaaabbbbaabbbbbbaaaabbbaababa
baaaababaabaaabababbbbaabbaaabbbaaaaaaab
aabbaabaaabbabbababbabbb
bbbaababaabaaabaaaababaa
aaabaaaabbbabaaababaabbababbbbbbbaaabaab
babbbbaaabbbaaabbaaaababbababaaabaabaaaababbbabbbbbbaaabbbbbaabb
bbaabbaababbbaaaaababbbbaaabbbaaaaaaaaab
babababaabaaaaabbaaabbbb
baabaababbaababbaaaabaabbbbaaabbabaaabaa
aaaaabbbbabbbabbaabbababbababbaa
bbbbbabbbababbbabaaabbab
babbabababbbaaabbaaababbbababaabbbaaaaba
bbababbabbbaabbbabaaaababbbaababbaaaabbaaabbbbabbababaaaaaaabbbabbaaaababaaababb
abbbbabbbababbaaababbaaa
babbabaaaabbbbbaabbbaaba
baaabaaabbbbabbaaabbbbab
bbaaaabbbbbababababbbbbababbaabaabaaaabaaaaaaaaa
baabbbbabababbaaaaabaaab
bbaaaabbabbaabaaaabababaabbaabababbabbabbbabababaaaabaaa
abbbbbababbbbabbbbaababb
aabbabbabaaaaaabaabbaabaabaaabaaabbabaaaaabbbbbb
aaabababaaabaaaabbbaaaabaabbbbbaabbbaaba
aababbabbabaababaaabbbaa
baaaaaaabbaabbaababbbbab
bababbbaaabbbbaabbababaa
aabbaaababbbbabbababbaaa
bbbaabababbbabbababaabbb
aaaababababaabababbbaaaa
babbbbbaabababbabbaaabab
abbbabbaabaaaaaaababaaaababaaaababbbbbabbaaabbbb
babaabaabaaaabaaaabbbaabaabaaababababbaaabbbababbbaababbabbabaab
bbbabaaabbbaabbaabaabbbbababbbbaabaaaababbaaabaaabbaaaaaaaabbabb
bababbabbabaabbaabbabbaa
bbbbabbbbbaaabaaabbabbbb
bbababbaaababbbabbbaaaaa
babaaabbaababbabaabbbaaa
aababbbababbbabbbbaabbbb
baaaabbbbaaaababbbbbbabbaabababaaabbbbaabbbaabaabbbbaaabaaababaaaaabbbba
babbaaaababbbaaaabbbbaab
bababbaabaababaaabaabbbaabbabbbabaabbbabaabababbbbabbabbaabaaabbbaabbbbabaabaaab
aabbabbaaabbbaabaabbbbbb
abaababababbbbaabbaaaaba
bbbabaabaaaabbaabbbbaabb
bababbabaabbababaaaababaaaabbbaa
aabbbaabbaabbaaabbbbaaaa
aababbaabbbbbbaabaabbbbaaabaabba
baababbbababbbbaabbbbaab
babbaabaaabbaaabbaaaabbabaaaabba
abbabbabbaabaabbabbbbaab
abababaaabaaaaaaabbbbaaaabbababbaabaabbababbaaab
abbbbabbaaabababaaaabaab
bbbbbabbbbbaababbaaababa
aabbaabbbaaaabbaaabbabbb
aaabaaabbbaaaababaabbbabbaabababaaababaabbababaabbabbabbaababaaabbabbbababbbbbbaaaababaa
aaaaabbabababbaabaabaaabbababababbabbaabaabbbaaa
bababaaababbbbbaaaabababbbabbbba
bbbabaaabbbabababababbaaabbbaaabbbbbaaab
bbbabababaabbbaaabbbaabb
aaabbbbbbabaabababbbbbbb
babbaabaaaaaababbbbaaabb
baaaaaababaabaabbaabaaba
bababbaaaababbbbabbbaaabbabbaaaabbabaabbaabaabaa
aabababaababbaabaabbbaba
bbabbbabaabbbbaaabbaaaaaababbbbaaaababaa
abaabbbababaabbaabbaabbb
babaababbbaaaabbaaabbbabaaaabaaa
aabababaabbbbbbaaabbbaaa
babaababbbabaabbbaabbabaaaabaaababbbbaab
aaaaabbbbabbbaabbbabbaba
aaaababbbbbbbbaaaaaabaab
aabbbbaabbbaababbabababb
babbbaabaabbababbabbabbaaaabaabb
aaabaaaaabababaabaabaaaabaabbbabbbaabaaaaabaabbabaaaaabbaabbbbbbbaaababb
aaabbbabaabbbbbaabbabaab
abaaaaababbaabaabbabbbaa
bbbbbbaababbaaaabbbabbaa
abbababbbabbabaabbbaaaaa
baaaaaaaaababbbbababbbab
aaababbabbabbabbbaababbabbabaaabbaabbbbbaaabaaabaabbabab
aabbbbbababbaaaaaaaaabaaababbaaabbbabbbb
bbabbabbaaabaaaabbababab
baaaaaaaabbbabbaaaaaaaaa
abbbbabbbaabaaaaaaabbbaaaaaabaab`
