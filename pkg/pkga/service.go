package pkga

import (
	"context"
	"net/http"

	"github.com/kyawmyintthein/twirp-grpc-envoy-poc/protopbs/protos/svca"
	"github.com/kyawmyintthein/twirp-grpc-envoy-poc/protopbs/protos/svcb"
	"github.com/kyawmyintthein/twirp-grpc-envoy-poc/protopbs/protos/svcc"
	"github.com/twitchtv/twirp"
)

type svcA struct {
	clientB svcb.BService
	clientC svcc.CService
}

func NewSvcA() svca.AService {
	return &svcA{
		clientB: svcb.NewBServiceProtobufClient("https://localhost:3002", &http.Client{}, twirp.WithClientPathPrefix("/rz")),
		clientC: svcc.NewCServiceProtobufClient("https://localhost:3003", &http.Client{}, twirp.WithClientPathPrefix("/rz")),
	}
}

func (svc *svcA) CallServiceA(ctx context.Context, req *svca.GetServiceARequest) (*svca.GetServiceAResponse, error) {
	var (
		resp svca.GetServiceAResponse
		err  error
	)
	svcbResp, err := svc.clientB.CallServiceB(ctx, &svcb.GetServiceBRequest{Count: 2})
	if err != nil {
		return &resp, err
	}

	svccResp, err := svc.clientC.CallServiceC(ctx, &svcc.GetServiceCRequest{Count: 2})
	if err != nil {
		return &resp, err
	}

	var responses []*svca.ServiceAResponse
	responses = append(responses, &svca.ServiceAResponse{
		ServiceName: "A",
		Status:      "Ok",
	})
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
