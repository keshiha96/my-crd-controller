package apim1

import (
	//appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	//"k8s.io/apimachinery/pkg/api/resource"
	apimv1alpha1 "github.com/wso2-incubator/wso2am-k8s-operator/pkg/apis/apim/v1alpha1"
)

// newService creates a new Service for a Apimanager resource.
// It expose the service with Nodeport type with minikube ip as the externel ip.
func Apim1Service(apimanager *apimv1alpha1.APIManager) *corev1.Service {
	labels := map[string]string{
		"deployment": "wso2am-pattern-1-am",
		"node": "wso2am-pattern-1-am-1",
	}
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			// Name: apimanager.Spec.ServiceName,
			Name:      "apim-1-svc",
			Namespace: apimanager.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(apimanager, apimv1alpha1.SchemeGroupVersion.WithKind("APIManager")),
			},
		},
		Spec: corev1.ServiceSpec{
			Selector: labels,
			Type:     "NodePort",
			ExternalIPs: []string{"192.168.99.101"},
			Ports: []corev1.ServicePort{
				{
					Name:       "pass-through-http",
					Protocol:   corev1.ProtocolTCP,
					Port:       8280,
					TargetPort: intstr.IntOrString{Type: intstr.Int, IntVal: 8280},
					NodePort:   32004,
					//NodePort:9611,
				},
				{
					Name:       "pass-through-https",
					Protocol:   corev1.ProtocolTCP,
					Port:       8243,
					TargetPort: intstr.IntOrString{Type: intstr.Int, IntVal: 8243},
					NodePort:   32003,
					//NodePort:   9711,

				},
				{
					Name:       "servlet-http",
					Protocol:   corev1.ProtocolTCP,
					Port:       9763,
					TargetPort: intstr.IntOrString{Type: intstr.Int, IntVal: 9763},
					NodePort:   32002,
					//NodePort:   5672,

				},
				{
					Name:       "servlet-https",
					Protocol:   corev1.ProtocolTCP,
					Port:       9443,
					TargetPort: intstr.IntOrString{Type: intstr.Int, IntVal: 9443},
					NodePort:   32001,
					//NodePort:   9443,

				},
				{
					Name:       "jms-provider",
					Protocol:   corev1.ProtocolTCP,
					Port:       5672,
					TargetPort: intstr.IntOrString{Type: intstr.Int, IntVal: 5672},
					NodePort:   32005,
					//NodePort:   5672,

				},
				//{
				//	Name:       "binary",
				//	Protocol:   corev1.ProtocolTCP,
				//	Port:       9611,
				//	TargetPort: intstr.IntOrString{Type: intstr.Int, IntVal: 9611},
				//	NodePort:   32002,
				//},
				//{
				//	Name:       "binary-secure",
				//	Protocol:   corev1.ProtocolTCP,
				//	Port:       9711,
				//	TargetPort: intstr.IntOrString{Type: intstr.Int, IntVal: 9711},
				//	NodePort:   32003,
				//},
				//{
				//	Name:       "jms-tcp",
				//	Protocol:   corev1.ProtocolTCP,
				//	Port:       5672,
				//	TargetPort: intstr.IntOrString{Type: intstr.Int, IntVal: 5672},
				//	NodePort:   32004,
				//},
				//{
				//	Name:       "servlet-https",
				//	Protocol:   corev1.ProtocolTCP,
				//	Port:       9443,
				//	TargetPort: intstr.IntOrString{Type: intstr.Int, IntVal: 9443},
				//	NodePort:   32001,
				//},
			},
		},
	}
}