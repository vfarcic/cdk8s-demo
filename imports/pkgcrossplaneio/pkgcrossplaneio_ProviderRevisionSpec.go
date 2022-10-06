// pkgcrossplaneio
package pkgcrossplaneio


// PackageRevisionSpec specifies the desired state of a PackageRevision.
type ProviderRevisionSpec struct {
	// DesiredState of the PackageRevision.
	//
	// Can be either Active or Inactive.
	DesiredState *string `field:"required" json:"desiredState" yaml:"desiredState"`
	// Package image used by install Pod to extract package contents.
	Image *string `field:"required" json:"image" yaml:"image"`
	// Revision number.
	//
	// Indicates when the revision will be garbage collected based on the parent's RevisionHistoryLimit.
	Revision *float64 `field:"required" json:"revision" yaml:"revision"`
	// ControllerConfigRef references a ControllerConfig resource that will be used to configure the packaged controller Deployment.
	ControllerConfigRef *ProviderRevisionSpecControllerConfigRef `field:"optional" json:"controllerConfigRef" yaml:"controllerConfigRef"`
	// IgnoreCrossplaneConstraints indicates to the package manager whether to honor Crossplane version constrains specified by the package.
	//
	// Default is false.
	IgnoreCrossplaneConstraints *bool `field:"optional" json:"ignoreCrossplaneConstraints" yaml:"ignoreCrossplaneConstraints"`
	// PackagePullPolicy defines the pull policy for the package.
	//
	// It is also applied to any images pulled for the package, such as a provider's controller image. Default is IfNotPresent.
	PackagePullPolicy *string `field:"optional" json:"packagePullPolicy" yaml:"packagePullPolicy"`
	// PackagePullSecrets are named secrets in the same namespace that can be used to fetch packages from private registries.
	//
	// They are also applied to any images pulled for the package, such as a provider's controller image.
	PackagePullSecrets *[]*ProviderRevisionSpecPackagePullSecrets `field:"optional" json:"packagePullSecrets" yaml:"packagePullSecrets"`
	// SkipDependencyResolution indicates to the package manager whether to skip resolving dependencies for a package.
	//
	// Setting this value to true may have unintended consequences. Default is false.
	SkipDependencyResolution *bool `field:"optional" json:"skipDependencyResolution" yaml:"skipDependencyResolution"`
	// WebhookTLSSecretName is the name of the TLS Secret that will be used by the provider to serve a TLS-enabled webhook server.
	//
	// The certificate will be injected to webhook configurations as well as CRD conversion webhook strategy if needed. If it's not given, provider will not have a certificate mounted to its filesystem, webhook configurations won't be deployed and if there is a CRD with webhook conversion strategy, the installation will fail.
	WebhookTlsSecretName *string `field:"optional" json:"webhookTlsSecretName" yaml:"webhookTlsSecretName"`
}

