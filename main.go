/*
Copyright 2019 The xridge kubestone contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"

	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func main() {
	config, err := rest.InClusterConfig()
	if err == rest.ErrNotInCluster {
		kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	checkError(err)

	clientset, err := kubernetes.NewForConfig(config)
	checkError(err)
	metricsClientset, err := metricsv.NewForConfig(config)
	checkError(err)

	nodesInfo, err := getNodeInfos(clientset, metricsClientset)
	checkError(err)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(nodesInfo)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}