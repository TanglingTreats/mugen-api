// "challenges" route
package challenges

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/TanglingTreats/mugen-typer-api/dotenv"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type challenge struct {
	Data string `json:"data"`
}

var text = "metaverse web3 NFT crypto decentralized meme stock stonk hodl ape GameStop AMC Reddit Robinhood Dogecoin elon tesla Twitter Muskrat quiet quitting great resignation quiet firing layoff recession inflation cost of living supply chain chip shortage climate crisis heat wave drought fire season net zero green energy EV plant-based oat milk cauliflower gnocchi charcuterie grazing board cheugy cringe slay zaddy bussy thirst trap y'all cap no cap fr fr wig go off understood the assignment hot girl walk feral girl summer that's the tweet main character energy unalive sadfishing negging love-bombing gatekeeping cloutlighting sliving going goblin mode crisitunity ambient anxiety"

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

	chatPrompt.Content = "give me only a list of different top 100 popular words or phrases of 2023 separated by commas without indexes and quotes"
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
	res := challenge{
		Data: chatResultData,
	}

	render.JSON(w, r, res)
}
