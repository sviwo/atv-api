package amqp

import (
	"context"
	"fmt"
	"pack.ag/amqp"
	"time"
)

type AmqpManager struct {
	Address  string
	UserName string
	Password string
	Client   *amqp.Client
	Session  *amqp.Session
	Receiver *amqp.Receiver
}

// 业务函数。用户自定义实现，该函数被异步执行，请考虑系统资源消耗情况。
func (am *AmqpManager) processMessage(message *amqp.Message) {
	fmt.Println("data received:", string(message.GetData()), " properties:", message.ApplicationProperties)
}

func (am *AmqpManager) StartReceiveMessage(ctx context.Context) {

	childCtx, _ := context.WithCancel(ctx)
	err := am.generateReceiverWithRetry(childCtx)
	if nil != err {
		return
	}
	defer func() {
		am.Receiver.Close(childCtx)
		am.Session.Close(childCtx)
		am.Client.Close()
	}()

	for {
		//阻塞接受消息，如果ctx是background则不会被打断。
		message, err := am.Receiver.Receive(ctx)

		if nil == err {
			go am.processMessage(message)
			message.Accept()
		} else {
			fmt.Println("amqp receive data error:", err)

			//如果是主动取消，则退出程序。
			select {
			case <-childCtx.Done():
				return
			default:
			}

			//非主动取消，则重新建立连接。
			err := am.generateReceiverWithRetry(childCtx)
			if nil != err {
				return
			}
		}
	}
}

func (am *AmqpManager) generateReceiverWithRetry(ctx context.Context) error {
	//退避重连，从10ms依次x2，直到20s。
	duration := 10 * time.Millisecond
	maxDuration := 20000 * time.Millisecond
	times := 1

	//异常情况，退避重连。
	for {
		select {
		case <-ctx.Done():
			return amqp.ErrConnClosed
		default:
		}

		err := am.generateReceiver()
		if nil != err {
			time.Sleep(duration)
			if duration < maxDuration {
				duration *= 2
			}
			fmt.Println("amqp connect retry,times:", times, ",duration:", duration)
			times++
		} else {
			fmt.Println("amqp connect init success")
			return nil
		}
	}
}

// 由于包不可见，无法判断Connection和Session状态，重启连接获取。
func (am *AmqpManager) generateReceiver() error {

	if am.Session != nil {
		receiver, err := am.Session.NewReceiver(
			amqp.LinkSourceAddress("/queue-name"),
			amqp.LinkCredit(20),
		)
		//如果断网等行为发生，Connection会关闭导致Session建立失败，未关闭连接则建立成功。
		if err == nil {
			am.Receiver = receiver
			return nil
		}
	}

	//清理上一个连接。
	if am.Client != nil {
		am.Client.Close()
	}

	client, err := amqp.Dial(am.Address, amqp.ConnSASLPlain(am.UserName, am.Password))
	if err != nil {
		return err
	}
	am.Client = client

	session, err := client.NewSession()
	if err != nil {
		return err
	}
	am.Session = session

	receiver, err := am.Session.NewReceiver(
		amqp.LinkSourceAddress("/queue-name"),
		amqp.LinkCredit(20),
	)
	if err != nil {
		return err
	}
	am.Receiver = receiver

	return nil
}
