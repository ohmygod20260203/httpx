package runner

import "sync"

type ResumeCfg struct {
	ResumeFrom         string
	Index              int
	current            string
	currentIndex       int
	lastCompletedIndex int
	lastCompletedItem  string
	mu                 sync.Mutex
}
