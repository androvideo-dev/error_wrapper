package error_wrapper_test

import (
	"github.com/daniel-androvideo/error_wrapper"

	"errors"
	"testing"
)

func TestDoingMultiTask(t *testing.T) {
	w := error_wrapper.ErrorWrapper{}
	count := 0
	w.Do(func() error {
		count++
		return nil
	}).Do(func() error {
		count++
		return nil
	})
	if count != 2 {
		t.Error("task not work correctly")
	}
	if w.GetErr() != nil {
		t.Error("has unexpected error")
	}
}

func TestTaskFailWillInterruptExecute(t *testing.T) {
	w := error_wrapper.ErrorWrapper{}
	count := 0
	w.Do(func() error {
		count++
		return errors.New("oops")
	})
	w.Do(func() error {
		count++
		return nil
	})
	if count == 2 {
		t.Error("task not interrupt execute")
	}
	if w.GetErr() == nil {
		t.Error("expected error")
	}
}

func TestTaskOnErrorExecuteWhenTaskReturnError(t *testing.T) {
	w := error_wrapper.ErrorWrapper{}
	count := 0
	w.Do(func() error {
		count++
		return nil
	}).OnError(func(err error) {
		t.Error("shouldn't execute")
	}).Do(func() error {
		count++
		return errors.New("hello")
	}).OnError(func(err error) {
		t.Log("should execute")
	})
}

func TestTaskOnErrorWillCatchError(t *testing.T) {
	w := error_wrapper.ErrorWrapper{}
	count := 0
	w.Do(func() error {
		count++
		return errors.New("hello")
	}).OnError(func(err error) {
		t.Log("should execute")
	}).Do(func() error {
		count++
		return errors.New("hello2")
	}).OnError(func(err error) {
		t.Error("shouldn't execute")
	})
}
