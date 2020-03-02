# Backing Catalog Service

catalog service for the Backing-fulfillment-service
For an interactive documentation and code samples visit the [Catalog Services Docs](https://backingcatalogservice.docs.apiary.io/#)
To build and run:

- `go build` to build the executable.

- Run the application. Make sure that the fulfillment service is also running, otherwise requests for individual SKU details will fail.

# Service API

| Resource              | Method | Description                                                                                        |
| --------------------- | ------ | -------------------------------------------------------------------------------------------------- |
| /catalog              | GET    | Retrieves a summary of catalog items                                                               |
| /catalog/{productId } | GET    | Retrieves details for an individual catalog item. This will invoke the fulfillment backing service |
