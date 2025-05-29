package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/invopop/jsonschema" // required go1.18+

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
)

// A struct that will be converted to a Structured Outputs response schema
type HistoricalComputer struct {
	Origin       Origin   `json:"origin" jsonschema_description:"The origin of the computer"`
	Name         string   `json:"full_name" jsonschema_description:"The name of the device model"`
	Legacy       string   `json:"legacy" jsonschema:"enum=positive,enum=neutral,enum=negative" jsonschema_description:"Its influence on the field of computing"`
	NotableFacts []string `json:"notable_facts" jsonschema_description:"A few key facts about the computer"`
}

type Origin struct {
	YearBuilt    int64  `json:"year_of_construction" jsonschema_description:"The year it was made"`
	Organization string `json:"organization" jsonschema_description:"The organization that was in charge of its development"`
}

func GenerateSchema[T any]() interface{} {
	// Structured Outputs uses a subset of JSON schema
	// These flags are necessary to comply with the subset
	reflector := jsonschema.Reflector{
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}
	var v T
	schema := reflector.Reflect(v)
	return schema
}

// Generate the JSON schema at initialization time
var HistoricalComputerResponseSchema = GenerateSchema[HistoricalComputer]()

func main() {
	client := arkruntime.NewClientWithApiKey(os.Getenv("ARK_API_KEY"))
	ctx := context.Background()

	question := "What computer ran the first neural network?"

	print("> ")
	println(question)

	schemaParam := model.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        "biography",
		Description: "Notable information about a person",
		Schema:      HistoricalComputerResponseSchema,
		Strict:      true,
	}

	req := model.CreateChatCompletionRequest{
		Model: "${YOUR_ENDPOINT_ID}",
		Messages: []*model.ChatCompletionMessage{
			{
				Role: model.ChatMessageRoleUser,
				Content: &model.ChatCompletionMessageContent{
					StringValue: volcengine.String(question),
				},
			},
		},
		ResponseFormat: &model.ResponseFormat{
			Type:       model.ResponseFormatJSONSchema,
			JSONSchema: schemaParam,
		},
	}
	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		fmt.Printf("structured output chat error: %v\n", err)
		return
	}

	// The model responds with a JSON string, so parse it into a struct
	historicalComputer := HistoricalComputer{}
	err = json.Unmarshal([]byte(*resp.Choices[0].Message.Content.StringValue), &historicalComputer)
	if err != nil {
		panic(err.Error())
	}

	// Use the model's structured response with a native Go struct
	fmt.Printf("Name: %v\n", historicalComputer.Name)
	fmt.Printf("Year: %v\n", historicalComputer.Origin.YearBuilt)
	fmt.Printf("Org: %v\n", historicalComputer.Origin.Organization)
	fmt.Printf("Legacy: %v\n", historicalComputer.Legacy)
	fmt.Printf("Facts:\n")
	for i, fact := range historicalComputer.NotableFacts {
		fmt.Printf("%v. %v\n", i+1, fact)
	}
}
