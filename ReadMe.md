# ShopFast Microservices

A production-grade microservices application deployed on Amazon EKS.

## Services
| Service | Language | Port | Description |
|---------|----------|------|-------------|
| API Gateway | Node.js | 3000 | Single entry point, routes to services |
| Product Service | Python (FastAPI) | 5000 | Product catalog management |
| Order Service | Go | 8080 | Order creation and management |

## Tech Stack
- **Containerization:** Docker + Docker Hub
- **Orchestration:** Kubernetes (EKS)
- **Deployment:** Helm Charts
- **CI/CD:** GitHub Actions
- **Infra:** Provisioned via [Terraform EKS Project](https://github.com/SubodhaDash/terraform-eks-cluster)

## Architecture
\```
Internet → ALB Ingress → API Gateway → Product Service
                                     → Order Service
\```

## CI/CD Pipeline
\```
Code Push → GitHub Actions → Test → Build Docker Image → Push to Docker Hub → Helm Deploy to EKS
\```
EOF