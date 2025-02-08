# langGraph 构造 Agent
from typing import List, Literal
from typing_extensions import TypedDict

from langchain_core.prompts import ChatPromptTemplate
from langchain_core.pydantic_v1 import BaseModel, Field

from aiops.module_5.demo_7.rag_flow.env.llm import llm

from aiops.module_5.demo_7.rag_flow.setup.answer_grader import answer_grader
from aiops.module_5.demo_7.rag_flow.setup.hallucination_grader import hallucination_grader
from aiops.module_5.demo_7.rag_flow.setup.question_rewriter import question_rewriter
from aiops.module_5.demo_7.rag_flow.setup.retrieval_grader import retrieval_grader
from aiops.module_5.demo_7.rag_flow.setup.retriever import retriever
from aiops.module_5.demo_7.rag_flow.setup.rag_chain import rag_chain


class GraphState(TypedDict):
    """
    Represents the state of our graph.

    Attributes；
        question: question
        generation: LLM generation
        documents: list of documents
    """
    question: str
    generation: str
    documents: List[str]


def retrieve(state):
    question = state["question"]
    documents = retriever.get_relevant_documents(question)
    return {"documents": documents, "question": question}


def generate(state):
    question = state["question"]
    documents = state["documents"]
    generation = rag_chain.invoke({"context": documents, "question": question})
    return {"documents": documents, "question": question, "generation": generation}


# 判断检索到的文档是否和问题相关
def grade_document(state):
    question = state["question"]
    documents = state["documents"]
    filtered_docs = []
    for d in documents:
        score = retrieval_grader.invoke(
            {"question": question, "document": d.page_content}
        )
        grade = score.binary_score
        if grade == "yes":
            print("文档和用户问题相关")
            filtered_docs.append(d)
        else:
            print("文档和用户问题不相关")
            continue
    return {"documents": filtered_docs, "question": question}


def transform_query(state):
    """
    Transform the query to produce a better question.

    Args:
        state (dict): The current graph state

    Returns:
        state (dict): Updates question key with a re-phrased question
    """
    print("改写问题\n")
    question = state["question"]
    documents = state["documents"]

    # Re-write question
    better_question = question_rewriter.invoke({"question": question})
    print("LLM 改写优化后更好的问题:", better_question)
    return {"documents": documents, "question": better_question}


# Edge
def decide_to_generate(state):
    print("访问检索到的相关知识库\n")
    filtered_documents = state["documents"]
    if not filtered_documents:
        print("所有的文档都不相关，重写问题")
        return "transform_query"
    else:
        print("文档跟问题相关，生成回答")
        return "generate"


# 评估生成的回复是否基于知识库
def grade_generation_v_documents_and_question(state):
    print("评估生成的回答是否基于知识库事实(是否产生了幻觉)")
    question = state["question"]
    documents = state["documents"]
    generation = state["generation"]

    score = hallucination_grader.invoke({
        "documents": documents, "generation": generation
    })
    grade = score.binary_score
    if grade == "yes":
        print("生成的回复基于知识库，没有幻觉\n")
        # 评估 LLM 的回答是否解决了用户问题
        score = answer_grader.invoke({"question": question, "generation": generation})
        grade = score.binary_score
        if grade == "yes":
            print("问题得到解决\n")
            return "useful"
        else:
            print("问题没有得到解决\n")
            return "not useful"
    else:
        print("生成的回复不是基于知识库，继续重试...\n")
        return "not supported"


class RouteQuery(BaseModel):
    datasource: Literal["vectorstore", "web_search"] = Field(
        ...,
        description="Given a user question choose to route it to web search or a vectorstore."
    )


def route_question(state):
    structured_llm_grader = llm.with_structured_output(RouteQuery)

    system = """您是将用户问题路由到 vectorstore 或 web_search 的专家。
vectorstore 包含运维、工单、微服务、网关、工作负载、日志等相关内容，使用 vectorstore 来回答有关这些主题的问题。
否则，请使用 web_search
    """
    route_prompt = ChatPromptTemplate.from_messages(
        [
            ("system", system),
            ("human", "the question is {question}")
        ]
    )

    question_router = route_prompt | structured_llm_grader
    print("---开始分流问题---")
    question = state["question"]
    source = question_router.invoke({"question": question})
    if source.datasource == "web_search":
        print("---问题分流到网络搜索---")
        return "web_search"
    elif source.datasource == "vectorstore":
        print("---问题分流到知识库向量检索---")
        return "vectorstore"


# Search
from langchain_community.tools.tavily_search import TavilySearchResults
from langchain.schema import Document

web_search_tool = TavilySearchResults(k=3)


# 网络搜索
def web_search(state):
    """
    Web search based on the re-phrased question.

    Args:
        state (dict): The current graph state

    Returns:
        state (dict): Updates documents key with appended web results
    """

    print("---进行网络搜索---")
    question = state["question"]

    # Web search
    docs = web_search_tool.invoke({"query": question})
    web_results = "\n".join([d["content"] for d in docs])
    web_results = Document(page_content=web_results)

    return {"documents": web_results, "question": question}
