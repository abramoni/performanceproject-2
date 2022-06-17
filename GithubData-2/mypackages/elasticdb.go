package mypackages

import (
	"context"
	"encoding/json"
	"fmt"

	elastic "github.com/olivere/elastic/v7"
)

func GetESClient() (*elastic.Client, error) {

	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")

	return client, err

}

func StoreInElasticDb(body []byte, IndexName string) {
	ctx := context.Background()
	esclient, err := GetESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}
	//json.Unmarshal(body, &responsedata)

	//dataJSON, err := json.Marshal(responsedata)
	js := string(body)
	ind, err := esclient.Index().
		Index(IndexName).
		BodyJson(js).
		Do(ctx)
	_ = ind
	if err != nil {
		panic(err)
	}

	fmt.Println("[Elastic][InsertProduct]Insertion Successful")
}

func RetriveFronElasticDb(IndexName string) {
	ctx := context.Background()
	esclient, err := GetESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}

	searchSource := elastic.NewSearchSource()
	//searchSource.Query(elastic.NewMatchQuery("metricName", "kCreateFileOps"))

	/* this block will basically print out the es query */
	//queryStr, err1 := searchSource.Source()
	//queryJs, err2 := json.Marshal(queryStr)

	/*if err1 != nil || err2 != nil {
		fmt.Println("[esclient][GetResponse]err during query marshal=", err1, err2)
	}
	fmt.Println("[esclient]Final ESQuery=\n", string(queryJs))
	/* until this block */

	searchService := esclient.Search().Index(IndexName).SearchSource(searchSource)

	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("[ProductsES][GetPIds]Error=", err)
		return
	}

	for _, hit := range searchResult.Hits.Hits {

		err := json.Unmarshal(hit.Source, &responsedata)
		if err != nil {
			fmt.Println("[Getting Students][Unmarshal] Err=", err)
		}

	}

	if err != nil {
		fmt.Println("Fetching student fail: ", err)
	} else {
		fmt.Println("metric name = ", responsedata.Name)
		fmt.Println("instances = ", len(responsedata.Datavec))
		for i := 0; i < len(responsedata.Datavec); i++ {
			fmt.Print("key = ", responsedata.Datavec[i].Time)
			fmt.Println("    value = ", responsedata.Datavec[i].Value.Data)
		}
	}
}
