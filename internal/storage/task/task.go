package task

import (
	"fmt"
	"sort"
	"time"

	"github.com/c2micro/c2mcli/internal/utils"
	"github.com/c2micro/c2mshr/defaults"
	"github.com/fatih/color"
	"github.com/lrita/cmap"
)

type TaskData interface {
	GetCreatedAt() time.Time
}

// хранение биконов в рантайме
var TaskGroups = &taskGroupsMapper{
	sorted: &taskGroups{
		taskGroups: make([]*TaskGroup, 0),
	},
}

// ресет сторожа
func ResetStorage() {
	TaskGroups = &taskGroupsMapper{
		sorted: &taskGroups{
			taskGroups: make([]*TaskGroup, 0),
		},
	}
}

// структура для хранения информации об отдельно взятой таск группе
type TaskGroup struct {
	id        int64
	cmd       string
	createdAt time.Time
	closedAt  time.Time
	author    string
	data      *TaskGroupResults
}

// хренение результатов выполнения таск группы
type TaskGroupResults struct {
	messages cmap.Map[int64, *Message]
	tasks    cmap.Map[int64, *Task]
	sorted   []TaskData
}

// добавление сообщения
func (t *TaskGroupResults) AddMessage(m *Message) {
	t.messages.Store(m.GetId(), m)
	t.Fill()
}

// получение результатов
func (t *TaskGroupResults) Get() []TaskData {
	return t.sorted
}

// получение таска из группы по его id
func (t *TaskGroupResults) GetTaskById(id int64) *Task {
	if v, ok := t.tasks.Load(id); ok {
		return v
	}
	return nil
}

// добавление таска
func (t *TaskGroupResults) AddTask(task *Task) {
	t.tasks.Store(task.GetId(), task)
	t.Fill()
}

// обновление таска
func (t *TaskGroupResults) UpdateTask(v *Task) {
	t.tasks.Store(v.GetId(), v)
	t.Fill()
}

// заполнение массива с сортировкой
func (t *TaskGroupResults) Fill() {
	temp := make([]TaskData, 0)

	// переносим из мапы с сообщениями в список
	t.messages.Range(func(k int64, v *Message) bool {
		temp = append(temp, v)
		return true
	})

	// переносим из мапы с тасками в список
	t.tasks.Range(func(k int64, v *Task) bool {
		temp = append(temp, v)
		return true
	})

	// сортируем
	sort.SliceStable(temp, func(i, j int) bool {
		return temp[i].GetCreatedAt().Before(temp[j].GetCreatedAt())
	})

	t.sorted = temp
}

// структура для хранения информации о сообщении внутри таск группы
type Message struct {
	TaskData
	id        int64
	kind      defaults.TaskMessage
	message   string
	createdAt time.Time
}

func (m *Message) String() string {
	var data string
	switch m.kind {
	case defaults.NotifyMessage:
		data = fmt.Sprintf("[%s] %s", color.CyanString("*"), m.message)
	case defaults.InfoMessage:
		data = fmt.Sprintf("[%s] %s", color.GreenString("+"), m.message)
	case defaults.WarningMessage:
		data = fmt.Sprintf("[%s] %s", color.YellowString("!"), m.message)
	case defaults.ErrorMessage:
		data = fmt.Sprintf("[%s] %s", color.RedString("-"), m.message)
	}
	return data
}

// стркутра для хранения информации о таске внутри таск группы
type Task struct {
	TaskData
	id          int64
	isOutputBig bool
	isBinary    bool
	output      []byte
	outputLen   int64
	status      defaults.TaskStatus
	createdAt   time.Time
}

func (t *Task) StringStatus() string {
	var data string
	switch t.status {
	case defaults.StatusInProgress:
		data = fmt.Sprintf("[%s] (%d) recieved output with length %d bytes", color.CyanString("*"), t.id, t.outputLen)
	case defaults.StatusCancelled:
		data = fmt.Sprintf("[%s] (%d) recieved output with length %d bytes", color.YellowString("!"), t.id, t.outputLen)
	case defaults.StatusError:
		data = fmt.Sprintf("[%s] (%d) recieved output with length %d bytes", color.RedString("!"), t.id, t.outputLen)
	case defaults.StatusSuccess:
		data = fmt.Sprintf("[%s] (%d) recieved output with length %d bytes", color.GreenString("+"), t.id, t.outputLen)
	}
	return data
}

func (t *Task) GetId() int64 {
	return t.id
}

func (t *Task) GetIdHex() string {
	return fmt.Sprintf("%06x", t.id)[:6]
}

func (t *Task) SetId(id int64) {
	t.id = id
}

func (t *Task) GetIsOutputBig() bool {
	return t.isOutputBig
}

func (t *Task) SetIsOutputBig(flag bool) {
	t.isOutputBig = flag
}

func (t *Task) GetOutput() []byte {
	return t.output
}

func (t *Task) GetOutputString() string {
	return string(t.output)
}

func (t *Task) SetOutput(data []byte) {
	if !utils.IsAsciiPrintable(string(data)) {
		t.SetIsBinary(true)
	} else {
		t.SetIsBinary(false)
	}
	t.output = data
}

func (t *Task) GetOutputLen() int64 {
	return t.outputLen
}

func (t *Task) SetOutputLen(length int64) {
	t.outputLen = length
}

func (t *Task) GetIsBinary() bool {
	return t.isBinary
}

func (t *Task) SetIsBinary(flag bool) {
	t.isBinary = flag
}

func (t *Task) GetStatus() defaults.TaskStatus {
	return t.status
}

func (t *Task) SetStatus(status defaults.TaskStatus) {
	t.status = status
}

func (t *Task) GetCreatedAt() time.Time {
	return t.createdAt
}

func (t *Task) SetCreatedAt(createdAt time.Time) {
	t.createdAt = createdAt
}

func (m *Message) GetId() int64 {
	return m.id
}

func (m *Message) SetId(id int64) {
	m.id = id
}

func (m *Message) GetKind() defaults.TaskMessage {
	return m.kind
}

func (m *Message) SetKind(kind defaults.TaskMessage) {
	m.kind = kind
}

func (m *Message) GetMessage() string {
	return m.message
}

func (m *Message) SetMessage(message string) {
	m.message = message
}

func (m *Message) GetCreatedAt() time.Time {
	return m.createdAt
}

func (m *Message) SetCreatedAt(t time.Time) {
	m.createdAt = t
}

// получение таска из группы по его id
func (t *TaskGroup) GetTaskById(id int64) *Task {
	return t.data.GetTaskById(id)
}

func (t *TaskGroup) UpdateTask(task *Task) {
	t.data.UpdateTask(task)
}

func (t *TaskGroup) AddMessage(m *Message) {
	t.data.AddMessage(m)
}

func (t *TaskGroup) AddTask(task *Task) {
	t.data.AddTask(task)
}

func (t *TaskGroup) GetId() int64 {
	return t.id
}

func (t *TaskGroup) GetIdHex() string {
	return fmt.Sprintf("%06x", t.id)[:6]
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

func (t *TaskGroup) GetData() *TaskGroupResults {
	return t.data
}

type taskGroupsMapper struct {
	taskGroups cmap.Map[int64, *TaskGroup]
	sorted     *taskGroups
}

type taskGroups struct {
	taskGroups []*TaskGroup
}

// получение последней таск группы для бикона
func (t *taskGroupsMapper) GetLast() *TaskGroup {
	data := t.Get()
	if len(data) == 0 {
		return nil
	}
	return data[len(data)-1]
}

// добавление бикона в хранилище
func (t *taskGroupsMapper) Add(v *TaskGroup) {
	v.data = &TaskGroupResults{
		sorted: make([]TaskData, 0),
	}
	t.taskGroups.Store(v.GetId(), v)
	t.Fill()
}

// получение списка отсортированных таск групп
func (t *taskGroupsMapper) Get() []*TaskGroup {
	return t.sorted.taskGroups
}

// получение всех тасок в таск группах
func (t *taskGroupsMapper) GetTasks() []*Task {
	temp := make([]*Task, 0)
	t.taskGroups.Range(func(k int64, v *TaskGroup) bool {
		v.data.tasks.Range(func(key int64, value *Task) bool {
			temp = append(temp, value)
			return true
		})
		return true
	})
	return temp
}

// получение таск группы по id
func (t *taskGroupsMapper) GetById(id int64) *TaskGroup {
	if v, ok := t.taskGroups.Load(id); ok {
		return v
	}
	return nil
}

// получение количества таск групп в мапе
func (t *taskGroupsMapper) Count() int {
	return t.taskGroups.Count()
}

// сортировка списка с таск группами
func (t *taskGroups) Sort() {
	sort.SliceStable(t.taskGroups, func(i, j int) bool {
		return t.taskGroups[i].GetCreatedAt().Before(t.taskGroups[j].GetCreatedAt())
	})
}

// заполнение списка с таск группами на базе мапы с дефолтной сортировкой
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
