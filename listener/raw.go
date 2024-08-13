package listener

type RawMessage struct {
	Object string     `json:"object"`
	Entry  []rawEntry `json:"entry"`
}

type rawEntry struct {
	ID      string      `json:"id"`
	Changes []rawChange `json:"changes"`
}

type rawChange struct {
	Value rawValue `json:"value"`
	Field string   `json:"field"`
}

type rawValue struct {
	MessagingProduct string              `json:"messaging_product"`
	Metadata         rawMetadata         `json:"metadata"`
	Contacts         []rawContact        `json:"contacts"`
	Messages         []rawMessageContent `json:"messages"`
	Statuses         []rawStatus         `json:"statuses,omitempty"`
}

type rawMetadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberID      string `json:"phone_number_id"`
}

type rawContact struct {
	Profile rawProfile `json:"profile"`
	WaID    string     `json:"wa_id"`
}

type rawProfile struct {
	Name string `json:"name"`
}

type rawMessageContent struct {
	From        string               `json:"from"`
	ID          string               `json:"id"`
	Timestamp   string               `json:"timestamp"`
	Type        string               `json:"type"`
	Text        *rawText             `json:"text,omitempty"`
	Image       *rawImage            `json:"image,omitempty"`
	Sticker     *rawSticker          `json:"sticker,omitempty"`
	Location    *rawLocation         `json:"location,omitempty"`
	Contacts    *[]rawMessageContact `json:"contacts,omitempty"`
	Reaction    *rawReaction         `json:"reaction,omitempty"`
	Interactive *rawInteractive      `json:"interactive,omitempty"`
	Referral    *rawReferral         `json:"referral,omitempty"`
	Order       *rawOrder            `json:"order,omitempty"`
	System      *rawSystem           `json:"system,omitempty"`
	Errors      *[]rawError          `json:"errors,omitempty"`
	Context     *rawContext          `json:"context,omitempty"`
}

type rawText struct {
	Body string `json:"body"`
}

type rawImage struct {
	Caption  string `json:"caption"`
	MimeType string `json:"mime_type"`
	Sha256   string `json:"sha256"`
	ID       string `json:"id"`
}

type rawSticker struct {
	MimeType string `json:"mime_type"`
	Sha256   string `json:"sha256"`
	ID       string `json:"id"`
}

type rawLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Name      string  `json:"name"`
	Address   string  `json:"address"`
}

type rawMessageContact struct {
	Addresses []rawAddress `json:"addresses,omitempty"`
	Birthday  string       `json:"birthday,omitempty"`
	Emails    []rawEmail   `json:"emails,omitempty"`
	Name      rawName      `json:"name,omitempty"`
	Org       rawOrg       `json:"org,omitempty"`
	Phones    []rawPhone   `json:"phones,omitempty"`
	Urls      []rawUrl     `json:"urls,omitempty"`
}

type rawAddress struct {
	City        string `json:"city"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	Street      string `json:"street"`
	Type        string `json:"type"`
	Zip         string `json:"zip"`
}

type rawEmail struct {
	Email string `json:"email"`
	Type  string `json:"type"`
}

type rawName struct {
	FormattedName string `json:"formatted_name"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	MiddleName    string `json:"middle_name"`
	Suffix        string `json:"suffix"`
	Prefix        string `json:"prefix"`
}

type rawOrg struct {
	Company    string `json:"company"`
	Department string `json:"department"`
	Title      string `json:"title"`
}

type rawPhone struct {
	Phone string `json:"phone"`
	WaID  string `json:"wa_id"`
	Type  string `json:"type"`
}

type rawUrl struct {
	Url  string `json:"url"`
	Type string `json:"type"`
}

type rawReaction struct {
	MessageID string `json:"message_id"`
	Emoji     string `json:"emoji"`
}

type rawInteractive struct {
	Type        string          `json:"type"`
	ListReply   *rawListReply   `json:"list_reply,omitempty"`
	ButtonReply *rawButtonReply `json:"button_reply,omitempty"`
}

type rawListReply struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type rawButtonReply struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type rawReferral struct {
	SourceURL    string `json:"source_url"`
	SourceID     string `json:"source_id"`
	SourceType   string `json:"source_type"`
	Headline     string `json:"headline"`
	Body         string `json:"body"`
	MediaType    string `json:"media_type"`
	ImageURL     string `json:"image_url,omitempty"`
	VideoURL     string `json:"video_url,omitempty"`
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	CtwaClid     string `json:"ctwa_clid"`
}

type rawOrder struct {
	CatalogID    string           `json:"catalog_id"`
	ProductItems []rawProductItem `json:"product_items"`
	Text         string           `json:"text"`
}

type rawProductItem struct {
	ProductRetailerID string `json:"product_retailer_id"`
	Quantity          string `json:"quantity"`
	ItemPrice         string `json:"item_price"`
	Currency          string `json:"currency"`
}

type rawSystem struct {
	Body    string `json:"body"`
	NewWaID string `json:"new_wa_id"`
	Type    string `json:"type"`
}

type rawError struct {
	Code    int    `json:"code"`
	Details string `json:"details"`
	Title   string `json:"title"`
}

type rawContext struct {
	From            string              `json:"from"`
	ID              string              `json:"id"`
	ReferredProduct *rawReferredProduct `json:"referred_product,omitempty"`
}

type rawReferredProduct struct {
	CatalogID         string `json:"catalog_id"`
	ProductRetailerID string `json:"product_retailer_id"`
}

type rawStatus struct {
	ID           string           `json:"id"`
	Status       string           `json:"status"`
	Timestamp    string           `json:"timestamp"`
	RecipientID  string           `json:"recipient_id"`
	Conversation *rawConversation `json:"conversation,omitempty"`
	Pricing      *rawPricing      `json:"pricing,omitempty"`
	Errors       *[]rawError      `json:"errors,omitempty"`
}

type rawConversation struct {
	ID                  string    `json:"id"`
	ExpirationTimestamp string    `json:"expiration_timestamp,omitempty"`
	Origin              rawOrigin `json:"origin"`
}

type rawOrigin struct {
	Type string `json:"type"`
}

type rawPricing struct {
	Billable     bool   `json:"billable"`
	PricingModel string `json:"pricing_model"`
	Category     string `json:"category"`
}
