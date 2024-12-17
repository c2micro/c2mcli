package service

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/c2micro/c2mcli/internal/storage/beacon"
	"github.com/c2micro/c2mcli/internal/version"
	"github.com/c2micro/c2mshr/defaults"
	operatorv1 "github.com/c2micro/c2mshr/proto/gen/operator/v1"
	"github.com/go-faster/errors"
	"google.golang.org/grpc"
)

// первичное подключение оператора к control стриму
func HelloInit(ctx context.Context) (grpc.ServerStreamingClient[operatorv1.HelloResponse], error) {
	return getSvc().Hello(ctx, &operatorv1.HelloRequest{
		Version: version.Version(),
	})
}

// обработка хендшейка из control стрима
func HelloHandshake(ctx context.Context) error {
	// получение данных
	msg, err := operatorConn.ss.controlStream.Recv()
	if err != nil {
		return err
	}
	if msg.GetHandshake() == nil {
		return fmt.Errorf("unexpected hello response (no handshake data)")
	}
	operatorConn.metadata.username = msg.GetHandshake().GetUsername()
	operatorConn.metadata.cookie = msg.GetHandshake().GetCookie().GetValue()
	operatorConn.metadata.delta = time.Now().Sub(msg.GetHandshake().GetTime().AsTime())
	return nil
}

// поддержание подписки на control стрим
func HelloMonitor(ctx context.Context) error {
	for {
		if _, err := operatorConn.ss.controlStream.Recv(); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
	}
	return nil
}

// подписка на получение биконов
func SubscribeBeacons(ctx context.Context) error {
	stream, err := getSvc().SubscribeBeacons(ctx, &operatorv1.SubscribeBeaconsRequest{
		Cookie: &operatorv1.SessionCookie{
			Value: operatorConn.metadata.cookie,
		},
	})
	if err != nil {
		return errors.Wrap(err, "open beacon subscription stream")
	}

	for {
		msg, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		// добавление биконов
		if msg.GetBeacons() != nil {
			for _, v := range msg.GetBeacons().GetBeacons() {
				b := &beacon.Beacon{}
				b.SetId(v.GetBid())
				b.SetListenerId(v.GetLid())
				b.SetExtIp(v.GetExtIp().GetValue())
				b.SetIntIp(v.GetIntIp().GetValue())
				b.SetOs(defaults.BeaconOS(v.GetOs()))
				b.SetOsMeta(v.GetOsMeta().GetValue())
				b.SetHostname(v.GetHostname().GetValue())
				b.SetUsername(v.GetUsername().GetValue())
				b.SetDomain(v.GetDomain().GetValue())
				b.SetIsPrivileged(v.GetPrivileged().GetValue())
				b.SetProcessName(v.GetProcName().GetValue())
				b.SetPid(v.GetPid().GetValue())
				b.SetArch(defaults.BeaconArch(v.GetArch()))
				b.SetSleep(v.GetSleep())
				b.SetJitter(uint8(v.GetJitter()))
				b.SetCaps(v.GetCaps())
				b.SetColor(v.GetColor().GetValue())
				b.SetNote(v.GetNote().GetValue())
				b.SetFirst(v.GetFirst().AsTime().Add(operatorConn.metadata.delta))
				b.SetLast(v.GetLast().AsTime().Add(operatorConn.metadata.delta))
				// добавление бикона
				beacon.Beacons.Add(b)
			}
			continue
		}
		// добавление бикона
		if msg.GetBeacon() != nil {
			b := &beacon.Beacon{}
			v := msg.GetBeacon()
			b.SetId(v.GetBid())
			b.SetListenerId(v.GetLid())
			b.SetExtIp(v.GetExtIp().GetValue())
			b.SetIntIp(v.GetIntIp().GetValue())
			b.SetOs(defaults.BeaconOS(v.GetOs()))
			b.SetOsMeta(v.GetOsMeta().GetValue())
			b.SetHostname(v.GetHostname().GetValue())
			b.SetUsername(v.GetUsername().GetValue())
			b.SetDomain(v.GetDomain().GetValue())
			b.SetIsPrivileged(v.GetPrivileged().GetValue())
			b.SetProcessName(v.GetProcName().GetValue())
			b.SetPid(v.GetPid().GetValue())
			b.SetArch(defaults.BeaconArch(v.GetArch()))
			b.SetSleep(v.GetSleep())
			b.SetJitter(uint8(v.GetJitter()))
			b.SetCaps(v.GetCaps())
			b.SetColor(v.GetColor().GetValue())
			b.SetNote(v.GetNote().GetValue())
			b.SetFirst(v.GetFirst().AsTime().Add(operatorConn.metadata.delta))
			b.SetLast(v.GetLast().AsTime().Add(operatorConn.metadata.delta))
			// добавление бикона
			beacon.Beacons.Add(b)
			continue
		}
		// обновление заметки
		if msg.GetNote() != nil {
			v := msg.GetNote()
			if b := beacon.Beacons.GetById(v.GetBid()); b != nil {
				b.SetNote(v.GetNote().GetValue())
			}
			continue
		}
		// обновление цвета
		if msg.GetColor() != nil {
			v := msg.GetColor()
			if b := beacon.Beacons.GetById(v.GetBid()); b != nil {
				b.SetColor(v.GetColor().GetValue())
			}
			continue
		}
		// обновление времени последнего чекаута
		if msg.GetLast() != nil {
			v := msg.GetLast()
			if b := beacon.Beacons.GetById(v.GetBid()); b != nil {
				b.SetLast(v.GetLast().AsTime().Add(operatorConn.metadata.delta))
			}
			continue
		}
		// обновление sleep бикона
		if msg.GetSleep() != nil {
			v := msg.GetSleep()
			if b := beacon.Beacons.GetById(v.GetBid()); b != nil {
				b.SetSleep(v.GetSleep())
				b.SetJitter(uint8(v.GetJitter()))
			}
			continue
		}
	}
	return nil
}

// подписка на получение обновлений по таскам
func SubscribeTasks(ctx context.Context) error {
	stream, err := getSvc().SubscribeTasks(ctx)
	if err != nil {
		return errors.Wrap(err, "open tasks subscription stream")
	}
	// авторизационное сообщение от оператора
	if err = stream.Send(&operatorv1.SubscribeTasksRequest{
		Cookie: &operatorv1.SessionCookie{
			Value: operatorConn.metadata.cookie,
		},
		Type: &operatorv1.SubscribeTasksRequest_Hello{
			Hello: &operatorv1.SubscribeTasksHelloRequest{},
		},
	}); err != nil {
		return errors.Wrap(err, "send hello message to tasks topic")
	}
	// сохраняем стрим
	operatorConn.ss.tasksStream = stream
	for {
		msg, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		// новая таск группа
		if msg.GetGroup() != nil {
			continue
		}
		// новое сообщение в таск группе
		if msg.GetMessage() != nil {
			continue
		}
		// новый таск в таск группе
		if msg.GetTask() != nil {
			continue
		}
		// обновление статуса таска
		if msg.GetTaskStatus() != nil {
			continue
		}
		// получение результатов выполненного таска
		if msg.GetTaskDone() != nil {
			continue
		}
	}
	return nil
}

// подписка на получение тасков для определенного бикона
func PollBeaconTasks(b *beacon.Beacon) error {
	if err := operatorConn.ss.tasksStream.Send(&operatorv1.SubscribeTasksRequest{
		Cookie: &operatorv1.SessionCookie{
			Value: operatorConn.metadata.cookie,
		},
		Type: &operatorv1.SubscribeTasksRequest_Start{
			Start: &operatorv1.StartPollBeaconRequest{
				Bid: b.GetId(),
			},
		},
	}); err != nil {
		return errors.Wrapf(err, "poll tasks for beacon %s", b.GetIdHex())
	}
	return nil
}

// стоп подписки на получение тасков для определенного бикона
func UnpollBeaconTasks(b *beacon.Beacon) error {
	if err := operatorConn.ss.tasksStream.Send(&operatorv1.SubscribeTasksRequest{
		Cookie: &operatorv1.SessionCookie{
			Value: operatorConn.metadata.cookie,
		},
		Type: &operatorv1.SubscribeTasksRequest_Stop{
			Stop: &operatorv1.StopPollBeaconRequest{
				Bid: b.GetId(),
			},
		},
	}); err != nil {
		return errors.Wrapf(err, "unpoll tasks for beacon %s", b.GetIdHex())
	}
	return nil
}

// открытие новой таск группы
func NewTaskGroup(id uint32, cmd string, visible bool) error {
	stream, err := getSvc().NewGroup(context.Background())
	if err != nil {
		return errors.Wrap(err, "open task group submition stream")
	}
	// создаем новую таск группу
	if err = stream.Send(&operatorv1.NewGroupRequest{
		Cookie: &operatorv1.SessionCookie{
			Value: operatorConn.metadata.cookie,
		},
		Type: &operatorv1.NewGroupRequest_Group{
			Group: &operatorv1.NewTaskGroupRequest{
				Bid:     id,
				Cmd:     cmd,
				Visible: visible,
			},
		},
	}); err != nil {
		return errors.Wrap(err, "open task group")
	}
	// сохраняем стрим
	operatorConn.ss.groupStreams.Store(id, stream)
	return nil
}

// закрытие таск группы
func CloseTaskGroup(id uint32) error {
	stream, ok := operatorConn.ss.groupStreams.Load(id)
	if !ok {
		return fmt.Errorf("unable load stream for beacon %d", id)
	}
	defer func() {
		// удаление стрима из мапы по ID бикона
		operatorConn.ss.groupStreams.Delete(id)
	}()
	if _, err := stream.CloseAndRecv(); err != nil {
		if !errors.Is(err, io.EOF) {
			return errors.Wrap(err, "close stream")
		}
	}
	return nil
}

// отправка сообщения в таск группу
func NewTaskGroupMessage(id uint32, tm defaults.TaskMessage, message string) error {
	stream, ok := operatorConn.ss.groupStreams.Load(id)
	if !ok {
		return fmt.Errorf("unable load stream for beacon %d", id)
	}
	return stream.Send(&operatorv1.NewGroupRequest{
		Cookie: &operatorv1.SessionCookie{
			Value: operatorConn.metadata.cookie,
		},
		Type: &operatorv1.NewGroupRequest_Message{
			Message: &operatorv1.NewTaskMessageRequest{
				Type: uint32(tm),
				Msg:  message,
			},
		},
	})
}

// создание нового таска
func NewTask(id uint32, v *operatorv1.NewTaskRequest) error {
	stream, ok := operatorConn.ss.groupStreams.Load(id)
	if !ok {
		return fmt.Errorf("unable load stream for beacon %d", id)
	}
	return stream.Send(&operatorv1.NewGroupRequest{
		Cookie: &operatorv1.SessionCookie{
			Value: operatorConn.metadata.cookie,
		},
		Type: &operatorv1.NewGroupRequest_Task{
			Task: v,
		},
	})
}
