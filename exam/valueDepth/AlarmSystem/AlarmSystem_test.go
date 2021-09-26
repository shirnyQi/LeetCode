package test1

import (
	"fmt"
	"sort"
	"testing"
)

type AlarmSystem struct {
	alarmMap map[int]alarm
	weekMap  map[int][]int
}

type alarm struct {
	id       int
	hour     int
	minute   int
	weekdays []int
	typeId   int
}

func Constructor() AlarmSystem {
	return AlarmSystem{
		alarmMap: make(map[int]alarm, 0),
		weekMap:  make(map[int][]int, 0),
	}
}

func (alarmSystem *AlarmSystem) addAlarm(id int, hour int, minute int, weekdays []int, typeId int) bool {
	if _, ok := alarmSystem.alarmMap[id]; ok {
		return false
	}
	alarmSystem.alarmMap[id] = alarm{
		id:       id,
		hour:     hour,
		minute:   minute,
		weekdays: weekdays,
		typeId:   typeId,
	}

	for _, v := range weekdays {
		alarmSystem.weekMap[v] = append(alarmSystem.weekMap[v], id)
	}

	return true
}

func (alarmSystem *AlarmSystem) deleteAlarm(id int) bool {
	if _, ok := alarmSystem.alarmMap[id]; !ok {
		return false
	}

	weekdays := alarmSystem.alarmMap[id].weekdays
	for _, v := range weekdays {
		idList := alarmSystem.weekMap[v]
		i := -1
		for j, v := range idList {
			if v == id {
				i = j
			}
		}
		if i == -1 {
			continue
		}
		alarmSystem.weekMap[v] = append(alarmSystem.weekMap[v][:i], alarmSystem.weekMap[v][i+1:]...)
	}
	delete(alarmSystem.alarmMap, id)
	return true
}

func (alarmSystem *AlarmSystem) queryAlarm(weekday int, hour int, startMinute int, endMinute int) []int {
	alarmIdList := alarmSystem.weekMap[weekday]
	var alarmList []alarm
	for _, id := range alarmIdList {
		a := alarmSystem.alarmMap[id]
		if a.hour == hour && a.minute >= startMinute && a.minute <= endMinute {
			alarmList = append(alarmList, a)
		}
	}

	sort.Slice(alarmList, func(i, j int) bool {
		a, b := alarmList[i], alarmList[j]
		if a.minute < b.minute {
			return true
		} else if a.minute == b.minute {
			if a.typeId < b.typeId {
				return true
			} else if a.typeId == b.typeId {
				return a.id < b.id
			}
		}
		return false
	})
	var res []int

	for _, a := range alarmList {
		res = append(res, a.id)
	}
	return res
}

func TestName(t *testing.T) {
	obj1 := Constructor()
	fmt.Println(obj1.addAlarm(7, 22, 30, []int{3, 2, 6}, 0))
	fmt.Println(obj1.addAlarm(2, 22, 30, []int{1, 2}, 2))
	fmt.Println(obj1.queryAlarm(2, 22, 0, 30))
	fmt.Println(obj1.queryAlarm(3, 22, 17, 36))
	fmt.Println(obj1.deleteAlarm(7))
}
