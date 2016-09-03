# shameboardAPI

**Basic REST API server written in go.**

_Supported endpoints:_
- GET /allshames - Returns all Shames
- GET /shame/{ida} - Returns shame with provided ID
- POST /shame/generate - Params: Body must contain valid Shame object - Creates a shame object

- GET /allusers - Returns all Users
- GET /user/{id} - Returns user with provided ID
- POST /user/generate - Params: Body must contain valid User object - Creates a shame object
- DELETE /user/{id} - Deletes user with provided ID
