package workerpool

import (
	"sync"
	"fmt"
	"time"
	"goload/metrics"
	"goload/models"
)

type Result struct {
	JobID int
	Data  interface{}
	Elapsed time.Duration
}

func worker(id int, jobs <-chan models.Job, events chan<- models.Event, wg *sync.WaitGroup, fn models.WorkerFunc) {
    defer wg.Done()
    for job := range jobs {
        _, _, err, elapsed := fn(job)

        //events <- metrics.Event{Type: metrics.EventLatency, JobID: job.ID, Latency: elapsed}
		
        if err != nil {
            events <- models.Event{Type: metrics.EventError, JobID: job.ID, Err: err, Latency: elapsed}
        } else {
            events <- models.Event{Type: metrics.EventCompleted, JobID: job.ID, Latency: elapsed}
        }
    }
}

func collectResults(results <-chan Result, wg *sync.WaitGroup){
	defer wg.Done()
	for result := range results {
		fmt.Printf("Result: %+v\n", result)
	}
}

func Dispatcher(jobCount, workerCount int, fn models.WorkerFunc, job models.Job) {
    jobs := make(chan models.Job, jobCount)
    events := make(chan models.Event, jobCount)
    done := make(chan models.Metrics)

    var wg sync.WaitGroup
    wg.Add(workerCount)
    for w := 1; w <= workerCount; w++ {
        go worker(w, jobs, events, &wg, fn)
    }

    go metrics.Collector(events, done)

    for j := 1; j <= jobCount; j++ {
        jobs <- models.Job{ID: j, API: job.API, Iter: job.Iter}
    }
    close(jobs)

    wg.Wait()
    close(events)

    metrics := <-done
    fmt.Printf("Time: %s, Completed: %d, Errors: %d, Avg Latency: %v\n",
        time.Now().Format(time.RFC3339),
        metrics.Completed, metrics.Errors,
        time.Duration(int64(metrics.TotalLatency)/int64(metrics.Completed + metrics.Errors)))
}