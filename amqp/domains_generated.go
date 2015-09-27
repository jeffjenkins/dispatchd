
package amqp

var ReadNoWait = ReadBit
var WriteNoWait = WriteBit
var ReadNoLocal = ReadBit
var WriteNoLocal = WriteBit
var ReadMethodId = ReadShort
var WriteMethodId = WriteShort
var ReadReplyText = ReadShortstr
var WriteReplyText = WriteShortstr
var ReadNoAck = ReadBit
var WriteNoAck = WriteBit
var ReadPeerProperties = ReadTable
var WritePeerProperties = WriteTable
var ReadRedelivered = ReadBit
var WriteRedelivered = WriteBit
var ReadClassId = ReadShort
var WriteClassId = WriteShort
var ReadMessageCount = ReadLong
var WriteMessageCount = WriteLong
var ReadReplyCode = ReadShort
var WriteReplyCode = WriteShort
var ReadQueueName = ReadShortstr
var WriteQueueName = WriteShortstr
var ReadConsumerTag = ReadShortstr
var WriteConsumerTag = WriteShortstr
var ReadPath = ReadShortstr
var WritePath = WriteShortstr
var ReadExchangeName = ReadShortstr
var WriteExchangeName = WriteShortstr
var ReadDeliveryTag = ReadLonglong
var WriteDeliveryTag = WriteLonglong
