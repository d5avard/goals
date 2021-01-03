# goals

## Curl commands for tests
``` zsh
curl -X GET localhost:3000/api/goals/
curl -X GET localhost:3000/api/goals/ -v
curl -X GET localhost:3000/api/goals/1
curl -X POST localhost:3000/api/goals/ -d '{"id":123}'
curl -X PUT localhost:3000/api/goals/1 -d '{"id":123}'
curl -X DELETE localhost:3000/api/goals/1 -d '{"id":123}'
```

### Go commands
``` zsh
go test
```

## Git commands
``` zsh
> git status
> git checkout -b feature/example
> git add .
> git commit -m "Comments"
> git push origin feature/example
> git checkout main
> git pull origin main
```