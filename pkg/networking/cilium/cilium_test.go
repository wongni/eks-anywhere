package cilium_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"

	"github.com/aws/eks-anywhere/internal/test"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/networking/cilium"
	"github.com/aws/eks-anywhere/pkg/networking/cilium/mocks"
)

type ciliumtest struct {
	*WithT
	ctx    context.Context
	client mocks.MockClient
	h      *mocks.MockHelm
	spec   *cluster.Spec
	cilium *cilium.Cilium
}

func newCiliumTest(t *testing.T) *ciliumtest {
	ctrl := gomock.NewController(t)
	h := mocks.NewMockHelm(ctrl)
	client := mocks.NewMockClient(ctrl)
	return &ciliumtest{
		WithT:  NewWithT(t),
		ctx:    context.Background(),
		client: *client,
		h:      h,
		spec: test.NewClusterSpec(func(s *cluster.Spec) {
			s.VersionsBundle.Cilium.Version = "v1.9.11-eksa.1"
		}),
		cilium: cilium.NewCilium(client, h),
	}
}

func TestCiliumGenerateManifestSuccess(t *testing.T) {
	tt := newCiliumTest(t)
	// templater tests already test whether templater.GenerateManifest returns expected values or not. This test ensures that cilium.GenerateManifest
	// calls the templater and does not try to load the static manifest like earlier version
	tt.h.EXPECT().Template(
		tt.ctx, gomock.AssignableToTypeOf(""), gomock.AssignableToTypeOf(""), gomock.AssignableToTypeOf(""), gomock.AssignableToTypeOf(map[string]interface{}{}),
	).Return([]byte("manifest"), nil)
	_, err := tt.cilium.GenerateManifest(tt.ctx, tt.spec)
	tt.Expect(err).To(Not(HaveOccurred()), "GenerateManifest() should succeed")
}
