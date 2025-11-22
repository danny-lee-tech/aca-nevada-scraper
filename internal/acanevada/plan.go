package acanevada

import (
	"fmt"
	"strings"
)

type Plan struct {
	name           string
	company        string
	network        string
	tier           string
	planType       string // HMO or EPO
	isHsa          bool
	premiumMonthly string

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
	fmt.Printf("Company: %s\n", p.company)
	fmt.Printf("Network: %s\n", p.network)
	fmt.Printf("Tier: %s\n", p.tier)
	fmt.Printf("Plan Type: %s\n", p.tier)
	fmt.Printf("HSA Qualified: %s\n", p.tier)
	fmt.Printf("Base Monthly Premium: %s\n", p.premiumMonthly)

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

func PrintPlanCSVHeader() string {
	headers := []string{
		"Plan Name",
		"Company",
		"Network",
		"Tier",
		"Plan Type",
		"HSA Qualified?",
		"Base Monthly Premium",
		"Deductible",
		"Out Of Pocket Max",
		"Primary Care Visit Cost",
		"Specialist Visit Cost",
		"Other Practitioner Visit Cost",
		"Preventative Care Coverage",
		"Lab Tests Cost",
		"X-Ray Cost",
		"Imaging Cost",
		"Generic Drugs Cost",
		"Preferred Brand Drugs Cost",
		"Non-Preferred Brand Drugs Cost",
		"Specialty Drugs Cost",
		"Outpatient Facility Cost",
		"Outpatient Surgery Cost",
		"Emergency Room Cost",
		"Emergency Transport Cost",
		"Urgent Care Cost",
		"Inpatient Hospital Services Cost",
		"Inpatient Physician Cost",
	}

	return fmt.Sprintln(strings.Join(headers, ","))
}

func (p Plan) PrintPlanCSVRow() string {
	return fmt.Sprintf("%s,%s,%s,%s,%s,%t,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s\n",
		p.name,
		p.company,
		p.network,
		p.tier,
		p.planType,
		p.isHsa,
		p.premiumMonthly,
		p.deductible,
		p.outOfPocketMax,
		p.primaryCareVisit,
		p.specialistVisit,
		p.otherPractictionerVisit,
		p.preventativeCare,
		p.lab,
		p.xray,
		p.imaging,
		p.genericDrugs,
		p.preferredBrandDrugs,
		p.nonPreferredBrandDrugs,
		p.specialtyDrugs,
		p.outpatientFacility,
		p.outpatientSurgery,
		p.emergencyRoom,
		p.emergencyTransport,
		p.urgentCare,
		p.inpatientHospitalServices,
		p.inpatientPhysician,
	)
}
