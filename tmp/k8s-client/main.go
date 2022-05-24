/*
Copyright 2022.

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
	"context"
	"math/rand"
	"os"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/clientcmd"
	v1 "my.domain/actiNode-api/api/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	k8s, err := clientcmd.BuildConfigFromFlags("", os.Getenv("HOME")+"/.kube/config")
	if err != nil {
		panic(err)
	}

	name := "actinode1"
	mapping := "test mapping sample"
	pods_desc := "PodA1, PodC2, PodC4"

	act := &v1.ActiNode{
		Spec: v1.ActiNodeSpec{
			Cpus:    rand.Int(),
			Pods:    pods_desc,
			Mapping: mapping,
		},
	}

	act.SetNamespace(os.Getenv("KUBE_NAMESPACE"))
	act.SetName(name)

	s := runtime.NewScheme()
	v1.AddToScheme(s)
	k8sClient, err := client.New(k8s, client.Options{Scheme: s})
	if err != nil {
		panic(err)
	}

	if err := k8sClient.Create(context.Background(), act); err != nil {
		panic(err)
	}
}
