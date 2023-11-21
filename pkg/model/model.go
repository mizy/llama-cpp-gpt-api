package model

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-skynet/go-llama.cpp"
	"github.com/mizy/llama-cpp-gpt-api/internal/config"
	"github.com/mizy/llama-cpp-gpt-api/internal/types"
)

var LlamaInstance *llama.LLama

func LoadModel() {
	l, err := llama.New(config.C.ModelPath, func(p *llama.ModelOptions) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("panic occurred:", err)
				os.Exit(1)
			}
		}()
		for k, v := range config.C.ModelOption {
			ReflectVal(k, v, p)
		}
		log.Print("Model options: ", p)
	})
	if err != nil {
		fmt.Println("Loading the model failed:", err.Error())
		os.Exit(1)
	}
	LlamaInstance = l
	fmt.Println("Model loaded successfully")
}

func ReflectVal(k string, v interface{}, p *llama.ModelOptions) {
	intValue := 0
	floatValue := float32(0.0)
	if num, ok := v.(json.Number); ok {
		intVal, err := num.Int64()
		if err == nil {
			intValue = int(intVal)
		}
		floatVal, err := strconv.ParseFloat(string(num), 32)
		if err != nil {
			floatValue = float32(floatVal)
		}
	}
	switch k {
	case "ContextSize":
		p.ContextSize = intValue
	case "Seed":
		p.Seed = intValue
	case "NBatch":
		p.NBatch = intValue
	case "F16Memory":
		p.F16Memory = v.(bool)
	case "MLock":
		p.MLock = v.(bool)
	case "MMap":
		p.MMap = v.(bool)
	case "LowVRAM":
		p.LowVRAM = v.(bool)
	case "Embeddings":
		p.Embeddings = v.(bool)
	case "NUMA":
		p.NUMA = v.(bool)
	case "NGPULayers":
		p.NGPULayers = intValue
	case "MainGPU":
		p.MainGPU = v.(string)
	case "TensorSplit":
		p.TensorSplit = v.(string)
	case "FreqRopeBase":
		p.FreqRopeBase = floatValue
	case "FreqRopeScale":
		p.FreqRopeScale = floatValue
	case "LoraBase":
		p.LoraBase = v.(string)
	case "LoraAdapter":
		p.LoraAdapter = v.(string)
	case "Perplexity":
		p.Perplexity = v.(bool)
	}
}

// types.Message{Role,Content}
func ConvertMessages2Text(messages []types.Message) string {
	text := fmt.Sprintf("You are %s.\n\n", config.C.UserName)
	for _, m := range messages {
		if m.Role != "" {
			text += m.Role + ":\n" + m.Content + "\n"
		} else {
			text += m.Content + "\n"
		}
	}
	return text + fmt.Sprintf("\n%s:", config.C.UserName)
}
