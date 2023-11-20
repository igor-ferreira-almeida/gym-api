package observer

func Test() {
	eObserver := ElementObserver{}
	sObserver := SegmentationObserver{}

	eObservable := ElementSubject{}
	eObservable.Add(eObserver)
	eObservable.Add(sObserver)

	eObservable.Notify()
}
