package domain

type Motivation struct {
	Id         int64
	AuthorTgId int64
	TgFileId   string
}

type GoydaResponse struct {
	TgFileId       *string
	ChipsAccounted bool
}
