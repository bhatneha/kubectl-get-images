package images

import (
	"testing"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
)

func Test_unique(t *testing.T) {
	t.Run("Should return unique itmes of an array", func (t *testing.T)  {
		expected := []image{
			{Namespace: "jenkins",Image: "jenkins/jenkins:2.289.1-jdk11"},
			{Namespace: "kube-system",Image: "rancher/klipper-helm:v0.5.0-build20210505"},
			{Namespace: "jenkins", Image: "bhatneha/jenkins-jnlp-agent:latest"},
		}

		input := []image{
			{Namespace: "jenkins",Image: "jenkins/jenkins:2.289.1-jdk11"},
			{Namespace: "kube-system",Image: "rancher/klipper-helm:v0.5.0-build20210505"},
			{Namespace: "jenkins", Image: "bhatneha/jenkins-jnlp-agent:latest"},
			{Namespace: "jenkins",Image: "jenkins/jenkins:2.289.1-jdk11"},
		}
		actual := unique(input)
		assert.Equal(t,expected,actual)	
		assert.Equal(t,len(expected),len(actual))
	})
	t.Run("Should fail",func(t *testing.T) {
		input := []image{
			{Namespace: "jenkins",Image: "jenkins/jenkins:2.289.1-jdk11"},
			{Namespace: "kube-system",Image: "rancher/klipper-helm:v0.5.0-build20210505"},
			{Namespace: "jenkins", Image: "bhatneha/jenkins-jnlp-agent:latest"},
			{Namespace: "jenkins",Image: "jenkins/jenkins:2.289.1-jdk11"},
		}
		actual := unique(input)
		assert.NotEqual(t,len(input),len(actual))	
	})
}

func Test_getImagesFromContainers(t *testing.T) {
	t.Run("should return all the images the containers", func(t *testing.T) {
		expected := []image{
			{Namespace: "jenkins",Image: "jenkins/jenkins:2.289.1-jdk11"},
			{Namespace: "kube-system",Image: "rancher/klipper-helm:v0.5.0-build20210505"},
			{Namespace: "jenkins", Image: "bhatneha/jenkins-jnlp-agent:latest"},
			{Namespace: "jenkins",Image: "jenkins/jenkins:2.289.1-jdk11"},
			{Namespace: "kube-system",Image: "rancher/klipper-lb:v0.2.0"},
		}

		input1 := []corev1.Container{
			{Name: "JenkinsPod",Image: "jenkins/jenkins:2.289.1-jdk11"},
			{Name: "JenkinsPOD",Image: "bhatneha/jenkins-jnlp-agent:latest"},
			{Name: "JenkinsPO",Image: "jenkins/jenkins:2.289.1-jdk11"},
		}

		input2 := []corev1.Container{
			{Name: "KlipperPod",Image: "rancher/klipper-helm:v0.5.0-build20210505"},
			{Name: "KlipperPOD",Image: "rancher/klipper-lb:v0.2.0"},
		}
		actual1 := getImagesFromContainers("jenkins",input1)
		actual2 := getImagesFromContainers("kube-system",input2)
		actual1 = append(actual1, actual2...)
		assert.ElementsMatch(t,expected,actual1)
	})
}
