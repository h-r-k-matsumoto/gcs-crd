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
	"fmt"

	storagev1alpha1 "github.com/h-r-k-matsumoto/gcs-crd/pkg/apis/storage/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
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
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	myFinalizerName := "bucket.finalizer.gcs.storage.matsumo.dev"

	if !instance.ObjectMeta.DeletionTimestamp.IsZero() {
		log.Info("delete object.")
		// The object is being deleted
		if containsString(instance.ObjectMeta.Finalizers, myFinalizerName) {
			// our finalizer is present, so lets handle our external dependency
			if err := r.deleteExternalDependency(instance); err != nil {
				return reconcile.Result{}, err
			}

			// remove our finalizer from the list and update it.
			instance.ObjectMeta.Finalizers = removeString(instance.ObjectMeta.Finalizers, myFinalizerName)
			if err := r.Update(context.Background(), instance); err != nil {
				return reconcile.Result{}, err
			}
		}
		// if deleted, logic finished.
		return reconcile.Result{}, nil
	}

	// update finalizer,and bucket name.
	modified := false
	if !containsString(instance.ObjectMeta.Finalizers, myFinalizerName) {
		instance.ObjectMeta.Finalizers = append(instance.ObjectMeta.Finalizers, myFinalizerName)
		modified = true
	}
	if instance.Spec.BucketName != instance.Status.BucketName {
		instance.Status.BucketName = instance.Spec.BucketName
		instance.Status.BucketFullName = GenerateBacketFullName(instance.Spec.BucketName)
		modified = true
	}

	client, err := NewGcsClient(context.Background())
	if err != nil {
		return reconcile.Result{}, err
	}

	//bucket operation.
	exists := IfExistsBucket(context.Background(), client, instance.Spec.ProjectID, instance.Status.BucketFullName)
	if !exists {
		log.Info(fmt.Sprintf("create bucket[ %s ]", instance.Status.BucketFullName))
		err := CreateBucket(context.Background(), client, instance.Spec.ProjectID, instance.Status.BucketFullName)
		if err != nil {
			fatal := DeleteBucket(context.Background(), client, instance.Status.BucketFullName)
			if fatal != nil {
				log.Error(fatal, "bucket cleanup error.")
			}
			return reconcile.Result{}, err
		}
	}

	if modified {
		if err := r.Update(context.Background(), instance); err != nil {
			return reconcile.Result{}, err
		}
	}

	return reconcile.Result{}, nil
}

//  delete dependency bucket.
func (r *ReconcileGcs) deleteExternalDependency(instance *storagev1alpha1.Gcs) error {
	log.Info(fmt.Sprintf("delete bucket[ %s ]", instance.Status.BucketFullName))
	client, err := NewGcsClient(context.Background())
	if err != nil {
		return err
	}

	err = DeleteBucket(context.Background(), client, instance.Status.BucketFullName)
	return err
}

// Helper functions to check string from a slice of strings.
func containsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

// Helper functions to remove string from slice.
func removeString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}
