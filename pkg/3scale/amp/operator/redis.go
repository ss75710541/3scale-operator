package operator

import (
	"fmt"

	"github.com/3scale/3scale-operator/pkg/3scale/amp/component"
	"github.com/3scale/3scale-operator/pkg/3scale/amp/product"
)

func (o *OperatorRedisOptionsProvider) GetRedisOptions() (*component.RedisOptions, error) {
	optProv := component.RedisOptionsBuilder{}
	productVersion := o.APIManagerSpec.ProductVersion
	imageProvider, err := product.NewImageProvider(productVersion)
	if err != nil {
		return nil, err
	}

	optProv.AppLabel(*o.APIManagerSpec.AppLabel)

	backendSpec := o.APIManagerSpec.BackendSpec
	systemSpec := o.APIManagerSpec.SystemSpec
	if backendSpec != nil && backendSpec.RedisImage != nil {
		optProv.BackendImage(*backendSpec.RedisImage)
	} else {
		optProv.BackendImage(imageProvider.GetBackendRedisImage())
	}
	if systemSpec != nil && systemSpec.RedisImage != nil {
		optProv.SystemImage(*systemSpec.RedisImage)
	} else {
		optProv.SystemImage(imageProvider.GetSystemRedisImage())
	}

	res, err := optProv.Build()
	if err != nil {
		return nil, fmt.Errorf("unable to create Redis Options - %s", err)
	}
	return res, nil
}
