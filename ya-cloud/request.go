package main

type Request struct {
	Meta *struct {
		ClientId   string `json:"client_id"`
		Interfaces struct {
			AccountLinking struct {
			} `json:"account_linking"`
			Payments struct {
			} `json:"payments"`
			Screen struct {
			} `json:"screen"`
		} `json:"interfaces"`
		Locale   string `json:"locale"`
		Timezone string `json:"timezone"`
	} `json:"meta"`
	Request *struct {
		OriginalUtterance string `json:"original_utterance"`
		Command           string `json:"command"`
		Nlu               struct {
			Entities []interface{} `json:"entities"`
			Tokens   []interface{} `json:"tokens"`
			Intents  struct {
			} `json:"intents"`
		} `json:"nlu"`
		Markup struct {
			DangerousContext bool `json:"dangerous_context"`
		} `json:"markup"`
		Type string `json:"type"`
	} `json:"request"`
	Session *struct {
		MessageId int    `json:"message_id"`
		New       bool   `json:"new"`
		SessionId string `json:"session_id"`
		SkillId   string `json:"skill_id"`
		UserId    string `json:"user_id"`
		User      struct {
			UserId string `json:"user_id"`
		} `json:"user"`
		Application struct {
			ApplicationId string `json:"application_id"`
		} `json:"application"`
	} `json:"session"`
	State *struct {
		Session struct {
		} `json:"session"`
		User struct {
		} `json:"user"`
		Application struct {
		} `json:"application"`
	} `json:"state"`
	Version string `json:"version"`
}
