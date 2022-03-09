package chat

import "time"

type UserID string
type TaskID string
type CommentID string
type ReplyID string
type AssignID string
type ContractID string
type OrderID string

// lookuptabe
type ContractType string
type UserType string

// validation table

// Contract Change History
type ContractChangeHistory struct {
}

type ContractReviewTask struct {
	TaskID       TaskID // PK
	CreatorID    UserID // FK
	ContractID   ContractID
	OrderIDs     []OrderID
	Participants []UserID
	CreatedAt    time.Time
	Attaches     [][]byte // TODO: to check this way work.
}

// embed or referecing

// query table

// tags

type ContractReviewTaskAssign struct {
	AssignID   AssignID `json:"assignID" db:"task_id"` // PK
	Dispatcher UserID
	TaskID     TaskID `json:"taskID" db:"task_id"` // CPK
	UserID     UserID `json:"userID" db:"user_id"` // CPK
}

// TODO: considering applying responsible, accountable, consulted, and informed

type ContractReviewComment struct {
	CommentID   CommentID // PK
	TaskID      TaskID    // FK
	CreatorID   UserID    // FK
	InformedIDs []UserID  //
	Message     string    `json:"message" maxLength:"1024"`
	CreatedAt   time.Time

	QuoteID string // pk
}

//type Reply struct {
//	ReplyID   ReplyID   // PK
//	CommentID CommentID // FK
//	AuthorID  UserID    // FK
//	Message   string
//	CreatedAt time.Time
//}

// list, group, at

// quote, reminder, @

// task management
// add sub-task, topit, give up, tag, attach, task moment, create copy, assigned to,
// priority: high, meddle, low, none

// notice
// task moment
