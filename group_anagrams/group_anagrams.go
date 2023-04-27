package group_anagrams

type Group struct {
	Definition map[rune]int
	Anagrams   []string
}

func (g *Group) AppendAnagram(str string) {
	g.Anagrams = append(g.Anagrams, str)
}

func groupAnagrams(strs []string) [][]string {
	return GroupAnagrams(strs)
}

func GroupAnagrams(strs []string) [][]string {
	var baseGroups []*Group
	for _, str := range strs {
		definition := readAnagramDefinition(str)
		group, isNew := getGroupForDefinition(definition, baseGroups)
		if isNew {
			baseGroups = append(baseGroups, group)
		}
		group.AppendAnagram(str)
	}

	var result [][]string
	for _, group := range baseGroups {
		result = append(result, group.Anagrams)
	}
	return result
}

func getGroupForDefinition(definition map[rune]int, groups []*Group) (grp *Group, isNew bool) {
	for _, group := range groups {
		if equals(definition, group.Definition) {
			grp = group
			isNew = false
			return grp, isNew
		}
	}
	if grp == nil {
		grp = &Group{
			Definition: definition,
		}
	}
	return grp, true
}

func equals(g1, g2 map[rune]int) bool {
	if len(g1) != len(g2) {
		return false
	}

	for k, v := range g1 {
		if v != g2[k] {
			return false
		}
	}

	return true
}

func readAnagramDefinition(str string) map[rune]int {
	runes := []rune(str)
	anagramDefinition := map[rune]int{}
	for _, r := range runes {
		if _, exists := anagramDefinition[r]; exists {
			anagramDefinition[r] = anagramDefinition[r] + 1
		} else {
			anagramDefinition[r] = 1
		}
	}
	return anagramDefinition
}
