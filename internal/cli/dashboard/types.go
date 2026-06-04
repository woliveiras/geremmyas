package dashboard

// Warning records a non-fatal parse issue.
type Warning struct {
	Path    string
	Message string
}

// TaskStats counts checkbox states in tasks.md.
type TaskStats struct {
	Total      int
	Done       int
	InProgress int
	Pending    int
}

// SpecSummary is parsed metadata for one spec folder.
type SpecSummary struct {
	Number     int
	Slug       string
	Title      string
	Family     string
	Phase      int
	Status     string
	Owner      string
	DependsOn  []int
	Origin     string
	Body       string
	HasPlan    bool
	PlanBody   string
	TasksBody  string
	Tasks      TaskStats
	Deprecated bool
	Dir        string
}

// Phase groups specs within a family.
type Phase struct {
	Number int
	Specs  []SpecSummary
}

// Family groups phases.
type Family struct {
	Name   string
	Phases []Phase
}

// PRD is a product requirements document.
type PRD struct {
	Slug  string
	Title string
	Date  string
	Path  string
	Body  string
}

// Bugfix is a bugfix document.
type Bugfix struct {
	Slug  string
	Title string
	Date  string
	Path  string
	Body  string
}

// SpecDates holds git-derived timestamps for a spec.
type SpecDates struct {
	Number        int
	CreatedAt     string
	ApprovedAt    string
	ImplementedAt string
}

// Metrics holds computed delivery metrics.
type Metrics struct {
	VelocityByMonth   map[string]int
	LeadTimeByMonth   map[string]float64
	AvgLeadTimeDays   float64
	AvgReviewTimeDays float64
	AvgImplTimeDays   float64
	InProgressCount   int
	PhaseBreakdown    map[string]map[int]map[string]int // family -> phase -> status -> count
	FamilyProgress    map[string]float64
	GitAvailable      bool
}

// LinkIndex supports artifact linking (spec 0006).
type LinkIndex struct {
	SpecsByNumber map[int]*SpecSummary
	PRDByPath   map[string]*PRD
	OriginToSpecs map[string][]int
	Dependents  map[int][]int
}

// DashboardData is the full parsed dashboard model.
type DashboardData struct {
	Families  []Family
	PRDs      []PRD
	Bugfixes  []Bugfix
	Warnings  []Warning
	Dates     map[int]SpecDates
	Metrics   Metrics
	Links     LinkIndex
}

const ungroupedFamily = "Ungrouped"
