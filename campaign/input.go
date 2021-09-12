package campaign

import (
	"fundraising/user"
)

type GetCampaignDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateCampaignInput struct {
	Name             string `json:"name" binding:"required"` //PIC yang bertanggung jawab
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"` //jumlah pengajuan uang
	Perks            string `json:"perks" binding:"required"`       // tujuan acara
	User             user.User

	// TODO: implementasikan
	// EventName string `json:"event_name"`
	// HospitalName string `json:"hospital_name"`
	// Contact int `json:"contact"`
	// ActualCost int `json:"actual_cost"`
	// EventDate time.Time `json:"event_date"`
	// SubmissionDate time.Time `json:"submission_date"`
	// RequiredDocument bool `json:"required_document"`
	// surat komunikasi dari THC
	// surat komunikasi dariRS materi presentasi/training
	// marketing form
	// approval kegiatan

	// daftar hadir
	// dokumentasi acara
	// invoice resmi

	// BankName string `json:"bank_name"`
	// Branch string `json:"branch"`
	// AccountNumber int `json:"account_number"`
	// AccountName string `json:"account_name"`

	// Approval bool `json:"approval"`
}
