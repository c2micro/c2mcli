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
	msg, err := operatorConn.controlStream.Recv()
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
		if _, err := operatorConn.controlStream.Recv(); err != nil {
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
