package operator

import (
	"fmt"

	"github.com/3scale/3scale-operator/pkg/3scale/amp/component"
	"github.com/3scale/3scale-operator/pkg/3scale/amp/product"
)

func (o *OperatorMemcachedOptionsProvider) GetMemcachedOptions() (*component.MemcachedOptions, error) {

	productVersion := o.APIManagerSpec.ProductVersion
	imageProvider, err := product.NewImageProvider(productVersion)
	if err != nil {
		return nil, err
	}

	optProv := component.MemcachedOptionsBuilder{}
	optProv.AppLabel(*o.APIManagerSpec.AppLabel)

	systemSpec := o.APIManagerSpec.SystemSpec
	if systemSpec != nil && systemSpec.MemcachedImage != nil {
		optProv.Image(*systemSpec.MemcachedImage)
	} else {
		optProv.Image(imageProvider.GetSystemMemcachedImage())
	}

	res, err := optProv.Build()
	if err != nil {
		return nil, fmt.Errorf("unable to create Memcached Options - %s", err)
	}
	return res, nil
}
