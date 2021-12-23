package images

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Config struct {
	KubeConfig string
	NameSpace  string
	All        bool
	client     kubernetes.Interface
}

type image struct {
	Namespace string
	Image     string
}

var namespace string

func (con *Config) GetImages(c *cobra.Command, args []string) error {
	config, err := con.getKubeConfig()
	if err != nil {
		return err
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	con.client = client
	podImages, err := con.fetchImagesWithClient()
	if err != nil {
		return err
	}
	for _, img := range unique(podImages) {
		fmt.Printf("Namespace:%s\tImage:%s\n", img.Namespace, img.Image)
	}
	return nil
}

func (c *Config) fetchImagesWithClient() ([]image, error) {
	if c.All {
		namespace = ""
	} else {
		namespace = c.NameSpace
	}

	pods, err := c.client.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	podImages := make([]image, 0)
	for _, podnames := range pods.Items {
		initImages := getImagesFromContainers(podnames.Namespace, podnames.Spec.InitContainers)
		containerImages := getImagesFromContainers(podnames.Namespace, podnames.Spec.Containers)
		podImages = append(podImages, initImages...)
		podImages = append(podImages, containerImages...)
	}

	return podImages, nil
}

func getImagesFromContainers(namespace string, containers []corev1.Container) []image {
	images := make([]image, 0)
	for _, container := range containers {
		images = append(images, image{Namespace: namespace, Image: container.Image})
	}
	return images
}

func unique(s []image) []image {
	m := make(map[image]bool)
	for _, item := range s {
		if _, ok := m[item]; !ok {
			m[item] = true
		}
	}
	var result []image
	for key := range m {
		result = append(result, key)
	}
	return result
}

func (c *Config) getKubeConfig() (*rest.Config, error) {
	config, err := clientcmd.BuildConfigFromFlags("", c.KubeConfig)
	if err != nil {
		return nil, err
	}
	return config, nil
}
