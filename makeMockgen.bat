mockgen -source .\repository\todo_repository.go -destination mock/mock_todo_repository.go -package mock TodoRepository
mockgen -source .\service\todo_service.go -destination mock/mock_todo_service.go -package mock TodoService
mockgen -source .\controllers\todo_controller.go -destination mock/mock_todo_controller.go -package mock TodoController

go test -coverprofile=cover.txt ./...
go tool cover -html=cover.txt -o cover.html