package acanevada

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

const InitialUrl string = "https://enroll.nevadahealthlink.com/prescreener/"

func RetrievePlans() ([]Plan, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	c, _ := chromedp.NewContext(ctx)
	defer func() {
		if err := chromedp.Cancel(c); err != nil {
			panic("chromedp could not be cancelled")
		}
	}()

	fmt.Println("Scraping: " + InitialUrl)

	navigateToPlanList(&c)

	planList := []Plan{}

	count, err := getTotalNumberOfPlans(&c)
	if err != nil {
		return nil, err
	}

	totalPages := count/12 + 1

	for page := 1; page <= totalPages; page++ {
		if page > 1 {
			err := nextPage(&c)
			if err != nil {
				return nil, err
			}
		}

		planCountOnPage, err := countPlansOnPage(&c)
		if err != nil {
			return nil, err
		}

		for i := 1; i <= planCountOnPage; i++ {
			plan := Plan{}
			navigateToPlan(&c, &plan, i)
			planList = append(planList, plan)

			err := goToPage(&c, page)
			if err != nil {
				return nil, err
			}
		}
	}

	takeScreenshot(&c)

	return planList, nil
}

func takeScreenshot(c *context.Context) error {
	var buf []byte

	err := chromedp.Run(*c,
		chromedp.FullScreenshot(&buf, 90),
	)
	if err != nil {
		return err
	}

	if err := os.WriteFile("screenshot.png", buf, 0644); err != nil {
		log.Fatal(err)
	}

	return nil
}

func getTotalNumberOfPlans(c *context.Context) (int, error) {
	var countStr string
	err := chromedp.Run(*c,
		chromedp.TextContent(`#filteredPlanCount`, &countStr, chromedp.ByQuery),
	)
	if err != nil {
		return 0, nil
	}

	count, err := strconv.Atoi(countStr)

	if err != nil {
		return 0, nil
	}
	return count, nil
}

func countPlansOnPage(c *context.Context) (int, error) {
	var nodes []*cdp.Node
	err := chromedp.Run(*c,
		chromedp.Nodes(`#mainSummary .cp-tile`, &nodes, chromedp.ByQueryAll),
	)
	if err != nil {
		return 0, err
	}
	count := len(nodes)

	return count, nil
}

func navigateToPlanList(c *context.Context) error {
	err := chromedp.Run(*c,
		chromedp.Navigate(InitialUrl),
		chromedp.SendKeys(`input[data-testid="zip-code-input"]`, "89118", chromedp.ByQuery),
		chromedp.SendKeys(`input[id="household-member-0-birthdate-picker"]`, "01012005", chromedp.ByQuery),
		chromedp.SendKeys(`input[data-testid="household-income-input"]`, "65000", chromedp.ByQuery),
		chromedp.Click("body"),
		chromedp.Sleep(2*time.Second),
		chromedp.Click(`button[data-testid="btn-see-savings"]`),
		chromedp.WaitVisible(`button[data-testid="btn-next"]`),
		chromedp.Click(`button[data-testid="btn-next"]`),
		chromedp.WaitVisible(`input[id="skipButton"]`),
		chromedp.Click(`input[id="skipButton"]`),
		chromedp.WaitVisible(`input[id="premiumAfterCredit"]`),
		chromedp.Click(`input[id="premiumAfterCredit"]`),
		//chromedp.Click(`input[id="filter_checkbox_BRONZE"]`),
		//chromedp.Click(`input[id="filter_checkbox_SILVER"]`),
		chromedp.Sleep(2*time.Second),
	)
	if err != nil {
		return err
	}

	return nil
}

func navigateToPlan(c *context.Context, plan *Plan, index int) error {
	fmt.Println("Navigating to Plan " + strconv.Itoa(index))
	err := chromedp.Run(*c,
		chromedp.Click(`#mainSummary > div:nth-child(`+strconv.Itoa(index)+` of .cp-tile) .cp-tile__body > a`, chromedp.ByQuery),
		chromedp.WaitVisible(`#simplifiedDeductibleDetail .details p`),
	)
	if err != nil {
		return err
	}

	fmt.Println("Setting fields for Plan " + strconv.Itoa(index))

	setName(c, plan)
	setCompanyAndNetwork(c, plan)
	setTierAndTypes(c, plan)
	setPremiumMonthly(c, plan)

	setDeductible(c, plan)
	setOutOfPocketMax(c, plan)

	setPrimaryCareVisit(c, plan)
	setSpecialistVisit(c, plan)
	setOtherPractictionerVisit(c, plan)
	setPreventativeCare(c, plan)

	setLab(c, plan)
	setXray(c, plan)
	setImaging(c, plan)

	setGenericDrugs(c, plan)
	setPreferredBrandDrugs(c, plan)
	setNonPreferredBrandDrugs(c, plan)
	setSpecialtyDrugs(c, plan)

	setOutpatientFacility(c, plan)
	setOutpatientSurgery(c, plan)

	setEmergencyRoom(c, plan)
	setEmergencyTransport(c, plan)
	setUrgentCare(c, plan)

	setInpatientHospitalServices(c, plan)
	setInpatientPhysician(c, plan)

	fmt.Println("Going back to list of plans")

	err = chromedp.Run(*c,
		chromedp.Click(`a.back-to-all-plans-link-detail`, chromedp.ByQuery),
		chromedp.WaitVisible(`#mainSummary > div:nth-child(`+strconv.Itoa(index)+` of .cp-tile) .cp-tile__body > a`),
		chromedp.Sleep(1*time.Second/10),
	)
	if err != nil {
		return err
	}

	return nil
}

func nextPage(c *context.Context) error {
	err := chromedp.Run(*c,
		chromedp.Click(`.cp-pagination__btn--right`, chromedp.ByQuery),
		chromedp.Sleep(1*time.Second/10),
		chromedp.Evaluate(`window.scrollTo(0, document.body.scrollHeight)`, nil),
	)
	if err != nil {
		return err
	}

	return nil
}

func goToPage(c *context.Context, page int) error {
	for i := 1; i < page; i++ {
		err := nextPage(c)
		if err != nil {
			return err
		}
	}

	takeScreenshot(c)

	return nil
}

func setName(c *context.Context, plan *Plan) error {
	return setTextSimple(c, &plan.name, `.ps-detail__highlights-table tbody tr:nth-child(1) td`)
}

func setCompanyAndNetwork(c *context.Context, plan *Plan) error {
	err := setTextSimple(c, &plan.network, `tr > td.cp-tile__value a:not([role])`)
	if err != nil {
		return err
	}

	switch plan.network {
	case "NVBBSP", "Premier":
		plan.company = "Ambetter"
	case "Anthem Battle Born with RX Choice Tiered Network", "HMO and Rx Choice Tiered Network", "Anthem Battle Born with Pharmacy Base Network", "Pathway X - HMO Dental Prime and Rx Choice Tiered Network":
		plan.company = "Anthem"
	case "HPN Ind On Ex BBSP Select", "HPN Ind On Exchange", "HPN Ind On Exchange Select":
		plan.company = "Health Plan of Nevada"
	case "Imperial Value":
		plan.company = "Imperial Insurance Companies"
	case "Med Network", "Value Network":
		plan.company = "Select Health"
	case "Nevada":
		plan.company = "CareSource"
	case "Molina Healthcare of Nevada Inc.":
		plan.company = "Molina"
	}

	return nil
}

func setPremiumMonthly(c *context.Context, plan *Plan) error {
	return setTextSimple(c, &plan.premiumMonthly, `.cp-tile__premium-amount`)
}

func setTierAndTypes(c *context.Context, plan *Plan) error {
	var bannerString string
	err := setTextSimple(c, &bannerString, `div.cp-tile__metal-tier`)
	if err != nil {
		return err
	}

	if strings.Contains(bannerString, "Bronze") {
		plan.tier = "Bronze"
	} else if strings.Contains(bannerString, "Silver") {
		plan.tier = "Silver"
	} else if strings.Contains(bannerString, "Gold") {
		plan.tier = "Gold"
	} else if strings.Contains(bannerString, "Catastrophic") {
		plan.tier = "Catastrophic (Age < 30)"
	}

	if strings.Contains(bannerString, "HSA") {
		plan.isHsa = true
	}

	if strings.Contains(bannerString, "HMO") {
		plan.planType = "HMO"
	} else if strings.Contains(bannerString, "EPO") {
		plan.planType = "EPO"
	}

	return nil
}

func setDeductible(c *context.Context, plan *Plan) error {
	return setTextSimple(c, &plan.deductible, `#simplifiedDeductibleDetail .details p`)
}

func setOutOfPocketMax(c *context.Context, plan *Plan) error {
	return setTextSimple(c, &plan.outOfPocketMax, `#simplifiedOOPMaxDetail .details p`)
}

func setPrimaryCareVisit(c *context.Context, plan *Plan) error {
	return setTextRemoveDivSibling(c, &plan.primaryCareVisit, "#doctorVisit1Detail > div:nth-child(1 of .details)")
}

func setSpecialistVisit(c *context.Context, plan *Plan) error {
	return setTextRemoveDivSibling(c, &plan.specialistVisit, "#doctorVisit2Detail > div:nth-child(1 of .details)")
}

func setOtherPractictionerVisit(c *context.Context, plan *Plan) error {
	return setTextRemoveDivSibling(c, &plan.otherPractictionerVisit, "#doctorVisit3Detail > div:nth-child(1 of .details)")
}

func setPreventativeCare(c *context.Context, plan *Plan) error {
	return setTextRemoveDivSibling(c, &plan.preventativeCare, "#doctorVisit4Detail > div:nth-child(1 of .details)")
}

func setLab(c *context.Context, plan *Plan) error {
	return setTextRemoveDivSibling(c, &plan.lab, "#test1Detail > div:nth-child(1 of .details)")
}

func setXray(c *context.Context, plan *Plan) error {
	return setTextRemoveDivSibling(c, &plan.xray, "#test2Detail > div:nth-child(1 of .details)")
}

func setImaging(c *context.Context, plan *Plan) error {
	return setTextRemoveDivSibling(c, &plan.imaging, "#test3Detail > div:nth-child(1 of .details)")
}

func setGenericDrugs(c *context.Context, plan *Plan) error {
	return setTextRemoveDivSibling(c, &plan.genericDrugs, "#drug1Detail > div:nth-child(1 of .details)")
}

func setPreferredBrandDrugs(c *context.Context, plan *Plan) error {
	return setTextRemoveDivSibling(c, &plan.preferredBrandDrugs, "#drug2Detail > div:nth-child(1 of .details)")
}

func setNonPreferredBrandDrugs(c *context.Context, plan *Plan) error {
	return setTextRemoveDivSibling(c, &plan.nonPreferredBrandDrugs, "#drug3Detail > div:nth-child(1 of .details)")
}

func setSpecialtyDrugs(c *context.Context, plan *Plan) error {
	return setTextRemoveDivSibling(c, &plan.specialtyDrugs, "#drug4Detail > div:nth-child(1 of .details)")
}

func setOutpatientFacility(c *context.Context, plan *Plan) error {
	return setTextRemoveDivSibling(c, &plan.outpatientFacility, "#outpatient1Detail > div:nth-child(1 of .details)")
}

func setOutpatientSurgery(c *context.Context, plan *Plan) error {
	return setTextRemoveDivSibling(c, &plan.outpatientSurgery, "#outpatient2Detail > div:nth-child(1 of .details)")
}

func setEmergencyRoom(c *context.Context, plan *Plan) error {
	return setTextRemoveDivSibling(c, &plan.emergencyRoom, "#urgent1Detail > div:nth-child(1 of .details)")
}

func setEmergencyTransport(c *context.Context, plan *Plan) error {
	return setTextRemoveDivSibling(c, &plan.emergencyTransport, "#urgent2Detail > div:nth-child(1 of .details)")
}

func setUrgentCare(c *context.Context, plan *Plan) error {
	return setTextRemoveDivSibling(c, &plan.urgentCare, "#urgent2Detail > div:nth-child(1 of .details)")
}

func setInpatientHospitalServices(c *context.Context, plan *Plan) error {
	return setTextRemoveDivSibling(c, &plan.inpatientHospitalServices, "#hospital1Detail > div:nth-child(1 of .details)")
}

func setInpatientPhysician(c *context.Context, plan *Plan) error {
	return setTextRemoveDivSibling(c, &plan.inpatientPhysician, "#hospital2Detail > div:nth-child(1 of .details)")
}

func setTextSimple(c *context.Context, field *string, cssSelector string) error {
	err := chromedp.Run(*c,
		chromedp.TextContent(cssSelector, field, chromedp.ByQuery),
	)

	*field = strings.TrimSpace(*field)
	*field = strings.ReplaceAll(*field, ",", " ")
	*field = strings.ReplaceAll(*field, "  ", " ")

	if err != nil {
		return err
	}

	return nil
}

func setTextRemoveDivSibling(c *context.Context, field *string, cssSelector string) error {
	err := chromedp.Run(*c,
		chromedp.Evaluate(`
			var element = document.querySelector("`+cssSelector+` div");
			if (element) {
				element.parentNode.removeChild(element);
			}
		`, nil),
		chromedp.TextContent(cssSelector, field, chromedp.ByQuery),
	)

	if err != nil {
		return err
	}

	*field = strings.TrimSpace(*field)
	*field = strings.ReplaceAll(*field, ",", " ")
	*field = strings.ReplaceAll(*field, "  ", " ")

	return nil
}
