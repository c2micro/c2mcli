package beacon

import (
	"fmt"
	"sort"
	"time"

	"github.com/c2micro/c2mshr/defaults"
	"github.com/lrita/cmap"
)

// активный бикон для поллинга
var ActiveBeacon *Beacon

// хранение биконов в рантайме
var Beacons = &beaconsMapper{
	sorted: &beacons{
		beacons: make([]*Beacon, 0),
	},
}

// структура для хранения информации об отдельно взятом биконе
type Beacon struct {
	id           uint32
	listenerId   int64
	extIp        string
	intIp        string
	os           defaults.BeaconOS
	osMeta       string
	hostname     string
	username     string
	domain       string
	isPrivileged bool
	processName  string
	pid          uint32
	arch         defaults.BeaconArch
	sleep        uint32
	jitter       uint8
	caps         uint32
	color        uint32
	note         string
	first        time.Time
	last         time.Time
}

// массив из биконов
type beacons struct {
	beacons []*Beacon
}

type beaconsMapper struct {
	beacons cmap.Map[uint32, *Beacon]
	sorted  *beacons
}

func (b *Beacon) GetId() uint32 {
	return b.id
}

func (b *Beacon) GetIdHex() string {
	return fmt.Sprintf("%06x", b.id)[:6]
}

func (b *Beacon) SetId(id uint32) {
	b.id = id
}

func (b *Beacon) GetListenerId() int64 {
	return b.listenerId
}

func (b *Beacon) SetListenerId(id int64) {
	b.listenerId = id
}

func (b *Beacon) GetExtIp() string {
	return b.extIp
}

func (b *Beacon) SetExtIp(data string) {
	b.extIp = data
}

func (b *Beacon) GetIntIp() string {
	return b.intIp
}

func (b *Beacon) SetIntIp(data string) {
	b.intIp = data
}

func (b *Beacon) GetOs() defaults.BeaconOS {
	return b.os
}

func (b *Beacon) SetOs(os defaults.BeaconOS) {
	b.os = os
}

func (b *Beacon) GetOsMeta() string {
	return b.osMeta
}

func (b *Beacon) SetOsMeta(data string) {
	b.osMeta = data
}

func (b *Beacon) GetHostname() string {
	return b.hostname
}

func (b *Beacon) SetHostname(data string) {
	b.hostname = data
}

func (b *Beacon) GetUsername() string {
	return b.username
}

func (b *Beacon) SetUsername(data string) {
	b.username = data
}

func (b *Beacon) GetDomain() string {
	return b.domain
}

func (b *Beacon) SetDomain(data string) {
	b.domain = data
}

func (b *Beacon) GetIsPrivileged() bool {
	return b.isPrivileged
}

func (b *Beacon) SetIsPrivileged(flag bool) {
	b.isPrivileged = flag
}

func (b *Beacon) GetProcessName() string {
	return b.processName
}

func (b *Beacon) SetProcessName(data string) {
	b.processName = data
}

func (b *Beacon) GetPid() uint32 {
	return b.pid
}

func (b *Beacon) SetPid(pid uint32) {
	b.pid = pid
}

func (b *Beacon) GetArch() defaults.BeaconArch {
	return b.arch
}

func (b *Beacon) SetArch(arch defaults.BeaconArch) {
	b.arch = arch
}

func (b *Beacon) GetSleep() uint32 {
	return b.sleep
}

func (b *Beacon) SetSleep(sleep uint32) {
	b.sleep = sleep
}

func (b *Beacon) GetJitter() uint8 {
	return b.jitter
}

func (b *Beacon) SetJitter(jitter uint8) {
	b.jitter = jitter
}

func (b *Beacon) GetCaps() uint32 {
	return b.caps
}

func (b *Beacon) SetCaps(caps uint32) {
	b.caps = caps
}

func (b *Beacon) GetColor() uint32 {
	return b.color
}

func (b *Beacon) SetColor(color uint32) {
	b.color = color
}

func (b *Beacon) GetNote() string {
	return b.note
}

func (b *Beacon) SetNote(data string) {
	b.note = data
}

func (b *Beacon) GetFirst() time.Time {
	return b.first
}

func (b *Beacon) SetFirst(t time.Time) {
	b.first = t
}

func (b *Beacon) GetLast() time.Time {
	return b.last
}

func (b *Beacon) SetLast(t time.Time) {
	b.last = t
}

// задерживается ли отстук от бикона (1x sleep + sleep * jitter)
func (b *Beacon) IsDelay(delta time.Duration) bool {
	sleep := int(b.sleep * 1000)
	jitter := int(b.jitter / 100)
	return time.Now().After(b.GetLast().Add(time.Duration(sleep+sleep*jitter) * time.Millisecond))
}

// умер ли бикон (3x sleep + sleep * jitter)
func (b *Beacon) IsDead(delta time.Duration) bool {
	sleep := int(b.sleep * 1000)
	jitter := int(b.jitter / 100)
	return time.Now().After(b.GetLast().Add(time.Duration(3*(sleep+sleep*jitter)) * time.Millisecond))
}

// сортировка списка с биконами
func (b *beacons) Sort() {
	sort.SliceStable(b.beacons, func(i, j int) bool {
		return b.beacons[i].GetLast().Before(b.beacons[j].GetLast())
	})
}

// добавление бикона в хранилище
func (b *beaconsMapper) Add(v *Beacon) {
	b.beacons.Store(v.GetId(), v)
	b.Fill()
}

// получение списка отсортированных биконов
func (b *beaconsMapper) Get() []*Beacon {
	return b.sorted.beacons
}

// получение бикона по id
func (b *beaconsMapper) GetById(id uint32) *Beacon {
	if v, ok := b.beacons.Load(id); ok {
		return v
	}
	return nil
}

// получение количества биконов в мапе
func (b *beaconsMapper) Count() int {
	return b.beacons.Count()
}

// заполнение списка с биконами на базе мапы с дефолтной сортировкой
func (b *beaconsMapper) Fill() {
	temp := &beacons{
		beacons: make([]*Beacon, 0),
	}

	// переносим из мапы в список
	b.beacons.Range(func(k uint32, v *Beacon) bool {
		temp.beacons = append(temp.beacons, v)
		return true
	})

	// сортрируем список
	temp.Sort()

	b.sorted = temp
}
