package calendar

import (
	"errors"
	"time"
)

type CalendarEvent struct {
	day *CalendarDay

	eventID   string
	startTime time.Time
	endTime   time.Time
	title     string

	subEvents *subEventList
}

type subEventList struct {
	head   *subEvent
	tail   *subEvent
	length int
}

type subEvent struct {
	index       int
	discription string
	finished    bool

	nextSubEvent *subEvent
}

func createNewCalendarEvent(start, end time.Time, title string) *CalendarEvent {
	cale := &CalendarEvent{
		day:       nil,
		eventID:   "",
		startTime: start,
		endTime:   end,
		title:     title,
		subEvents: nil,
	}
	return cale
}

func (cale *CalendarEvent) AppendNewSubEvent(discription string) {
	// first append sub event
	if cale.subEvents == nil {
		subevent := &subEvent{
			index:        1,
			discription:  discription,
			finished:     false,
			nextSubEvent: nil,
		}

		cale.subEvents = &subEventList{
			head:   subevent,
			tail:   subevent,
			length: 1,
		}
		return
	}

	subevent := &subEvent{
		index:        cale.subEvents.length + 1,
		discription:  discription,
		finished:     false,
		nextSubEvent: nil,
	}
	cale.subEvents.tail.nextSubEvent = subevent
	cale.subEvents.tail = subevent
	cale.subEvents.length++
}

func (cale *CalendarEvent) DeleteSubEvent(index int) error {
	if cale.subEvents == nil || cale.subEvents.length == 0 || index > cale.subEvents.length {
		return errors.New("ERROR: Delete failed, because there is not subevent can be deleted.")
	}

	// only one node.
	if cale.subEvents.head == cale.subEvents.tail {
		cale.subEvents = nil
		return nil
	}

	var pre *subEvent
	var iter *subEvent = cale.subEvents.head

	// delete last node
	if index == cale.subEvents.length {
		for i := 1; i < index; i++ {
			iter = iter.nextSubEvent
			cale.subEvents.tail = iter
			cale.subEvents.length--
		}
		return nil
	}

	for iter.index != index {
		pre = iter
		iter = iter.nextSubEvent
	}
	pre.nextSubEvent = iter.nextSubEvent
	pre = pre.nextSubEvent
	for pre != nil {
		pre.index--
		pre = pre.nextSubEvent
	}
	cale.subEvents.length--
	return nil
}

func (cale *CalendarEvent) CompleteSubEvent(index int) error {
	if cale.subEvents == nil || cale.subEvents.length == 0 || index > cale.subEvents.length {
		return errors.New("ERROR: Complete failed, because there is not subevent can be completed.")
	}

	iter := cale.subEvents.head
	for iter.index != index {
		iter = iter.nextSubEvent
	}

	iter.finished = true
	return nil
}
