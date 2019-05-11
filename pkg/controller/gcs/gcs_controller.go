/*
Copyright 2019 Hiroki Matsumoto.

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

package gcs

import (
	"context"

	storagev1alpha1 "github.com/h-r-k-matsumoto/gcs-crd/pkg/apis/storage/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new Gcs Controller and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileGcs{Client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("gcs-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to Gcs
	err = c.Watch(&source.Kind{Type: &storagev1alpha1.Gcs{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

var _ reconcile.Reconciler = &ReconcileGcs{}

const secretKeyName = "secret.json"

// ReconcileGcs reconciles a Gcs object
type ReconcileGcs struct {
	client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a Gcs object and makes changes based on the state read
// and what is in the Gcs.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  The scaffolding writes
// a Deployment as an example
// Automatically generate RBAC rules to allow the Controller to read and write Deployments
// +kubebuilder:rbac:groups=apps,resources=secret,verbs=get
// +kubebuilder:rbac:groups=storage.matsumo.dev,resources=gcs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=storage.matsumo.dev,resources=gcs/status,verbs=get;update;patch
func (r *ReconcileGcs) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	var err error

	// Fetch the Gcs instance
	instance := &storagev1alpha1.Gcs{}

	err = r.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.

			//TODO: DELETE GCS Bucket?
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	if instance.Spec.BucketName != instance.Status.BucketName {
		instance.Status.BucketName = instance.Spec.BucketName
		instance.Status.BucketFullName = GenerateBacketFullName(instance.Spec.BucketName)

		log.Info("*********** Status Updating...")
		err = r.Update(context.Background(), instance)
		if err != nil {
			return reconcile.Result{}, err
		}
		log.Info("*********** Status Updated")
	}

	// initialize gcp service account credential.
	secret := &corev1.Secret{}
	err = r.Get(context.TODO(), types.NamespacedName{Name: instance.Spec.SecretName, Namespace: instance.Namespace}, secret)
	if err != nil {
		return reconcile.Result{}, err
	}
	credential := secret.Data[secretKeyName]
	ctx := context.Background()
	client, err := NewGcsClient(ctx, credential)
	if err != nil {
		return reconcile.Result{}, err
	}

	//bucket operation.
	exists := IfExistsBucket(ctx, client, instance.Spec.ProjectID, instance.Status.BucketFullName)
	if !exists {
		err := CreateBucket(ctx, client, instance.Spec.ProjectID, instance.Status.BucketFullName)
		if err != nil {
			return reconcile.Result{}, err
		}
	}
	return reconcile.Result{}, nil
}
