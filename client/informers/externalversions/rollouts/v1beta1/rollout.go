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

package v1beta1

import (
	"context"
	time "time"

	versioned "github.com/tiancheng92/kruise-rollout-api/client/clientset/versioned"
	internalinterfaces "github.com/tiancheng92/kruise-rollout-api/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/tiancheng92/kruise-rollout-api/client/listers/rollouts/v1beta1"
	rolloutsv1beta1 "github.com/tiancheng92/kruise-rollout-api/rollouts/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// RolloutInformer provides access to a shared informer and lister for
// Rollouts.
type RolloutInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.RolloutLister
}

type rolloutInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewRolloutInformer constructs a new informer for Rollout type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRolloutInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRolloutInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredRolloutInformer constructs a new informer for Rollout type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRolloutInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RolloutsV1beta1().Rollouts(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RolloutsV1beta1().Rollouts(namespace).Watch(context.TODO(), options)
			},
		},
		&rolloutsv1beta1.Rollout{},
		resyncPeriod,
		indexers,
	)
}

func (f *rolloutInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRolloutInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *rolloutInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&rolloutsv1beta1.Rollout{}, f.defaultInformer)
}

func (f *rolloutInformer) Lister() v1beta1.RolloutLister {
	return v1beta1.NewRolloutLister(f.Informer().GetIndexer())
}
