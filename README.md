# gin-starter

Based on this template, we can build our application faster.

+ Based on Gin web framework
+ Simple CRUD application
+ Provide database migration
+ Provide model validation and i18n error code translate
+ Provide business error codes for our application domain
+ Provide middleware handle global error and response

## How to run

### local build and run

1. install postgres and golang in your computer

2. ```bash
    make local-build
    make local-run
    ```

### how to build,test,codegen,run

1. build 
```bash
./dockerbuild.sh
```
2. test
```bash
./dockerbuild.sh test
```
3. codegen
```bash
./dockerbuild.sh codegen
```

4. run
```bash
docker-compose up server
```

## Contribution
+ provide your idea about this template in issue or raise PR request

## Roadmap

1. implement more useful middleware
2. add unit test
3. add CI/CD