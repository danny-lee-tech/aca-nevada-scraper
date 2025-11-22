package acanevada

import "fmt"

type Plan struct {
	name           string
	deductible     string
	outOfPocketMax string

	primaryCareVisit        string
	specialistVisit         string
	otherPractictionerVisit string
	preventativeCare        string

	lab     string
	xray    string
	imaging string

	genericDrugs           string
	preferredBrandDrugs    string
	nonPreferredBrandDrugs string
	specialtyDrugs         string

	outpatientFacility string
	outpatientSurgery  string

	emergencyRoom      string
	emergencyTransport string
	urgentCare         string

	inpatientHospitalServices string
	inpatientPhysician        string
}

func (p Plan) Print() {
	fmt.Printf("Plan Name: %s\n", p.name)
	fmt.Printf("Deductible: %s\n", p.deductible)
	fmt.Printf("Out of Pocket Max: %s\n\n", p.outOfPocketMax)

	fmt.Printf("Primary Care Visit: %s\n", p.primaryCareVisit)
	fmt.Printf("Specialist Visit: %s\n", p.specialistVisit)
	fmt.Printf("Other Practitioner Visit: %s\n", p.otherPractictionerVisit)
	fmt.Printf("Preventative Care: %s\n\n", p.preventativeCare)

	fmt.Printf("Lab: %s\n", p.lab)
	fmt.Printf("X-Ray: %s\n", p.xray)
	fmt.Printf("Imaging: %s\n\n", p.imaging)

	fmt.Printf("Generic Drugs: %s\n", p.genericDrugs)
	fmt.Printf("Preferred Brand Drugs: %s\n", p.preferredBrandDrugs)
	fmt.Printf("Non-Preferred Brand Drugs: %s\n\n", p.nonPreferredBrandDrugs)

	fmt.Printf("Outpatient Facility: %s\n", p.outpatientFacility)
	fmt.Printf("Outpatient Surgery: %s\n\n", p.outpatientSurgery)

	fmt.Printf("Emergency Room: %s\n", p.emergencyRoom)
	fmt.Printf("Emergency Transport: %s\n", p.emergencyTransport)
	fmt.Printf("Urgent Care: %s\n\n", p.urgentCare)

	fmt.Printf("Inpatient Hospital Services: %s\n", p.inpatientHospitalServices)
	fmt.Printf("Inpatient Physician: %s\n", p.inpatientPhysician)
}
