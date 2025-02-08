# 重写问题
from langchain_core.output_parsers import StrOutputParser
from langchain_core.prompts import ChatPromptTemplate

from aiops.module_5.demo_7.rag_flow.env.llm import llm

system = """
您有一个问题重写器，可将输入问题转换为针对 vectorstore 检索进行了优化的更好版本 \n
查看输入并尝试推断底层语义意图/含义，使用用户语言回复。
"""
re_write_prompt = ChatPromptTemplate.from_messages(
    [
        ("system", system),
        ("human", "Here is the initial question: \n\n {question} \n Formulate an improved question")
    ]
)

question_rewriter = re_write_prompt | llm | StrOutputParser()

if __name__ == '__main__':
    question = ""
    question_rewriter.invoke({"question": question})
