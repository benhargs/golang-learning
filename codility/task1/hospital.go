package solution

import(
    "fmt"
)

A := ["Cardiology", "Cardiology", "Cardiology"]
// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func patientCountPerDepartment(patientsDept []string) map[string]int {
    //creating a map for the different departments
    dept := make(map[string]int)

    dept["Cardiology"] = 0
    dept["Orthopaedics"] = 0
    dept["Neurology"] = 0
    dept["Oncology"] = 0
    dept["Gynaecology"] = 0
    
    for index, patientDept := range len(dept) {
        if department == dept[index] {
            dept[patientDept]++
    }

    return dept
}

func mostPatients(patientsDept []string]) int {
    departmentsPatientTotals := patientCountPerDepartment(patientsDept)
    mostPatients := 0

    for dept, patients := range departmentsPatientTotals {
        if patients > mostPatients {
            mostpatients = patients 
        }
    }


    return mostpatients
}


func Solution(A []string) int {
    patientInDepts := patientCountPerDepartment(A)
    highestPatients := mostPatients(patientInDepts)
    fmt.Println(highestPatients)
    // Implement your solution here
    return highestPatients
}