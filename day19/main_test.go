package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testInput = `0: 4 1 5
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

func TestCreateGrammar(t *testing.T) {
	rules, _, err := parseInput(testInput)
	require.NoError(t, err)

	gram, err := parseGrammar(rules)
	require.NoError(t, err)

	assert.Equal(t, []termRule{
		termRule{From: 4, To: "a"},
		termRule{From: 5, To: "b"},
	}, gram.terminals)

	assert.Len(t, gram.rules, 8)
}

func TestGrammar(t *testing.T) {
	rules, tests, err := parseInput(testInput)
	require.NoError(t, err)

	gram, err := parseGrammar(rules)
	require.NoError(t, err)

	s := countMatches(gram, tests)
	assert.Equal(t, 2, s)
}

const testInput2 = `42: 9 14 | 10 1
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

func TestPartTwoNoReplace(t *testing.T) {
	rules, tests, err := parseInput(testInput2)
	require.NoError(t, err)

	g, err := parseGrammar(rules)
	require.NoError(t, err)

	s := countMatches(g, tests)

	assert.Equal(t, 3, s)
}

func TestPartTwo(t *testing.T) {
	rules, tests, err := parseInput(testInput2)
	require.NoError(t, err)

	replaceRules(rules)

	g, err := parseGrammar(rules)
	require.NoError(t, err)

	g = g.Minimalise()

	s := countMatches(g, tests)

	assert.Equal(t, 12, s)
}
