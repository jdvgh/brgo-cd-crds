/*
Copyright 2023.

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

package controller

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	brgov1alpha1 "github.com/jdvgh/brgo-cd-crds/api/v1alpha1"
)

// SyncTargetReconciler reconciles a SyncTarget object
type SyncTargetReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=brgo.jdvgh.com,resources=synctargets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=brgo.jdvgh.com,resources=synctargets/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=brgo.jdvgh.com,resources=synctargets/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the SyncTarget object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.0/pkg/reconcile
func (r *SyncTargetReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logCtx := log.FromContext(ctx)

	// TODO(user): your logic here

	var syncTargetInfo brgov1alpha1.SyncTarget

	if err := r.Get(ctx, req.NamespacedName, &syncTargetInfo); err != nil {
		if client.IgnoreNotFound(err) != nil {
			logCtx.Error(err, "unable to get SyncTarget: '%v' ", err)
		}
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	fmt.Printf("syncTarget: repo %s, path %s\n", syncTargetInfo.Spec.RepoUrl, syncTargetInfo.Spec.Path)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *SyncTargetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&brgov1alpha1.SyncTarget{}).
		Complete(r)
}
