package academy

import (
	"math"
)

type Student struct {
	Name       string
	Grades     []int
	Project    int
	Attendance []bool
}

// AverageGrade returns an average grade given a
// slice containing all grades received during a
// semester, rounded to the nearest integer.
func AverageGrade(grades []int) int {
	result := 0
	for i := 0; i < cap(grades); i++ {
		result = result + grades[i]
	}
	if result == 0 {
		return result
	}
	var final float64 = float64(result) / float64(cap(grades))
	return int(math.Round(final))
}

// AttendancePercentage returns a percentage of class
// attendance, given a slice containing information
// whether a student was present (true) of absent (false).
//
// The percentage of attendance is represented as a
// floating-point number ranging from 0 to 1.
func AttendancePercentage(attendance []bool) float64 {
	arrived := 0

	var result float64 = 0
	for i := 0; i < cap(attendance); i++ {
		if attendance[i] == true {
			arrived = arrived + 1
		}
	}
	if arrived == 0 {
		return result
	}
	result = float64(arrived) / float64(cap(attendance))

	return result
}

// FinalGrade returns a final grade achieved by a student,
// ranging from 1 to 5.
//
// The final grade is calculated as the average of a project grade
// and an average grade from the semester, with adjustments based
// on the student's attendance. The final grade is rounded
// to the nearest integer.

// If the student's attendance is below 80%, the final grade is
// decreased by 1. If the student's attendance is below 60%, average
// grade is 1 or project grade is 1, the final grade is 1.
func FinalGrade(s Student) int {
	a := AverageGrade(s.Grades)
	b := AttendancePercentage(s.Attendance)
	c := s.Project
	result := math.Round((float64(a) + float64(c)) / 2.0)
	if result < 1 {
		result := 1
		return result
	}
	if result > 5 {
		result := 5
		return result
	}
	if b < 0.8 {
		if b < 0.6 {
			result := 1
			return result
		}
		result := result - 1
		if result < 1 {
			result := 1
			return result
		}
		return int(result)
	}
	if c == 1 {
		result := 1
		return result
	}
	if a == 1 {
		result := 1
		return result
	}

	return int(result)
}

// GradeStudents returns a map of final grades for a given slice of
// Student structs. The key is a student's name and the value is a
// final grade.
func GradeStudents(students []Student) map[string]uint8 {
	expected := map[string]uint8{}

	for i := 0; i < cap(students); i++ {
		name := students[i].Name
		finalGrade := FinalGrade(students[i])
		expected[name] = uint8(finalGrade)
	}

	return expected
}
