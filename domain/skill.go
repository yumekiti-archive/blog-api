package domain

import (
	"time"

	"gorm.io/gorm"
)

type Skill struct {
	gorm.Model
	ID          int `gorm:"primaryKey"`
	Name        string
	Explain     string
	Link        string
	Category    Category
	Tags        []Tag `gorm:"many2many:skill_tags;"`
	SkillDetail SkillDetail
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type SkillDetail struct {
	gorm.Model
	ID        int `gorm:"primaryKey"`
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Skill) Set(skill *Skill) error {
	t.Name = skill.Name
	t.Explain = skill.Explain
	t.Link = skill.Link
	t.Category = skill.Category
	t.Tags = skill.Tags
	t.SkillDetail = skill.SkillDetail

	return nil
}
