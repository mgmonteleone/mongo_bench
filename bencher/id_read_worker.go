package bencher

import (
	"log"
	"math/rand"
	"time"

	"github.com/pterm/pterm"
	"go.mongodb.org/mongo-driver/bson"
)

type IDReadWorker struct {
	bencher *Bencher
}

func StartIDReadWorker(bencher *Bencher) *IDReadWorker {
	worker := &IDReadWorker{
		bencher: bencher,
	}
	go worker.ReadThread()
	return worker
}

func (worker *IDReadWorker) ReadThread() {
	ticker := time.NewTicker(time.Duration(worker.bencher.statTickSpeedMillis) * time.Millisecond)
	numOps := 0
	totalTimeMicros := 0
	collection := worker.bencher.Collection()

	for {
		select {
		case <-ticker.C:
			worker.bencher.returnChannel <- FuncResult{
				numOps:     numOps,
				timeMicros: totalTimeMicros,
				opType:     "id_read",
			}
			numOps = 0
			totalTimeMicros = 0
		default:
			start := time.Now()
			workerId := rand.Intn(worker.bencher.numInsertWorkers)
			insertWorker := worker.bencher.workerMap[workerId]
			if insertWorker.lastId == 0 {
				pterm.Printfln("Waiting for insert worker to start before reading....")
				time.Sleep(1 * time.Second)
			} else {
				docId := rand.Intn(worker.bencher.workerMap[workerId].lastId) + 1 + (workerId * 100_000_000_000)
				doc := collection.FindOne(worker.bencher.ctx, bson.M{"_id": docId})
				if doc.Err() != nil {
					log.Fatal("Bad find...", doc.Err())
				}
				totalTimeMicros += int(time.Since(start).Microseconds())
				numOps++
			}
		}
	}
}