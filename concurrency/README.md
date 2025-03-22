````mermaid
graph TD;
    A[Start] -->|Check if list is empty| B{List Empty?}
    B -- Yes --> C[Return Error]
    B -- No --> D[Initialize variables]
    
    D --> E[Create job & result channels]
    E --> F[Create WaitGroup]
    F --> G[Start worker goroutines]
    
    subgraph Workers
        G1[Worker 1: Read job from channel] --> G2[Perform Reduction]
        G2 --> G3[Send result to result channel]
        
        H1[Worker 2: Read job from channel] --> H2[Perform Reduction]
        H2 --> H3[Send result to result channel]
        
        I1[Worker N: Read job from channel] --> I2[Perform Reduction]
        I2 --> I3[Send result to result channel]
    end
    
    G --> G1
    G --> H1
    G --> I1

    G1 --> G2
    H1 --> H2
    I1 --> I2

    G3 --> K[Wait for all workers]
    H3 --> K
    I3 --> K

    K --> L[Close result channel]
    L --> M[Final Reduction]
    M --> N[Return final result]

`````