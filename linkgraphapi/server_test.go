package linkgraphapi_test

import (
	"github.com/google/uuid"
	"github.com/igavrysh/linksrus/linkgraph/graph"
	"github.com/igavrysh/linksrus/linkgraphapi/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	gc "gopkg.in/check.v1"
)

var _ = gc.Suite(new(ServerTestSuite))
var minUUID = uuid.Nil
var maxUUID = uuid.MustParse("ffffffff-ffff-ffff-ffff-ffffffffffff")

type ServerTestSuite struct {
	g graph.Graph

	netListener *bufconn.Listener
	grpcSrv     *grpc.Server

	cliConn *grpc.ClientConn
	cli     proto.LinkGraphClient
}