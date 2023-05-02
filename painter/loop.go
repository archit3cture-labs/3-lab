package painter

import (
	"image"

	"golang.org/x/exp/shiny/screen"
)

// Receiver ������ ��������, ��� ���� ����������� � ��������� ��������� ������ � ���� ����.
type Receiver interface {
	Update(t screen.Texture)
}

// Loop ������ ���� ���� ��� ���������� �������� �������� ����� ��������� �������� ��������� � ��������� �����.
type Loop struct {
	Receiver Receiver

	next screen.Texture // ��������, ��� ����� ���������
	prev screen.Texture // ��������, ��� ���� ���������� ���������� ���� � Receiver

	MsgQueue messageQueue
}

var size = image.Pt(800, 800)

// Start ������� ���� ����. ��� ����� ������� ��������� �� ����, �� ��������� �� ����� ����-�� ���� ������.
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

// Post ���� ���� �������� � �������� �����.
func (l *Loop) Post(op Operation) {
	// TODO: ���������� ��������� �������� � �����. ������� �������������
	if op != nil {
		l.MsgQueue.Push(op)
	}
}

// StopAndWait ��������
func (l *Loop) StopAndWait() {

}

// TODO: ���������� ������ ����� ����������.
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