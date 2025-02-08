import os
from dotenv import load_dotenv
from langchain_openai import ChatOpenAI

# 初始化 langsmith
os.environ["LANGCHAIN_TRACING_V2"] = "true"
os.environ["LANGCHAIN_ENDPOINT"] = "https://api.smith.langchain.com"
os.environ["LANGCHAIN_API_KEY"] = "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"
os.environ["LANGCHAIN_PROJECT"] = "default"

# https://smith.langchain.com

# Tavily API Key
os.environ["TAVILY_API_KEY"] = "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"

# https://app.tavily.com/

# 加载 .env 文件
load_dotenv()

llm = ChatOpenAI(model="gpt-4o-mini", temperature=0, openai_api_key=os.getenv("OPENAI_API_KEY"),
                 openai_api_base=os.getenv("OPENAI_API_BASE"))