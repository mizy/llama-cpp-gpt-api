refer github.com/go-skynet/go-llama.cpp

# usage

generate binding code
```
cd go-llama.cpp
make libbinding.a

# for mac m1
BUILD_TYPE=metal make libbinding.a
cp go-llama.cpp/ggml-metal.metal ./
```

start server
```
make gen
make run
```

config model path in etc/gpt-api.yaml