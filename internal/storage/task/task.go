package task

import (
	"sort"
	"time"

	"github.com/lrita/cmap"
)

// хранение биконов в рантайме
var TaskGroups = &taskGroupsMapper{
	sorted: &taskGroups{
		taskGroups: make([]*TaskGroup, 0),
	},
}

// структура для хранения информации об отдельно взятой таск группе
type TaskGroup struct {
	id        int64
	cmd       string
	createdAt time.Time
	closedAt  time.Time
	author    string
}

func (t *TaskGroup) GetId() int64 {
	return t.id
}

func (t *TaskGroup) SetId(id int64) {
	t.id = id
}

func (t *TaskGroup) GetCmd() string {
	return t.cmd
}

func (t *TaskGroup) SetCmd(cmd string) {
	t.cmd = cmd
}

func (t *TaskGroup) GetCreatedAt() time.Time {
	return t.createdAt
}

func (t *TaskGroup) SetCreatedAt(createdAt time.Time) {
	t.createdAt = createdAt
}

func (t *TaskGroup) GetClosedAt() time.Time {
	return t.closedAt
}

func (t *TaskGroup) SetClosedAt(closedAt time.Time) {
	t.closedAt = closedAt
}

func (t *TaskGroup) GetAuthor() string {
	return t.author
}

func (t *TaskGroup) SetAuthor(author string) {
	t.author = author
}

type taskGroupsMapper struct {
	taskGroups cmap.Map[int64, *TaskGroup]
	sorted     *taskGroups
}

type taskGroups struct {
	taskGroups []*TaskGroup
}

// добавление бикона в хранилище
func (t *taskGroupsMapper) Add(v *TaskGroup) {
	t.taskGroups.Store(v.GetId(), v)
	t.Fill()
}

// получение списка отсортированных биконов
func (t *taskGroupsMapper) Get() []*TaskGroup {
	return t.sorted.taskGroups
}

// получение бикона по id
func (t *taskGroupsMapper) GetById(id int64) *TaskGroup {
	if v, ok := t.taskGroups.Load(id); ok {
		return v
	}
	return nil
}

// получение количества биконов в мапе
func (t *taskGroupsMapper) Count() int {
	return t.taskGroups.Count()
}

// сортировка списка с биконами
func (t *taskGroups) Sort() {
	sort.SliceStable(t.taskGroups, func(i, j int) bool {
		return t.taskGroups[i].GetCreatedAt().Before(t.taskGroups[j].GetCreatedAt())
	})
}

// заполнение списка с биконами на базе мапы с дефолтной сортировкой
func (t *taskGroupsMapper) Fill() {
	temp := &taskGroups{
		taskGroups: make([]*TaskGroup, 0),
	}

	// переносим из мапы в список
	t.taskGroups.Range(func(k int64, v *TaskGroup) bool {
		temp.taskGroups = append(temp.taskGroups, v)
		return true
	})

	// сортрируем список
	temp.Sort()

	t.sorted = temp
}
