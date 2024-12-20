package models

import "gorm.io/gorm"

type Ticket struct {
    gorm.Model
    UserID     uint   // Référence à l'utilisateur
    EventID    uint   // Référence à l'événement
    IssueDate  string
    TicketCode string `gorm:"unique"` // Code unique du billet
    Status     string // Statut du billet: 'valid', 'cancelled', 'used'
    CreatedAt  string `gorm:"not null"` // Date de création
    UpdatedAt  string `gorm:"not null"` // Date de mise à jour
}
