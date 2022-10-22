package dto

type WebhookInput struct {
	Entry  []Entry `json:"entry"`
	Object string  `json:"object"`
}

type Entry struct {
	Changes []Change `json:"changes"`
	ID      string   `json:"id"`
}

type Change struct {
	Value Value  `json:"value"`
	Field string `json:"field"`
}

type Value struct {
	Contacts         []Contact `json:"contacts"`
	Messages         []Message `json:"messages"`
	Statuses         []Status  `json:"statuses"`
	Metadata         Metadata  `json:"metadata"`
	MessagingProduct string    `json:"messaging_product"`
}

type Contact struct {
	Profile Profile `json:"profile"`
	WaID    string  `json:"wa_id"`
}

type Profile struct {
	Name string `json:"name"`
}

type Message struct {
	Context     Context     `json:"context"`
	Text        Text        `json:"text"`
	Interactive Interactive `json:"interactive"`
	From        string      `json:"from"`
	ID          string      `json:"id"`
	Timestamp   string      `json:"timestamp"`
	Type        string      `json:"type"`
}

type Text struct {
	Body string `json:"body"`
}

type Metadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberID      string `json:"phone_number_id"`
}

type Status struct {
	Conversation Conversation `json:"conversation"`
	Pricing      Pricing      `json:"pricing"`
	ID           string       `json:"id"`
	Status       string       `json:"status"`
	Timestamp    string       `json:"timestamp"`
	RecipientID  string       `json:"recipient_id"`
	Errors       []Error      `json:"errors"`
}

type Conversation struct {
	Origin Origin `json:"origin"`
	ID     string `json:"id"`
}

type Origin struct {
	Type string `json:"type"`
}

type Pricing struct {
	Billable     bool   `json:"billable"`
	PricingModel string `json:"pricing_model"`
	Category     string `json:"category"`
}

type Context struct {
	From string `json:"from"`
	ID   string `json:"id"`
}

type Interactive struct {
	ListReply ListReply `json:"list_reply"`
	Type      string    `json:"type"`
}

type ListReply struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Error struct {
	Code  int64  `json:"code"`
	Title string `json:"title"`
	Href  string `json:"href"`
}
