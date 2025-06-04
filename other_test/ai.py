# Please install OpenAI SDK first: `pip3 install openai`

from openai import OpenAI

client = OpenAI(api_key="sk-f4a03a09438443319513e5f242b7e7df", base_url="https://api.deepseek.com")

response = client.chat.completions.create(
   model="deepseek-chat",
#   model="deepseek-reasoner",
    messages=[
        {"role": "system", "content": "You are a helpful assistant"},
        {"role": "user", "content": "如果你有1000万现金，你会存在哪个银行？"},
    ],
    stream=False
)

# 逐步打印流内容
print(response.choices[0].message.content)
