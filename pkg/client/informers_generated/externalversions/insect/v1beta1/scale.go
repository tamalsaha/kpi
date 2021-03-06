/*
Copyright 2018 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package v1beta1

import (
	insect_v1beta1 "github.com/tamalsaha/apiserver-builder-demo/pkg/apis/insect/v1beta1"
	clientset "github.com/tamalsaha/apiserver-builder-demo/pkg/client/clientset_generated/clientset"
	internalinterfaces "github.com/tamalsaha/apiserver-builder-demo/pkg/client/informers_generated/externalversions/internalinterfaces"
	v1beta1 "github.com/tamalsaha/apiserver-builder-demo/pkg/client/listers_generated/insect/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// ScaleInformer provides access to a shared informer and lister for
// Scales.
type ScaleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ScaleLister
}

type scaleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewScaleInformer constructs a new informer for Scale type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewScaleInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredScaleInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredScaleInformer constructs a new informer for Scale type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredScaleInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.InsectV1beta1().Scales(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.InsectV1beta1().Scales(namespace).Watch(options)
			},
		},
		&insect_v1beta1.Scale{},
		resyncPeriod,
		indexers,
	)
}

func (f *scaleInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredScaleInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *scaleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&insect_v1beta1.Scale{}, f.defaultInformer)
}

func (f *scaleInformer) Lister() v1beta1.ScaleLister {
	return v1beta1.NewScaleLister(f.Informer().GetIndexer())
}
