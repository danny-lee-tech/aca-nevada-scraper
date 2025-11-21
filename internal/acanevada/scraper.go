package acanevada

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

const InitialUrl string = "https://enroll.nevadahealthlink.com/prescreener/"

func RetrievePlans() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	c, _ := chromedp.NewContext(ctx)
	defer func() {
		if err := chromedp.Cancel(c); err != nil {
			panic("chromedp could not be cancelled")
		}
	}()

	fmt.Println("Scraping: " + InitialUrl)

	navigateToPlanList(&c)

	plan := &Plan{}

	//var buf []byte

	err := chromedp.Run(c,
		chromedp.Click(`.cp-tile__body .detail:nth-child(1)`, chromedp.ByQuery),
		chromedp.WaitVisible(`#simplifiedDeductibleDetail .details p`),
	)
	if err != nil {
		return "", err
	}

	setName(&c, plan)
	setDeductible(&c, plan)
	setOutOfPocketMax(&c, plan)
	setPrimaryCareVisit(&c, plan)

	plan.primaryCareVisit = strings.TrimSpace(plan.primaryCareVisit)

	log.Println(plan.name)
	log.Println(plan.deductible)
	log.Println(plan.outOfPocketMax)
	log.Println(plan.primaryCareVisit)

	// if err := os.WriteFile("screenshot.png", buf, 0644); err != nil {
	// 	log.Fatal(err)
	// }

	return "", nil
}

func setName(c *context.Context, plan *Plan) error {
	err := chromedp.Run(*c,
		chromedp.TextContent(`.ps-detail__highlights-table tbody tr:nth-child(1) td`, &plan.name, chromedp.ByQuery),
	)

	if err != nil {
		return err
	}

	plan.name = strings.TrimSpace(plan.name)

	return nil
}

func setDeductible(c *context.Context, plan *Plan) error {
	err := chromedp.Run(*c,
		chromedp.TextContent(`#simplifiedDeductibleDetail .details p`, &plan.deductible, chromedp.ByQuery),
	)

	if err != nil {
		return err
	}

	return nil
}

func setOutOfPocketMax(c *context.Context, plan *Plan) error {
	err := chromedp.Run(*c,
		chromedp.TextContent(`#simplifiedOOPMaxDetail .details p`, &plan.outOfPocketMax, chromedp.ByQuery),
	)

	if err != nil {
		return err
	}

	return nil
}

func setPrimaryCareVisit(c *context.Context, plan *Plan) error {
	err := chromedp.Run(*c,
		chromedp.Evaluate(`
			var element = document.querySelector("#doctorVisit1Detail > div:nth-child(1 of .details) div");
			if (element) {
				element.parentNode.removeChild(element);
			}
		`, nil),
		chromedp.TextContent(`#doctorVisit1Detail > div:nth-child(1 of .details)`, &plan.primaryCareVisit, chromedp.ByQuery),
	)

	if err != nil {
		return err
	}

	return nil
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
		chromedp.WaitVisible(`input[id="filter_checkbox_BRONZE"]`),
		chromedp.Click(`input[id="premiumAfterCredit"]`),
		chromedp.Click(`input[id="filter_checkbox_BRONZE"]`),
		chromedp.Sleep(2*time.Second),
	)
	if err != nil {
		return err
	}

	return nil
}
