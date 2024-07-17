package main

import (
	"errors"
	"fmt"
)

var pl = fmt.Println

type DepartmentOfHospital struct {
	department string
	patients   int
}

// "Cardiology", "Orthopaedics", "Neurology", "Gynaecology", "Oncology"

//Return the department with the most patients in the department, by N patients (N being an int)

func CountPatientsByDepartments(patientInDepartments []string) ([]DepartmentOfHospital, error) {
	patientsByDepartment := []DepartmentOfHospital{
		{department: "Cardiology", patients: 0},
		{department: "Neurology", patients: 0},
		{department: "Orthopaedics", patients: 0},
		{department: "Gynaecology", patients: 0},
		{department: "Oncology", patients: 0},
	}

	/*
		patientsByDept := make(map[string]int)

		patientsByDept["Cardiology"] = 0
		patientsByDept["Neurology"] = 0
		patientsByDept["Orthopaedics"] = 0
		patientsByDept["Gynaecology"] = 0
		patientsByDept["Oncology"] = 0
	*/

	//iterate over each patient in department
	for patient_number := 0; patient_number < len(patientInDepartments); patient_number++ {
		currentPatientDept := patientInDepartments[patient_number]

		isValid := checkIfDepartmentIsValid(currentPatientDept, patientsByDepartment)

		if !isValid {
			return patientsByDepartment, errors.New("unknown department")
		}

		patientsByDepartment = addPatientToDept(currentPatientDept, patientsByDepartment)
		//Append string array for department, based on the patient issue.
	}
	return patientsByDepartment, nil
}

func checkIfDepartmentIsValid(name string, validDepartments []DepartmentOfHospital) bool {

	for _, value := range validDepartments {

		if name == value.department {
			return true
		}

	}
	return false
}

// With a pointer solution
/*
func addPatientToDept(deptName string, patientsByDepartment []DepartmentOfHospital) []DepartmentOfHospital {
	//increase the patient count in the department by one and return the new count on patients in the slice of structs.

	for deptNumber, _ := range patientsByDepartment {
		currentDept := &patientsByDepartment[deptNumber]

		if deptName == currentDept.department {
			currentDept.patients++
		}
	}

	return patientsByDepartment
}
*/

// With returning a different struct solution
func addPatientToDept(deptName string, patientsByDepartment []DepartmentOfHospital) []DepartmentOfHospital {
	//increase the patient count in the department by one and return the new count on patients in the slice of structs.
	updated := []DepartmentOfHospital{}

	for _, currentDept := range patientsByDepartment {

		if deptName == currentDept.department {
			currentDept.patients++
		}
		updated = append(updated, currentDept)
	}

	return updated
}

func mostPatients(departments []string) (int, error) {
	patientsByDepartment, err := CountPatientsByDepartments(departments)

	if err != nil {
		return 0, errors.New("unknown department")
	}
	//Call function to give slice of strings to return a slice of department hospital

	result := findMostPatients(patientsByDepartment)
	return result, nil
}

// Calculate max patient department.
func findMostPatients(hospital []DepartmentOfHospital) int {
	most_patients := 0

	for _, dept := range hospital { //Dept is assigning the individual structs a name in the slice of structs in hospital.

		if dept.patients > most_patients {
			most_patients = dept.patients
		}
	}
	return most_patients
}

func main() {
}
