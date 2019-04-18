package product

type upstream struct{}

func (u *upstream) GetApicastImage() string {
	return "quay.io/3scale/apicast:nightly"
}

func (u *upstream) GetBackendImage() string {
	return "quay.io/3scale/apisonator:nightly"
}

func (u *upstream) GetBackendRedisImage() string {
	return "docker.io/centos/redis-32-centos7:3.2"
}

func (u *upstream) GetSystemImage() string {
	return "quay.io/3scale/porta:nightly"
}

func (u *upstream) GetSystemRedisImage() string {
	return "docker.io/centos/redis-32-centos7:3.2"
}

func (u *upstream) GetSystemMySQLImage() string {
	return "docker.io/centos/mysql-57-rhel7:5.7"
}

func (u *upstream) GetSystemPostgreSQLImage() string {
	return "docker.io/centos/postgresql-10-centos7"
}

func (u *upstream) GetSystemMemcachedImage() string {
	return "docker.io/library/memcached"
}

func (u *upstream) GetWildcardRouterImage() string {
	return "quay.io/3scale/wildcard-router:nightly"
}

func (u *upstream) GetZyncImage() string {
	return "quay.io/3scale/zync:nightly"
}

func (u *upstream) GetZyncPostgreSQLImage() string {
	return "docker.io/centos/postgresql-10-centos7"
}
