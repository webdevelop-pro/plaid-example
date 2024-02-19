package main

import (
	"context"
	"fmt"
	"os"

	plaid "github.com/plaid/plaid-go/v20/plaid"
)

func main() {
	configuration := plaid.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", os.Getenv("PLAID_CLIENT_ID"))
	configuration.AddDefaultHeader("PLAID-SECRET", os.Getenv("PLAID_SECRET"))
	configuration.UseEnvironment(os.Getenv("PLAID_ENV"))
	client := plaid.NewAPIClient(configuration)

	user := plaid.LinkTokenCreateRequestUser{
		ClientUserId: "6666-6666-6666-6666",
	}
	request := plaid.NewLinkTokenCreateRequest(
		"Plaid Test",
		"en",
		[]plaid.CountryCode{plaid.COUNTRYCODE_US},
		user,
	)
	request.SetProducts([]plaid.Products{plaid.PRODUCTS_AUTH})
	request.SetLinkCustomizationName("default")
	request.SetWebhook("https://webhook.site/#!/view/XXXXX") // replace with UUID code from webhook.site
	request.SetRedirectUri("https://domainname.com/oauth-page.html")
	request.SetAccountFilters(plaid.LinkTokenAccountFilters{
		Depository: &plaid.DepositoryFilter{
			AccountSubtypes: []plaid.AccountSubtype{plaid.ACCOUNTSUBTYPE_CHECKING, plaid.ACCOUNTSUBTYPE_SAVINGS},
		},
	})
	ctx := context.Background()
	resp, _, err := client.PlaidApi.LinkTokenCreate(ctx).LinkTokenCreateRequest(*request).Execute()
	linkToken := resp.GetLinkToken()
	fmt.Println(linkToken, resp, err)
}
