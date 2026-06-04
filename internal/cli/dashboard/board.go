package dashboard

import "strings"

// BoardColumn groups specs by status for kanban view.
type BoardColumn struct {
	Status string
	Specs  []SpecSummary
}

var boardStatuses = []string{"Draft", "In Review", "Approved", "Implemented"}

// BoardColumns builds kanban columns for a family's specs.
func BoardColumns(specs []SpecSummary) []BoardColumn {
	buckets := map[string][]SpecSummary{}
	for _, st := range boardStatuses {
		buckets[st] = nil
	}
	buckets["Deprecated"] = nil
	for _, spec := range specs {
		st := spec.Status
		if spec.Deprecated || strings.EqualFold(st, "Deprecated") {
			buckets["Deprecated"] = append(buckets["Deprecated"], spec)
			continue
		}
		if _, ok := buckets[st]; !ok {
			st = "Draft"
		}
		buckets[st] = append(buckets[st], spec)
	}
	var cols []BoardColumn
	for _, st := range append(boardStatuses, "Deprecated") {
		cols = append(cols, BoardColumn{Status: st, Specs: buckets[st]})
	}
	return cols
}

// FamilySpecs flattens all specs in a family.
func FamilySpecs(fam Family) []SpecSummary {
	var out []SpecSummary
	for _, ph := range fam.Phases {
		out = append(out, ph.Specs...)
	}
	return out
}
