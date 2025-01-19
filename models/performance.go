package models

type Performance struct {
	EmployeeID string
	Score      int
	Feedback   string
}

func NewPerformance(employeeID string, score int, feedback string) *Performance {
	return &Performance{
		EmployeeID: employeeID,
		Score:      score,
		Feedback:   feedback,
	}
}
