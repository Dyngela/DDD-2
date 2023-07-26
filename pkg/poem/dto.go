package poem

type CreatePoemRequest struct {
	Author      string
	ReleaseDate string
	Content     string
}

type GetPoemByIdResponse struct {
	Author      string
	ReleaseDate string
	Content     string
}
