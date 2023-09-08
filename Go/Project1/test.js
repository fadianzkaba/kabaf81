mermaid
sequenceDiagram
    autonumber
    Cloud Scheduler->>PubSub: Schedule
    PubSub->>+Dataflow Job: Trigger
    loop
        Dataflow Job->>+DB: Read
        DB-->>-Dataflow Job: <Data>
    end
    Dataflow Job->>-GCS: Write
    opt
        GCS->>DB: Restore
    end