package pkgc

import (
	"context"

	"github.com/kyawmyintthein/twirp-grpc-envoy-poc/protopbs/protos/svcc"
)

type svcC struct {
}

func NewSvcC() svcc.CService {
	return &svcC{}
}

func (svc *svcC) CallServiceC(ctx context.Context, req *svcc.GetServiceCRequest) (*svcc.GetServiceCResponse, error) {
	var (
		resp svcc.GetServiceCResponse
		err  error
	)
	var responses []*svcc.ServiceCResponse
	responses = append(responses, &svcc.ServiceCResponse{
		ServiceName: "C",
		Status:      "Ok",
	})
	resp.Responses = responses
	return &resp, err
}
