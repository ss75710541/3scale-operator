package operator

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/ss75710541/3scale-operator/pkg/3scale/amp/component"
	"github.com/ss75710541/3scale-operator/pkg/3scale/amp/product"
	oprand "github.com/ss75710541/3scale-operator/pkg/crypto/rand"
	"k8s.io/apimachinery/pkg/api/errors"
)

func (o *OperatorMysqlOptionsProvider) GetMysqlOptions() (*component.MysqlOptions, error) {
	optProv := component.MysqlOptionsBuilder{}
	optProv.AppLabel(*o.APIManagerSpec.AppLabel)
	productVersion := o.APIManagerSpec.ProductVersion
	imageProvider, err := product.NewImageProvider(productVersion)
	if err != nil {
		return nil, err
	}

	optProv.AppLabel(*o.APIManagerSpec.AppLabel)
	if o.APIManagerSpec.SystemSpec != nil && o.APIManagerSpec.SystemSpec.DatabaseSpec != nil &&
		o.APIManagerSpec.SystemSpec.DatabaseSpec.MySQLSpec != nil && o.APIManagerSpec.SystemSpec.DatabaseSpec.MySQLSpec.Image != nil {
		optProv.Image(*o.APIManagerSpec.SystemSpec.DatabaseSpec.MySQLSpec.Image)
	} else {
		optProv.Image(imageProvider.GetSystemMySQLImage())
	}

	err = o.setSecretBasedOptions(&optProv)
	if err != nil {
		return nil, err
	}

	res, err := optProv.Build()
	if err != nil {
		return nil, fmt.Errorf("unable to create Mysql Options - %s", err)
	}
	return res, nil
}

func (o *OperatorMysqlOptionsProvider) setSecretBasedOptions(builder *component.MysqlOptionsBuilder) error {
	err := o.setSystemDatabaseOptions(builder)
	if err != nil {
		return fmt.Errorf("unable to create System Database secret options - %s", err)
	}

	return nil
}

func (o *OperatorMysqlOptionsProvider) setSystemDatabaseOptions(builder *component.MysqlOptionsBuilder) error {
	currSecret, err := getSecret(component.SystemSecretSystemDatabaseSecretName, o.Namespace, o.Client)
	defaultDatabaseName := "system"
	defaultDatabaseRootPassword := oprand.String(8)
	defaultDatabaseUsername := "mysql"
	defaultDatabasePassword := oprand.String(8)
	// TODO is this correct?? in templates the user provides dbname and rootpassword
	// but the secret is only the URL.
	defaultDatabaseURL := "mysql2://root:" + defaultDatabaseRootPassword + "@system-mysql/" + defaultDatabaseName
	if err != nil {
		if errors.IsNotFound(err) {
			// Set options defaults
			builder.DatabaseName(defaultDatabaseName)
			builder.User(defaultDatabaseUsername)
			builder.Password(defaultDatabasePassword)
			builder.RootPassword(defaultDatabaseRootPassword)
			builder.DatabaseURL(defaultDatabaseURL)
		} else {
			return err
		}
	} else {
		// If a field of a secret already exists in the deployed secret then
		// We do not modify it. Otherwise we set a default value
		secretData := currSecret.Data
		builder.User(getSecretDataValueOrDefault(secretData, component.SystemSecretSystemDatabaseUserFieldName, defaultDatabaseUsername))
		builder.Password(getSecretDataValueOrDefault(secretData, component.SystemSecretSystemDatabasePasswordFieldName, defaultDatabasePassword))
		err := o.parseAndSetDatabaseURLAndParts(builder, secretData, defaultDatabaseURL)
		if err != nil {
			return err
		}
	}
	return nil
}

func (o *OperatorMysqlOptionsProvider) parseAndSetDatabaseURLAndParts(builder *component.MysqlOptionsBuilder, secretData map[string][]byte, defaultDatabaseURL string) error {
	resultURLStr := getSecretDataValueOrDefault(secretData, component.SystemSecretSystemDatabaseURLFieldName, defaultDatabaseURL)
	resultURL, err := o.systemDatabaseURLIsValid(resultURLStr)
	if err != nil {
		return err
	}
	builder.DatabaseURL(resultURLStr)
	builder.DatabaseName(strings.TrimPrefix(resultURL.Path, "/")) // Remove possible leading slash in URL Path
	dbRootPassword, _ := resultURL.User.Password()
	builder.RootPassword(dbRootPassword)
	return nil
}

func (o *OperatorMysqlOptionsProvider) systemDatabaseURLIsValid(rawURL string) (*url.URL, error) {
	resultURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("'%s' field of '%s' secret must have 'scheme://user:password@host/path' format", component.SystemSecretSystemDatabaseURLFieldName, component.SystemSecretSystemDatabaseSecretName)
	}
	if resultURL.Scheme != "mysql2" {
		return nil, fmt.Errorf("'%s' field of '%s' secret must contain 'mysql2' as the scheme part", component.SystemSecretSystemDatabaseURLFieldName, component.SystemSecretSystemDatabaseSecretName)
	}

	if resultURL.User == nil {
		return nil, fmt.Errorf("authentication information in '%s' field of '%s' secret must be provided", component.SystemSecretSystemDatabaseURLFieldName, component.SystemSecretSystemDatabaseSecretName)
	}
	if resultURL.User.Username() == "" {
		return nil, fmt.Errorf("authentication information in '%s' field of '%s' secret must contain a username", component.SystemSecretSystemDatabaseURLFieldName, component.SystemSecretSystemDatabaseSecretName)
	}
	if resultURL.User.Username() != "root" {
		return nil, fmt.Errorf("authentication information in '%s' field of '%s' secret must contain 'root' as the username", component.SystemSecretSystemDatabaseURLFieldName, component.SystemSecretSystemDatabaseSecretName)
	}
	if _, set := resultURL.User.Password(); !set {
		return nil, fmt.Errorf("authentication information in '%s' field of '%s' secret must contain a password", component.SystemSecretSystemDatabaseURLFieldName, component.SystemSecretSystemDatabaseSecretName)
	}

	if resultURL.Host == "" {
		return nil, fmt.Errorf("host information in '%s' field of '%s' secret must be provided", component.SystemSecretSystemDatabaseURLFieldName, component.SystemSecretSystemDatabaseSecretName)
	}
	if resultURL.Path == "" {
		return nil, fmt.Errorf("database name in '%s' field of '%s' secret must be provided", component.SystemSecretSystemDatabaseURLFieldName, component.SystemSecretSystemDatabaseSecretName)
	}

	return resultURL, nil
}
