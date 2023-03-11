package cmd

type TranslateResponse struct {
	Ok             bool   `json:"ok"`
	TextLang       string `json:"text_lang"`
	TranslatedText string `json:"translated_text"`
}
