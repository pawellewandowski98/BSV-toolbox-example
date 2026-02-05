package monitor

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/defs"
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/monitor"
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/storage"
	"github.com/go-softwarelab/common/pkg/must"
)

func RunMonitor(
	ctx context.Context,
	storage *storage.Provider,
	onTxBroadcasted, onTxProven chan defs.TransactionStatusUpdate,
	logger *slog.Logger,
) (daemon *monitor.Daemon, cleanup func(), err error) {

	opts := []monitor.DaemonEventOption{
		monitor.WithBroadcastedTxChannel(onTxBroadcasted),
		monitor.WithProvenTxChannel(onTxProven),
	}

	daemon, err = monitor.NewDaemonWithGORMLocker(ctx, logger, storage, storage.Database.DB, opts...)
	if err != nil {
		return nil, func() {
			close(onTxBroadcasted)
			close(onTxProven)
		}, fmt.Errorf("failed to create daemon: %w", err)
	}

	monitorConfig := defs.Monitor{
		Tasks: defs.TasksConfig{
			CheckForProofs: defs.TaskConfig{
				Enabled:         true,
				IntervalSeconds: must.ConvertToUInt((1 * time.Minute).Seconds()),
			},
			SendWaiting: defs.TaskConfig{
				Enabled:          true,
				IntervalSeconds:  must.ConvertToUInt((5 * time.Minute).Seconds()),
				StartImmediately: true,
			},
			FailAbandoned: defs.TaskConfig{
				Enabled:         true,
				IntervalSeconds: must.ConvertToUInt((5 * time.Minute).Seconds()),
			},
			UnFail: defs.TaskConfig{
				Enabled:         true,
				IntervalSeconds: must.ConvertToUInt((10 * time.Minute).Seconds()),
			},
		},
	}

	// Run consumers
	go consumeTxBroadcasted(onTxBroadcasted, logger)
	go consumeTxProven(onTxProven, logger)

	if err = daemon.Start(monitorConfig.Tasks.EnabledTasks()); err != nil {
		return nil, func() {
			close(onTxBroadcasted)
			close(onTxProven)
		}, fmt.Errorf("failed to start storage monitor: %w", err)
	}

	return daemon, func() {
		close(onTxBroadcasted)
		close(onTxProven)
		_ = daemon.Stop()
	}, nil
}

func PrepareChannels(capacity int) (txBroadcastedCh, txProvenCh chan defs.TransactionStatusUpdate) {
	txBroadcastedCh = make(chan defs.TransactionStatusUpdate, capacity)
	txProvenCh = make(chan defs.TransactionStatusUpdate, capacity)

	return
}

func consumeTxBroadcasted(channel chan defs.TransactionStatusUpdate, logger *slog.Logger) {
	for msg := range channel {
		fmt.Println("<---------------------------- BROADCASTED")
		fmt.Println("MSG: ", msg)
		logger.Info(
			"new tx broadcasted",
			slog.String("tx_id", msg.TxID),
			slog.String("status", msg.Status.String()),
		)
	}
}

func consumeTxProven(channel chan defs.TransactionStatusUpdate, logger *slog.Logger) {
	for msg := range channel {
		fmt.Println("<---------------------------- PROVEN")
		fmt.Println("MSG: ", msg)
		logger.Info(
			"new tx proven",
			slog.String("tx_id", msg.TxID),
			slog.String("status", msg.Status.String()),
		)
	}
}
