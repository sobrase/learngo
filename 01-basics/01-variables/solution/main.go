package main

import "fmt"

func main() {
	fmt.Println("Exercise 1: Student Grade Report System")
	fmt.Println("======================================")

	// Task 1: Create student information variables
	var studentName string = "John Doe"
	var studentID uint64 = 123456789
	var age int = 20
	var height float64 = 1.75
	var weight float32 = 70.5
	var isEnrolled bool = true
	var gradeLevel byte = 2
	var currentGPA float64 = 3.75

	// Task 2: Create course grades variables
	courseNames := []string{"Mathematics", "Physics", "Chemistry", "English", "History"}
	var grades []int8 = []int8{92, 88, 95, 85, 78}
	var creditHours []uint8 = []uint8{4, 4, 3, 3, 3}
	var passStatus []bool = []bool{true, true, true, true, true}

	// Task 3: Calculate grade statistics
	var totalGrade int = 0
	var totalCredits uint = 0

	for i := 0; i < len(grades); i++ {
		totalGrade += int(grades[i])
		totalCredits += uint(creditHours[i])
	}

	averageGrade := float64(totalGrade) / float64(len(grades))

	// Convert grades to letter grades
	letterGrades := make([]string, len(grades))
	for i, grade := range grades {
		switch {
		case grade >= 90:
			letterGrades[i] = "A"
		case grade >= 80:
			letterGrades[i] = "B"
		case grade >= 70:
			letterGrades[i] = "C"
		case grade >= 60:
			letterGrades[i] = "D"
		default:
			letterGrades[i] = "F"
		}
	}

	// Calculate GPA (4.0 scale)
	var totalGPA float64 = 0
	for i, letter := range letterGrades {
		var gradePoints float64
		switch letter {
		case "A":
			gradePoints = 4.0
		case "B":
			gradePoints = 3.0
		case "C":
			gradePoints = 2.0
		case "D":
			gradePoints = 1.0
		default:
			gradePoints = 0.0
		}
		totalGPA += gradePoints * float64(creditHours[i])
	}
	gpa := totalGPA / float64(totalCredits)

	// Determine academic standing
	var academicStanding string
	switch {
	case gpa >= 3.5:
		academicStanding = "Dean's List"
	case gpa >= 3.0:
		academicStanding = "Good Standing"
	case gpa >= 2.0:
		academicStanding = "Academic Warning"
	default:
		academicStanding = "Academic Probation"
	}

	// Task 4: Print the beautiful grade report
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════════════════════════╗")
	fmt.Println("║                    STUDENT GRADE REPORT                      ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════╣")
	fmt.Printf("║  Student: %s (ID: %d)                          ║\n", studentName, studentID)
	fmt.Printf("║  Age: %d | Height: %.2fm | Weight: %.1fkg | Enrolled: %s   ║\n",
		age, height, weight, func() string {
			if isEnrolled {
				return "Yes"
			} else {
				return "No"
			}
		}())
	fmt.Printf("║  Grade Level: %d | Current GPA: %.2f                         ║\n", gradeLevel, currentGPA)
	fmt.Println("╠══════════════════════════════════════════════════════════════╣")
	fmt.Println("║  COURSE GRADES:                                              ║")
	fmt.Println("║  ┌─────────────────┬─────────┬──────────┬─────────┬────────┐ ║")
	fmt.Println("║  │ Course          │ Grade   │ Letter   │ Credits │ Pass   │ ║")
	fmt.Println("║  ├─────────────────┼─────────┼──────────┼─────────┼────────┤ ║")

	for i := 0; i < len(courseNames); i++ {
		passStr := "Yes"
		if !passStatus[i] {
			passStr = "No"
		}
		fmt.Printf("║  │ %-15s │   %3d    │    %s     │   %2d    │  %s   │ ║\n",
			courseNames[i], grades[i], letterGrades[i], creditHours[i], passStr)
	}

	fmt.Println("║  └─────────────────┴─────────┴──────────┴─────────┴────────┘ ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════╣")
	fmt.Println("║  SUMMARY:                                                    ║")
	fmt.Printf("║  • Average Grade: %.1f                                       ║\n", averageGrade)
	fmt.Printf("║  • Total Credits: %d                                         ║\n", totalCredits)
	fmt.Printf("║  • Academic Standing: %s                          ║\n", academicStanding)
	fmt.Printf("║  • GPA: %.2f                                                 ║\n", gpa)
	fmt.Println("╚══════════════════════════════════════════════════════════════╝")

	// Demonstrate type conversions
	fmt.Println("\n📊 Type Conversion Examples:")
	floatValue := 3.7
	intFromFloat := int(floatValue)
	floatFromInt := float64(10)
	byteFromInt := byte(65) // ASCII 'A'

	fmt.Printf("Float %.1f converted to int: %d\n", floatValue, intFromFloat)
	fmt.Printf("Int %d converted to float64: %.1f\n", 10, floatFromInt)
	fmt.Printf("Int %d converted to byte (ASCII): %c\n", 65, byteFromInt)
}
