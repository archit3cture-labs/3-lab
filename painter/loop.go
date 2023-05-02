package painter

import (
	"image"

	"golang.org/x/exp/shiny/screen"
)

type Receiver interface {
	Update(t screen.Texture)
}

type Loop struct {
	Receiver Receiver

	next screen.Texture 
	prev screen.Texture 

	MsgQueue messageQueue
}

var size = image.Pt(800, 800)

func (l *Loop) Start(s screen.Screen) {
	l.next, _ = s.NewTexture(size)
	l.prev, _ = s.NewTexture(size)

	l.MsgQueue = messageQueue{}
	go l.eventProcess()
}

func (l *Loop) eventProcess() {
	for {
		if op := l.MsgQueue.Pull(); op != nil {
			if update := op.Do(l.next); update {
				l.Receiver.Update(l.next)
				l.next, l.prev = l.prev, l.next
			}
		}
	}
}

func (l *Loop) Post(op Operation) {
	if op != nil {
		l.MsgQueue.Push(op)
	}
}

func (l *Loop) StopAndWait() {

}

type messageQueue struct {
	Queue []Operation
}

func (MsgQueue *messageQueue) Push(op Operation) {
	MsgQueue.Queue = append(MsgQueue.Queue, op)
}

func (MsgQueue *messageQueue) Pull() Operation {
	if len(MsgQueue.Queue) == 0 {
		return nil
	}

	op := MsgQueue.Queue[0]
	MsgQueue.Queue = MsgQueue.Queue[1:]
	return op
}