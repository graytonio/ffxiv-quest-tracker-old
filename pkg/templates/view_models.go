package templates

import (
	"sort"

	"github.com/graytonio/ffxivquesttracker/pkg/db"
)

type CategorizedQuests map[string]map[string][]db.Quest

func CategorizeQuests(quests []db.Quest) (out [][][]db.Quest) {
	results := make(map[string]map[string][]db.Quest)

	for _, q := range quests {
		if _, ok := results[q.Category.Section]; !ok {
			results[q.Category.Section] = make(map[string][]db.Quest)
		}

		results[q.Category.Section][q.Category.Category] = append(results[q.Category.Section][q.Category.Category], q)
	}
	
	return categorizedToSorted(results)
}

func categorizedToSorted(data CategorizedQuests) [][][]db.Quest {
	// Sort keys of the outer map
	outerKeys := make([]string, 0, len(data))
	for k := range data {
		outerKeys = append(outerKeys, k)
	}
	sort.Strings(outerKeys)

	// Initialize the triple slice
	tripleSlice := make([][][]db.Quest, len(outerKeys))

	// Iterate over sorted outer keys
	for i, outerKey := range outerKeys {
		innerMap := data[outerKey]

		// Sort keys of the inner map
		innerKeys := make([]string, 0, len(innerMap))
		for k := range innerMap {
			innerKeys = append(innerKeys, k)
		}
		sort.Strings(innerKeys)

		// Construct the inner slices
		innerSlices := make([][]db.Quest, len(innerKeys))
		for j, innerKey := range innerKeys {
			innerSlices[j] = innerMap[innerKey]
		}

		tripleSlice[i] = innerSlices
	}

	return tripleSlice
}