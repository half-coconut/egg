# 评估 LLM 回答是否解决了用户的问题
from langchain_core.prompts import ChatPromptTemplate
from langchain_core.pydantic_v1 import BaseModel, Field

from aiops.module_5.demo_7.rag_flow.env.llm import llm


class GraderAnswer(BaseModel):
    """Binary score to access answer addresses question."""

    binary_score: str = Field(description="Documents are relevant to the question, 'yes' or 'no'")


structured_llm_grader = llm.with_structured_output(GraderAnswer)

# prompt
system = """
您是一名评分员，评估答案是否解决了某个问题。\n
给出二进制分数"yes"或"no"，"yes"表示答案解决了问题。
"""

answer_prompt = ChatPromptTemplate.from_messages(
    [
        ("system", system),
        ("human", "User question: \n\n {question} \n\n LLM generations: {generation}")
    ]
)

answer_grader = answer_prompt | structured_llm_grader

if __name__ == '__main__':
    question = ""
    generation = ""
    answer_grader.invoke({"question": question, "generation": generation})
