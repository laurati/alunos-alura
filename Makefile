#-------------------- Run 
run-api:
	go run main.go


#-------------------- Test 
test:
	go test


#-------------------- Swagger 
swagger:
	swag init  --parseDependency --parseInternal --parseDepth 1