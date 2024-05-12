package queue

import "testing"

func TestOffer(t *testing.T) {
	q := New()
	q.Offer(1)
	q.Offer(2)
	q.Offer(3)

	if q.Len != 3 {
		t.Errorf(`Expected 3 but got %d`, q.Len)
	}

	if q.Front() != 1 {
		t.Errorf(`Expected 1 but got %d`, q.Front())
	}

	if q.Rear() != 3 {
		t.Errorf(`Expected 3 but got %d`, q.Rear())
	}

}

func TestPoll(t *testing.T) {
	q := New()
	q.Offer(1)
	q.Offer(2)
	q.Offer(3)

	if q.Len != 3 {
		t.Errorf(`Expected 3 but got %d`, q.Len)
	}

	if q.Poll() != 1 {
		t.Errorf(`Expected 1 but got %d`, q.Poll())
	}

	if q.Len != 2 {
		t.Errorf(`Expected 2 but got %d`, q.Len)
	}

	if q.Front() != 2 {
		t.Errorf(`Expected 2 but got %d`, q.Front())
	}

	if q.Rear() != 3 {
		t.Errorf(`Expected 3 but got %d`, q.Rear())
	}
}

func TestPollEmpty(t *testing.T) {
	q := New()
	if q.Poll() != nil {
		t.Errorf(`Expected nil but got %d`, q.Poll())
	}
}

func TestIsEmpty(t *testing.T) {
	q := New()
	if !q.IsEmpty() {
		t.Errorf(`Expected true but got false`)
	}
	q.Offer(1)
	if q.IsEmpty() {
		t.Errorf(`Expected false but got true`)
	}
}

func TestClear(t *testing.T) {
	q := New()
	q.Offer(1)
	q.Offer(2)
	q.Offer(3)
	q.Clear()
	if q.Len != 0 {
		t.Errorf(`Expected 0 but got %d`, q.Len)
	}
}
