/*
Copyright 2024 The Kruise Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	versioned "github.com/tiancheng92/kruise-rollout-api/client/clientset/versioned"
	internalinterfaces "github.com/tiancheng92/kruise-rollout-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/tiancheng92/kruise-rollout-api/client/listers/rollouts/v1alpha1"
	rolloutsv1alpha1 "github.com/tiancheng92/kruise-rollout-api/rollouts/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TrafficRoutingInformer provides access to a shared informer and lister for
// TrafficRoutings.
type TrafficRoutingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.TrafficRoutingLister
}

type trafficRoutingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTrafficRoutingInformer constructs a new informer for TrafficRouting type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTrafficRoutingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTrafficRoutingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTrafficRoutingInformer constructs a new informer for TrafficRouting type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTrafficRoutingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RolloutsV1alpha1().TrafficRoutings(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RolloutsV1alpha1().TrafficRoutings(namespace).Watch(context.TODO(), options)
			},
		},
		&rolloutsv1alpha1.TrafficRouting{},
		resyncPeriod,
		indexers,
	)
}

func (f *trafficRoutingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTrafficRoutingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *trafficRoutingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&rolloutsv1alpha1.TrafficRouting{}, f.defaultInformer)
}

func (f *trafficRoutingInformer) Lister() v1alpha1.TrafficRoutingLister {
	return v1alpha1.NewTrafficRoutingLister(f.Informer().GetIndexer())
}
