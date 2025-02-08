from langgraph.graph import END, StateGraph, START
from IPython.display import Image, display

from aiops.module_5.demo_7.rag_flow.node import GraphState, web_search, retrieve, grade_document, generate, \
    transform_query, route_question, decide_to_generate, grade_generation_v_documents_and_question


def workflow():
    workflow = StateGraph(GraphState)

    workflow.add_node("web_search", web_search)
    workflow.add_node("retrieve", retrieve)
    workflow.add_node("grade_documents", grade_document)
    workflow.add_node("generate", generate)
    workflow.add_node("transform_query", transform_query)

    # Graph
    workflow.add_conditional_edges(
        START,
        route_question,
        {
            "web_search": "web_search",
            "vectorstore": "retrieve",
        }
    )
    workflow.add_edge("web_search", "generate")
    workflow.add_edge("retrieve", "grade_documents")
    workflow.add_conditional_edges(
        "grade_documents",
        decide_to_generate,
        {
            "transform_query": "transform_query",
            "generate": "generate",
        }
    )

    workflow.add_edge("transform_query", "retrieve")
    workflow.add_conditional_edges(
        "generate",
        grade_generation_v_documents_and_question,
        {
            "not supported": "generate",
            "useful": END,
            "not useful": "transform_query",
        }
    )
    app = workflow.compile()
    display(Image(app.get_graph().draw_mermaid_png()))
    return app


app = workflow()
