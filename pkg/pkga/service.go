package pkga

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kyawmyintthein/twirp-grpc-envoy-poc/protopbs/protos/svca"
	"github.com/kyawmyintthein/twirp-grpc-envoy-poc/protopbs/protos/svcb"
	"github.com/kyawmyintthein/twirp-grpc-envoy-poc/protopbs/protos/svcc"
	"github.com/twitchtv/twirp"
	"golang.org/x/net/http2"
)

type svcA struct {
	clientB svcb.BService
	clientC svcc.CService
}

func NewSvcA() svca.AService {
	return &svcA{
		clientB: svcb.NewBServiceProtobufClient("http://localhost:3002", &http.Client{
			//	Transport: transport2(),
		}, twirp.WithClientPathPrefix("/rz")),
		clientC: svcc.NewCServiceProtobufClient("http://localhost:3003", &http.Client{
			//Transport: transport2(),
		}, twirp.WithClientPathPrefix("/rz")),
	}
}

func (svc *svcA) CallServiceA(ctx context.Context, req *svca.GetServiceARequest) (*svca.GetServiceAResponse, error) {
	var (
		resp svca.GetServiceAResponse
		err  error
	)
	svcbResp, err := svc.clientB.CallServiceB(ctx, &svcb.GetServiceBRequest{Count: req.Count})
	if err != nil {
		return &resp, err
	}

	svccResp, err := svc.clientC.CallServiceC(ctx, &svcc.GetServiceCRequest{Count: req.Count})
	if err != nil {
		return &resp, err
	}

	var responses []*svca.ServiceAResponse
	for i := int64(1); i <= req.Count; i++ {
		responses = append(responses, &svca.ServiceAResponse{
			ServiceName: "A",
			Status:      "Ok",
		})
	}
	for _, resp := range svcbResp.Responses {
		responses = append(responses, &svca.ServiceAResponse{
			ServiceName: resp.ServiceName,
			Status:      resp.Status,
		})
	}
	for _, resp := range svccResp.Responses {
		responses = append(responses, &svca.ServiceAResponse{
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
