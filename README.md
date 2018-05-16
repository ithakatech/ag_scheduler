# agscheduler
Advanced Golang scheduler
A full featured task scheduler built in golang. Exposes a simple to use API for scheduling tasks that run at a specific time or at regular intervals or for a certain number of repetitions.

Some potential use cases:

```golang
scheduler := NewScheduler()

func task() {
    // ... do something
}

scheduler.After(3).Seconds().Run(task)
scheduler.After(10).Milliseconds().Run(task)

scheduler.At(datetime).Run(task)

scheduler.Every(5).Minutes().Run(task)
scheduler.Every("Thursday").At(time).Run(task)
scheduler.Each().Month().On(date).At(time).Run(task)

scheduler.Upcoming("Sunday").At(time).Run(task)
```

Each invocation of the task happens on a separate goroutine.

## TODO:
- Adding more time units to **After** style scheduling
- Building **At** style scheduling
- Building **Every** style scheduling
- Adding support for disk persistence and fault tolerance
- Network interface which allows building and running the library as a standalone server (TBD)