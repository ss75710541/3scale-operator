package operator

import (
	"fmt"

	"github.com/3scale/3scale-operator/pkg/3scale/amp/component"
	"github.com/3scale/3scale-operator/pkg/3scale/amp/product"
)

func (o *OperatorProductizedOptionsProvider) GetProductizedOptions() (*component.ProductizedOptions, error) {
	pob := component.ProductizedOptionsBuilder{}

	productVersion := o.APIManagerSpec.ProductVersion
	imageProvider, err := product.NewImageProvider(productVersion)
	if err != nil {
		return nil, err
	}

	pob.AmpRelease(string(productVersion))

	apicastSpec := o.APIManagerSpec.ApicastSpec
	backendSpec := o.APIManagerSpec.BackendSpec
	wildcardRouterSpec := o.APIManagerSpec.WildcardRouterSpec
	systemSpec := o.APIManagerSpec.SystemSpec
	zyncSpec := o.APIManagerSpec.ZyncSpec

	if apicastSpec != nil && apicastSpec.Image != nil {
		pob.ApicastImage(*apicastSpec.Image)
	} else {
		pob.ApicastImage(imageProvider.GetApicastImage())
	}

	if backendSpec != nil && backendSpec.Image != nil {
		pob.BackendImage(*backendSpec.Image)
	} else {
		pob.BackendImage(imageProvider.GetBackendImage())
	}

	if wildcardRouterSpec != nil && wildcardRouterSpec.Image != nil {
		pob.RouterImage(*wildcardRouterSpec.Image)
	} else {
		pob.RouterImage(imageProvider.GetWildcardRouterImage())
	}

	if systemSpec != nil && systemSpec.Image != nil {
		pob.SystemImage(*systemSpec.Image)
	} else {
		pob.SystemImage(imageProvider.GetSystemImage())
	}

	if zyncSpec != nil && zyncSpec.Image != nil {
		pob.ZyncImage(*zyncSpec.Image)
	} else {
		pob.ZyncImage(imageProvider.GetZyncImage())
	}

	res, err := pob.Build()
	if err != nil {
		return nil, fmt.Errorf("unable to create Productized Options - %s", err)
	}
	return res, nil
}
