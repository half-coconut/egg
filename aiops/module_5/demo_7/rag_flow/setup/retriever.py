from langchain import hub as langchain_hub
from langchain.schema import StrOutputParser
from langchain_openai import ChatOpenAI
from langchain.schema.runnable import RunnablePassthrough
from langchain.text_splitter import MarkdownHeaderTextSplitter
from langchain_openai import OpenAIEmbeddings
import os
from langchain_community.vectorstores.chroma import Chroma
from langchain_core.prompts import PromptTemplate
from string import Template
import uuid


def retriever():
    # 读取 ./data/data.md 文件作为知识库
    file_path = os.path.join('/Users/chenchen/Desktop/egg/aiops/module_5/demo_7/rag_flow/data', 'data.md')
    with open(file_path, 'r', encoding='utf-8') as f:
        docs_string = f.read()

    # split the document into chunks base on markdown headers
    headers_to_split_on = [
        ("#", "Header 1"),
        ("##", "Header 2"),
        ("###", "Header 3"),
    ]
    text_splitter = MarkdownHeaderTextSplitter(headers_to_split_on=headers_to_split_on)
    splits = text_splitter.split_text(docs_string)
    print("Length of splits: ", len(splits))
    print(splits)

    # 向量化
    # 保存到随机目录里
    random_directory = "./" + str(uuid.uuid4())
    # embedding = OpenAIEmbeddings(model="text-embedding-3-small",
    #                              openai_api_key=os.getenv("OPENAI_API_KEY"), openai_api_base=os.getenv("OPENAI_API_BASE"))

    embedding = OpenAIEmbeddings()
    vectorstore = Chroma.from_documents(documents=splits, embedding=embedding, persist_directory=random_directory)
    vectorstore.persist()
    return vectorstore.as_retriever()


retriever = retriever()
