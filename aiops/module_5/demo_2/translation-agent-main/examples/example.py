from openai import OpenAI

client = OpenAI(api_key="sk-bzNskWdx7jT08w7KfxA8pAnAoU8fkN6pFCdGkWCs3rKxxWyy",
                base_url="https://sg.uiuiapi.com/v1")

completion = client.chat.completions.create(
    # model="gpt-4o-mini",
    model="gpt-4-turbo",
    messages=[
        {
            "role": "system",
            "content": "将以下英文翻译成中文"
        },
        {
            "role": "user",
            "content": "Quiet Night Thoughts\nThe bright moonlight before my bed,\nI suspect it is frost on the ground.\nI raise my head to gaze at the bright moon,\nAnd lower it to think of my hometown."
        }
    ]
)

print(completion.choices[0].message.content)
