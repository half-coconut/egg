# 评估检索的文档与用户提出的问题是否相关
from langchain_core.prompts import ChatPromptTemplate
from langchain_core.pydantic_v1 import BaseModel, Field

from aiops.module_5.demo_7.rag_flow.env.llm import llm
from aiops.module_5.demo_7.rag_flow.setup.retriever import retriever


class GradeDocuments(BaseModel):
    binary_score: str = Field(description="Documents are relevant to the question, 'yes' or 'no'")


structured_llm_grader = llm.with_structured_output(GradeDocuments)

# prompt
system = """
您是一名评分员，负责评估检索到的文档与用户问题的相关性。\n
测试不需要很严格，目标是过滤掉错误的检索。\n
如果文档包含与用户问题相关的关键字或语义含义，则将其评为相关。\n
给出二进制分数"yes"或"no"，以指示文档是否与问题相关。
"""
grade_prompt = ChatPromptTemplate.from_messages(
    [
        ("system", system),
        ("human", "Retrieved document: \n\n {document} \n\n User question: \n\n {question}")
    ]
)

retrieval_grader = grade_prompt | structured_llm_grader

if __name__ == '__main__':
    # 相关问题
    question = "payment_backend 服务是谁维护的"
    # 不相关问题
    # question = "北京天气如何"

    docs = retriever._get_relevant_documents(question, run_manager=None)

    # 观察每一个文档块的相关性判断结果
    for doc in docs:
        print("doc: \n", doc.page_content, "\n")
        print(retrieval_grader.invoke({"question": question, "document": doc.page_content}))
        print("\n")
