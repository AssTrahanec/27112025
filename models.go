package main

type CheckLinksRequest struct {
	Links []string `json:"links" binding:"required"`
}

type CheckLinksResponse struct {
	Links    map[string]string `json:"links"`
	LinksNum int               `json:"links_num"`
}

type ReportRequest struct {
	LinksList []int `json:"links_list" binding:"required"`
}

type LinkStatus struct {
	URL    string `json:"url"`
	Status string `json:"status"`
}
