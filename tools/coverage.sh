go test -coverprofile=cover.txt  ./controllers ./repository ./server ./service
go tool cover -html=cover.txt -o cover.html