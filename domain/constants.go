package domain

import "time"

//DefaultConnectionTimeout is the timeout for http requests
const DefaultConnectionTimeout time.Duration = 1 * time.Minute

//constants to interage with safra api
const (
	SafraClientID        string = "9dffe873bf3b44b3ad067e87f354bef4"
	SafraSecrect         string = "86a1a12e-595d-474e-8280-d0d26cbb53b7"
	SafraTokenURL        string = "https://idcs-902a944ff6854c5fbe94750e48d66be5.identity.oraclecloud.com/oauth2/v1/token"
	SafraMorningCallsURL string = "https://af3tqle6wgdocsdirzlfrq7w5m.apigateway.sa-saopaulo-1.oci.customer-oci.com/fiap-sandbox/media/v1/youtube?fromData=2020-07-09&toData=2020-07-14&playlist=morningCalls&channel=safra"
)
