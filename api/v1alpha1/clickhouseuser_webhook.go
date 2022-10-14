// Copyright (c) 2022 Aiven, Helsinki, Finland. https://aiven.io/

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var clickhouseuserlog = logf.Log.WithName("clickhouseuser-resource")

func (r *ClickhouseUser) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-aiven-io-v1alpha1-clickhouseuser,mutating=true,failurePolicy=fail,sideEffects=None,groups=aiven.io,resources=clickhouseusers,verbs=create;update,versions=v1alpha1,name=mclickhouseuser.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &ClickhouseUser{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *ClickhouseUser) Default() {
	clickhouseuserlog.Info("default", "name", r.Name)
}

//+kubebuilder:webhook:path=/validate-aiven-io-v1alpha1-clickhouseuser,mutating=false,failurePolicy=fail,sideEffects=None,groups=aiven.io,resources=clickhouseusers,verbs=create;update,versions=v1alpha1,name=vclickhouseuser.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &ClickhouseUser{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *ClickhouseUser) ValidateCreate() error {
	clickhouseuserlog.Info("validate create", "name", r.Name)
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *ClickhouseUser) ValidateUpdate(old runtime.Object) error {
	clickhouseuserlog.Info("validate update", "name", r.Name)

	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *ClickhouseUser) ValidateDelete() error {
	clickhouseuserlog.Info("validate delete", "name", r.Name)

	return nil
}
