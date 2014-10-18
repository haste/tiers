package queue

import (
	"log"

	"tiers/model"
	"tiers/ocr"
	"time"
)

var Queue = make(chan bool, 1)

func ProcessQueue() {
	for {
		<-Queue
		log.Println("Queue: Processing.")

		queue := model.GetPendingQueues()

		for queue.Next() {
			var start = time.Now()

			var id, user_id, timestamp int
			var file string
			var processed bool

			if err := queue.Scan(&id, &user_id, &timestamp, &file, &processed); err != nil {
				log.Fatal(err)
			}

			p := ocr.OCR(file)
			model.InsertProfile(user_id, timestamp, p)

			processTime := time.Now().Sub(start).Nanoseconds() / 1e6
			model.SetQueueProcessed(id, processTime)

			log.Printf("Queue: Entry processed in %dms: %s L%d %dAP", processTime, p.Nick, p.Level, p.AP)
		}

		log.Println("Queue: Done.")

		queue.Close()
	}
}
