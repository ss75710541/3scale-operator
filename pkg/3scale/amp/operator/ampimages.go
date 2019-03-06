package operator

import (
	"fmt"

	"github.com/3scale/3scale-operator/pkg/3scale/amp/product"

	"github.com/3scale/3scale-operator/pkg/3scale/amp/component"
)

func (o *OperatorAmpImagesOptionsProvider) GetAmpImagesOptions() (*component.AmpImagesOptions, error) {
	optProv := component.AmpImagesOptionsBuilder{}

	productVersion := o.APIManagerSpec.ProductVersion
	imageProvider, err := product.NewImageProvider(productVersion)
	if err != nil {
		return nil, err
	}

	optProv.AppLabel(*o.APIManagerSpec.AppLabel)
	optProv.AMPRelease(string(productVersion))

	apicastSpec := o.APIManagerSpec.ApicastSpec
	backendSpec := o.APIManagerSpec.BackendSpec
	wildcardRouterSpec := o.APIManagerSpec.WildcardRouterSpec
	systemSpec := o.APIManagerSpec.SystemSpec
	zyncSpec := o.APIManagerSpec.ZyncSpec
	if apicastSpec != nil && apicastSpec.Image != nil {
		optProv.ApicastImage(*apicastSpec.Image)
	} else {
		optProv.ApicastImage(imageProvider.GetApicastImage())
	}

	if backendSpec != nil && backendSpec.Image != nil {
		optProv.BackendImage(*backendSpec.Image)
	} else {
		optProv.BackendImage(imageProvider.GetBackendImage())
	}

	if wildcardRouterSpec != nil && wildcardRouterSpec.Image != nil {
		optProv.RouterImage(*wildcardRouterSpec.Image)
	} else {
		optProv.RouterImage(imageProvider.GetWildcardRouterImage())
	}

	if systemSpec != nil && systemSpec.Image != nil {
		optProv.SystemImage(*systemSpec.Image)
	} else {
		optProv.SystemImage(imageProvider.GetSystemImage())
	}

	if zyncSpec != nil && zyncSpec.Image != nil {
		optProv.ZyncImage(*zyncSpec.Image)
	} else {
		optProv.ZyncImage(imageProvider.GetZyncImage())
	}

	if zyncSpec != nil && zyncSpec.PostgreSQLImage != nil {
		optProv.PostgreSQLImage(*zyncSpec.PostgreSQLImage)
	} else {
		optProv.PostgreSQLImage(imageProvider.GetZyncPostgreSQLImage())
	}

	optProv.InsecureImportPolicy(*o.APIManagerSpec.ImageStreamTagImportInsecure)
	res, err := optProv.Build()
	if err != nil {
		return nil, fmt.Errorf("unable to create AMPImages Options - %s", err)
	}
	return res, nil
}
