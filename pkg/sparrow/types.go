package sparrow

import "time"

type Payload struct {
	Data Data `json:"data,omitempty"`
}

type Book struct {
	ID               int       `json:"id,omitempty"`
	Type             string    `json:"type,omitempty"`
	Slug             string    `json:"slug,omitempty"`
	Name             string    `json:"name,omitempty"`
	UserID           int       `json:"user_id,omitempty"`
	Description      string    `json:"description,omitempty"`
	CreatorID        int       `json:"creator_id,omitempty"`
	Public           int       `json:"public,omitempty"`
	ItemsCount       int       `json:"items_count,omitempty"`
	LikesCount       int       `json:"likes_count,omitempty"`
	WatchesCount     int       `json:"watches_count,omitempty"`
	ContentUpdatedAt time.Time `json:"content_updated_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	User             User      `json:"user,omitempty"`
	Serializer       string    `json:"_serializer,omitempty"`
}

type User struct {
	ID               int       `json:"id,omitempty"`
	Type             string    `json:"type,omitempty"`
	Login            string    `json:"login,omitempty"`
	Name             string    `json:"name,omitempty"`
	Description      string    `json:"description,omitempty"`
	AvatarURL        string    `json:"avatar_url,omitempty"`
	BooksCount       int       `json:"books_count,omitempty"`
	PublicBooksCount int       `json:"public_books_count,omitempty"`
	FollowersCount   int       `json:"followers_count,omitempty"`
	FollowingCount   int       `json:"following_count,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
	Serializer       string    `json:"_serializer,omitempty"`
}

type Actor struct {
	ID             int         `json:"id,omitempty"`
	Type           string      `json:"type,omitempty"`
	Login          string      `json:"login,omitempty"`
	Name           string      `json:"name,omitempty"`
	Avatar         string      `json:"avatar,omitempty"`
	Scene          interface{} `json:"scene,omitempty"`
	AvatarURL      string      `json:"avatar_url,omitempty"`
	Role           int         `json:"role,omitempty"`
	IsPaid         bool        `json:"isPaid,omitempty"`
	MemberLevel    int         `json:"member_level,omitempty"`
	FollowersCount int         `json:"followers_count,omitempty"`
	FollowingCount int         `json:"following_count,omitempty"`
	Description    string      `json:"description,omitempty"`
	Status         int         `json:"status,omitempty"`
	ExpiredAt      time.Time   `json:"expired_at,omitempty"`
	Serializer     string      `json:"_serializer,omitempty"`
}

type Data struct {
	ID                 int       `json:"id,omitempty"`
	Slug               string    `json:"slug,omitempty"`
	Title              string    `json:"title,omitempty"`
	BookID             int       `json:"book_id,omitempty"`
	Book               Book      `json:"book,omitempty"`
	UserID             int       `json:"user_id,omitempty"`
	User               User      `json:"user,omitempty"`
	Format             string    `json:"format,omitempty"`
	Body               string    `json:"body,omitempty"`
	BodyDraft          string    `json:"body_draft,omitempty"`
	BodyHTML           string    `json:"body_html,omitempty"`
	Public             int       `json:"public,omitempty"`
	Status             int       `json:"status,omitempty"`
	ViewStatus         int       `json:"view_status,omitempty"`
	ReadStatus         int       `json:"read_status,omitempty"`
	LikesCount         int       `json:"likes_count,omitempty"`
	CommentsCount      int       `json:"comments_count,omitempty"`
	ContentUpdatedAt   time.Time `json:"content_updated_at,omitempty"`
	DeletedAt          time.Time `json:"deleted_at,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
	PublishedAt        time.Time `json:"published_at,omitempty"`
	FirstPublishedAt   time.Time `json:"first_published_at,omitempty"`
	WordCount          int       `json:"word_count,omitempty"`
	Serializer         string    `json:"_serializer,omitempty"`
	Path               string    `json:"path,omitempty"`
	Publish            bool      `json:"publish,omitempty"`
	ActionType         string    `json:"action_type,omitempty"`
	WebhookSubjectType string    `json:"webhook_subject_type,omitempty"`
	ActorID            int       `json:"actor_id,omitempty"`
	Actor              Actor     `json:"actor,omitempty"`
}
