package types

// ChatRequest doc link https://api-docs.deepseek.com/zh-cn/api/create-chat-completion
type ChatRequest struct {
	Messages         []Message      `json:"messages"`
	Model            string         `json:"model"`             //  使用的模型的 ID。您可以使用 deepseek-chat。 [deepseek-chat, deepseek-reasoner]
	FrequencyPenalty float64        `json:"frequency_penalty"` //介于 -2.0 和 2.0 之间的数字。如果该值为正，那么新 token 会根据其在已有文本中的出现频率受到相应的惩罚，降低模型重复相同内容的可能性。
	MaxTokens        int            `json:"max_tokens"`        // 介于 1 到 8192 间的整数，限制一次请求中模型生成 completion 的最大 token 数。输入 token 和输出 token 的总长度受模型的上下文长度的限制。 如未指定 max_tokens参数，默认使用 4096。
	PresencePenalty  float64        `json:"presence_penalty"`  // 介于 -2.0 和 2.0 之间的数字。如果该值为正，那么新 token 会根据其是否已在已有文本中出现受到相应的惩罚，从而增加模型谈论新主题的可能性。
	ResponseFormat   ResponseFormat `json:"response_format"`   // 一个 object，指定模型必须输出的格式。 设置为 { "type": "json_object" } 以启用 JSON 模式，该模式保证模型生成的消息是有效的 JSON。 注意: 使用 JSON 模式时，你还必须通过系统或用户消息指示模型生成 JSON。否则，模型可能会生成不断的空白字符，直到生成达到令牌限制，从而导致请求长时间运行并显得“卡住”。此外，如果 finish_reason="length"，这表示生成超过了 max_tokens 或对话超过了最大上下文长度，消息内容可能会被部分截断。
	Stop             []string       `json:"stop"`              // 一个 string 或最多包含 16 个 string 的 list，在遇到这些词时，API 将停止生成更多的 token。
	Stream           bool           `json:"stream"`            // 如果设置为 True，将会以 SSE（server-sent events）的形式以流式发送消息增量。消息流以 data: [DONE] 结尾。
	StreamOptions    interface{}    `json:"stream_options"`    // 流式输出相关选项。只有在 stream 参数为 true 时，才可设置此参数。 include_usage  boolean 如果设置为 true，在流式消息最后的 data: [DONE] 之前将会传输一个额外的块。此块上的 usage 字段显示整个请求的 token 使用统计信息，而 choices 字段将始终是一个空数组。所有其他块也将包含一个 usage 字段，但其值为 null。
	Temperature      float64        `json:"temperature"`       //  采样温度，介于 0 和 2 之间。更高的值，如 0.8，会使输出更随机，而更低的值，如 0.2，会使其更加集中和确定。 我们通常建议可以更改这个值或者更改 top_p，但不建议同时对两者进行修改。
	TopP             float64        `json:"top_p"`             // 作为调节采样温度的替代方案，模型会考虑前 top_p 概率的 token 的结果。所以 0.1 就意味着只有包括在最高 10% 概率中的 token 会被考虑。 我们通常建议修改这个值或者更改 temperature，但不建议同时对两者进行修改。
	Tools            interface{}    `json:"tools"`             // 模型可能会调用的 tool 的列表。目前，仅支持 function 作为工具。使用此参数来提供以 JSON 作为输入参数的 function 列表。最多支持 128 个 function。
	ToolChoice       string         `json:"tool_choice"`       // 控制模型调用 tool 的行为。 none 意味着模型不会调用任何 tool，而是生成一条消息。 auto 意味着模型可以选择生成一条消息或调用一个或多个 tool。 required 意味着模型必须调用一个或多个 tool。 通过 {"type": "function", "function": {"name": "my_function"}} 指定特定 tool，会强制模型调用该 tool。 当没有 tool 时，默认值为 none。如果有 tool 存在，默认值为 auto。
	Logprobs         bool           `json:"logprobs"`          // 是否返回所输出 token 的对数概率。如果为 true，则在 message 的 content 中返回每个输出 token 的对数概率。
	TopLogprobs      *int           `json:"top_logprobs"`      // 一个介于 0 到 20 之间的整数 N，指定每个输出位置返回输出概率 top N 的 token，且返回这些 token 的对数概率。指定此参数时，logprobs 必须为 true。
}

type Message struct {
	Content          string `json:"content"`
	Role             string `json:"role"`                        //  system, user, assistant, tool
	Name             string `json:"name,omitempty"`              // 可以选填的参与者的名称，为模型提供信息以区分相同角色的参与者。
	Prefix           string `json:"prefix,omitempty"`            // (Beta) 设置此参数为 true，来强制模型在其回答中以此 assistant 消息中提供的前缀内容开始。 您必须设置 base_url="https://api.deepseek.com/beta" 来使用此功能
	ReasoningContent string `json:"reasoning_content,omitempty"` // (Beta) 用于 deepseek-reasoner 模型在对话前缀续写功能下，作为最后一条 assistant 思维链内容的输入。使用此功能时，prefix 参数必须设置为 true。
	ToolCallId       string `json:"tool_call_id,omitempty"`      // role is tool 此消息所响应的 tool call 的 ID。
}

type ResponseFormat struct {
	Type string `json:"type"` // response format type:  json_object or text
}

type ChatResponse struct {
	ID                string   `json:"id"`                 // 该对话的唯一标识符。
	Choices           []Choice `json:"choices"`            // 模型生成的 completion 的选择列表。
	Created           int64    `json:"created"`            // 创建聊天完成时的 Unix 时间戳（以秒为单位）。
	Model             string   `json:"model"`              // 生成该 completion 的模型名。
	SystemFingerprint string   `json:"system_fingerprint"` // This fingerprint represents the backend configuration that the model runs with.
	Object            string   `json:"object"`             // 对象的类型, 其值为 chat.completion。
	Usage             Usage    `json:"usage"`              // 该对话补全请求的用量信息
}

type Choice struct {
	Delta        Delta               `json:"delta"`         // 流式返回的一个 completion 增量。
	FinishReason string              `json:"finish_reason"` // 模型停止生成 token 的原因。 stop：模型自然停止生成，或遇到 stop 序列中列出的字符串。 length ：输出长度达到了模型上下文长度限制，或达到了 max_tokens 的限制。 content_filter：输出内容因触发过滤策略而被过滤。 insufficient_system_resource：系统推理资源不足，生成被打断。
	Index        int                 `json:"index"`         // 该 completion 在模型生成的 completion 的选择列表中的索引。
	Message      ChatResponseMessage `json:"message"`       // 模型生成的 completion 消息。
	Logprobs     Logprobs            `json:"logprobs"`      // 该 choice 的对数概率信息。
}

type Delta struct {
	Content          string `json:"content"`           // completion 增量的内容。
	ReasoningContent string `json:"reasoning_content"` // 仅适用于 deepseek-reasoner 模型。内容为 assistant 消息中在最终答案之前的推理内容
	Role             string `json:"role"`              //产生这条消息的角色
}

type ChatResponseMessage struct {
	Content          string     `json:"content"`           // 该 completion 的内容。
	ReasoningContent string     `json:"reasoning_content"` // 仅适用于 deepseek-reasoner 模型。内容为 assistant 消息中在最终答案之前的推理内容。
	ToolCalls        []ToolCall `json:"tool_calls"`        // 模型生成的 tool 调用，例如 function 调用。
	Role             string     `json:"role"`              // 生成这条消息的角色。
}

type ToolCall struct {
	ID       string   `json:"id"`       // tool 调用的 ID。
	Type     string   `json:"type"`     // tool 的类型。目前仅支持 function。
	Function Function `json:"function"` // 模型调用的 function。
}

type Function struct {
	Name      string `json:"name"`      // 模型调用的 function 名。
	Arguments string `json:"arguments"` // 要调用的 function 的参数，由模型生成，格式为 JSON。请注意，模型并不总是生成有效的 JSON，并且可能会臆造出你函数模式中未定义的参数。在调用函数之前，请在代码中验证这些参数。
}

type Logprobs struct {
	Content []LogprobContent `json:"content"` // 一个包含输出 token 对数概率信息的列表。
}

type LogprobContent struct {
	Token       string       `json:"token"`        // 输出的 token。
	Logprob     float64      `json:"logprob"`      // 该 token 的对数概率。-9999.0 代表该 token 的输出概率极小，不在 top 20 最可能输出的 token 中。
	Bytes       []int        `json:"bytes"`        // 一个包含该 token UTF-8 字节表示的整数列表。一般在一个 UTF-8 字符被拆分成多个 token 来表示时有用。如果 token 没有对应的字节表示，则该值为 null。
	TopLogprobs []TopLogprob `json:"top_logprobs"` // 一个包含在该输出位置上，输出概率 top N 的 token 的列表，以及它们的对数概率。在罕见情况下，返回的 token 数量可能少于请求参数中指定的 top_logprobs 值。
}

type TopLogprob struct {
	Token   string  `json:"token"`   // 输出的 token。
	Logprob float64 `json:"logprob"` // 该 token 的对数概率。-9999.0 代表该 token 的输出概率极小，不在 top 20 最可能输出的 token 中。
	Bytes   []int   `json:"bytes"`   // 一个包含该 token UTF-8 字节表示的整数列表。一般在一个 UTF-8 字符被拆分成多个 token 来表示时有用。如果 token 没有对应的字节表示，则该值为 null。
}

type Usage struct {
	CompletionTokens        int                     `json:"completion_tokens"`         // 模型 completion 产生的 token 数。
	PromptTokens            int                     `json:"prompt_tokens"`             // 用户 prompt 所包含的 token 数。该值等于 prompt_cache_hit_tokens + prompt_cache_miss_tokens
	PromptCacheHitTokens    int                     `json:"prompt_cache_hit_tokens"`   // 用户 prompt 中，命中上下文缓存的 token 数。
	PromptCacheMissTokens   int                     `json:"prompt_cache_miss_tokens"`  // 用户 prompt 中，未命中上下文缓存的 token 数。
	TotalTokens             int                     `json:"total_tokens"`              // 该请求中，所有 token 的数量（prompt + completion）。
	CompletionTokensDetails CompletionTokensDetails `json:"completion_tokens_details"` // completion tokens 的详细信息。
}

type CompletionTokensDetails struct {
	ReasoningTokens int `json:"reasoning_tokens"` // 推理模型所产生的思维链 token 数量
}
