package domain

type validationResult struct {
	IsSuccess bool
	Messages  map[string]string
}

func NewValidationResult(messages map[string]string) *validationResult {
	vr := &validationResult{}
	if messages == nil || len(messages) == 0 {
		vr.IsSuccess = true
		return vr
	}
	vr.IsSuccess = false
	vr.Messages = messages
	return vr
}

func (result *validationResult) ToMessagesSlice() []string {
	var messages []string
	for _, item := range result.Messages {
		messages = append(messages, item)
	}
	return messages
}
