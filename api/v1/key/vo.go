package key

type Status uint8

const (
	Normal Status = iota + 1
	Used
	Discard
	Expire
)

var statusMap = map[Status]string{
	Normal:  "正常",
	Used:    "已核销",
	Discard: "已作废",
	Expire:  "已过期",
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

func (s Status) IsNormal() bool {
	return s == Normal
}

func (s Status) IsUsed() bool {
	return s == Used
}

func (s Status) IsDiscard() bool {
	return s == Discard
}

func (s Status) IsExpire() bool {
	return s == Expire
}
