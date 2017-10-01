package main

func initWorkers() {
	initWorkerReset()

	initWorkerStart()
	initWorkerIO()
	initWorkerEnd()
}

func shutdownWorkers() {
	close(workerResetJobs)
	workerResetJobswg.Wait()

	close(workerStartJobs)
	workerStartJobswg.Wait()
	workerIOJobswg.Wait()
	workerEndJobswg.Wait()
}
