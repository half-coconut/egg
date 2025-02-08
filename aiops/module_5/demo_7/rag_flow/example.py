from pprint import pprint

from aiops.module_5.demo_7.rag_flow.workflow import app

if __name__ == '__main__':
    # inputs = {"question": "谁管理的服务最多"}
    inputs = {"question": "北京最近一周天气如何"}

    for output in app.stream(inputs):
        for key, value in output.items():
            pprint(f"Node '{key}': ")

        pprint("\n--\n")

    pprint(value["generation"])