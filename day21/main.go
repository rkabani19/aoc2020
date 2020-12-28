package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/rkabani19/aoc2020/utils"
)

type Food struct {
	ingredients StringSet
	allergens   StringSet
}

type StringSet map[string]bool

func main() {
	input, _ := utils.ReadFile("day21.txt")

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %s\n", part2(input))
}

func part2(input []string) string {
	foods := parseInput(input)
	candidates := getCandidates(foods)
	allergens := findAllergens(candidates)

	values := make([]string, 0, len(allergens))
	for _, v := range allergens {
		values = append(values, v)
	}
	sort.Strings(values)

	allergens2 := make(map[string]string)
	for k, v := range allergens {
		allergens2[v] = k
	}

	keys := make([]string, 0, len(allergens))
	for _, v := range values {
		keys = append(keys, allergens2[v])
	}

	return strings.Join(keys, ",")
}

func part1(input []string) int {
	foods := parseInput(input)
	candidates := getCandidates(foods)
	allergens := findAllergens(candidates)
	safeIngredients := getSafeIngredients(allergens, foods)

	return countSafeIngredients(safeIngredients, foods)
}

func countSafeIngredients(safeIngredients StringSet, foods []Food) int {
	count := 0
	for _, f := range foods {
		for i := range f.ingredients {
			if safeIngredients[i] {
				count++
			}
		}
	}
	return count
}

func getSafeIngredients(allergens map[string]string, foods []Food) StringSet {
	out := make(StringSet)
	for _, f := range foods {
		for i := range f.ingredients {
			if _, ok := allergens[i]; !ok {
				out[i] = true
			}
		}
	}
	return out
}

func getCandidates(foods []Food) map[string]StringSet {
	candidates := make(map[string]StringSet)
	for _, f := range foods {
		for k := range f.allergens {
			if len(candidates[k]) == 0 {
				candidates[k] = f.ingredients
			} else {
				candidates[k] = intersect(candidates[k], f.ingredients)
			}
		}
	}
	return candidates
}

func findAllergens(candidates map[string]StringSet) map[string]string {
	allergens := make(map[string]string)
allergenLoop:
	for len(candidates) > 0 {
		for a, c := range candidates {
			i, ok := singleton(c)
			if !ok {
				continue
			}
			for _, c := range candidates {
				delete(c, i)
			}
			allergens[i] = a
			delete(candidates, a)
			continue allergenLoop
		}
		break
	}
	return allergens
}

func singleton(s StringSet) (e string, b bool) {
	for e := range s {
		return e, len(s) == 1
	}
	return "", false
}

func intersect(s1 StringSet, s2 StringSet) StringSet {
	res := make(StringSet)
	for k := range s1 {
		if s2[k] {
			res[k] = true
		}
	}
	return res
}

func parseInput(input []string) []Food {
	res := make([]Food, 0)
	for _, line := range input {
		ss := make(StringSet)
		sp := strings.Split(line, " (contains ")
		for _, s := range strings.Split(sp[0], " ") {
			ss[s] = true
		}
		f := Food{ingredients: ss}

		ss = make(StringSet)
		sp = strings.Split(sp[1], ", ")
		for i, a := range sp {
			if i == len(sp)-1 {
				ss[a[:len(a)-1]] = true
				continue
			}
			ss[a] = true
		}
		f.allergens = ss
		res = append(res, f)
	}
	return res
}
