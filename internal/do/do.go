package do

import (
	"context"

	"github.com/digitalocean/godo"

	"github.com/AnshulAngaria/ClusterOP/api/v1alpha1"
)

func Create(token string, spec v1alpha1.ClusterSpec) (string, error) {

	client := godo.NewFromToken(token)

	request := &godo.KubernetesClusterCreateRequest{
		Name:        spec.Name,
		RegionSlug:  spec.Region,
		VersionSlug: spec.Version,
		NodePools: []*godo.KubernetesNodePoolCreateRequest{
			&godo.KubernetesNodePoolCreateRequest{
				Size:  spec.NodePools[0].Size,
				Name:  spec.NodePools[0].Name,
				Count: spec.NodePools[0].Count,
			},
		},
	}

	cluster, _, err := client.Kubernetes.Create(context.Background(), request)
	if err != nil {
		return "", err
	}

	return cluster.ID, nil
}

// func ClusterState(c kubernetes.Interface, spec v1alpha1.ClusterSpec, id string) (string, error) {
// 	token, err := getToken(c, spec.TokenSecret)
// 	if err != nil {
// 		return "", err
// 	}

// 	client := godo.NewFromToken(token)

// 	cluster, _, err := client.Kubernetes.Get(context.Background(), id)
// 	return string(cluster.Status.State), err
// }

// func getToken(c *controller.ClusterReconciler, sec string) (string, error) {
// 	namespace := strings.Split(sec, "/")[0]
// 	name := strings.Split(sec, "/")[1]

// 	secret := &corev1.Secret{}

// 	err := c.List(ctx, secret, client.InNamespace(req.Namespace), client.MatchingLabels(rs.Spec.Template.Labels))
// 	if err != nil {
// 		return ctrl.Result{}, err
// 	}

// 	s, err := client.CoreV1().Secrets(namespace).Get(context.Background(), name, metav1.GetOptions{})
// 	if err != nil {
// 		return "", err
// 	}

// 	return string(s.Data["token"]), nil
// }
