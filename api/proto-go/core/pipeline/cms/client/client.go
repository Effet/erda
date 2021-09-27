// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: cms.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/pipeline/cms/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// CmsService cms.proto
	CmsService() pb.CmsServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		cmsService: pb.NewCmsServiceClient(cc),
	}
}

type serviceClients struct {
	cmsService pb.CmsServiceClient
}

func (c *serviceClients) CmsService() pb.CmsServiceClient {
	return c.cmsService
}

type cmsServiceWrapper struct {
	client pb.CmsServiceClient
	opts   []grpc1.CallOption
}

func (s *cmsServiceWrapper) CreateNs(ctx context.Context, req *pb.CmsCreateNsRequest) (*pb.CmsCreateNsResponse, error) {
	return s.client.CreateNs(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *cmsServiceWrapper) ListCmsNs(ctx context.Context, req *pb.CmsListNsRequest) (*pb.CmsListNsResponse, error) {
	return s.client.ListCmsNs(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *cmsServiceWrapper) UpdateCmsNsConfigs(ctx context.Context, req *pb.CmsNsConfigsUpdateRequest) (*pb.CmsNsConfigsUpdateResponse, error) {
	return s.client.UpdateCmsNsConfigs(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *cmsServiceWrapper) DeleteCmsNsConfigs(ctx context.Context, req *pb.CmsNsConfigsDeleteRequest) (*pb.CmsNsConfigsDeleteResponse, error) {
	return s.client.DeleteCmsNsConfigs(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *cmsServiceWrapper) GetCmsNsConfigs(ctx context.Context, req *pb.CmsNsConfigsGetRequest) (*pb.CmsNsConfigsGetResponse, error) {
	return s.client.GetCmsNsConfigs(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}