package main

// import (
// 	"context"
// 	"flag"
// 	"fmt"
// 	"path/filepath"

// 	v1 "k8s.io/api/core/v1"
// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// 	"k8s.io/client-go/kubernetes"
// 	"k8s.io/client-go/tools/clientcmd"
// 	"k8s.io/client-go/util/homedir"
// )

// func BuildImage(name string, version string) {
// 	// Create a Kubernetes config object from the local kubeconfig file

// 	var kubeconfig string = "/Users/michael/.kube/config"
// 	home := homedir.HomeDir()

// 	if home != "" {
// 		kubeconfig = *flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
// 	} else {
// 		kubeconfig = *flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
// 	}

// 	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Create a Kubernetes client
// 	clientset, err := kubernetes.NewForConfig(config)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Define a pod specification
// 	podSpec := &v1.Pod{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name: "kaniko-" + name,
// 		},
// 		Spec: v1.PodSpec{
// 			Containers: []v1.Container{
// 				{
// 					Name:  "kaniko",
// 					Image: "gcr.io/kaniko-project/executor:debug",
// 					Args: []string{
// 						"--dockerfile=Dockerfile",
// 						"--context=git://github.com/michaelhtm/aws-docker.git",
// 						"--destination=399481058530.dkr.ecr.us-west-2.amazonaws.com/prow:" + version,
// 					},
// 					VolumeMounts: []v1.VolumeMount{
// 						{
// 							Name:      "kaniko-secret",
// 							MountPath: "/kaniko/.docker",
// 						},
// 					},
// 				},
// 			},
// 			Volumes: []v1.Volume{
// 				{
// 					Name: "kaniko-secret",
// 					VolumeSource: v1.VolumeSource{
// 						Secret: &v1.SecretVolumeSource{
// 							SecretName: "regcred",
// 							Items: []v1.KeyToPath{
// 								{
// 									Key:  ".dockerconfigjson",
// 									Path: "config.json",
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 			RestartPolicy: v1.RestartPolicyNever,
// 		},
// 	}

// 	// Create the pod
// 	pod, err := clientset.CoreV1().Pods("default").Create(context.Background(), podSpec, metav1.CreateOptions{})
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("Pod '%s' created\n", pod.Name)
// }
