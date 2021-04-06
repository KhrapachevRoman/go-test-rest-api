# go-test-rest-api
Simple REST API server that will expose endpoints to allow accessing and manipulating ‘products’.

API Specification
Application should:

Create a new product in response to a valid POST request at /product,
Update a product in response to a valid PUT request at /product/{id},
Delete a product in response to a valid DELETE request at /product/{id},
Fetch a product in response to a valid GET request at /product/{id}, and
Fetch a list of products in response to a valid GET request at /products.
The {id} in some of the endpoints above will determine which product the request will work with.