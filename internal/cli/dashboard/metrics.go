package dashboard

import (
	"math"
	"sort"
	"strings"
)

// ComputeMetrics fills data.Metrics from specs and git dates.
func ComputeMetrics(data *DashboardData) {
	m := &data.Metrics
	m.VelocityByMonth = map[string]int{}
	m.LeadTimeByMonth = map[string]float64{}
	m.PhaseBreakdown = map[string]map[int]map[string]int{}
	m.FamilyProgress = map[string]float64{}

	var leadTimes []float64
	var reviewTimes []float64
	var implTimes []float64
	inProgress := 0

	for _, fam := range data.Families {
		m.PhaseBreakdown[fam.Name] = map[int]map[string]int{}
		total, done := 0, 0
		for _, ph := range fam.Phases {
			if m.PhaseBreakdown[fam.Name][ph.Number] == nil {
				m.PhaseBreakdown[fam.Name][ph.Number] = map[string]int{}
			}
			for _, spec := range ph.Specs {
				if spec.Deprecated {
					continue
				}
				total++
				st := spec.Status
				m.PhaseBreakdown[fam.Name][ph.Number][st]++
				if strings.EqualFold(st, "Implemented") {
					done++
				} else if st != "Draft" {
					inProgress++
				}
				dates, ok := data.Dates[spec.Number]
				if !ok || !m.GitAvailable {
					continue
				}
				created, okC := parseISODate(dates.CreatedAt)
				approved, okA := parseISODate(dates.ApprovedAt)
				impl, okI := parseISODate(dates.ImplementedAt)
				if okI {
					key := impl.Format("2006-01")
					m.VelocityByMonth[key]++
				}
				if okC && okI {
					days := impl.Sub(created).Hours() / 24
					leadTimes = append(leadTimes, days)
					m.LeadTimeByMonth[impl.Format("2006-01")] = median(append(monthLead(m.LeadTimeByMonth, impl.Format("2006-01")), days))
				}
				if okC && okA {
					reviewTimes = append(reviewTimes, approved.Sub(created).Hours()/24)
				}
				if okA && okI {
					implTimes = append(implTimes, impl.Sub(approved).Hours()/24)
				}
			}
		}
		if total > 0 {
			m.FamilyProgress[fam.Name] = float64(done) / float64(total) * 100
		}
	}
	m.AvgLeadTimeDays = median(leadTimes)
	m.AvgReviewTimeDays = median(reviewTimes)
	m.AvgImplTimeDays = median(implTimes)
	m.InProgressCount = inProgress
}

func monthLead(m map[string]float64, key string) []float64 {
	if v, ok := m[key]; ok {
		return []float64{v}
	}
	return nil
}

func median(vals []float64) float64 {
	if len(vals) == 0 {
		return 0
	}
	sort.Float64s(vals)
	mid := len(vals) / 2
	if len(vals)%2 == 0 {
		return (vals[mid-1] + vals[mid]) / 2
	}
	return vals[mid]
}

// ChartJSON returns velocity labels and data for Chart.js.
func ChartJSON(m Metrics) (labels []string, values []int) {
	keys := make([]string, 0, len(m.VelocityByMonth))
	for k := range m.VelocityByMonth {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		labels = append(labels, k)
		values = append(values, m.VelocityByMonth[k])
	}
	return labels, values
}

func roundDays(f float64) float64 {
	return math.Round(f*10) / 10
}
