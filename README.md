# fumidai

A stepping-stone server that provides APIs while hiding access tokens

## example

For example, when you want the OpenAI API to hide the access token and make the API available.

Give fumidai a header, a method, and a URL to start.

```
 $ export OPEN_API_KEY=YOUTR_ACCESS_TOKEN
 $ ./fumidai \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $OPEN_API_KEY" \
  -X "POST"  \
  https://api.openai.com
```

POST the Content-Type and request body to fumidai.

```
curl http://localhost:8080/v1/chat/completions \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gpt-3.5-turbo",
    "messages": [{"role": "user", "content": "Hello!"}]
  }'
```

`{"id":"chatcmpl-74sN3THp5Mh4Xo7err5RzWlZeylHy","object":"chat.completion","created":1681396669,"model":"gpt-3.5-turbo-0301","usage":{"prompt_tokens":10,"completion_tokens":10,"total_tokens":20},"choices":[{"message":{"role":"assistant","content":"Hello there! How may I assist you today?"},"finish_reason":"stop","index":0}]}`

The request thrown to the springboard is thrown to OpenAI's API, and return.