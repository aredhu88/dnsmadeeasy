package main

import (
	"os"
	"testing"

	"github.com/cert-manager/cert-manager/test/acme/dns"

	"github.com/cert-manager/webhook-example/example"
)

var (
	zone = os.Getenv("TEST_ZONE_NAME")
)

func TestRunsSuite(t *testing.T) {
	// The manifest path should contain a file named config.json that is a
	// snippet of valid configuration that should be included on the
	// ChallengeRequest passed as part of the test cases.

  //Options from https://github.com/jetstack/cert-manager/blob/master/test/acme/dns/options.go
	fixture := dns.NewFixture(&DNSMadeEasyProviderSolver{},
		dns.SetResolvedZone(zone),
		dns.SetAllowAmbientCredentials(false),
		dns.SetManifestPath("testdata/dnsmadeeasy"),
		dns.SetBinariesPath("_out/kubebuilder/bin"),
    //dns.SetDNSServer("ns1.sandbox.dnsmadeeasy.com:53"),
	)

	fixture.RunConformance(t)
}
