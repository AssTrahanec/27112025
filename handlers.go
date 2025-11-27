package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
)

type Handler struct {
	storage *Storage
}

func NewHandler(storage *Storage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) CheckLinks(c *gin.Context) {
	var req CheckLinksRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := make(map[string]string)
	var linkStatuses []LinkStatus

	for _, link := range req.Links {
		status := CheckLink(link)
		result[link] = status
		linkStatuses = append(linkStatuses, LinkStatus{
			URL:    link,
			Status: status,
		})
	}

	linksNum := h.storage.Add(linkStatuses)

	c.JSON(http.StatusOK, CheckLinksResponse{
		Links:    result,
		LinksNum: linksNum,
	})
}

func (h *Handler) GenerateReport(c *gin.Context) {
	var req ReportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	links := h.storage.Get(req.LinksList)

	if len(links) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no links found for provided numbers"})
		return
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, "Links Status Report")
	pdf.Ln(12)

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(140, 10, "URL")
	pdf.Cell(40, 10, "Status")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 10)
	for _, link := range links {
		pdf.Cell(140, 8, link.URL)
		pdf.Cell(40, 8, link.Status)
		pdf.Ln(8)
	}

	tmpFile := "report.pdf"
	if err := pdf.OutputFileAndClose(tmpFile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to generate PDF: %v", err)})
		return
	}

	c.File(tmpFile)
}
