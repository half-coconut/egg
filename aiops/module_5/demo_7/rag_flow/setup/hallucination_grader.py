# 评估 LLM 的回复是否基于事实，有没有产生幻觉
from langchain_core.prompts import ChatPromptTemplate
from langchain_core.pydantic_v1 import BaseModel, Field

from aiops.module_5.demo_7.rag_flow.env.llm import llm


class GraderHallucinations(BaseModel):
    """Binary score for hallucinations present in generation answer."""

    binary_score: str = Field(description="Documents are relevant to the question, 'yes' or 'no'")


structured_llm_grader = llm.with_structured_output(GraderHallucinations)

# prompt
system = """
您是一名评分员，正在评估 LLM 生成是否基于一组检索的事实/由一组检索到的事实支持。\n
给出二进制分数"yes"或"no"，"yes"表示答案基于一组事实/由一组事实支持。
"""

hallucination_prompt = ChatPromptTemplate.from_messages(
    [
        ("system", system),
        ("human", "Set of facts: \n\n {documents} \n\n LLM generations: {generation}")
    ]
)

hallucination_grader = hallucination_prompt | structured_llm_grader

if __name__ == '__main__':
    docs = ""
    generation = ""
    hallucination_grader.invoke({"documents": docs, "generation": generation})
