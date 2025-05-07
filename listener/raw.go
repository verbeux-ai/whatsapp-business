package listener

type RawMessage struct {
	Object string     `json:"object"`
	Entry  []RawEntry `json:"entry"`
}

type RawEntry struct {
	ID      string      `json:"id"`
	Changes []RawChange `json:"changes"`
}

type RawChange struct {
	Value RawValue `json:"value"`
	Field string   `json:"field"`
}

type RawValue struct {
	MessagingProduct string              `json:"messaging_product"`
	Metadata         RawMetadata         `json:"metadata"`
	Contacts         []RawContact        `json:"contacts"`
	Messages         []RawMessageContent `json:"messages"`
	Statuses         []RawStatus         `json:"statuses,omitempty"`
}

type RawMetadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberID      string `json:"phone_number_id"`
}

type RawContact struct {
	Profile RawProfile `json:"profile"`
	WaID    string     `json:"wa_id"`
}

type RawProfile struct {
	Name string `json:"name"`
}

type RawMessageContent struct {
	From        string               `json:"from"`
	ID          string               `json:"id"`
	Timestamp   string               `json:"timestamp"`
	Type        string               `json:"type"`
	Text        *RawText             `json:"text,omitempty"`
	Audio       *RawAudio            `json:"audio,omitempty"`
	Document    *RawDocument         `json:"document,omitempty"`
	Image       *RawImage            `json:"image,omitempty"`
	Sticker     *RawSticker          `json:"sticker,omitempty"`
	Location    *RawLocation         `json:"location,omitempty"`
	Contacts    *[]RawMessageContact `json:"contacts,omitempty"`
	Reaction    *RawReaction         `json:"reaction,omitempty"`
	Interactive *RawInteractive      `json:"interactive,omitempty"`
	Referral    *RawReferral         `json:"referral,omitempty"`
	Order       *RawOrder            `json:"order,omitempty"`
	System      *RawSystem           `json:"system,omitempty"`
	Errors      *[]RawError          `json:"errors,omitempty"`
	Context     *RawContext          `json:"context,omitempty"`
}

type RawText struct {
	Body string `json:"body"`
}

type RawDocument struct {
	Caption  string `json:"caption"`
	Filename string `json:"filename"`
	ID       string `json:"id"`
	MimeType string `json:"mime_type"`
	Sha256   string `json:"sha256"`
}

type RawAudio struct {
	ID       string `json:"id"`
	MimeType string `json:"mime_type"`
	Sha256   string `json:"sha256"`
	Voice    bool   `json:"voice"`
}

type RawImage struct {
	Caption  string `json:"caption"`
	MimeType string `json:"mime_type"`
	Sha256   string `json:"sha256"`
	ID       string `json:"id"`
}

type RawSticker struct {
	MimeType string `json:"mime_type"`
	Sha256   string `json:"sha256"`
	ID       string `json:"id"`
}

type RawLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Name      string  `json:"name"`
	Address   string  `json:"address"`
}

type RawMessageContact struct {
	Addresses []RawAddress `json:"addresses,omitempty"`
	Birthday  string       `json:"birthday,omitempty"`
	Emails    []RawEmail   `json:"emails,omitempty"`
	Name      RawName      `json:"name,omitempty"`
	Org       RawOrg       `json:"org,omitempty"`
	Phones    []RawPhone   `json:"phones,omitempty"`
	Urls      []RawUrl     `json:"urls,omitempty"`
}

type RawAddress struct {
	City        string `json:"city"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	Street      string `json:"street"`
	Type        string `json:"type"`
	Zip         string `json:"zip"`
}

type RawEmail struct {
	Email string `json:"email"`
	Type  string `json:"type"`
}

type RawName struct {
	FormattedName string `json:"formatted_name"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	MiddleName    string `json:"middle_name"`
	Suffix        string `json:"suffix"`
	Prefix        string `json:"prefix"`
}

type RawOrg struct {
	Company    string `json:"company"`
	Department string `json:"department"`
	Title      string `json:"title"`
}

type RawPhone struct {
	Phone string `json:"phone"`
	WaID  string `json:"wa_id"`
	Type  string `json:"type"`
}

type RawUrl struct {
	Url  string `json:"url"`
	Type string `json:"type"`
}

type RawReaction struct {
	MessageID string `json:"message_id"`
	Emoji     string `json:"emoji"`
}

type RawInteractive struct {
	Type        string          `json:"type"`
	ListReply   *RawListReply   `json:"list_reply,omitempty"`
	ButtonReply *RawButtonReply `json:"button_reply,omitempty"`
}

type RawListReply struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type RawButtonReply struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type RawReferral struct {
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

type RawOrder struct {
	CatalogID    string           `json:"catalog_id"`
	ProductItems []RawProductItem `json:"product_items"`
	Text         string           `json:"text"`
}

type RawProductItem struct {
	ProductRetailerID string `json:"product_retailer_id"`
	Quantity          string `json:"quantity"`
	ItemPrice         string `json:"item_price"`
	Currency          string `json:"currency"`
}

type RawSystem struct {
	Body    string `json:"body"`
	NewWaID string `json:"new_wa_id"`
	Type    string `json:"type"`
}

type RawError struct {
	Code    int    `json:"code"`
	Details string `json:"details"`
	Title   string `json:"title"`
}

type RawContext struct {
	From            string              `json:"from"`
	ID              string              `json:"id"`
	ReferredProduct *RawReferredProduct `json:"referred_product,omitempty"`
}

type RawReferredProduct struct {
	CatalogID         string `json:"catalog_id"`
	ProductRetailerID string `json:"product_retailer_id"`
}

type RawStatus struct {
	ID           string           `json:"id"`
	Status       string           `json:"status"`
	Timestamp    string           `json:"timestamp"`
	RecipientID  string           `json:"recipient_id"`
	Conversation *RawConversation `json:"conversation,omitempty"`
	Pricing      *RawPricing      `json:"pricing,omitempty"`
	Errors       *[]RawError      `json:"errors,omitempty"`
}

type RawConversation struct {
	ID                  string    `json:"id"`
	ExpirationTimestamp string    `json:"expiration_timestamp,omitempty"`
	Origin              RawOrigin `json:"origin"`
}

type RawOrigin struct {
	Type string `json:"type"`
}

type RawPricing struct {
	Billable     bool   `json:"billable"`
	PricingModel string `json:"pricing_model"`
	Category     string `json:"category"`
}
