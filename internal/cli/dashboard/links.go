package dashboard

import "strings"

// BlockedBy returns dependency specs not yet implemented.
func BlockedBy(spec SpecSummary, index LinkIndex) []SpecSummary {
	var out []SpecSummary
	for _, n := range spec.DependsOn {
		dep, ok := index.SpecsByNumber[n]
		if !ok {
			continue
		}
		if !strings.EqualFold(dep.Status, "Implemented") {
			out = append(out, *dep)
		}
	}
	return out
}

// Blocks returns specs that depend on this one.
func Blocks(specNum int, index LinkIndex) []SpecSummary {
	var out []SpecSummary
	for _, n := range index.Dependents[specNum] {
		if s, ok := index.SpecsByNumber[n]; ok {
			out = append(out, *s)
		}
	}
	return out
}

// PRDLink resolves origin to a dashboard path.
func PRDLink(origin string, index LinkIndex) (href string, found bool) {
	origin = normalizePath(origin)
	if origin == "" {
		return "", false
	}
	if _, ok := index.PRDByPath[origin]; ok {
		return "/prds/" + prdSlugFromPath(origin) + "/index.html", true
	}
	return "", false
}

func prdSlugFromPath(path string) string {
	_, slug, ok := ParseDateFromFilename(path[strings.LastIndex(path, "/")+1:])
	if ok {
		return slug
	}
	return strings.TrimSuffix(path, ".md")
}

// DetectCycles returns warning messages for circular depends_on chains.
func DetectCycles(index LinkIndex) []string {
	var warnings []string
	visiting := map[int]bool{}
	visited := map[int]bool{}
	var dfs func(int, []int)
	dfs = func(n int, stack []int) {
		if visiting[n] {
			warnings = append(warnings, formatCycle(stack, n))
			return
		}
		if visited[n] {
			return
		}
		visiting[n] = true
		stack = append(stack, n)
		if s, ok := index.SpecsByNumber[n]; ok {
			for _, dep := range s.DependsOn {
				dfs(dep, stack)
			}
		}
		delete(visiting, n)
		visited[n] = true
	}
	for n := range index.SpecsByNumber {
		dfs(n, nil)
	}
	return warnings
}

func formatCycle(stack []int, repeat int) string {
	return "circular dependency detected involving spec " + itoa(repeat)
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var b [12]byte
	i := len(b)
	for n > 0 {
		i--
		b[i] = byte('0' + n%10)
		n /= 10
	}
	return string(b[i:])
}
