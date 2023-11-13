refer github.com/go-skynet/go-llama.cpp

# usage

generate binding code
```
cd go-llama.cpp
make libbinding.a
```

start server
```
make gen
make run
```

config model path in etc/gpt-api.yaml