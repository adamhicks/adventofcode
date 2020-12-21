package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

type food struct {
	ingredients []string
	allergens   []string
}

func parseInput(content string) []food {
	var ret []food
	for _, l := range strings.Split(content, "\n") {
		if l == "" {
			continue
		}
		l = strings.TrimRight(l, ")")
		parts := strings.Split(l, " (contains ")
		ret = append(ret, food{
			ingredients: strings.Split(parts[0], " "),
			allergens:   strings.Split(parts[1], ", "),
		})
	}
	return ret
}

type common map[string]int

func (i common) Unique() string {
	var max int
	var withMax []string
	for k, v := range i {
		if v > max {
			max = v
			withMax = []string{k}
		} else if v == max {
			withMax = append(withMax, k)
		}
	}
	if len(withMax) > 1 {
		return ""
	}
	return withMax[0]
}

func findAllergenCommonIngredients(flist []food) map[string]common {
	allCommon := make(map[string]common)
	for _, f := range flist {
		for _, a := range f.allergens {
			allCommon[a] = make(common)
		}
	}
	for _, f := range flist {
		for _, a := range f.allergens {
			for _, i := range f.ingredients {
				allCommon[a][i]++
			}
		}
	}
	return allCommon
}

func matchIngredients(allCommon map[string]common) map[string]string {
	matched := make(map[string]string)
	for len(allCommon) > 0 {
		for all, common := range allCommon {
			i := common.Unique()
			if i == "" {
				continue
			}
			matched[i] = all
			delete(allCommon, all)
			for all2 := range allCommon {
				delete(allCommon[all2], i)
			}
		}
	}
	return matched
}

func countUnmatched(flist []food, allMap map[string]string) int {
	var ret int
	for _, f := range flist {
		for _, i := range f.ingredients {
			_, ok := allMap[i]
			if !ok {
				ret++
			}
		}
	}
	return ret
}

func runPartOne() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	flist := parseInput(string(content))
	allCommon := findAllergenCommonIngredients(flist)
	matched := matchIngredients(allCommon)

	i := countUnmatched(flist, matched)
	fmt.Println(i)
	return nil
}

func arrangeIngredients(matched map[string]string) string {
	reverse := make(map[string]string, len(matched))
	for k, v := range matched {
		reverse[v] = k
	}

	aller := make([]string, 0, len(matched))
	for _, v := range matched {
		aller = append(aller, v)
	}
	sort.Strings(aller)

	in := make([]string, 0, len(matched))
	for _, a := range aller {
		in = append(in, reverse[a])
	}

	return strings.Join(in, ",")
}

func runPartTwo() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	flist := parseInput(string(content))
	allCommon := findAllergenCommonIngredients(flist)
	matched := matchIngredients(allCommon)

	s := arrangeIngredients(matched)
	fmt.Println(s)
	return nil
}

func main() {
	if err := runPartOne(); err != nil {
		log.Fatal(err)
	}
	if err := runPartTwo(); err != nil {
		log.Fatal(err)
	}
}
