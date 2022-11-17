# Finna <!-- omit in toc -->

# Import Vanguard Transactions to Stock Events

## Download transaction history
Visit the Download Center in Vanguard

https://personal.vanguard.com/us/OfxWelcome

![download-center.png](images/download-center.png)

## Run `format` task to fix up data

```bash
➜ task format
task: [format] go run cmd/finna/main.go format -i input.csv -o output.csv
{"level":"info","msg":"Running finna","time":"2022-11-17T12:01:45-06:00","version":"latest"}
{"args":{"OutputFormat":"stock-events","InputFile":"input.csv","OutputFile":"output.csv"},"level":"info","msg":"Running format command","time":"2022-11-17T12:01:45-06:00"}
{"level":"info","msg":"Found 133 CSV rows","time":"2022-11-17T12:01:45-06:00"}
```

This task will filter out transactions that don't affect holdings, sort by date, and dump into a CSV format that is required by Stock Events.

![stock-events-instructions.jpeg](images/stock-events-instructions.jpeg)

Note that duplicate transactions will be ignored, so this operation should be safe to run as many times as necessary.

## Send output file to iPhone
Using iCloud, send the output file to your phone and import it from the Stock Events App.
