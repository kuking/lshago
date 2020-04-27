
build:
	go build -o lsha

run:
	go run lsha.go

compile:
	GOOS=linux GOARCH=amd64 go build -o lsha-linux-amd64
	GOOS=linux GOARCH=386   go build -o lsha-linux-386
	GOOS=linux GOARCH=arm	go build -o lsha-linux-arm

release: compile
	strip lsha-linux-amd64
	upx -9 lsha-linux-amd64
	strip lsha-linux-386
	upx -9 lsha-linux-386
	# strip lsha-linux-arm
	upx -9 lsha-linux-arm

clean:
	rm lsha-linux-amd64
	rm lsha-linux-386
	rm lsha-linux-arm

