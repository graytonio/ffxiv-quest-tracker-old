package templates

import (
	"sort"

	"github.com/graytonio/ffxivquesttracker/pkg/db"
)

func SortKeys[T any](data map[string]T) []string {
	keys := []string{}
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

type CategorizedQuests map[string]map[string][]db.Quest

func CategorizeQuests(quests []db.Quest) (results CategorizedQuests) {
	results = make(CategorizedQuests)

	for _, q := range quests {
		if _, ok := results[q.Category.Section]; !ok {
			results[q.Category.Section] = make(map[string][]db.Quest)
		}

		results[q.Category.Section][q.Category.Category] = append(results[q.Category.Section][q.Category.Category], q)
	}
	
	return results
}

func SectionComplete(section map[string][]db.Quest) bool {
	for _, category := range section {
		if !CategoryComplete(category) {
			return false
		}
	}
	return true
}

func CategoryComplete(category []db.Quest) bool {
	for _, quest := range category {
		if !quest.Completed {
			return false
		}
	}

	return true
}