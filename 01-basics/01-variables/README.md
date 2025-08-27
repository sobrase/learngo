# Exercise 1: Student Grade Report System

## 🎯 **Objective**
Create a student grade report system that demonstrates Go's variable declaration, basic types, and type inference.

## 🎨 **Goal: Print a Beautiful Grade Report**

You will create a program that generates a formatted grade report for students, showing their personal information, grades, and statistics.

## 📋 **Requirements**
1. Declare variables using `var` keyword with explicit types
2. Use short variable declaration (`:=`)
3. Work with different basic types:
   - `int`, `int8`, `int16`, `int32`, `int64`
   - `uint`, `uint8`, `uint16`, `uint32`, `uint64`
   - `float32`, `float64`
   - `string`
   - `bool`
   - `byte` (alias for uint8)
   - `rune` (alias for int32)
4. Demonstrate type conversion between numeric types
5. Use `fmt` package to print formatted output

## 🎯 **Specific Tasks**

### **Task 1: Student Information**
Create variables for a student's information:
- Name (string)
- Student ID (uint64)
- Age (int)
- Height in meters (float64)
- Weight in kg (float32)
- Is enrolled (bool)
- Grade level (byte)
- GPA (float64)

### **Task 2: Course Grades**
Create variables for 5 different courses:
- Course names (strings)
- Grades (int8, 0-100)
- Credit hours (uint8)
- Pass/Fail status (bool)

### **Task 3: Grade Calculations**
- Calculate average grade (float64)
- Convert grades to letter grades (A, B, C, D, F)
- Calculate GPA (4.0 scale)
- Determine academic standing (string)

### **Task 4: Print the Report**
Create a beautiful, formatted grade report that looks like this:

```
╔══════════════════════════════════════════════════════════════╗
║                    STUDENT GRADE REPORT                      ║
╠══════════════════════════════════════════════════════════════╣
║  Student: John Doe (ID: 123456789)                          ║
║  Age: 20 | Height: 1.75m | Weight: 70.5kg | Enrolled: Yes   ║
║  Grade Level: 2 | Current GPA: 3.75                         ║
╠══════════════════════════════════════════════════════════════╣
║  COURSE GRADES:                                              ║
║  ┌─────────────────┬─────────┬──────────┬─────────┬────────┐ ║
║  │ Course          │ Grade   │ Letter   │ Credits │ Pass   │ ║
║  ├─────────────────┼─────────┼──────────┼─────────┼────────┤ ║
║  │ Mathematics     │   92    │    A     │    4    │  Yes   │ ║
║  │ Physics         │   88    │    B     │    4    │  Yes   │ ║
║  │ Chemistry       │   95    │    A     │    3    │  Yes   │ ║
║  │ English         │   85    │    B     │    3    │  Yes   │ ║
║  │ History         │   78    │    C     │    3    │  Yes   │ ║
║  └─────────────────┴─────────┴──────────┴─────────┴────────┘ ║
╠══════════════════════════════════════════════════════════════╣
║  SUMMARY:                                                    ║
║  • Average Grade: 87.6                                       ║
║  • Total Credits: 17                                         ║
║  • Academic Standing: Good Standing                          ║
║  • GPA: 3.75                                                 ║
╚══════════════════════════════════════════════════════════════╝
```

## 🎯 **Expected Output**
Your program should produce a beautifully formatted grade report with:
- Student's personal information
- Course grades in a table format
- Grade calculations and statistics
- Academic standing assessment

## 💡 **Hints**
- Use `fmt.Printf` with format specifiers for alignment
- Use Unicode box-drawing characters for the borders
- Convert numeric grades to letter grades using type conversion
- Use different numeric types to demonstrate type safety
- Calculate GPA using the 4.0 scale (A=4.0, B=3.0, C=2.0, D=1.0, F=0.0)

## 🚀 **Bonus Challenges**
1. Add color to the output using ANSI escape codes
2. Include grade trends (improving/declining)
3. Add course difficulty ratings
4. Calculate class rank percentile
5. Include attendance statistics

## 🎨 **Learning Outcomes**
By completing this exercise, you will:
- Master Go's variable declaration syntax
- Understand type safety and conversion
- Learn formatted output techniques
- Create visually appealing console applications
- Practice real-world data representation

This exercise transforms abstract variable concepts into a practical, visual application that you can be proud of! 🎓
