package parserFile

import (
	"encoding/csv"
	"fmt"
	"io"
	"microservice-process-file-go/pkg/models"
	"os"
	"strconv"
	"sync"
)

func ParserFile(filePath string) []*models.Sales {
	f1, _ := os.Open(filePath)
	defer func(f1 *os.File) {
		err := f1.Close()
		if err != nil {
			fmt.Println("ERROR: ", err.Error())
		}
	}(f1)
	return concuRSwWP(f1)
}

// with Worker pools
func concuRSwWP(f *os.File) []*models.Sales {
	fcsv := csv.NewReader(f)
	rs := make([]*models.Sales, 0)
	numWps := 100
	jobs := make(chan []string, numWps)
	res := make(chan *models.Sales)

	var wg sync.WaitGroup
	worker := func(jobs <-chan []string, results chan<- *models.Sales) {
		for {
			select {
			case job, ok := <-jobs: // you must check for readable state of the channel.
				if !ok {
					return
				}
				results <- parseStruct(job)
			}
		}
	}

	// init workers
	for w := 0; w < numWps; w++ {
		wg.Add(1)
		go func() {
			// this line will exec when chan `res` processed output at line 107 (func worker: line 71)
			defer wg.Done()
			worker(jobs, res)
		}()
	}

	go func() {
		for {
			rStr, err := fcsv.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("ERROR: ", err.Error())
				break
			}
			jobs <- rStr
		}
		close(jobs) // close jobs to signal workers that no more job are incoming.
	}()

	go func() {
		wg.Wait()
		close(res) // when you close(res) it breaks the below loop.
	}()

	for r := range res {
		rs = append(rs, r)
	}

	fmt.Println("Total records: ", len(rs))
	return rs
}

func parseStruct(data []string) *models.Sales {
	id, _ := strconv.ParseInt(data[6], 10, 64)
	unitSold, _ := strconv.ParseInt(data[8], 10, 64)
	unitPrice, _ := strconv.ParseFloat(data[9], 64)
	unitCost, _ := strconv.ParseFloat(data[10], 64)
	totalRev, _ := strconv.ParseFloat(data[11], 64)
	totalCost, _ := strconv.ParseFloat(data[12], 64)
	totalProfit, _ := strconv.ParseFloat(data[13], 64)
	return &models.Sales{
		Region:        data[0],
		Country:       data[1],
		ItemType:      data[2],
		SaleChannel:   data[3],
		OrderPriority: data[4],
		OrderDate:     data[5],
		OrderId:       id,
		ShipDate:      data[7],
		UnitSold:      unitSold,
		UnitPrice:     unitPrice,
		UnitCost:      unitCost,
		TotalRevenue:  totalRev,
		TotalCost:     totalCost,
		TotalProfit:   totalProfit,
	}
}
