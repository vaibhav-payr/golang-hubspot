package hubspot

type ContactField struct {
	Name  string `json:"property"`
	Value string `json:"value"`
}

type CreateOrUpdateContactRequest struct {
	Properties []ContactField `json:"properties"`
}

type ListBody struct {
	Emails []string `json:"emails"`
}

