package types

// JSend Response Types

type SuccessResponse struct {
  Status  string      `json:"status"`
  Message string      `json:"message"`
  Data    interface{} `json:"data"`
}

type FailResponse struct {
  Status  string      `json:"status"`
  Message string      `json:"message"`
  Data    interface{} `json:"data"`
}

type ErrorResponse struct {
  Status  string      `json:"status"`
  Message string      `json:"message"`
  Code    int         `json:"code,omitempty"`
  Data    interface{} `json:"data,omitempty"`
}

// Pagination Meta
type PaginationMeta struct {
  Page        int   `json:"page"`
  Limit       int   `json:"limit"`
  Total       int64 `json:"total"`
  TotalPages  int   `json:"total_pages"`
  HasNextPage bool  `json:"has_next_page"`
  HasPrevPage bool  `json:"has_prev_page"`
}
