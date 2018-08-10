# vault

## Stateless secure password generator in go

[Scrypt](https://en.wikipedia.org/wiki/Scrypt)

[Master Password](https://en.wikipedia.org/wiki/Master_Password)

## TODO

- goji -> chi @DONE
- rename generator in master_password
- run in minikube

## Vegeta

```echo "POST http://0.0.0.0:3000/password" | vegeta attack -duration=5s -body=body.json -http2 -rate=5000 | tee results.bin | vegeta report```

1. Goji

    Requests      [total, rate]            25000, 5000.20
    Duration      [total, attack, wait]    5.000605241s, 4.9997995s, 805.741µs
    Latencies     [mean, 50, 95, 99, max]  14.086706ms, 2.84644ms, 89.886483ms, 195.154983ms, 285.813935ms
    Bytes In      [total, mean]            875000, 35.00
    Bytes Out     [total, mean]            2375000, 95.00
    Success       [ratio]                  100.00%
    Status Codes  [code:count]             200:25000  
    Error Set:

1. Chi

    Requests      [total, rate]            25000, 4999.88
    Duration      [total, attack, wait]    5.001429649s, 5.000119764s, 1.309885ms
    Latencies     [mean, 50, 95, 99, max]  10.524365ms, 829.808µs, 90.547445ms, 145.97906ms, 1.058563681s
    Bytes In      [total, mean]            875000, 35.00
    Bytes Out     [total, mean]            2375000, 95.00
    Success       [ratio]                  100.00%
    Status Codes  [code:count]             200:25000  
    Error Set:
