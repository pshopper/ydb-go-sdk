// Code generated by gtrace. DO NOT EDIT.

package trace

import (
	"context"
)

// topicComposeOptions is a holder of options
type topicComposeOptions struct {
	panicCallback func(e interface{})
}

// TopicOption specified Topic compose option
type TopicComposeOption func(o *topicComposeOptions)

// WithTopicPanicCallback specified behavior on panic
func WithTopicPanicCallback(cb func(e interface{})) TopicComposeOption {
	return func(o *topicComposeOptions) {
		o.panicCallback = cb
	}
}

// Compose returns a new Topic which has functional fields composed both from t and x.
func (t Topic) Compose(x Topic, opts ...TopicComposeOption) (ret Topic) {
	options := topicComposeOptions{}
	for _, opt := range opts {
		opt(&options)
	}
	{
		h1 := t.OnReaderReconnect
		h2 := x.OnReaderReconnect
		ret.OnReaderReconnect = func(startInfo TopicReaderReconnectStartInfo) func(TopicReaderReconnectDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TopicReaderReconnectDoneInfo)
			if h1 != nil {
				r = h1(startInfo)
			}
			if h2 != nil {
				r1 = h2(startInfo)
			}
			return func(doneInfo TopicReaderReconnectDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(doneInfo)
				}
				if r1 != nil {
					r1(doneInfo)
				}
			}
		}
	}
	{
		h1 := t.OnReaderReconnectRequest
		h2 := x.OnReaderReconnectRequest
		ret.OnReaderReconnectRequest = func(info TopicReaderReconnectRequestInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			if h1 != nil {
				h1(info)
			}
			if h2 != nil {
				h2(info)
			}
		}
	}
	{
		h1 := t.OnReaderPartitionReadStartResponse
		h2 := x.OnReaderPartitionReadStartResponse
		ret.OnReaderPartitionReadStartResponse = func(startInfo TopicReaderPartitionReadStartResponseStartInfo) func(TopicReaderPartitionReadStartResponseDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TopicReaderPartitionReadStartResponseDoneInfo)
			if h1 != nil {
				r = h1(startInfo)
			}
			if h2 != nil {
				r1 = h2(startInfo)
			}
			return func(doneInfo TopicReaderPartitionReadStartResponseDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(doneInfo)
				}
				if r1 != nil {
					r1(doneInfo)
				}
			}
		}
	}
	{
		h1 := t.OnReaderPartitionReadStopResponse
		h2 := x.OnReaderPartitionReadStopResponse
		ret.OnReaderPartitionReadStopResponse = func(startInfo TopicReaderPartitionReadStopResponseStartInfo) func(TopicReaderPartitionReadStopResponseDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TopicReaderPartitionReadStopResponseDoneInfo)
			if h1 != nil {
				r = h1(startInfo)
			}
			if h2 != nil {
				r1 = h2(startInfo)
			}
			return func(doneInfo TopicReaderPartitionReadStopResponseDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(doneInfo)
				}
				if r1 != nil {
					r1(doneInfo)
				}
			}
		}
	}
	{
		h1 := t.OnReaderCommit
		h2 := x.OnReaderCommit
		ret.OnReaderCommit = func(startInfo TopicReaderCommitStartInfo) func(TopicReaderCommitDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TopicReaderCommitDoneInfo)
			if h1 != nil {
				r = h1(startInfo)
			}
			if h2 != nil {
				r1 = h2(startInfo)
			}
			return func(doneInfo TopicReaderCommitDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(doneInfo)
				}
				if r1 != nil {
					r1(doneInfo)
				}
			}
		}
	}
	{
		h1 := t.OnReaderSendCommitMessage
		h2 := x.OnReaderSendCommitMessage
		ret.OnReaderSendCommitMessage = func(startInfo TopicReaderSendCommitMessageStartInfo) func(TopicReaderSendCommitMessageDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TopicReaderSendCommitMessageDoneInfo)
			if h1 != nil {
				r = h1(startInfo)
			}
			if h2 != nil {
				r1 = h2(startInfo)
			}
			return func(doneInfo TopicReaderSendCommitMessageDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(doneInfo)
				}
				if r1 != nil {
					r1(doneInfo)
				}
			}
		}
	}
	{
		h1 := t.OnReaderCommittedNotify
		h2 := x.OnReaderCommittedNotify
		ret.OnReaderCommittedNotify = func(info TopicReaderCommittedNotifyInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			if h1 != nil {
				h1(info)
			}
			if h2 != nil {
				h2(info)
			}
		}
	}
	{
		h1 := t.OnReaderClose
		h2 := x.OnReaderClose
		ret.OnReaderClose = func(startInfo TopicReaderCloseStartInfo) func(TopicReaderCloseDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TopicReaderCloseDoneInfo)
			if h1 != nil {
				r = h1(startInfo)
			}
			if h2 != nil {
				r1 = h2(startInfo)
			}
			return func(doneInfo TopicReaderCloseDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(doneInfo)
				}
				if r1 != nil {
					r1(doneInfo)
				}
			}
		}
	}
	{
		h1 := t.OnReaderInit
		h2 := x.OnReaderInit
		ret.OnReaderInit = func(startInfo TopicReaderInitStartInfo) func(TopicReaderInitDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TopicReaderInitDoneInfo)
			if h1 != nil {
				r = h1(startInfo)
			}
			if h2 != nil {
				r1 = h2(startInfo)
			}
			return func(doneInfo TopicReaderInitDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(doneInfo)
				}
				if r1 != nil {
					r1(doneInfo)
				}
			}
		}
	}
	{
		h1 := t.OnReaderError
		h2 := x.OnReaderError
		ret.OnReaderError = func(info TopicReaderErrorInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			if h1 != nil {
				h1(info)
			}
			if h2 != nil {
				h2(info)
			}
		}
	}
	{
		h1 := t.OnReaderUpdateToken
		h2 := x.OnReaderUpdateToken
		ret.OnReaderUpdateToken = func(startInfo OnReadUpdateTokenStartInfo) func(OnReadUpdateTokenMiddleTokenReceivedInfo) func(OnReadStreamUpdateTokenDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(OnReadUpdateTokenMiddleTokenReceivedInfo) func(OnReadStreamUpdateTokenDoneInfo)
			if h1 != nil {
				r = h1(startInfo)
			}
			if h2 != nil {
				r1 = h2(startInfo)
			}
			return func(updateTokenInfo OnReadUpdateTokenMiddleTokenReceivedInfo) func(OnReadStreamUpdateTokenDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				var r2, r3 func(OnReadStreamUpdateTokenDoneInfo)
				if r != nil {
					r2 = r(updateTokenInfo)
				}
				if r1 != nil {
					r3 = r1(updateTokenInfo)
				}
				return func(doneInfo OnReadStreamUpdateTokenDoneInfo) {
					if options.panicCallback != nil {
						defer func() {
							if e := recover(); e != nil {
								options.panicCallback(e)
							}
						}()
					}
					if r2 != nil {
						r2(doneInfo)
					}
					if r3 != nil {
						r3(doneInfo)
					}
				}
			}
		}
	}
	{
		h1 := t.OnReaderSentDataRequest
		h2 := x.OnReaderSentDataRequest
		ret.OnReaderSentDataRequest = func(startInfo TopicReaderSentDataRequestInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			if h1 != nil {
				h1(startInfo)
			}
			if h2 != nil {
				h2(startInfo)
			}
		}
	}
	{
		h1 := t.OnReaderReceiveDataResponse
		h2 := x.OnReaderReceiveDataResponse
		ret.OnReaderReceiveDataResponse = func(startInfo TopicReaderReceiveDataResponseStartInfo) func(TopicReaderReceiveDataResponseDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TopicReaderReceiveDataResponseDoneInfo)
			if h1 != nil {
				r = h1(startInfo)
			}
			if h2 != nil {
				r1 = h2(startInfo)
			}
			return func(doneInfo TopicReaderReceiveDataResponseDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(doneInfo)
				}
				if r1 != nil {
					r1(doneInfo)
				}
			}
		}
	}
	{
		h1 := t.OnReaderReadMessages
		h2 := x.OnReaderReadMessages
		ret.OnReaderReadMessages = func(startInfo TopicReaderReadMessagesStartInfo) func(TopicReaderReadMessagesDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TopicReaderReadMessagesDoneInfo)
			if h1 != nil {
				r = h1(startInfo)
			}
			if h2 != nil {
				r1 = h2(startInfo)
			}
			return func(doneInfo TopicReaderReadMessagesDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(doneInfo)
				}
				if r1 != nil {
					r1(doneInfo)
				}
			}
		}
	}
	{
		h1 := t.OnReaderUnknownGrpcMessage
		h2 := x.OnReaderUnknownGrpcMessage
		ret.OnReaderUnknownGrpcMessage = func(info OnReadUnknownGrpcMessageInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			if h1 != nil {
				h1(info)
			}
			if h2 != nil {
				h2(info)
			}
		}
	}
	{
		h1 := t.OnWriterReconnect
		h2 := x.OnWriterReconnect
		ret.OnWriterReconnect = func(startInfo TopicWriterReconnectStartInfo) func(TopicWriterReconnectDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TopicWriterReconnectDoneInfo)
			if h1 != nil {
				r = h1(startInfo)
			}
			if h2 != nil {
				r1 = h2(startInfo)
			}
			return func(doneInfo TopicWriterReconnectDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(doneInfo)
				}
				if r1 != nil {
					r1(doneInfo)
				}
			}
		}
	}
	{
		h1 := t.OnWriterInitStream
		h2 := x.OnWriterInitStream
		ret.OnWriterInitStream = func(startInfo TopicWriterInitStreamStartInfo) func(TopicWriterInitStreamDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TopicWriterInitStreamDoneInfo)
			if h1 != nil {
				r = h1(startInfo)
			}
			if h2 != nil {
				r1 = h2(startInfo)
			}
			return func(doneInfo TopicWriterInitStreamDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(doneInfo)
				}
				if r1 != nil {
					r1(doneInfo)
				}
			}
		}
	}
	{
		h1 := t.OnWriterClose
		h2 := x.OnWriterClose
		ret.OnWriterClose = func(startInfo TopicWriterCloseStartInfo) func(TopicWriterCloseDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TopicWriterCloseDoneInfo)
			if h1 != nil {
				r = h1(startInfo)
			}
			if h2 != nil {
				r1 = h2(startInfo)
			}
			return func(doneInfo TopicWriterCloseDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(doneInfo)
				}
				if r1 != nil {
					r1(doneInfo)
				}
			}
		}
	}
	{
		h1 := t.OnWriterCompressMessages
		h2 := x.OnWriterCompressMessages
		ret.OnWriterCompressMessages = func(startInfo TopicWriterCompressMessagesStartInfo) func(TopicWriterCompressMessagesDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TopicWriterCompressMessagesDoneInfo)
			if h1 != nil {
				r = h1(startInfo)
			}
			if h2 != nil {
				r1 = h2(startInfo)
			}
			return func(doneInfo TopicWriterCompressMessagesDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(doneInfo)
				}
				if r1 != nil {
					r1(doneInfo)
				}
			}
		}
	}
	{
		h1 := t.OnWriterSendMessages
		h2 := x.OnWriterSendMessages
		ret.OnWriterSendMessages = func(startInfo TopicWriterSendMessagesStartInfo) func(TopicWriterSendMessagesDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TopicWriterSendMessagesDoneInfo)
			if h1 != nil {
				r = h1(startInfo)
			}
			if h2 != nil {
				r1 = h2(startInfo)
			}
			return func(doneInfo TopicWriterSendMessagesDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(doneInfo)
				}
				if r1 != nil {
					r1(doneInfo)
				}
			}
		}
	}
	{
		h1 := t.OnWriterReadUnknownGrpcMessage
		h2 := x.OnWriterReadUnknownGrpcMessage
		ret.OnWriterReadUnknownGrpcMessage = func(info TopicOnWriterReadUnknownGrpcMessageInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			if h1 != nil {
				h1(info)
			}
			if h2 != nil {
				h2(info)
			}
		}
	}
	return ret
}
func (t Topic) onReaderReconnect(startInfo TopicReaderReconnectStartInfo) func(doneInfo TopicReaderReconnectDoneInfo) {
	fn := t.OnReaderReconnect
	if fn == nil {
		return func(TopicReaderReconnectDoneInfo) {
			return
		}
	}
	res := fn(startInfo)
	if res == nil {
		return func(TopicReaderReconnectDoneInfo) {
			return
		}
	}
	return res
}
func (t Topic) onReaderReconnectRequest(info TopicReaderReconnectRequestInfo) {
	fn := t.OnReaderReconnectRequest
	if fn == nil {
		return
	}
	fn(info)
}
func (t Topic) onReaderPartitionReadStartResponse(startInfo TopicReaderPartitionReadStartResponseStartInfo) func(doneInfo TopicReaderPartitionReadStartResponseDoneInfo) {
	fn := t.OnReaderPartitionReadStartResponse
	if fn == nil {
		return func(TopicReaderPartitionReadStartResponseDoneInfo) {
			return
		}
	}
	res := fn(startInfo)
	if res == nil {
		return func(TopicReaderPartitionReadStartResponseDoneInfo) {
			return
		}
	}
	return res
}
func (t Topic) onReaderPartitionReadStopResponse(startInfo TopicReaderPartitionReadStopResponseStartInfo) func(doneInfo TopicReaderPartitionReadStopResponseDoneInfo) {
	fn := t.OnReaderPartitionReadStopResponse
	if fn == nil {
		return func(TopicReaderPartitionReadStopResponseDoneInfo) {
			return
		}
	}
	res := fn(startInfo)
	if res == nil {
		return func(TopicReaderPartitionReadStopResponseDoneInfo) {
			return
		}
	}
	return res
}
func (t Topic) onReaderCommit(startInfo TopicReaderCommitStartInfo) func(doneInfo TopicReaderCommitDoneInfo) {
	fn := t.OnReaderCommit
	if fn == nil {
		return func(TopicReaderCommitDoneInfo) {
			return
		}
	}
	res := fn(startInfo)
	if res == nil {
		return func(TopicReaderCommitDoneInfo) {
			return
		}
	}
	return res
}
func (t Topic) onReaderSendCommitMessage(startInfo TopicReaderSendCommitMessageStartInfo) func(doneInfo TopicReaderSendCommitMessageDoneInfo) {
	fn := t.OnReaderSendCommitMessage
	if fn == nil {
		return func(TopicReaderSendCommitMessageDoneInfo) {
			return
		}
	}
	res := fn(startInfo)
	if res == nil {
		return func(TopicReaderSendCommitMessageDoneInfo) {
			return
		}
	}
	return res
}
func (t Topic) onReaderCommittedNotify(info TopicReaderCommittedNotifyInfo) {
	fn := t.OnReaderCommittedNotify
	if fn == nil {
		return
	}
	fn(info)
}
func (t Topic) onReaderClose(startInfo TopicReaderCloseStartInfo) func(doneInfo TopicReaderCloseDoneInfo) {
	fn := t.OnReaderClose
	if fn == nil {
		return func(TopicReaderCloseDoneInfo) {
			return
		}
	}
	res := fn(startInfo)
	if res == nil {
		return func(TopicReaderCloseDoneInfo) {
			return
		}
	}
	return res
}
func (t Topic) onReaderInit(startInfo TopicReaderInitStartInfo) func(doneInfo TopicReaderInitDoneInfo) {
	fn := t.OnReaderInit
	if fn == nil {
		return func(TopicReaderInitDoneInfo) {
			return
		}
	}
	res := fn(startInfo)
	if res == nil {
		return func(TopicReaderInitDoneInfo) {
			return
		}
	}
	return res
}
func (t Topic) onReaderError(info TopicReaderErrorInfo) {
	fn := t.OnReaderError
	if fn == nil {
		return
	}
	fn(info)
}
func (t Topic) onReaderUpdateToken(startInfo OnReadUpdateTokenStartInfo) func(updateTokenInfo OnReadUpdateTokenMiddleTokenReceivedInfo) func(doneInfo OnReadStreamUpdateTokenDoneInfo) {
	fn := t.OnReaderUpdateToken
	if fn == nil {
		return func(OnReadUpdateTokenMiddleTokenReceivedInfo) func(OnReadStreamUpdateTokenDoneInfo) {
			return func(OnReadStreamUpdateTokenDoneInfo) {
				return
			}
		}
	}
	res := fn(startInfo)
	if res == nil {
		return func(OnReadUpdateTokenMiddleTokenReceivedInfo) func(OnReadStreamUpdateTokenDoneInfo) {
			return func(OnReadStreamUpdateTokenDoneInfo) {
				return
			}
		}
	}
	return func(updateTokenInfo OnReadUpdateTokenMiddleTokenReceivedInfo) func(OnReadStreamUpdateTokenDoneInfo) {
		res := res(updateTokenInfo)
		if res == nil {
			return func(OnReadStreamUpdateTokenDoneInfo) {
				return
			}
		}
		return res
	}
}
func (t Topic) onReaderSentDataRequest(startInfo TopicReaderSentDataRequestInfo) {
	fn := t.OnReaderSentDataRequest
	if fn == nil {
		return
	}
	fn(startInfo)
}
func (t Topic) onReaderReceiveDataResponse(startInfo TopicReaderReceiveDataResponseStartInfo) func(doneInfo TopicReaderReceiveDataResponseDoneInfo) {
	fn := t.OnReaderReceiveDataResponse
	if fn == nil {
		return func(TopicReaderReceiveDataResponseDoneInfo) {
			return
		}
	}
	res := fn(startInfo)
	if res == nil {
		return func(TopicReaderReceiveDataResponseDoneInfo) {
			return
		}
	}
	return res
}
func (t Topic) onReaderReadMessages(startInfo TopicReaderReadMessagesStartInfo) func(doneInfo TopicReaderReadMessagesDoneInfo) {
	fn := t.OnReaderReadMessages
	if fn == nil {
		return func(TopicReaderReadMessagesDoneInfo) {
			return
		}
	}
	res := fn(startInfo)
	if res == nil {
		return func(TopicReaderReadMessagesDoneInfo) {
			return
		}
	}
	return res
}
func (t Topic) onReaderUnknownGrpcMessage(info OnReadUnknownGrpcMessageInfo) {
	fn := t.OnReaderUnknownGrpcMessage
	if fn == nil {
		return
	}
	fn(info)
}
func (t Topic) onWriterReconnect(startInfo TopicWriterReconnectStartInfo) func(doneInfo TopicWriterReconnectDoneInfo) {
	fn := t.OnWriterReconnect
	if fn == nil {
		return func(TopicWriterReconnectDoneInfo) {
			return
		}
	}
	res := fn(startInfo)
	if res == nil {
		return func(TopicWriterReconnectDoneInfo) {
			return
		}
	}
	return res
}
func (t Topic) onWriterInitStream(startInfo TopicWriterInitStreamStartInfo) func(doneInfo TopicWriterInitStreamDoneInfo) {
	fn := t.OnWriterInitStream
	if fn == nil {
		return func(TopicWriterInitStreamDoneInfo) {
			return
		}
	}
	res := fn(startInfo)
	if res == nil {
		return func(TopicWriterInitStreamDoneInfo) {
			return
		}
	}
	return res
}
func (t Topic) onWriterClose(startInfo TopicWriterCloseStartInfo) func(doneInfo TopicWriterCloseDoneInfo) {
	fn := t.OnWriterClose
	if fn == nil {
		return func(TopicWriterCloseDoneInfo) {
			return
		}
	}
	res := fn(startInfo)
	if res == nil {
		return func(TopicWriterCloseDoneInfo) {
			return
		}
	}
	return res
}
func (t Topic) onWriterCompressMessages(startInfo TopicWriterCompressMessagesStartInfo) func(doneInfo TopicWriterCompressMessagesDoneInfo) {
	fn := t.OnWriterCompressMessages
	if fn == nil {
		return func(TopicWriterCompressMessagesDoneInfo) {
			return
		}
	}
	res := fn(startInfo)
	if res == nil {
		return func(TopicWriterCompressMessagesDoneInfo) {
			return
		}
	}
	return res
}
func (t Topic) onWriterSendMessages(startInfo TopicWriterSendMessagesStartInfo) func(doneInfo TopicWriterSendMessagesDoneInfo) {
	fn := t.OnWriterSendMessages
	if fn == nil {
		return func(TopicWriterSendMessagesDoneInfo) {
			return
		}
	}
	res := fn(startInfo)
	if res == nil {
		return func(TopicWriterSendMessagesDoneInfo) {
			return
		}
	}
	return res
}
func (t Topic) onWriterReadUnknownGrpcMessage(info TopicOnWriterReadUnknownGrpcMessageInfo) {
	fn := t.OnWriterReadUnknownGrpcMessage
	if fn == nil {
		return
	}
	fn(info)
}
func TopicOnReaderReconnect(t Topic) func(error) {
	var p TopicReaderReconnectStartInfo
	res := t.onReaderReconnect(p)
	return func(e error) {
		var p TopicReaderReconnectDoneInfo
		p.Error = e
		res(p)
	}
}
func TopicOnReaderReconnectRequest(t Topic, reason error, wasSent bool) {
	var p TopicReaderReconnectRequestInfo
	p.Reason = reason
	p.WasSent = wasSent
	t.onReaderReconnectRequest(p)
}
func TopicOnReaderPartitionReadStartResponse(t Topic, readerConnectionID string, partitionContext context.Context, topic string, partitionID int64, partitionSessionID int64) func(readOffset *int64, commitOffset *int64, _ error) {
	var p TopicReaderPartitionReadStartResponseStartInfo
	p.ReaderConnectionID = readerConnectionID
	p.PartitionContext = partitionContext
	p.Topic = topic
	p.PartitionID = partitionID
	p.PartitionSessionID = partitionSessionID
	res := t.onReaderPartitionReadStartResponse(p)
	return func(readOffset *int64, commitOffset *int64, e error) {
		var p TopicReaderPartitionReadStartResponseDoneInfo
		p.ReadOffset = readOffset
		p.CommitOffset = commitOffset
		p.Error = e
		res(p)
	}
}
func TopicOnReaderPartitionReadStopResponse(t Topic, readerConnectionID string, partitionContext context.Context, topic string, partitionID int64, partitionSessionID int64, committedOffset int64, graceful bool) func(error) {
	var p TopicReaderPartitionReadStopResponseStartInfo
	p.ReaderConnectionID = readerConnectionID
	p.PartitionContext = partitionContext
	p.Topic = topic
	p.PartitionID = partitionID
	p.PartitionSessionID = partitionSessionID
	p.CommittedOffset = committedOffset
	p.Graceful = graceful
	res := t.onReaderPartitionReadStopResponse(p)
	return func(e error) {
		var p TopicReaderPartitionReadStopResponseDoneInfo
		p.Error = e
		res(p)
	}
}
func TopicOnReaderCommit(t Topic, requestContext context.Context, topic string, partitionID int64, partitionSessionID int64, startOffset int64, endOffset int64) func(error) {
	var p TopicReaderCommitStartInfo
	p.RequestContext = requestContext
	p.Topic = topic
	p.PartitionID = partitionID
	p.PartitionSessionID = partitionSessionID
	p.StartOffset = startOffset
	p.EndOffset = endOffset
	res := t.onReaderCommit(p)
	return func(e error) {
		var p TopicReaderCommitDoneInfo
		p.Error = e
		res(p)
	}
}
func TopicOnReaderSendCommitMessage(t Topic, commitsInfo TopicReaderStreamSendCommitMessageStartMessageInfo) func(error) {
	var p TopicReaderSendCommitMessageStartInfo
	p.CommitsInfo = commitsInfo
	res := t.onReaderSendCommitMessage(p)
	return func(e error) {
		var p TopicReaderSendCommitMessageDoneInfo
		p.Error = e
		res(p)
	}
}
func TopicOnReaderCommittedNotify(t Topic, readerConnectionID string, topic string, partitionID int64, partitionSessionID int64, committedOffset int64) {
	var p TopicReaderCommittedNotifyInfo
	p.ReaderConnectionID = readerConnectionID
	p.Topic = topic
	p.PartitionID = partitionID
	p.PartitionSessionID = partitionSessionID
	p.CommittedOffset = committedOffset
	t.onReaderCommittedNotify(p)
}
func TopicOnReaderClose(t Topic, readerConnectionID string, closeReason error) func(closeError error) {
	var p TopicReaderCloseStartInfo
	p.ReaderConnectionID = readerConnectionID
	p.CloseReason = closeReason
	res := t.onReaderClose(p)
	return func(closeError error) {
		var p TopicReaderCloseDoneInfo
		p.CloseError = closeError
		res(p)
	}
}
func TopicOnReaderInit(t Topic, preInitReaderConnectionID string, initRequestInfo TopicReadStreamInitRequestInfo) func(readerConnectionID string, _ error) {
	var p TopicReaderInitStartInfo
	p.PreInitReaderConnectionID = preInitReaderConnectionID
	p.InitRequestInfo = initRequestInfo
	res := t.onReaderInit(p)
	return func(readerConnectionID string, e error) {
		var p TopicReaderInitDoneInfo
		p.ReaderConnectionID = readerConnectionID
		p.Error = e
		res(p)
	}
}
func TopicOnReaderError(t Topic, readerConnectionID string, e error) {
	var p TopicReaderErrorInfo
	p.ReaderConnectionID = readerConnectionID
	p.Error = e
	t.onReaderError(p)
}
func TopicOnReaderUpdateToken(t Topic, readerConnectionID string) func(tokenLen int, _ error) func(error) {
	var p OnReadUpdateTokenStartInfo
	p.ReaderConnectionID = readerConnectionID
	res := t.onReaderUpdateToken(p)
	return func(tokenLen int, e error) func(error) {
		var p OnReadUpdateTokenMiddleTokenReceivedInfo
		p.TokenLen = tokenLen
		p.Error = e
		res := res(p)
		return func(e error) {
			var p OnReadStreamUpdateTokenDoneInfo
			p.Error = e
			res(p)
		}
	}
}
func TopicOnReaderSentDataRequest(t Topic, readerConnectionID string, requestBytes int, localBufferSizeAfterSent int) {
	var p TopicReaderSentDataRequestInfo
	p.ReaderConnectionID = readerConnectionID
	p.RequestBytes = requestBytes
	p.LocalBufferSizeAfterSent = localBufferSizeAfterSent
	t.onReaderSentDataRequest(p)
}
func TopicOnReaderReceiveDataResponse(t Topic, readerConnectionID string, localBufferSizeAfterReceive int, dataResponse TopicReaderDataResponseInfo) func(error) {
	var p TopicReaderReceiveDataResponseStartInfo
	p.ReaderConnectionID = readerConnectionID
	p.LocalBufferSizeAfterReceive = localBufferSizeAfterReceive
	p.DataResponse = dataResponse
	res := t.onReaderReceiveDataResponse(p)
	return func(e error) {
		var p TopicReaderReceiveDataResponseDoneInfo
		p.Error = e
		res(p)
	}
}
func TopicOnReaderReadMessages(t Topic, requestContext context.Context, minCount int, maxCount int, freeBufferCapacity int) func(messagesCount int, topic string, partitionID int64, partitionSessionID int64, offsetStart int64, offsetEnd int64, freeBufferCapacity int, _ error) {
	var p TopicReaderReadMessagesStartInfo
	p.RequestContext = requestContext
	p.MinCount = minCount
	p.MaxCount = maxCount
	p.FreeBufferCapacity = freeBufferCapacity
	res := t.onReaderReadMessages(p)
	return func(messagesCount int, topic string, partitionID int64, partitionSessionID int64, offsetStart int64, offsetEnd int64, freeBufferCapacity int, e error) {
		var p TopicReaderReadMessagesDoneInfo
		p.MessagesCount = messagesCount
		p.Topic = topic
		p.PartitionID = partitionID
		p.PartitionSessionID = partitionSessionID
		p.OffsetStart = offsetStart
		p.OffsetEnd = offsetEnd
		p.FreeBufferCapacity = freeBufferCapacity
		p.Error = e
		res(p)
	}
}
func TopicOnReaderUnknownGrpcMessage(t Topic, readerConnectionID string, e error) {
	var p OnReadUnknownGrpcMessageInfo
	p.ReaderConnectionID = readerConnectionID
	p.Error = e
	t.onReaderUnknownGrpcMessage(p)
}
func TopicOnWriterReconnect(t Topic, writerInstanceID string, topic string, producerID string, attempt int) func(error) {
	var p TopicWriterReconnectStartInfo
	p.WriterInstanceID = writerInstanceID
	p.Topic = topic
	p.ProducerID = producerID
	p.Attempt = attempt
	res := t.onWriterReconnect(p)
	return func(e error) {
		var p TopicWriterReconnectDoneInfo
		p.Error = e
		res(p)
	}
}
func TopicOnWriterInitStream(t Topic, writerInstanceID string, topic string, producerID string) func(sessionID string, _ error) {
	var p TopicWriterInitStreamStartInfo
	p.WriterInstanceID = writerInstanceID
	p.Topic = topic
	p.ProducerID = producerID
	res := t.onWriterInitStream(p)
	return func(sessionID string, e error) {
		var p TopicWriterInitStreamDoneInfo
		p.SessionID = sessionID
		p.Error = e
		res(p)
	}
}
func TopicOnWriterClose(t Topic, writerInstanceID string, reason error) func(error) {
	var p TopicWriterCloseStartInfo
	p.WriterInstanceID = writerInstanceID
	p.Reason = reason
	res := t.onWriterClose(p)
	return func(e error) {
		var p TopicWriterCloseDoneInfo
		p.Error = e
		res(p)
	}
}
func TopicOnWriterCompressMessages(t Topic, writerInstanceID string, sessionID string, codec int32, firstSeqNo int64, messagesCount int, reason TopicWriterCompressMessagesReason) func(error) {
	var p TopicWriterCompressMessagesStartInfo
	p.WriterInstanceID = writerInstanceID
	p.SessionID = sessionID
	p.Codec = codec
	p.FirstSeqNo = firstSeqNo
	p.MessagesCount = messagesCount
	p.Reason = reason
	res := t.onWriterCompressMessages(p)
	return func(e error) {
		var p TopicWriterCompressMessagesDoneInfo
		p.Error = e
		res(p)
	}
}
func TopicOnWriterSendMessages(t Topic, writerInstanceID string, sessionID string, codec int32, firstSeqNo int64, messagesCount int) func(error) {
	var p TopicWriterSendMessagesStartInfo
	p.WriterInstanceID = writerInstanceID
	p.SessionID = sessionID
	p.Codec = codec
	p.FirstSeqNo = firstSeqNo
	p.MessagesCount = messagesCount
	res := t.onWriterSendMessages(p)
	return func(e error) {
		var p TopicWriterSendMessagesDoneInfo
		p.Error = e
		res(p)
	}
}
func TopicOnWriterReadUnknownGrpcMessage(t Topic, writerInstanceID string, sessionID string, e error) {
	var p TopicOnWriterReadUnknownGrpcMessageInfo
	p.WriterInstanceID = writerInstanceID
	p.SessionID = sessionID
	p.Error = e
	t.onWriterReadUnknownGrpcMessage(p)
}
