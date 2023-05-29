package main

//go:generate rm -rf openapi/model openapi/restapi openapi/client
//go:generate swagger generate server -t ./openapi/ -f ./openapi/swagger.yaml -P models.Principal --strict-responders --strict-additional-properties --exclude-main
//go:generate swagger generate client -t ./openapi/ -f ./openapi/swagger.yaml -P models.Principal --strict-responders --strict-additional-properties
///go:generate find restapi -maxdepth 1 -name "configure_*.go" -exec sed -i -e "/go:generate/d" {} ;
