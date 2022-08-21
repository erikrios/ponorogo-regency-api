package rpc

import (
	"context"

	"github.com/erikrios/ponorogo-regency-api/pb"
	"github.com/erikrios/ponorogo-regency-api/service"
)

type ProvinceServer struct {
	pb.UnimplementedProvinceServiceServer
	service service.ProvinceService
}

func NewProvinceServer(service service.ProvinceService) *ProvinceServer {
	return &ProvinceServer{
		service: service,
	}
}

func (p *ProvinceServer) GetProvinces(
	ctx context.Context,
	req *pb.GetProvincesRequest,
) (res *pb.GetProvincesResponse, err error) {
	filter := req.GetFilter()

	responses, serviceErr := p.service.GetAll(ctx, filter.GetName())
	if serviceErr != nil {
		err = handleError(serviceErr)
		return
	}

	res = &pb.GetProvincesResponse{}

	for _, response := range responses {
		province := &pb.Province{
			Id:   response.ID,
			Name: response.Name,
		}
		res.Provinces = append(res.Provinces, province)
	}

	return
}
