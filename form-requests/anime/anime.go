package formRequest

type StoreRequest struct {
  Name string `form:"name" json:"name" binding:"required"`
  Description string `form:"description" json:"description" binding:"required"`
  CoverImage string `form:"coverImage" json:"coverImage" binding:"required,url"`
  Status string `form:"status" json:"status" binding:"required"`
  PublishedYear int `form:"publishedYear" json:"publishedYear" binding:"required,numeric,min=1900"`
}