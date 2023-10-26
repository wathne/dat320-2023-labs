package wordcount

import (
	"log"
	"os"
	"runtime"
	"sync"
	"unicode"
)

func loadMoby() []byte {
	moby, err := os.ReadFile("mobydick.txt")
	if err != nil {
		log.Fatal(err)
	}
	return moby
}

func wordCount(b []byte) (words int) {
	inWord := false
	for _, v := range b {
		r := rune(v)
		if unicode.IsSpace(r) && inWord {
			words++
		}
		inWord = unicode.IsLetter(r)
	}
	return
}

func shardSlice(input []byte, numShards int) (shards [][]byte) {
	shards = make([][]byte, numShards)
	if numShards < 2 {
		shards[0] = input[:]
		return
	}
	shardSize := len(input) / numShards
	start, end := 0, shardSize
	for i := 0; i < numShards; i++ {
		for j := end; j < len(input); j++ {
			char := rune(input[j])
			if unicode.IsSpace(char) {
				// split slice at position j, where there is a space
				// note: need to include the space in the shard to get accurate count
				end = j + 1
				shards[i] = input[start:end]
				start = end
				end += shardSize
				break
			}
		}
	}
	shards[numShards-1] = input[start:]
	return
}

func parallelWordCount(input []byte) (words int) {
	return doParallelWordCount(input, runtime.NumCPU())
}

func doParallelWordCount(input []byte, numShards int) (words int) {
	// Implement parallel word count.
	// WaitGroup
	return doParallelWordCountWaitGroup(input, numShards)
	// Channel
	//return doParallelWordCountChannel(input, numShards)
}

func doParallelWordCountWaitGroup(input []byte, numShards int) (words int) {
	words = 0
	var shards [][]byte = shardSlice(input, numShards)
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var i int
	for i = range shards {
		wg.Add(1)
		go func(
			wg *sync.WaitGroup,
			mutex *sync.Mutex,
			shard *[]byte,
			words *int,
		) {
			wordCount := wordCount(*shard)
			mutex.Lock()
			*words += wordCount
			mutex.Unlock()
			wg.Done()
		}(&wg, &mutex, &shards[i], &words)
	}
	wg.Wait()
	return words
}

func doParallelWordCountChannel(input []byte, numShards int) (words int) {
	words = 0
	var shards [][]byte = shardSlice(input, numShards)
	var channel chan int = make(chan int, numShards) // Buffered channel.
	var sender chan<- int = channel                  // Send-only channel.
	var receiver <-chan int = channel                // Receive-only channel.
	var i int
	for i = range shards {
		go func(
			sender chan<- int,
			shard *[]byte,
		) {
			sender <- wordCount(*shard)
		}(sender, &shards[i])
	}
	var wordCount int
	for wordCount = range receiver {
		words += wordCount
		if i == 0 {
			close(sender)
		} else {
			i--
		}
	}
	return words
}
