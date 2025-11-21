package acanevada

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

	outpatientFacility string
	outpatientSurgery  string

	emergencyRoom      string
	emergencyTransport string
	urgentCare         string

	inpatientHospitalServices string
	inpatientPhysician        string
}
