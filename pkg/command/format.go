package command

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

type OutputType string

const (
	FormatFlagOutputFormat  = "format"
	FormatFlagInputFile     = "input-file"
	FormatFlagOutputFile    = "output-file"
	OutputFormatStockEvents = "stock-events"
)

type FormatCommand struct{}

func NewFormatCommand() *FormatCommand {
	return &FormatCommand{}
}

type FormatCommandArgs struct {
	OutputFormat string
	InputFile    string
	OutputFile   string
}

type VanguardTransaction struct {
	AccountNumber          string
	TradeDate              time.Time
	SettlementDate         time.Time
	TransactionType        string
	TransactionDescription string
	InvestmentName         string
	Symbol                 string
	Shares                 float64
	SharePrice             float64
	PrinicpalAmount        float64
	CommissionFees         float64
	NetAmount              float64
	AccruedInterest        float64
	AccountType            string
}

func (v *VanguardTransaction) ToStockEvent() StockEvent {
	// If the transaction is a sell, we need to flip the sign of the quantity
	quantity := v.Shares
	if strings.Contains(strings.ToLower(v.TransactionType), "sell") && quantity > 0 {
		quantity = quantity * -1
	}
	return StockEvent{
		Date:     v.TradeDate,
		Symbol:   v.Symbol,
		Quantity: quantity,
		Price:    v.SharePrice,
	}
}

type VanguardTransactions struct {
	Transactions []VanguardTransaction
}

func (t *VanguardTransactions) ToStockEvents() StockEvents {

	transactionTypeKeywords := []string{
		"buy",
		"sell",
		"reinvestment",
	}

	stockEvents := StockEvents{Events: []StockEvent{}}
	for _, transaction := range t.Transactions {
		// We only care about transactions that change holdings
		if transaction.Symbol == "" {
			logrus.WithField("transaction", transaction).Debug("Skipping transaction with no symbol")
			continue
		}

		shouldInclude := false
		for _, transactionType := range transactionTypeKeywords {
			if strings.Contains(strings.ToLower(transaction.TransactionType), strings.ToLower(transactionType)) {
				shouldInclude = true
			}
		}
		if !shouldInclude {
			logrus.WithField("transaction", transaction).Debug("Skipping transaction that doesn't change holdings")
			continue
		}

		stockEvents.Events = append(stockEvents.Events, transaction.ToStockEvent())
	}
	return stockEvents
}

type StockEvent struct {
	Symbol   string
	Date     time.Time
	Quantity float64
	Price    float64
}

type StockEvents struct {
	Events []StockEvent
}

func (s *StockEvents) ToCSVSlice() [][]string {
	csvSlice := [][]string{
		{"Symbol", "Date", "Quantity", "Price"},
	}
	for _, event := range s.Events {
		csvSlice = append(csvSlice, []string{
			event.Symbol,
			event.Date.Format("2006-01-02"),
			strconv.FormatFloat(event.Quantity, 'f', -1, 64),
			strconv.FormatFloat(event.Price, 'f', 2, 64),
		})
	}
	return csvSlice
}

func (c *FormatCommand) ToCliCommand() *cli.Command {
	args := FormatCommandArgs{}
	return &cli.Command{
		Name:  "format",
		Usage: "Format transaction data from Vanguard to be consumed elsewhere",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        FormatFlagOutputFormat,
				Aliases:     []string{"f"},
				Usage:       "Output format",
				Value:       OutputFormatStockEvents,
				Destination: &args.OutputFormat,
			},
			&cli.PathFlag{
				Name:        FormatFlagInputFile,
				Aliases:     []string{"i"},
				Usage:       "Input Vanguard CSV file",
				Destination: &args.InputFile,
				Required:    true,
			},
			&cli.PathFlag{
				Name:        FormatFlagOutputFile,
				Aliases:     []string{"o"},
				Usage:       "Output CSV file",
				Destination: &args.OutputFile,
				Value:       "output.csv",
			},
		},
		Action: func(ctx *cli.Context) error {
			logrus.WithField("args", args).Info("Running format command")
			transactions, err := readCSVToVanguardRecords(args.InputFile)
			if err != nil {
				return err
			}
			var csvData [][]string
			switch args.OutputFormat {
			case OutputFormatStockEvents:
				stockEvents := transactions.ToStockEvents()
				csvData = stockEvents.ToCSVSlice()
			default:
				return fmt.Errorf("unknown output format: %s", args.OutputFormat)
			}
			return writeCSV(args.OutputFile, csvData)
		},
	}
}

func readCSVToVanguardRecords(path string) (VanguardTransactions, error) {
	transactions := VanguardTransactions{Transactions: []VanguardTransaction{}}
	f, err := os.Open(path)
	if err != nil {
		return transactions, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.FieldsPerRecord = -1
	rows, err := csvReader.ReadAll()
	if err != nil {
		return transactions, err
	}
	logrus.Infof("Found %d CSV rows", len(rows))

	// Skip to the start of the transaction table
	headerRow := 0
	foundHeader := false
	for i, row := range rows {
		if row[0] == "Account Number" && row[1] == "Trade Date" {
			headerRow = i
			foundHeader = true
			break
		}
	}
	if !foundHeader {
		return transactions, fmt.Errorf("could not find header row in CSV")
	}

	// Parse transactions to a struct
	for _, row := range rows[headerRow+1:] {
		tradeDate, err := parseVanguardDate(row[1])
		if err != nil {
			return transactions, err
		}
		settlementDate, err := parseVanguardDate(row[2])
		if err != nil {
			return transactions, err
		}
		shares, err := parseFloat64(row[7])
		if err != nil {
			return transactions, err
		}
		sharePrice, err := parseFloat64(row[8])
		if err != nil {
			return transactions, err
		}
		principalAmount, err := parseFloat64(row[9])
		if err != nil {
			return transactions, err
		}
		commissionFees, err := parseFloat64(row[10])
		if err != nil {
			return transactions, err
		}
		netAmount, err := parseFloat64(row[11])
		if err != nil {
			return transactions, err
		}
		accruedInterest, err := parseFloat64(row[12])
		if err != nil {
			return transactions, err
		}
		transactions.Transactions = append(
			transactions.Transactions,
			VanguardTransaction{
				AccountNumber:          row[0],
				TradeDate:              tradeDate,
				SettlementDate:         settlementDate,
				TransactionType:        row[3],
				TransactionDescription: row[4],
				InvestmentName:         row[5],
				Symbol:                 row[6],
				Shares:                 shares,
				SharePrice:             sharePrice,
				PrinicpalAmount:        principalAmount,
				CommissionFees:         commissionFees,
				NetAmount:              netAmount,
				AccruedInterest:        accruedInterest,
				AccountType:            row[13],
			},
		)
	}

	// Sort transactions by trade date
	sort.Slice(
		transactions.Transactions, func(i, j int) bool {
			return transactions.Transactions[i].TradeDate.Before(transactions.Transactions[j].TradeDate)
		},
	)

	return transactions, nil
}

func parseVanguardDate(date string) (time.Time, error) {
	parsedTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		// Sometimes the report has a different date format
		parsedTime, err = time.Parse("01/02/2006", date)
	}
	return parsedTime, err
}

func parseFloat64(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}

func writeCSV(path string, data [][]string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	csvWriter := csv.NewWriter(f)
	_ = csvWriter.WriteAll(data)
	return csvWriter.Error()
}
