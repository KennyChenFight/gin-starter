# gin-starter

Based on this template, we can build our application faster.

+ Based on Gin web framework
+ Simple CRUD application
+ Use route-service-dao three layers
+ Provide database migration
+ Provide model validation and i18n error code translate
+ Provide bussiness error codes for our application domain

## How to run

### local build and run

+ Requirement: install postgres and golang in your computer

```bash
make local-build
make local-run
```

### docker build and run

todo...

## Roadmap

1. Implement auth middleware
2. business error codes should design better
3. add more script in makefile to build and run, like docker, docker-compose
4. add unit test