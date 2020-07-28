package progressbar

import (
	"testing"
	"time"
)

func TestProgressBar_Play(t *testing.T) {
	bar := NewProgressBar(50)
	for i:=1;i<=50;i++{
		bar.Play(int64(i))
		time.Sleep(100*time.Millisecond)
	}
	bar.Finish()
}
