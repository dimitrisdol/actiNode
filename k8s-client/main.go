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

	name1 := "actinode-termi7"
	cpus1 := 4
	mapping1 := "test mapping sample"
	pods_desc1 := "PodA1, PodC2, PodC4"

	act1 := &v1.ActiNode{
		Spec: v1.ActiNodeSpec{
			Cpus:    cpus1,
			Pods:    pods_desc1,
			Mapping: mapping1,
		},
	}

	act1.SetNamespace(os.Getenv("KUBE_NAMESPACE"))
	act1.SetName(name1)

	s1 := runtime.NewScheme()
	v1.AddToScheme(s1)
	k8sClient, err := client.New(k8s, client.Options{Scheme: s1})
	if err != nil {
		panic(err)
	}

	if err := k8sClient.Create(context.Background(), act1); err != nil {
		panic(err)
	}
	
	name2 := "actinode-termi8"
	cpus2 := 4
	mapping2 := "test mapping sample"
	pods_desc2 := "PodA2, PodB2, PodD4"

	act2 := &v1.ActiNode{
		Spec: v1.ActiNodeSpec{
			Cpus:    cpus2,
			Pods:    pods_desc2,
			Mapping: mapping2,
		},
	}

	act2.SetNamespace(os.Getenv("KUBE_NAMESPACE"))
	act2.SetName(name2)

	s2 := runtime.NewScheme()
	v1.AddToScheme(s2)
	k8sClient2, err2 := client.New(k8s, client.Options{Scheme: s2})
	if err2 != nil {
		panic(err2)
	}

	if err2 := k8sClient2.Create(context.Background(), act2); err2 != nil {
		panic(err2)
	}
	
	name3 := "actinode-termi9"
	cpus3 := 4
	mapping3 := "test mapping sample"
	pods_desc3 := "PodB4, PodA3, PodC1"

	act3 := &v1.ActiNode{
		Spec: v1.ActiNodeSpec{
			Cpus:    cpus3,
			Pods:    pods_desc3,
			Mapping: mapping3,
		},
	}

	act3.SetNamespace(os.Getenv("KUBE_NAMESPACE"))
	act3.SetName(name3)

	s3 := runtime.NewScheme()
	v1.AddToScheme(s3)
	k8sClient3, err3 := client.New(k8s, client.Options{Scheme: s3})
	if err3 != nil {
		panic(err3)
	}

	if err3 := k8sClient3.Create(context.Background(), act3); err3 != nil {
		panic(err3)
	}
	
	name4 := "actinode-termi10"
	cpus4 := 4
	mapping4 := "test mapping sample"
	pods_desc4 := "PodD1, PodD2, PodB4"

	act4 := &v1.ActiNode{
		Spec: v1.ActiNodeSpec{
			Cpus:    cpus4,
			Pods:    pods_desc4,
			Mapping: mapping4,
		},
	}

	act4.SetNamespace(os.Getenv("KUBE_NAMESPACE"))
	act4.SetName(name4)

	s4 := runtime.NewScheme()
	v1.AddToScheme(s4)
	k8sClient4, err4 := client.New(k8s, client.Options{Scheme: s4})
	if err4 != nil {
		panic(err4)
	}

	if err4 := k8sClient4.Create(context.Background(), act4); err4 != nil {
		panic(err4)
	}
}
