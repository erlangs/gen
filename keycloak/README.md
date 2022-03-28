[comment]: <> (This is a generated file please edit source in ./templates)
[comment]: <> (All modification will be lost, you have been warned)
[comment]: <> ()
### Sample CRUD API for the postgres database user=postgres password=postgres dbname=keycloak sslmode=disable

## Example
The project is a RESTful api for accessing the postgres database user=postgres password=postgres dbname=keycloak sslmode=disable.

## Project Files
The generated project will contain the following code under the `./example` directory.
* Makefile
  * useful Makefile for installing tools building project etc. Issue `make` to display help
* .gitignore
  * git ignore for go project
* go.mod
  * go module setup, pass `--module` flag for setting the project module default `example.com/example`
* README.md
  * Project readme
* app/server/main.go
  * Sample Gin Server, with swagger init and comments
* api/*.go
  * REST crud controllers
* dao/*.go
  * DAO functions providing CRUD access to database
* model/*.go
  * Structs representing a row for each database table

The REST api server utilizes the Gin framework, GORM db api and Swag for providing swagger documentation
* [Gin](https://github.com/gin-gonic/gin)
* [Swaggo](https://github.com/swaggo/swag)
* [Gorm](https://github.com/jinzhu/gorm)

## Building
```.bash
make example
```
Will create a binary `./bin/example`

## Running
```.bash
./bin/example
```
This will launch the web server on localhost:8080

## Swagger
The swagger web ui contains the documentation for the http server, it also provides an interactive interface to exercise the api and view results.
http://localhost:8080/swagger/index.html

## REST urls for fetching data


* http://localhost:8080/adminevententity
* http://localhost:8080/associatedpolicy
* http://localhost:8080/authenticationexecution
* http://localhost:8080/authenticationflow
* http://localhost:8080/authenticatorconfig
* http://localhost:8080/authenticatorconfigentry
* http://localhost:8080/brokerlink
* http://localhost:8080/client
* http://localhost:8080/clientattributes
* http://localhost:8080/clientauthflowbindings
* http://localhost:8080/clientinitialaccess
* http://localhost:8080/clientnoderegistrations
* http://localhost:8080/clientscope
* http://localhost:8080/clientscopeattributes
* http://localhost:8080/clientscopeclient
* http://localhost:8080/clientscoperolemapping
* http://localhost:8080/clientsession
* http://localhost:8080/clientsessionauthstatus
* http://localhost:8080/clientsessionnote
* http://localhost:8080/clientsessionprotmapper
* http://localhost:8080/clientsessionrole
* http://localhost:8080/clientusersessionnote
* http://localhost:8080/component
* http://localhost:8080/componentconfig
* http://localhost:8080/compositerole
* http://localhost:8080/credential
* http://localhost:8080/databasechangelog
* http://localhost:8080/databasechangeloglock
* http://localhost:8080/defaultclientscope
* http://localhost:8080/evententity
* http://localhost:8080/feduserattribute
* http://localhost:8080/feduserconsent
* http://localhost:8080/feduserconsentclscope
* http://localhost:8080/fedusercredential
* http://localhost:8080/fedusergroupmembership
* http://localhost:8080/feduserrequiredaction
* http://localhost:8080/feduserrolemapping
* http://localhost:8080/federatedidentity
* http://localhost:8080/federateduser
* http://localhost:8080/groupattribute
* http://localhost:8080/grouprolemapping
* http://localhost:8080/identityprovider
* http://localhost:8080/identityproviderconfig
* http://localhost:8080/identityprovidermapper
* http://localhost:8080/idpmapperconfig
* http://localhost:8080/keycloakgroup
* http://localhost:8080/keycloakrole
* http://localhost:8080/migrationmodel
* http://localhost:8080/offlineclientsession
* http://localhost:8080/offlineusersession
* http://localhost:8080/policyconfig
* http://localhost:8080/protocolmapper
* http://localhost:8080/protocolmapperconfig
* http://localhost:8080/realm
* http://localhost:8080/realmattribute
* http://localhost:8080/realmdefaultgroups
* http://localhost:8080/realmenabledeventtypes
* http://localhost:8080/realmeventslisteners
* http://localhost:8080/realmlocalizations
* http://localhost:8080/realmrequiredcredential
* http://localhost:8080/realmsmtpconfig
* http://localhost:8080/realmsupportedlocales
* http://localhost:8080/redirecturis
* http://localhost:8080/requiredactionconfig
* http://localhost:8080/requiredactionprovider
* http://localhost:8080/resourceattribute
* http://localhost:8080/resourcepolicy
* http://localhost:8080/resourcescope
* http://localhost:8080/resourceserver
* http://localhost:8080/resourceserverpermticket
* http://localhost:8080/resourceserverpolicy
* http://localhost:8080/resourceserverresource
* http://localhost:8080/resourceserverscope
* http://localhost:8080/resourceuris
* http://localhost:8080/roleattribute
* http://localhost:8080/scopemapping
* http://localhost:8080/scopepolicy
* http://localhost:8080/user
* http://localhost:8080/userattribute
* http://localhost:8080/userconsent
* http://localhost:8080/userconsentclientscope
* http://localhost:8080/userentity
* http://localhost:8080/userfederationconfig
* http://localhost:8080/userfederationmapper
* http://localhost:8080/userfederationmapperconfig
* http://localhost:8080/userfederationprovider
* http://localhost:8080/usergroupmembership
* http://localhost:8080/userrequiredaction
* http://localhost:8080/userrolemapping
* http://localhost:8080/usersession
* http://localhost:8080/usersessionnote
* http://localhost:8080/usernameloginfailure
* http://localhost:8080/weborigins

## Project Generated Details
```.bash
/Users/kerry/Workspace/Github/gen/go_build_main_go \
    --sqltype=postgres \
    --connstr \
    user=postgres password=postgres dbname=keycloak sslmode=disable \
    --database \
    keycloak \
    --json \
    --gorm \
    --rest \
    --out \
    ./keycloak \
    --module \
    keycloak/rest/api \
    --mod \
    --server \
    --makefile \
    --json-fmt=snake \
    --generate-dao \
    --generate-proj \
    --overwrite
```











