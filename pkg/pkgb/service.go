package pkgb

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kyawmyintthein/twirp-grpc-envoy-poc/protopbs/protos/svcb"
	"github.com/kyawmyintthein/twirp-grpc-envoy-poc/protopbs/protos/svcc"
	"github.com/twitchtv/twirp"
	"golang.org/x/net/http2"
)

type svcB struct {
	clientC svcc.CService
}

func NewSvcB() svcb.BService {
	return &svcB{
		clientC: svcc.NewCServiceProtobufClient("https://localhost:3003", &http.Client{
			Transport: transport2(),
		}, twirp.WithClientPathPrefix("/rz")),
	}
}

func (svc *svcB) CallServiceB(ctx context.Context, req *svcb.GetServiceBRequest) (*svcb.GetServiceBResponse, error) {
	var (
		resp svcb.GetServiceBResponse
		err  error
	)
	svcbResp, err := svc.clientC.CallServiceC(ctx, &svcc.GetServiceCRequest{Count: req.Count})
	if err != nil {
		return &resp, err
	}
	var responses []*svcb.ServiceBResponse
	for i := int64(1); i <= req.Count; i++ {
		responses = append(responses, &svcb.ServiceBResponse{
			ServiceName: "B",
			Status:      "Ok",
		})
	}
	for _, resp := range svcbResp.Responses {
		responses = append(responses, &svcb.ServiceBResponse{
			ServiceName: resp.ServiceName,
			Status:      resp.Status,
		})
	}
	resp.Responses = responses
	return &resp, err
}

func transport2() *http2.Transport {
	return &http2.Transport{
		TLSClientConfig:    tlsConfig(),
		DisableCompression: true,
		AllowHTTP:          false,
	}
}

func tlsConfig() *tls.Config {
	crt, err := ioutil.ReadFile("../../server.crt")
	if err != nil {
		log.Fatal(err)
	}

	rootCAs := x509.NewCertPool()
	rootCAs.AppendCertsFromPEM(crt)

	return &tls.Config{
		RootCAs:            rootCAs,
		InsecureSkipVerify: false,
		ServerName:         "localhost",
	}
}
