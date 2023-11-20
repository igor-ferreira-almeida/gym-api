package observer

import (
	"fmt"
)

type Observable interface {
	Add(o Observer)
	Remove(o Observer)
	Notify()
}

type ElementSubject struct {
	observers []Observer
}

func (s *ElementSubject) Add(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *ElementSubject) Remove(o Observer) {

}

func (s *ElementSubject) Notify() {
	for _, o := range s.observers {
		o.Update()
	}
}

type Observer interface {
	Update()
}

type ElementObserver struct {
	cache string
}

func (c ElementObserver) Update() {
	fmt.Println("Element updated in cache")
}

type SegmentationObserver struct {
	cache string
}

func (c SegmentationObserver) Update() {
	fmt.Println("Segmentation updated in cache")
}
