package cluster_test

import (
	"testing"

	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
)

func TestSetDefaultFluxConfigPath(t *testing.T) {
	tests := []struct {
		name           string
		config         *cluster.Config
		wantConfigPath string
	}{
		{
			name: "self-managed cluster",
			config: &cluster.Config{
				Cluster: &anywherev1.Cluster{
					ObjectMeta: metav1.ObjectMeta{
						Name: "c-1",
					},
					Spec: anywherev1.ClusterSpec{
						ManagementCluster: anywherev1.ManagementCluster{
							Name: "c-1",
						},
					},
				},
				FluxConfig: &anywherev1.FluxConfig{},
			},
			wantConfigPath: "clusters/c-1",
		},
		{
			name: "managed cluster",
			config: &cluster.Config{
				Cluster: &anywherev1.Cluster{
					ObjectMeta: metav1.ObjectMeta{
						Name: "c-1",
					},
					Spec: anywherev1.ClusterSpec{
						ManagementCluster: anywherev1.ManagementCluster{
							Name: "c-m",
						},
					},
				},
				FluxConfig: &anywherev1.FluxConfig{},
			},
			wantConfigPath: "clusters/c-m",
		},
		{
			name: "config path is already set",
			config: &cluster.Config{
				Cluster: &anywherev1.Cluster{
					ObjectMeta: metav1.ObjectMeta{
						Name: "c-1",
					},
					Spec: anywherev1.ClusterSpec{
						ManagementCluster: anywherev1.ManagementCluster{
							Name: "c-m",
						},
					},
				},
				FluxConfig: &anywherev1.FluxConfig{
					Spec: anywherev1.FluxConfigSpec{
						ClusterConfigPath: "my-path",
					},
				},
			},
			wantConfigPath: "my-path",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)

			g.Expect(cluster.SetDefaultFluxConfigPath(tt.config)).To(Succeed())
			g.Expect(tt.config.FluxConfig.Spec.ClusterConfigPath).To(Equal(tt.wantConfigPath))
		})
	}
}

func TestSetConfigDefaults(t *testing.T) {
	g := NewWithT(t)
	c := clusterConfigFromFile(t, "testdata/cluster_1_19.yaml")
	originalC := clusterConfigFromFile(t, "testdata/cluster_1_19.yaml")
	g.Expect(cluster.SetConfigDefaults(c)).To(Succeed())
	g.Expect(c).NotTo(Equal(originalC))
}
