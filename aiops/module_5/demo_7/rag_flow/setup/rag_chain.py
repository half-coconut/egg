# 生成回复
from langchain import hub
from langchain_core.output_parsers import StrOutputParser

from aiops.module_5.demo_7.rag_flow.env.llm import llm

# Prompt
prompt = hub.pull("rlm/rag-prompt")


# Post-processing
def format_docs(docs):
    return "\n\n".join(doc.page_content for doc in docs)


# Chain
rag_chain = prompt | llm | StrOutputParser()

if __name__ == '__main__':
    docs = ""
    question = ""
    generation = rag_chain.invoke({"context": docs, "question": question})
    print(generation)
