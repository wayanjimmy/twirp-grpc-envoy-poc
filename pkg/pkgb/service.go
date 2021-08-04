package pkgb

import (
	"context"
	"net/http"

	"github.com/kyawmyintthein/twirp-grpc-envoy-poc/protopbs/protos/svcb"
	"github.com/kyawmyintthein/twirp-grpc-envoy-poc/protopbs/protos/svcc"
	"github.com/twitchtv/twirp"
)

type svcB struct {
	clientC svcc.CService
}

func NewSvcB() svcb.BService {
	return &svcB{
		clientC: svcc.NewCServiceProtobufClient("https://localhost:3003", &http.Client{}, twirp.WithClientPathPrefix("/rz")),
	}
}

func (svc *svcB) CallServiceB(ctx context.Context, req *svcb.GetServiceBRequest) (*svcb.GetServiceBResponse, error) {
	var (
		resp svcb.GetServiceBResponse
		err  error
	)
	svcbResp, err := svc.clientC.CallServiceC(ctx, &svcc.GetServiceCRequest{Count: 2})
	if err != nil {
		return &resp, err
	}
	var responses []*svcb.ServiceBResponse
	responses = append(responses, &svcb.ServiceBResponse{
		ServiceName: "B",
		Status:      "Ok",
	})
	for _, resp := range svcbResp.Responses {
		responses = append(responses, &svcb.ServiceBResponse{
			ServiceName: resp.ServiceName,
			Status:      resp.Status,
		})
	}
	resp.Responses = responses
	return &resp, err
}
