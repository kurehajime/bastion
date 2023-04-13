# fumidai

A stepping-stone server that provides APIs while hiding access tokens

## example
 $ export OPEN_API_KEY=YOUTR_ACCESS_TOKEN
 $ ./fumidai \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $OPEN_API_KEY" \
  -X "POST"  \
  https://api.openai.com/v1/chat/completions
