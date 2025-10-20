# `/api`

OpenAPI/Swagger specs, JSON schema files, protocol definition files.

Examples:

* https://github.com/kubernetes/kubernetes/tree/master/api
* https://github.com/moby/moby/tree/master/api

> [!NOTE]
> For semantic and idiomatic purposes of project structure organization, the *Interface layer* adopts the naming of `api` for consistency, mostly because alphabetically within file / directory sorting, `api` comes first, allowing us to sort our directories in order of the domain burger.
>
> ```shell
> # DDD Hamburger compliant
> 
> DOMAIN_MODULE
> ├── api
> ├── app
> ├── domain
> └── infra
> ```
>
> versus...
>
> ```shell
> ├── app
> ├── domain
> ├── infra
> └── interface
> ```

API Folder Structure:

```shell
/{ protocol }/{ module }/{ version }/
```
