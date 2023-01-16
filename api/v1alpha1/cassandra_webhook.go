// Copyright (c) 2022 Aiven, Helsinki, Finland. https://aiven.io/

package v1alpha1

import (
	"errors"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
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

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (in *Cassandra) ValidateCreate() error {
	cassandralog.Info("validate create", "name", in.Name)

	return in.Spec.Validate()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (in *Cassandra) ValidateUpdate(old runtime.Object) error {
	cassandralog.Info("validate update", "name", in.Name)

	if in.Spec.Project != old.(*Cassandra).Spec.Project {
		return errors.New("cannot update a Cassandra service, project field is immutable and cannot be updated")
	}

	if in.Spec.ConnInfoSecretTarget.Name != old.(*Cassandra).Spec.ConnInfoSecretTarget.Name {
		return errors.New("cannot update a Cassandra service, connInfoSecretTarget.name field is immutable and cannot be updated")
	}

	return in.Spec.Validate()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (in *Cassandra) ValidateDelete() error {
	cassandralog.Info("validate delete", "name", in.Name)

	if in.Spec.TerminationProtection {
		return errors.New("cannot delete Cassandra service, termination protection is on")
	}

	return nil
}
