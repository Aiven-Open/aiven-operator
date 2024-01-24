// Copyright (c) 2022 Aiven, Helsinki, Finland. https://aiven.io/

package v1alpha1

import (
	"context"
	"errors"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var cassandralog = logf.Log.WithName("cassandra-resource")

func (in *Cassandra) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(in).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-aiven-io-v1alpha1-cassandra,mutating=true,failurePolicy=fail,sideEffects=None,groups=aiven.io,resources=cassandras,verbs=create;update,versions=v1alpha1,name=mcassandra.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Cassandra{}

func (in *Cassandra) Default() {
	cassandralog.Info("default", "name", in.Name)
}

//+kubebuilder:webhook:verbs=create;update;delete,path=/validate-aiven-io-v1alpha1-cassandra,mutating=false,failurePolicy=fail,groups=aiven.io,resources=cassandras,versions=v1alpha1,name=vcassandra.kb.io,sideEffects=none,admissionReviewVersions=v1

var _ webhook.CustomValidator = &Cassandra{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type
func (in *Cassandra) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	cassandralog.Info("validate create", "name", in.Name)

	return nil, in.Spec.Validate()
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type
func (in *Cassandra) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	cassandralog.Info("validate update", "name", in.Name)

	if in.Spec.Project != oldObj.(*Cassandra).Spec.Project {
		return nil, errors.New("cannot update a Cassandra service, project field is immutable and cannot be updated")
	}

	if in.Spec.ConnInfoSecretTarget.Name != oldObj.(*Cassandra).Spec.ConnInfoSecretTarget.Name {
		return nil, errors.New("cannot update a Cassandra service, connInfoSecretTarget.name field is immutable and cannot be updated")
	}

	return nil, in.Spec.Validate()
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type
func (in *Cassandra) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	cassandralog.Info("validate delete", "name", in.Name)

	if in.Spec.TerminationProtection != nil && *in.Spec.TerminationProtection {
		return nil, errors.New("cannot delete Cassandra service, termination protection is on")
	}

	if in.Spec.ProjectVPCID != "" && in.Spec.ProjectVPCRef != nil {
		return nil, errors.New("cannot use both projectVpcId and projectVPCRef")
	}

	return nil, nil
}
