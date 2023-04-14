import openai
openai.api_base = "http://localhost:8080/v1"
completion = openai.ChatCompletion.create(
    model="gpt-3.5-turbo",
    messages=[
        {"role": "user", "content": "Think of three names for cats that seem smart."}
    ]
)

for cho in completion.choices:
    print(cho.message.content)
