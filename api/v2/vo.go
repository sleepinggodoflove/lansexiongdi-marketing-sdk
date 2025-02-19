package v2

const SuccessCode = 200

type Status uint8

const (
	StatusIng Status = iota + 1
	StatusSuccess
	StatusFailed
)

var statusMap = map[Status]string{
	StatusIng:     "生成中",
	StatusSuccess: "生成完成",
	StatusFailed:  "生成失败",
}

func (s Status) Value() uint8 {
	return uint8(s)
}

func (s Status) GetText() string {
	t, ok := statusMap[s]
	if !ok {
		return ""
	}
	return t
}

type NotifyEvent string

const (
	NotifyEventKeyCreate  NotifyEvent = "key.create"
	NotifyEventKeyUsage   NotifyEvent = "key.usage"
	NotifyEventKeyDiscard NotifyEvent = "key.discard"
)

func (s NotifyEvent) Value() string {
	return string(s)
}

func (s NotifyEvent) IsKeyCreate() bool {
	return NotifyEventKeyCreate == s
}

func (s NotifyEvent) IsKeyUsage() bool {
	return NotifyEventKeyUsage == s
}

func (s NotifyEvent) IsKeyDiscard() bool {
	return NotifyEventKeyDiscard == s
}

type KeyStatus uint8

const (
	KeyNormal KeyStatus = iota + 1
	KeyUsed
	KeyDiscard
	KeyExpire
)

var keyStatusMap = map[KeyStatus]string{
	KeyNormal:  "正常",
	KeyUsed:    "已核销",
	KeyDiscard: "已作废",
	KeyExpire:  "已过期",
}

func (s KeyStatus) Value() uint8 {
	return uint8(s)
}

func (s KeyStatus) GetText() string {
	t, ok := keyStatusMap[s]
	if !ok {
		return ""
	}
	return t
}

func (s KeyStatus) IsNormal() bool {
	return s == KeyNormal
}

func (s KeyStatus) IsUsed() bool {
	return s == KeyUsed
}

func (s KeyStatus) IsDiscard() bool {
	return s == KeyDiscard
}

func (s KeyStatus) IsExpire() bool {
	return s == KeyExpire
}
