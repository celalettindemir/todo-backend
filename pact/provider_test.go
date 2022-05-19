package pact

import (
	"backend/config"
	"backend/server"
	"fmt"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func TestProvider(t *testing.T) {
	port := 8081
	svr := server.NewServer()

	go svr.StartServer(port)
	pact := dsl.Pact{
		Host:                     "127.0.0.1",
		Provider:                 "Backend",
		Consumer:                 "Frontend",
		DisableToolValidityCheck: true,
	}

	request := types.VerifyRequest{
		ProviderBaseURL: fmt.Sprintf("%s:%d", "http://localhost", config.C.PactPort),
		BrokerToken:     "dUyaVq2xZheha-_mByqRNw",
		BrokerURL:       "https://modanisa-test1.pactflow.io",
		ProviderVersion: config.VERSION,
		PactURLs: []string{
			"https://modanisa-test1.pactflow.io/pacts/provider/Backend/consumer/Frontend/latest",
		},

		PublishVerificationResults: true,
	}

	verifyResponses, err := pact.VerifyProvider(t, request)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(verifyResponses), "pact tests run")
}
