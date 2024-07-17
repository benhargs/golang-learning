package main

import (
	"errors"
	"testing"
)

func Test2PatientsInNeurology(t *testing.T) {
	departments := []string{"Oncology", "Neurology", "Neurology"}
	got, _ := mostPatients(departments)

	want := 2

	if got != want {
		t.Errorf("Failed. want = %v, got = %v", want, got)
	}
}

func Test3PatientsInCardiology(t *testing.T) {
	departments := []string{"Cardiology", "Cardiology", "Cardiology"}
	got, _ := mostPatients(departments)

	want := 3

	if got != want {
		t.Errorf("Failed. want = %v, got = %v", want, got)
	}
}

func TestNoPatientsInHospital(t *testing.T) {
	departments := []string{}
	got, _ := mostPatients(departments)

	want := 0

	if got != want {
		t.Errorf("Failed. want = %v, got = %v", want, got)
	}
}

func Test4PatientInOncology(t *testing.T) {
	departments := []string{"Oncology", "Oncology", "Oncology", "Oncology", "Cardiology"}
	got, _ := mostPatients(departments)

	want := 4

	if got != want {
		t.Errorf("Failed. want = %v, got = %v", want, got)
	}
}

func Test6PatientInGynaecology(t *testing.T) {
	departments := []string{"Gynaecology", "Gynaecology", "Gynaecology", "Gynaecology", "Cardiology", "Gynaecology", "Gynaecology"}
	got, _ := mostPatients(departments)

	want := 6

	if got != want {
		t.Errorf("Failed. want = %v, got = %v", want, got)
	}
}

func Test7PatientInGynaecology(t *testing.T) {
	departments := []string{"Orthopaedics", "Orthopaedics", "Orthopaedics", "Orthopaedics", "Orthopaedics", "Orthopaedics", "Orthopaedics"}
	got, _ := mostPatients(departments)

	want := 7

	if got != want {
		t.Errorf("Failed. want = %v, got = %v", want, got)
	}
}

func TestNeurologyHasMostPatients(t *testing.T) {
	departments := []string{"Orthopaedics", "Orthopaedics", "Neurology", "Oncology", "Neurology", "Neurology", "Gynaecology", "Cardiology"}
	got, _ := mostPatients(departments)
	want := 3

	if got != want {
		t.Errorf("Failed. want = %v, got = %v", want, got)
	}
}

func TestReportErrorForUnknownDepartment(t *testing.T) {
	departments := []string{"invalid department name"}
	_, err := mostPatients(departments)

	want := errors.New("unknown department")

	if err.Error() != want.Error() {
		t.Errorf("got: %v, want: %v", err, want)
	}
}
