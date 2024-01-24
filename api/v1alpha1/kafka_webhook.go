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
var kafkalog = logf.Log.WithName("kafka-resource")

func (in *Kafka) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(in).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-aiven-io-v1alpha1-kafka,mutating=true,failurePolicy=fail,groups=aiven.io,resources=kafkas,verbs=create;update,versions=v1alpha1,name=mkafka.kb.io,sideEffects=none,admissionReviewVersions=v1

var _ webhook.Defaulter = &Kafka{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (in *Kafka) Default() {
	kafkalog.Info("default", "name", in.Name)
}

//+kubebuilder:webhook:verbs=create;update;delete,path=/validate-aiven-io-v1alpha1-kafka,mutating=false,failurePolicy=fail,groups=aiven.io,resources=kafkas,versions=v1alpha1,name=vkafka.kb.io,sideEffects=none,admissionReviewVersions=v1

var _ webhook.CustomValidator = &Kafka{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type
func (in *Kafka) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	kafkalog.Info("validate create", "name", in.Name)

	return nil, in.Spec.Validate()
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type
func (in *Kafka) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	kafkalog.Info("validate update", "name", in.Name)

	if in.Spec.Project != oldObj.(*Kafka).Spec.Project {
		return nil, errors.New("cannot update a Kafka service, project field is immutable and cannot be updated")
	}

	if in.Spec.ConnInfoSecretTarget.Name != oldObj.(*Kafka).Spec.ConnInfoSecretTarget.Name {
		return nil, errors.New("cannot update a Kafka service, connInfoSecretTarget.name field is immutable and cannot be updated")
	}

	return nil, in.Spec.Validate()
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type
func (in *Kafka) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	kafkalog.Info("validate delete", "name", in.Name)

	if in.Spec.TerminationProtection != nil && *in.Spec.TerminationProtection {
		return nil, errors.New("cannot delete Kafka service, termination protection is on")
	}

	if in.Spec.ProjectVPCID != "" && in.Spec.ProjectVPCRef != nil {
		return nil, errors.New("cannot use both projectVpcId and projectVPCRef")
	}

	return nil, nil
}
