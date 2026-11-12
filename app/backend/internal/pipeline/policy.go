package pipeline

type WorkerPoolConfig struct {
	MaxWorkers    int // batas atas
	MinWorkers    int // minimal worker kalau job > 1
	JobsPerWorker int // rasio job per worker untuk kasus medium+
}

func CalculateWorkerCount(jobCount int, cfg WorkerPoolConfig) int {
	switch {
	case jobCount <= 0:
		return 1
	case jobCount <= cfg.MinWorkers:
		// job sedikit, tapi minimal MinWorkers (kalau jobCount >= MinWorkers)
		return jobCount
	default:
		// medium ke atas: bagi rata, tapi dibatasi MaxWorkers
		workers := jobCount / cfg.JobsPerWorker
		if workers < cfg.MinWorkers {
			workers = cfg.MinWorkers
		}
		if workers > cfg.MaxWorkers {
			workers = cfg.MaxWorkers
		}
		return workers
	}
}
