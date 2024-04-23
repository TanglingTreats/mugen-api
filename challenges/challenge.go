// "challenges" route
package challenges

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/TanglingTreats/mugen-api/dotenv"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type challenge struct {
	Data string `json:"data"`
}

var chatEndpoint = "/chat/completions"

// "global" http client
var client = &http.Client{
	Timeout: 10 * time.Second,
}

var chatPromptReq = promptReq{Model: "gpt-3.5-turbo"}
var chatPrompt = prompt{Role: "user"}

func Routes(r chi.Router) {
	r.Get("/", getChallenge)
}

func getChallenge(w http.ResponseWriter, r *http.Request) {
	openaiUrl := dotenv.GetEnvVar("OPENAI_URL")

	chatPrompt.Content = "give me only a list of different top 50 popular words or phrases of 2023 separated by commas without indexes and quotes"
	promptMsgs := []prompt{chatPrompt}
	chatPromptReq.Messages = promptMsgs

	jsonBytes, jsonErr := json.Marshal(chatPromptReq)
	if jsonErr != nil {
		fmt.Println(jsonErr.Error())
	}

	// Create post request
	req, err := http.NewRequest("POST", openaiUrl+chatEndpoint, bytes.NewBuffer(jsonBytes))
	if err != nil {
		fmt.Println(err)
	}

	apiToken := dotenv.GetEnvVar("OPENAI_API_KEY")

	req.Header.Add("Authorization", "Bearer "+apiToken)
	req.Header.Add("Content-Type", "application/json")

	defer req.Body.Close()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	// Close connection
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)

	chatPromptRes := chatResponse{}
	if err := json.Unmarshal(data, &chatPromptRes); err != nil {
		fmt.Println(err.Error())
	}

	chatResultData := chatPromptRes.Choices[0].Msg.Content
	chatResultData = strings.ReplaceAll(chatResultData, ".", "")

	// Randomize function
	randomized := randomizeWords(chatResultData)

	fmt.Println(randomized)

	res := challenge{
		Data: randomized,
	}

	render.JSON(w, r, res)
}

func randomizeWords(input string) string {
	wordSlice := strings.Split(input, ", ")
	rand.Shuffle(len(wordSlice), func(i, j int) { wordSlice[i], wordSlice[j] = wordSlice[j], wordSlice[i] })

	return strings.Join(wordSlice, " ")
}
