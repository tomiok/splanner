# SPLANNER

---

#### A really nice pool worker

---

#### Config

* In your main file or something like that

```go
// init the shared data structure.
splanner.InitQueue(20)

// init the dispatcher & keep it listening.
splanner.NewDispatcher(15).Run(true)
```

* The caller should have something similar...
```go

type HeavyWork struct {
    Name   string
    number int
}

// implement the job method, and heavyWork is now implementing the Unit interface
func (p *HeavyWork) Job() error {
    time.Sleep(500 * time.Millisecond)
    fmt.Println(fmt.Sprintf("heavy job is running %d", p.number))
    return nil
}

work := HeavyWork{Name: q, number: a}
// add the work here, could be in a loop to add more than one or whatever you want
splanner.AddUnit(&work)
```