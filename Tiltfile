# 1. Database
k8s_yaml('infra/k8s/postgres.yaml')
k8s_resource('postgres', port_forwards=5432)

# 2. Build Independent Apps
# Each build points to the root context '.' so they can share libraries if needed
docker_build(
    'asset-guard-asset-service',
    context='./services/asset-service',
    dockerfile='services/asset-service/Dockerfile'
)
docker_build(
    'asset-guard-gateway',
    context='.',
    dockerfile='apps/gateway/Dockerfile'
)

docker_build(
    'asset-guard-frontend',
    context='.',
    dockerfile='apps/frontend/Dockerfile'
)

# 3. Deploy and Port Forward
k8s_yaml('infra/k8s/apps.yaml')

# Go Service
k8s_resource('asset-service', port_forwards=8080)

# NestJS Gateway
k8s_resource('gateway', port_forwards=3000)

# Next.js Frontend
k8s_resource('frontend', port_forwards=4200)

# 4. Watch and Test (Automatic Pass/Fail in Tilt UI)
local_resource(
    'gateway-unit-tests',
    cmd='npx nx test gateway',
    deps=['apps/gateway/src'],
    resource_deps=['gateway']
)

local_resource(
    'frontend-e2e',
    cmd='npx nx e2e frontend-e2e --headless',
    deps=['apps/frontend/src', 'apps/frontend-e2e/src'],
    resource_deps=['frontend', 'gateway']
)
