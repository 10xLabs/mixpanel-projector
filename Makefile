build:
	GOOS=linux GOARCH=amd64 go build -ldflags="-d -s -w" -o main -mod=vendor
stay:
	GOOS=linux go build -o main
	docker run --rm -e DOCKER_LAMBDA_STAY_OPEN=1 -p 9001:9001 -v $(PWD):/var/task:ro,delegated --env-file dev.env lambci/lambda:go1.x main
zip:
	make build
	zip main.zip main
dep:
	make zip
	aws lambda update-function-code --function-name staging-bookings-projector --zip-file fileb://main.zip
run:
	make build
	./runTestEvents BookingCreated
	./runTestEvents BookingDepartureTripAdded