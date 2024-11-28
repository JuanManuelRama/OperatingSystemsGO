package attend

import "container/list"

type memoryProcess struct {
	pid          int32
	instructions []string
}

var currentProcess *memoryProcess

var processList = list.New()

func findPID(pid int32) *memoryProcess {
	if e := find(pid); e != nil {
		return e.Value.(*memoryProcess)
	}
	return nil
}

func find(pid int32) *list.Element {
	for e := processList.Front(); e != nil; e = e.Next() {
		if e.Value.(*memoryProcess).pid == pid {
			return e
		}
	}
	return nil
}
