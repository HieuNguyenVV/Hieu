package main

// import (
// 	"fmt"
// 	"sync"
// )

// type Student struct {
// 	name string
// }
// type Course struct {
// 	name string
// }

// //Student after Registration
// type StudentRegistration struct {
// 	student Student
// 	course  []Course
// }

// type StudentRegistrationResult struct {
// 	registration StudentRegistration
// 	err          error
// }

// //List Student after register
// type RegisterStudentsResults struct {
// 	lisStudent []StudentRegistrationResult
// }

// //Funtion registers use goroutine
// func RegisterStudents(student []Student, course []Course) RegisterStudentsResults {
// 	//create channel
// 	output := make(chan RegisterStudentsResults)
// 	input := make(chan StudentRegistrationResult)
// 	var wg sync.WaitGroup
// 	defer close(output)
// 	//Add jop for channel input
// 	for _, student := range student {
// 		wg.Add(1)
// 		go ConcurrentRegisterStudent(student, course, input)
// 	}
// 	//receive value from channel and send value to channel output
// 	go handleResults(input, output, &wg)
// 	wg.Wait()
// 	close(input)
// 	return <-output
// }

// //Funtion receive value from channel and send value to channel output
// func handleResults(input chan StudentRegistrationResult, output chan RegisterStudentsResults, wg *sync.WaitGroup) {
// 	var resutls RegisterStudentsResults
// 	for resutl := range input {
// 		resutls.lisStudent = append(resutls.lisStudent, resutl)
// 		wg.Done()
// 	}
// 	output <- resutls
// }

// //funtion send value to channel input
// func ConcurrentRegisterStudent(student Student, course []Course, input chan StudentRegistrationResult) {
// 	result := RegisterStudent(student, course)
// 	input <- result
// }

// // Funtion Register for one student
// func RegisterStudent(student Student, course []Course) StudentRegistrationResult {
// 	return StudentRegistrationResult{
// 		registration: StudentRegistration{
// 			student: student,
// 			course:  course,
// 		},
// 	}
// }
// func main() {
// 	student := []Student{
// 		Student{name: "Hieu"},
// 		Student{name: "Van"},
// 	}
// 	course := []Course{
// 		Course{name: "Toan"},
// 		Course{name: "Li"},
// 	}
// 	results := RegisterStudents(student, course)
// 	fmt.Println(results)
// }
