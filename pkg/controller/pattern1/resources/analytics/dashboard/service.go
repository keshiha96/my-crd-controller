package dashboard

import (
	//appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	apimv1alpha1 "github.com/wso2-incubator/wso2am-k8s-operator/pkg/apis/apim/v1alpha1"
)

// newService creates a new Service for a Apimanager resource.
// It expose the service with Nodeport type with minikube ip as the externel ip.
func DashboardService(apimanager *apimv1alpha1.APIManager) *corev1.Service {
	labels := map[string]string{
		"deployment": "wso2am-pattern-1-analytics-dashboard",
	}
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			// Name: apimanager.Spec.ServiceName,
			Name:      "analytics-dash-svc",
			Namespace: apimanager.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(apimanager, apimv1alpha1.SchemeGroupVersion.WithKind("Apimanager")),
			},
		},
		Spec: corev1.ServiceSpec{
			Selector: labels,
			Type:     "NodePort",
			// values are fetched from wso2-apim.yaml file
			// Type: apimanager.Spec.ServType,
			ExternalIPs: []string{"192.168.99.101"},
			// ExternalIPs: apimanager.Spec.ExternalIps,
			Ports: []corev1.ServicePort{
				//{
				//	Name:       "analytics-dashboard",
				//	Protocol:   corev1.ProtocolTCP,
				//	Port:       9643,
				//	TargetPort: intstr.IntOrString{Type: intstr.Int, IntVal: 9643},
				//	NodePort:   32009,
				//},
				{
					Name:       "analytics-dashboard",
					Protocol:   corev1.ProtocolTCP,
					Port:       9643,
					TargetPort: intstr.IntOrString{Type: intstr.Int, IntVal: 9643},
					NodePort:   32201,
				},
			},
		},
	}
}














//package dashboard
//
//import (
//	//appsv1 "k8s.io/api/apps/v1"
//	corev1 "k8s.io/api/core/v1"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"k8s.io/apimachinery/pkg/util/intstr"
//
//	apimv1alpha1 "github.com/wso2-incubator/wso2am-k8s-operator/pkg/apis/apim/v1alpha1"
//)
//
//// newService creates a new Service for a Apimanager resource.
//// It expose the service with Nodeport type with minikube ip as the externel ip.
//func DashboardService(apimanager *apimv1alpha1.APIManager) *corev1.Service {
//	labels := map[string]string{
//		"deployment": "wso2am-pattern-1-analytics-dashboard",
//	}
//	return &corev1.Service{
//		ObjectMeta: metav1.ObjectMeta{
//			// Name: apimanager.Spec.ServiceName,
//			Name:      "analytics-dash-svc",
//			Namespace: apimanager.Namespace,
//			OwnerReferences: []metav1.OwnerReference{
//				*metav1.NewControllerRef(apimanager, apimv1alpha1.SchemeGroupVersion.WithKind("Apimanager")),
//			},
//		},
//		Spec: corev1.ServiceSpec{
//			Selector: labels,
//			Type:     "NodePort",
//			// values are fetched from wso2-apim.yaml file
//			// Type: apimanager.Spec.ServType,
//			ExternalIPs: []string{"192.168.99.101"},
//			// ExternalIPs: apimanager.Spec.ExternalIps,
//			Ports: []corev1.ServicePort{
//				//{
//				//	Name:       "analytics-dashboard",
//				//	Protocol:   corev1.ProtocolTCP,
//				//	Port:       9643,
//				//	TargetPort: intstr.IntOrString{Type: intstr.Int, IntVal: 9643},
//				//	NodePort:   32009,
//				//},
//				{
//					Name:       "analytics-dashboard",
//					Protocol:   corev1.ProtocolTCP,
//					Port:       9643,
//					TargetPort: intstr.IntOrString{Type: intstr.Int, IntVal: 9643},
//					NodePort:   32201,
//				},
//			},
//		},
//	}
//}
//
//
//
