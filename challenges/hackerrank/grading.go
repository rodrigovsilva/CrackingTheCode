package hackerrank

/*
 * Complete the 'gradingStudents' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY grades as parameter.
 */

func gradingStudents(grades []int32) []int32 {
	for i := 0; i < len(grades); i++ {
		grades[i] = roundGrade(grades[i])
	}

	return grades
}

func roundGrade(grade int32) int32 {
	rest := grade % 5

	if grade >= 38 && rest >= 3 {
		return grade + (5 - rest)
	}

	return grade
}
