package product

type release_2_5 struct{}

func (p *release_2_5) GetApicastImage() string {
	return "quay.io/3scale/apicast:nightly"
}

func (p *release_2_5) GetBackendImage() string {
	return "quay.io/3scale/apisonator:nightly"
}

func (p *release_2_5) GetBackendRedisImage() string {
	return "docker.io/centos/redis-32-centos7:3.2"
}

func (p *release_2_5) GetSystemImage() string {
	return "quay.io/3scale/porta:nightly"
}

func (p *release_2_5) GetSystemRedisImage() string {
	return "docker.io/centos/redis-32-centos7:3.2"
}

func (p *release_2_5) GetSystemMySQLImage() string {
	return "docker.io/centos/mysql-57-centos7:5.7"
}

func (p *release_2_5) GetSystemPostgreSQLImage() string {
	return "docker.io/centos/postgresql-10-centos7"
}

func (p *release_2_5) GetSystemMemcachedImage() string {
	return "docker.io/library/memcached"
}

func (p *release_2_5) GetWildcardRouterImage() string {
	return "quay.io/3scale/wildcard-router:nightly"
}

func (p *release_2_5) GetZyncImage() string {
	return "quay.io/3scale/zync:nightly"
}

func (p *release_2_5) GetZyncPostgreSQLImage() string {
	return "docker.io/centos/postgresql-10-centos7"
}
